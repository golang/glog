package glog

import (
	"flag"
	"io"
	"time"
)

var (
	logging     loggingT
	slogFlagSet = flag.NewFlagSet("slog", flag.ExitOnError)
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
	slogFlagSet.BoolVar(&logging.toStderr, "logtostderr", false, "log to standard error instead of files")
	slogFlagSet.BoolVar(&logging.alsoToStderr, "alsologtostderr", false, "log to standard error as well as files")
	slogFlagSet.Var(&logging.verbosity, "v", "log level for V logs")
	slogFlagSet.Var(&logging.stderrThreshold, "stderrthreshold", "logs at or above this threshold go to stderr")
	slogFlagSet.Var(&logging.vmodule, "vmodule", "comma-separated list of pattern=N settings for file-filtered logging")
	slogFlagSet.Var(&logging.traceLocation, "log_backtrace_at", "when logging hits line file:N, emit a stack trace")

	// Default stderrThreshold is ERROR.
	logging.stderrThreshold = errorLog

	logging.setVState(0, nil, false)
	go logging.flushDaemon()
}
