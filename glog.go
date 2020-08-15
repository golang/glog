package glog

import (
	"flag"
	"io"
	"time"
)

var (
	logging loggingT
)

const (
	flushInterval = 30 * time.Second
)

// flushSyncWriter is the interface satisfied by logging destinations.
type flushSyncWriter interface {
	Flush() error
	Sync() error
	io.Writer
}

func init() {
	initGlobalLogging()
	initLogFile()
}

func initGlobalLogging() {
	flag.BoolVar(&logging.toStderr, "logtostderr", false, "log to standard error instead of files")
	flag.BoolVar(&logging.alsoToStderr, "alsologtostderr", false, "log to standard error as well as files")
	flag.Var(&logging.verbosity, "v", "log level for V logs")
	flag.Var(&logging.stderrThreshold, "stderrthreshold", "logs at or above this threshold go to stderr")
	flag.Var(&logging.vmodule, "vmodule", "comma-separated list of pattern=N settings for file-filtered logging")
	flag.Var(&logging.traceLocation, "log_backtrace_at", "when logging hits line file:N, emit a stack trace")

	// Default stderrThreshold is ERROR.
	logging.stderrThreshold = errorLog

	logging.setVState(0, nil, false)
	go logging.flushDaemon()
}
