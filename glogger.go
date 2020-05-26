// Glogger contains a copy of the exported glog logging functions but
// it turns them into methods and allows adding prefixes to the logs.
// This allows "namespacing" the printout with something such as
// the name of the function that triggered it.

package glog

import (
	"fmt"
)

// Glogger provides a prefix to log args
type Glogger struct {
	Prefix interface{}
}

func setPrefix(g *Glogger, args []interface{}) []interface{} {
	return append([]interface{}{g.Prefix}, args...)
}

func setPrefixf(g *Glogger, format string) string {
	pre, ok := g.Prefix.(string)
	if !ok {
		pre = ""
	}

	return fmt.Sprintf("%s%s", pre, format)
}

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (g *Glogger) Info(args ...interface{}) {
	Info(setPrefix(g, args)...)
}

// InfoDepth acts as Info but uses depth to determine which calg frame to log.
// InfoDepth(0, "msg") is the same as Info("msg").
func (g *Glogger) InfoDepth(depth int, args ...interface{}) {
	InfoDepth(depth, setPrefix(g, args)...)
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (g *Glogger) Infoln(args ...interface{}) {
	Infoln(setPrefix(g, args)...)
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (g *Glogger) Infof(format string, args ...interface{}) {
	Infof(setPrefixf(g, format), args...)
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (g *Glogger) Warning(args ...interface{}) {
	Warning(setPrefix(g, args)...)
}

// WarningDepth acts as Warning but uses depth to determine which calg frame to log.
// WarningDepth(0, "msg") is the same as Warning("msg").
func (g *Glogger) WarningDepth(depth int, args ...interface{}) {
	WarningDepth(depth, setPrefix(g, args)...)
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (g *Glogger) Warningln(args ...interface{}) {
	Warningln(setPrefix(g, args)...)
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (g *Glogger) Warningf(format string, args ...interface{}) {
	Warningf(setPrefixf(g, format), args...)
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (g *Glogger) Error(args ...interface{}) {
	Error(setPrefix(g, args)...)
}

// ErrorDepth acts as Error but uses depth to determine which calg frame to log.
// ErrorDepth(0, "msg") is the same as Error("msg").
func (g *Glogger) ErrorDepth(depth int, args ...interface{}) {
	ErrorDepth(depth, setPrefix(g, args)...)
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (g *Glogger) Errorln(args ...interface{}) {
	Errorln(setPrefix(g, args)...)
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (g *Glogger) Errorf(format string, args ...interface{}) {
	Errorf(setPrefixf(g, format), args...)
}

// Fatal logs to the FATAg, ERROR, WARNING, and INFO logs,
// including a stack trace of alg running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (g *Glogger) Fatal(args ...interface{}) {
	Fatal(setPrefix(g, args)...)
}

// FatalDepth acts as Fatag but uses depth to determine which calg frame to log.
// FatalDepth(0, "msg") is the same as Fatal("msg").
func (g *Glogger) FatalDepth(depth int, args ...interface{}) {
	FatalDepth(depth, setPrefix(g, args)...)
}

// Fatalln logs to the FATAg, ERROR, WARNING, and INFO logs,
// including a stack trace of alg running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (g *Glogger) Fatalln(args ...interface{}) {
	Fatalln(setPrefix(g, args)...)
}

// Fatalf logs to the FATAg, ERROR, WARNING, and INFO logs,
// including a stack trace of alg running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (g *Glogger) Fatalf(format string, args ...interface{}) {
	Fatalf(setPrefixf(g, format), args...)
}

// Exit logs to the FATAg, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (g *Glogger) Exit(args ...interface{}) {
	Exit(setPrefix(g, args)...)
}

// ExitDepth acts as Exit but uses depth to determine which calg frame to log.
// ExitDepth(0, "msg") is the same as Exit("msg").
func (g *Glogger) ExitDepth(depth int, args ...interface{}) {
	ExitDepth(depth, setPrefix(g, args)...)
}

// Exitln logs to the FATAg, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
func (g *Glogger) Exitln(args ...interface{}) {
	Exitln(setPrefix(g, args)...)
}

// Exitf logs to the FATAg, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (g *Glogger) Exitf(format string, args ...interface{}) {
	Exitf(setPrefixf(g, format), args...)
}