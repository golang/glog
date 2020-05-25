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

func TestSetPrefixf(t *testing.T) {
	g := &Glogger{Prefix: "zero "}
	expected := "zero one two"
	og := "one two"
	withPre := setPrefixf(g, og)

	if withPre != expected {
		t.Errorf("Resulting string did not match expected: %q != %q", withPre, expected)
	}

	// Should fail to add prefix
	g.Prefix = 0
	withPre = setPrefixf(g, og)

	if withPre != og {
		t.Errorf("Resulting string was expected to be unchanged (%q). Intead we got: %q", og, withPre)
	}
}
