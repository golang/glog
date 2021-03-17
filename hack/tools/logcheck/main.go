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

				checkForFunctionExpr(fexpr.Fun, fexpr.Args, pass)
			}

			return true
		})
	}
	return nil, nil
}

// checkForFunctionExpr checks for unstructured logging function, prints error if found any.
func checkForFunctionExpr(fun ast.Expr, args []ast.Expr, pass *analysis.Pass) {

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

		if ok && pName.Name == "klog" && !isUnstructured((fName)) {
			if fName == "InfoS" {
				isKeysValid(args[1:], fun, pass, fName)
			} else if fName == "ErrorS" {
				isKeysValid(args[2:], fun, pass, fName)
			}
		}
		// Matching if package name is klog and any unstructured logging function is used.
		if ok && pName.Name == "klog" && isUnstructured((fName)) {

			msg := fmt.Sprintf("unstructured logging function %q should not be used", fName)
			pass.Report(analysis.Diagnostic{
				Pos:     fun.Pos(),
				Message: msg,
			})
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
			Message: fmt.Sprintf("Invalid Number of arguments for %s", funName),
		})
	}

	for index, arg := range keyValues {
		if index%2 != 0 {
			continue
		}
		if lit, ok := arg.(*ast.BasicLit); ok {
			if lit.Kind != token.STRING {
				pass.Report(analysis.Diagnostic{
					Pos:     fun.Pos(),
					Message: fmt.Sprintf("Invalid value type for key %v", lit.Value),
				})
				continue
			}
			isASCII := utf8string.NewString(lit.Value).IsASCII()
			if !isASCII {
				pass.Report(analysis.Diagnostic{
					Pos:     fun.Pos(),
					Message: fmt.Sprintf("Invalid value for key %v, it's must be ascii", lit.Value),
				})
			}
		} else {
			pass.Report(analysis.Diagnostic{
				Pos:     fun.Pos(),
				Message: "Invalid value type for key",
			})
		}
	}
}
