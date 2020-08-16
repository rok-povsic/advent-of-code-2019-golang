package main

import "testing"

func TestCompute(t *testing.T) {
	program1 := "3,9,8,9,10,9,4,9,99,-1,8"
	output := Compute(8, program1)
	if output != "1\n" {
		t.Errorf("Unexpected output %s", output)
	}
	output = Compute(7, program1)
	if output != "0\n" {
		t.Errorf("Unexpected output %s", output)
	}

	program2 := "3,9,7,9,10,9,4,9,99,-1,8"
	output = Compute(7, program2)
	if output != "1\n" {
		t.Errorf("Unexpected output %s", output)
	}
	output = Compute(9, program2)
	if output != "0\n" {
		t.Errorf("Unexpected output %s", output)
	}

	program3 := "3,3,1108,-1,8,3,4,3,99"
	output = Compute(8, program3)
	if output != "1\n" {
		t.Errorf("Unexpected output %s", output)
	}
	output = Compute(-8, program3)
	if output != "0\n" {
		t.Errorf("Unexpected output %s", output)
	}

	program4 := "3,3,1107,-1,8,3,4,3,99"
	output = Compute(7, program4)
	if output != "1\n" {
		t.Errorf("Unexpected output %s", output)
	}
	output = Compute(9, program4)
	if output != "0\n" {
		t.Errorf("Unexpected output %s", output)
	}

	program5 := "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9"
	output = Compute(2, program5)
	if output != "1\n" {
		t.Errorf("Unexpected output %s", output)
	}

	output = Compute(0, program5)
	if output != "0\n" {
		t.Errorf("Unexpected output %s", output)
	}

	program6 := "3,3,1105,-1,9,1101,0,0,12,4,12,99,1"
	output = Compute(2, program6)
	if output != "1\n" {
		t.Errorf("Unexpected output %s", output)
	}

	output = Compute(0, program6)
	if output != "0\n" {
		t.Errorf("Unexpected output %s", output)
	}

	program7 := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
	output = Compute(7, program7)
	if output != "999\n" {
		t.Errorf("Unexpected output %s", output)
	}
	output = Compute(8, program7)
	if output != "1000\n" {
		t.Errorf("Unexpected output %s", output)
	}
	output = Compute(9, program7)
	if output != "1001\n" {
		t.Errorf("Unexpected output %s", output)
	}
}
