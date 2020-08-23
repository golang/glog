package slog

import (
	"fmt"
	stdLog "log"
	"os"
	"runtime"
	"strings"
	"sync/atomic"
)

func beforeParseOutput(args ...interface{}) {
	var sb strings.Builder
	sb.WriteString("ERROR: slog before flag.Parse():")
	sb.WriteString(fmt.Sprint(args...))
	fmt.Fprintln(os.Stderr, sb.String())
}

func beforeParseOutputf(format string, args ...interface{}) {
	var sb strings.Builder
	sb.WriteString("ERROR: slog before flag.Parse():")
	sb.WriteString(fmt.Sprintf(format, args...))
	fmt.Fprintln(os.Stderr, sb.String())
}

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Info(args ...interface{}) {
	InitLogging()
	logging.print(infoLog, args...)
}

// InfoDepth acts as Info but uses depth to determine which call frame to log.
// InfoDepth(0, "msg") is the same as Info("msg").
func InfoDepth(depth int, args ...interface{}) {
	InitLogging()
	logging.printDepth(infoLog, depth, args...)
}

// InfoDepthf acts as Info but uses depth to determine which call frame to log,
// Arguments format and args act as Infof.
func InfoDepthf(depth int, format string, args ...interface{}) {
	InitLogging()
	logging.printDepthf(infoLog, depth, format, args...)
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Infoln(args ...interface{}) {
	InitLogging()
	logging.println(infoLog, args...)
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Infof(format string, args ...interface{}) {
	InitLogging()
	logging.printf(infoLog, format, args...)
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Warning(args ...interface{}) {
	InitLogging()
	logging.print(warningLog, args...)
}

// WarningDepth acts as Warning but uses depth to determine which call frame to log.
// WarningDepth(0, "msg") is the same as Warning("msg").
func WarningDepth(depth int, args ...interface{}) {
	InitLogging()
	logging.printDepth(warningLog, depth, args...)
}

// WarningDepthf acts as Warning but uses depth to determine which call frame to log,
// Arguments format and args act as Warningf.
func WarningDepthf(depth int, format string, args ...interface{}) {
	InitLogging()
	logging.printDepthf(warningLog, depth, format, args...)
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Warningln(args ...interface{}) {
	InitLogging()
	logging.println(warningLog, args...)
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Warningf(format string, args ...interface{}) {
	InitLogging()
	logging.printf(warningLog, format, args...)
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Error(args ...interface{}) {
	InitLogging()
	logging.print(errorLog, args...)
}

// ErrorDepth acts as Error but uses depth to determine which call frame to log.
// ErrorDepth(0, "msg") is the same as Error("msg").
func ErrorDepth(depth int, args ...interface{}) {
	InitLogging()
	logging.printDepth(errorLog, depth, args...)
}

// ErrorDepthf acts as Error but uses depth to determine which call frame to log.
// Arguments format and args act as Errorf.
func ErrorDepthf(depth int, format string, args ...interface{}) {
	InitLogging()
	logging.printDepthf(errorLog, depth, format, args...)
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Errorln(args ...interface{}) {
	InitLogging()
	logging.println(errorLog, args...)
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Errorf(format string, args ...interface{}) {
	InitLogging()
	logging.printf(errorLog, format, args...)
}

// Fatal logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Fatal(args ...interface{}) {
	InitLogging()
	logging.print(fatalLog, args...)
}

// FatalDepth acts as Fatal but uses depth to determine which call frame to log.
// FatalDepth(0, "msg") is the same as Fatal("msg").
func FatalDepth(depth int, args ...interface{}) {
	InitLogging()
	logging.printDepth(fatalLog, depth, args...)
}

// FatalDepthf acts as Fatal but uses depth to determine which call frame to log,
// Arguments format and args act as Exitf.
func FatalDepthf(depth int, format string, args ...interface{}) {
	InitLogging()
	logging.printDepthf(fatalLog, depth, format, args...)
}

// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Fatalln(args ...interface{}) {
	InitLogging()
	logging.println(fatalLog, args...)
}

// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Fatalf(format string, args ...interface{}) {
	InitLogging()
	logging.printf(fatalLog, format, args...)
}

// fatalNoStacks is non-zero if we are to exit without dumping goroutine stacks.
// It allows Exit and relatives to use the Fatal logs.
var fatalNoStacks uint32

// Exit logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Exit(args ...interface{}) {
	InitLogging()
	atomic.StoreUint32(&fatalNoStacks, 1)
	logging.print(fatalLog, args...)
}

// ExitDepth acts as Exit but uses depth to determine which call frame to log.
// ExitDepth(0, "msg") is the same as Exit("msg").
func ExitDepth(depth int, args ...interface{}) {
	InitLogging()
	atomic.StoreUint32(&fatalNoStacks, 1)
	logging.printDepth(fatalLog, depth, args...)
}

// ExitDepthf acts as Exit but uses depth to determine which call frame to log,
// Arguments format and args act as Exitf.
func ExitDepthf(depth int, format string, args ...interface{}) {
	InitLogging()
	logging.printDepthf(fatalLog, depth, format, args...)
}

// Exitln logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
func Exitln(args ...interface{}) {
	InitLogging()
	atomic.StoreUint32(&fatalNoStacks, 1)
	logging.println(fatalLog, args...)
}

// Exitf logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Exitf(format string, args ...interface{}) {
	InitLogging()
	atomic.StoreUint32(&fatalNoStacks, 1)
	logging.printf(fatalLog, format, args...)
}

// Flush flushes all pending log I/O.
func Flush() {
	InitLogging()
	// loggingOnce.Do(initLogging)
	logging.lockAndFlushAll()
}

// CopyStandardLogTo arranges for messages written to the Go "log" package's
// default logs to also appear in the Google logs for the named and lower
// severities.  Subsequent changes to the standard log's default output location
// or format may break this behavior.
//
// Valid names are "INFO", "WARNING", "ERROR", and "FATAL".  If the name is not
// recognized, CopyStandardLogTo panics.
func CopyStandardLogTo(name string) {
	InitLogging()
	// loggingOnce.Do(initLogging)
	sev, ok := severityByName(name)
	if !ok {
		panic(fmt.Sprintf("log.CopyStandardLogTo(%q): unrecognized severity name", name))
	}
	// Set a log format that captures the user's file and line:
	//   d.go:23: message
	stdLog.SetFlags(stdLog.Lshortfile)
	stdLog.SetOutput(logBridge(sev))
}

// Verbose is a boolean type that implements Infof (like Printf) etc.
// See the documentation of V for more information.
type Verbose bool

// V reports whether verbosity at the call site is at least the requested level.
// The returned value is a boolean of type Verbose, which implements Info, Infoln
// and Infof. These methods will write to the Info log if called.
// Thus, one may write either
//	if slog.V(2) { slog.Info("log this") }
// or
//	slog.V(2).Info("log this")
// The second form is shorter but the first is cheaper if logging is off because it does
// not evaluate its arguments.
//
// Whether an individual call to V generates a log record depends on the setting of
// the -v and --vmodule flags; both are off by default. If the level in the call to
// V is at least the value of -v, or of -vmodule for the source file containing the
// call, the V call will log.
func V(level Level) Verbose {
	InitLogging()
	// This function tries hard to be cheap unless there's work to do.
	// The fast path is two atomic loads and compares.

	// Here is a cheap but safe test to see if V logging is enabled globally.
	if logging.verbosity.get() >= level {
		return Verbose(true)
	}

	// It's off globally but it vmodule may still be set.
	// Here is another cheap but safe test to see if vmodule is enabled.
	if atomic.LoadInt32(&logging.filterLength) > 0 {
		// Now we need a proper lock to use the logging structure. The pcs field
		// is shared so we must lock before accessing it. This is fairly expensive,
		// but if V logging is enabled we're slow anyway.
		logging.mu.Lock()
		defer logging.mu.Unlock()
		if runtime.Callers(2, logging.pcs[:]) == 0 {
			return Verbose(false)
		}
		v, ok := logging.vmap[logging.pcs[0]]
		if !ok {
			v = logging.setV(logging.pcs[0])
		}
		return Verbose(v >= level)
	}
	return Verbose(false)
}

// Info is equivalent to the global Info function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Info(args ...interface{}) {
	InitLogging()
	if v {
		logging.print(infoLog, args...)
	}
}

// Infoln is equivalent to the global Infoln function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Infoln(args ...interface{}) {
	InitLogging()
	if v {
		logging.println(infoLog, args...)
	}
}

// Infof is equivalent to the global Infof function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Infof(format string, args ...interface{}) {
	InitLogging()
	if v {
		logging.printf(infoLog, format, args...)
	}
}
