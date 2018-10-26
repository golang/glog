package main

import (
	"flag"
	"github.com/dims/klog"
)

func main() {
	flag.Set("log_file", "myfile.log")
	flag.Parse()
	klog.Info("nice to meet you")
	klog.Flush()
}
