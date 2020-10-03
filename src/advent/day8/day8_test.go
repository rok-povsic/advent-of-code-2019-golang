package main

import (
	"testing"
)

func TestCompute(t *testing.T) {
	output := Compute("123456789012", 3, 2)
	if output != 1 {
		t.Errorf("Unexpeced output %d", output)
	}
}
