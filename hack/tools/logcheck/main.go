/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/exp/utf8string"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

// Doc explaining the tool.
const Doc = "Tool to check use of unstructured logging patterns."

// Analyzer runs static analysis.
var Analyzer = &analysis.Analyzer{
	Name: "logcheck",
	Doc:  Doc,
	Run:  run,
}

func main() {
	singlechecker.Main(Analyzer)
}

func run(pass *analysis.Pass) (interface{}, error) {

	for _, file := range pass.Files {

		ast.Inspect(file, func(n ast.Node) bool {

			// We are intrested in function calls, as we want to detect klog.* calls
			// passing all function calls to checkForFunctionExpr
			if fexpr, ok := n.(*ast.CallExpr); ok {
				checkForFunctionExpr(fexpr, pass)
			}

			return true
		})
	}
	return nil, nil
}

// checkForFunctionExpr checks for unstructured logging function, prints error if found any.
func checkForFunctionExpr(fexpr *ast.CallExpr, pass *analysis.Pass) {

	fun := fexpr.Fun
	args := fexpr.Args

	/* we are extracting external package function calls e.g. klog.Infof fmt.Printf
	   and eliminating calls like setLocalHost()
	   basically function calls that has selector expression like .
	*/
	if selExpr, ok := fun.(*ast.SelectorExpr); ok {
		// extracting function Name like Infof
		fName := selExpr.Sel.Name

		// for nested function cases klog.V(1).Infof scenerios
		// if selExpr.X contains one more caller expression which is selector expression
		// we are extracting klog and discarding V(1)
		if n, ok := selExpr.X.(*ast.CallExpr); ok {
			if _, ok = n.Fun.(*ast.SelectorExpr); ok {
				selExpr = n.Fun.(*ast.SelectorExpr)
			}
		}

		// extracting package name
		pName, ok := selExpr.X.(*ast.Ident)

		if ok && pName.Name == "klog" {
			// Matching if any unstructured logging function is used.
			if !isUnstructured((fName)) {
				// if format specifier is used, check for arg length will most probably fail
				// so check for format specifier first and skip if found
				if checkForFormatSpecifier(fexpr, pass) {
					return
				}
				if fName == "InfoS" {
					isKeysValid(args[1:], fun, pass, fName)
				} else if fName == "ErrorS" {
					isKeysValid(args[2:], fun, pass, fName)
				}
			} else {
				msg := fmt.Sprintf("unstructured logging function %q should not be used", fName)
				pass.Report(analysis.Diagnostic{
					Pos:     fun.Pos(),
					Message: msg,
				})
			}
		}
	}
}

func isUnstructured(fName string) bool {

	// List of klog functions we do not want to use after migration to structured logging.
	unstrucured := []string{
		"Infof", "Info", "Infoln", "InfoDepth",
		"Warning", "Warningf", "Warningln", "WarningDepth",
		"Error", "Errorf", "Errorln", "ErrorDepth",
		"Fatal", "Fatalf", "Fatalln", "FatalDepth",
	}

	for _, name := range unstrucured {
		if fName == name {
			return true
		}
	}

	return false
}

// isKeysValid check if all keys in keyAndValues is string type
func isKeysValid(keyValues []ast.Expr, fun ast.Expr, pass *analysis.Pass, funName string) {
	if len(keyValues)%2 != 0 {
		pass.Report(analysis.Diagnostic{
			Pos:     fun.Pos(),
			Message: fmt.Sprintf("Additional arguments to %s should always be Key Value pairs. Please check if there is any key or value missing.", funName),
		})
		return
	}

	for index, arg := range keyValues {
		if index%2 != 0 {
			continue
		}
		lit, ok := arg.(*ast.BasicLit)
		if !ok {
			pass.Report(analysis.Diagnostic{
				Pos:     fun.Pos(),
				Message: fmt.Sprintf("Key positional arguments are expected to be inlined constant strings. Please replace %v provided with string value", arg),
			})
			continue
		}
		if lit.Kind != token.STRING {
			pass.Report(analysis.Diagnostic{
				Pos:     fun.Pos(),
				Message: fmt.Sprintf("Key positional arguments are expected to be inlined constant strings. Please replace %v provided with string value", lit.Value),
			})
			continue
		}
		isASCII := utf8string.NewString(lit.Value).IsASCII()
		if !isASCII {
			pass.Report(analysis.Diagnostic{
				Pos:     fun.Pos(),
				Message: fmt.Sprintf("Key positional arguments %s are expected to be lowerCamelCase alphanumeric strings. Please remove any non-Latin characters.", lit.Value),
			})
		}
	}
}

func checkForFormatSpecifier(expr *ast.CallExpr, pass *analysis.Pass) bool {
	if selExpr, ok := expr.Fun.(*ast.SelectorExpr); ok {
		// extracting function Name like Infof
		fName := selExpr.Sel.Name
		if specifier, found := hasFormatSpecifier(expr.Args); found {
			msg := fmt.Sprintf("structured logging function %q should not use format specifier %q", fName, specifier)
			pass.Report(analysis.Diagnostic{
				Pos:     expr.Fun.Pos(),
				Message: msg,
			})
			return true
		}
	}
	return false
}

func hasFormatSpecifier(fArgs []ast.Expr) (string, bool) {
	formatSpecifiers := []string{
		"%v", "%+v", "%#v", "%T",
		"%t", "%b", "%c", "%d", "%o", "%O", "%q", "%x", "%X", "%U",
		"%e", "%E", "%f", "%F", "%g", "%G", "%s", "%q", "%p",
	}
	for _, fArg := range fArgs {
		if arg, ok := fArg.(*ast.BasicLit); ok {
			for _, specifier := range formatSpecifiers {
				if strings.Contains(arg.Value, specifier) {
					return specifier, true
				}
			}
		}
	}
	return "", false
}
