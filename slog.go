package slog

import (
	"flag"
	"io"
	"strconv"
	"sync"
	"time"

	_ "github.com/golang/glog" // load glog flag
)

var (
	// logging     *loggingT
	logging     loggingT
	loggingOnce sync.Once
)

const (
	flushInterval = 5 * time.Second
)

// flushSyncWriter is the interface satisfied by logging destinations.
type flushSyncWriter interface {
	Flush() error
	Sync() error
	io.Writer
}

// InitLogging 初始化 slog
func InitLogging() bool {
	if flag.Parsed() {
		loggingOnce.Do(initLogging)
		return true
	}
	return false
}

// initLogging
// 在API函数之前初始化一次
// 结合 loggingOnce 使用 loggingOnce.Do(initLogging)
// 为什么这么搞呢
// 原因是为了兼容和使用glog flags
// 像go test main函数，都会显示调用 flag.Parse()
// 等待glog go test XXXX 处理完flag后 slog才能Lookup已经处理的flags
func initLogging() {
	initLogFile()
	// logging = new(loggingT)
	logtostderrFlag := flag.Lookup("logtostderr")
	if logtostderrFlag == nil {
		logging.toStderr = false
	} else {
		logging.toStderr, _ = strconv.ParseBool(logtostderrFlag.Value.String())
	}
	alsologtostderrFlag := flag.Lookup("alsologtostderr")
	if alsologtostderrFlag == nil {
		logging.alsoToStderr = false
	} else {
		logging.alsoToStderr, _ = strconv.ParseBool(alsologtostderrFlag.Value.String())
	}

	stderrThresholdFlag := flag.Lookup("stderrthreshold")
	if stderrThresholdFlag != nil {
		logging.stderrThreshold.Set(stderrThresholdFlag.Value.String())
	} else {
		logging.stderrThreshold = errorLog
	}

	vFlag := flag.Lookup("v")
	if vFlag != nil {
		logging.verbosity.Set(vFlag.Value.String())
	} else {
		logging.setVState(0, nil, false)
	}

	vmoduleFlag := flag.Lookup("vmodule")
	if vmoduleFlag != nil {
		logging.vmodule.Set(vmoduleFlag.Value.String())
	}

	traceLocationFlag := flag.Lookup("log_backtrace_at")
	if traceLocationFlag != nil {
		logging.traceLocation.Set(traceLocationFlag.Value.String())
	}

	logDirFlag := flag.Lookup("log_dir")
	if logDirFlag == nil {
		logging.logDir = ""
	} else {
		logging.logDir = logDirFlag.Value.String()
	}
	go logging.flushDaemon()
}
