package main

import (
	"flag"
	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	flag.Set("log_file", "myfile.log")
	flag.Parse()
	klog.Info("nice to meet you")
	klog.Flush()
}
