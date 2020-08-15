package glog

import (
	"fmt"
	stdLog "log"
	"sync/atomic"
)

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Info(args ...interface{}) {
	logging.print(infoLog, args...)
}

// InfoDepth acts as Info but uses depth to determine which call frame to log.
// InfoDepth(0, "msg") is the same as Info("msg").
func InfoDepth(depth int, args ...interface{}) {
	logging.printDepth(infoLog, depth, args...)
}

// InfoDepthf acts as Info but uses depth to determine which call frame to log,
// Arguments format and args act as Infof.
func InfoDepthf(depth int, format string, args ...interface{}) {
	logging.printDepthf(infoLog, depth, format, args...)
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Infoln(args ...interface{}) {
	logging.println(infoLog, args...)
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Infof(format string, args ...interface{}) {
	logging.printf(infoLog, format, args...)
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Warning(args ...interface{}) {
	logging.print(warningLog, args...)
}

// WarningDepth acts as Warning but uses depth to determine which call frame to log.
// WarningDepth(0, "msg") is the same as Warning("msg").
func WarningDepth(depth int, args ...interface{}) {
	logging.printDepth(warningLog, depth, args...)
}

// WarningDepthf acts as Warning but uses depth to determine which call frame to log,
// Arguments format and args act as Warningf.
func WarningDepthf(depth int, format string, args ...interface{}) {
	logging.printDepthf(warningLog, depth, format, args...)
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Warningln(args ...interface{}) {
	logging.println(warningLog, args...)
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Warningf(format string, args ...interface{}) {
	logging.printf(warningLog, format, args...)
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Error(args ...interface{}) {
	logging.print(errorLog, args...)
}

// ErrorDepth acts as Error but uses depth to determine which call frame to log.
// ErrorDepth(0, "msg") is the same as Error("msg").
func ErrorDepth(depth int, args ...interface{}) {
	logging.printDepth(errorLog, depth, args...)
}

// ErrorDepthf acts as Error but uses depth to determine which call frame to log.
// Arguments format and args act as Errorf.
func ErrorDepthf(depth int, format string, args ...interface{}) {
	logging.printDepthf(errorLog, depth, format, args...)
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Errorln(args ...interface{}) {
	logging.println(errorLog, args...)
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Errorf(format string, args ...interface{}) {
	logging.printf(errorLog, format, args...)
}

// Fatal logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Fatal(args ...interface{}) {
	logging.print(fatalLog, args...)
}

// FatalDepth acts as Fatal but uses depth to determine which call frame to log.
// FatalDepth(0, "msg") is the same as Fatal("msg").
func FatalDepth(depth int, args ...interface{}) {
	logging.printDepth(fatalLog, depth, args...)
}

// FatalDepthf acts as Fatal but uses depth to determine which call frame to log,
// Arguments format and args act as Exitf.
func FatalDepthf(depth int, format string, args ...interface{}) {
	logging.printDepthf(fatalLog, depth, format, args...)
}

// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Fatalln(args ...interface{}) {
	logging.println(fatalLog, args...)
}

// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Fatalf(format string, args ...interface{}) {
	logging.printf(fatalLog, format, args...)
}

// fatalNoStacks is non-zero if we are to exit without dumping goroutine stacks.
// It allows Exit and relatives to use the Fatal logs.
var fatalNoStacks uint32

// Exit logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Exit(args ...interface{}) {
	atomic.StoreUint32(&fatalNoStacks, 1)
	logging.print(fatalLog, args...)
}

// ExitDepth acts as Exit but uses depth to determine which call frame to log.
// ExitDepth(0, "msg") is the same as Exit("msg").
func ExitDepth(depth int, args ...interface{}) {
	atomic.StoreUint32(&fatalNoStacks, 1)
	logging.printDepth(fatalLog, depth, args...)
}

// ExitDepthf acts as Exit but uses depth to determine which call frame to log,
// Arguments format and args act as Exitf.
func ExitDepthf(depth int, format string, args ...interface{}) {
	logging.printDepthf(fatalLog, depth, format, args...)
}

// Exitln logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
func Exitln(args ...interface{}) {
	atomic.StoreUint32(&fatalNoStacks, 1)
	logging.println(fatalLog, args...)
}

// Exitf logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Exitf(format string, args ...interface{}) {
	atomic.StoreUint32(&fatalNoStacks, 1)
	logging.printf(fatalLog, format, args...)
}

// Flush flushes all pending log I/O.
func Flush() {
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
	sev, ok := severityByName(name)
	if !ok {
		panic(fmt.Sprintf("log.CopyStandardLogTo(%q): unrecognized severity name", name))
	}
	// Set a log format that captures the user's file and line:
	//   d.go:23: message
	stdLog.SetFlags(stdLog.Lshortfile)
	stdLog.SetOutput(logBridge(sev))
}
