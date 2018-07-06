package main

import (
	"flag"

	"github.com/go-logr/glogr"
	"github.com/golang/glog"
)

type E struct {
	str string
}

func (e E) Error() string {
	return e.str
}

func main() {
	flag.Set("v", "3")
	flag.Parse()
	log := glogr.New().WithName("MyName").WithValues("user", "you")
	log.Info("hello", "val1", 1, "val2", map[string]int{"k": 1})
	log.V(3).Info("nice to meet you")
	log.Error(nil, "uh oh", "trouble", true, "reasons", []float64{0.1, 0.11, 3.14})
	log.Error(E{"an error occurred"}, "goodbye", "code", -1)
	glog.Flush()
}
