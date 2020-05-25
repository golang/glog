package glog

import (
	"testing"
)

const (
	zero int = iota
	one
)
const two = "two"

func TestSetPrefix(t *testing.T) {
	g := &Glogger{Prefix: zero}
	expected := []interface{}{zero, one, two}

	x := []interface{}{one, two}
	y := setPrefix(g, x)

	for i, v := range y {
		if v != expected[i] {
			t.Errorf("Resulting slice did not match expected: %q != %q at index %d", v, expected[i], i)
		}
	}
}
