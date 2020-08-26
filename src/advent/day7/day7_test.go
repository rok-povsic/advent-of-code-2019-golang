package main

import (
	"testing"
)

func TestCompute(t *testing.T) {
	program := "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
	phases := []int{4, 3, 2, 1, 0}
	output := MaxThrust(program, phases)
	if output != 43210 {
		t.Errorf("Unexpeced output %d", output)
	}

	program = "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
	phases = []int{0, 1, 2, 3, 4}
	output = MaxThrust(program, phases)
	if output != 54321 {
		t.Errorf("Unexpeced output %d", output)
	}

	program = "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
	phases = []int{1, 0, 4, 3, 2}
	output = MaxThrust(program, phases)
	if output != 65210 {
		t.Errorf("Unexpeced output %d", output)
	}
}
