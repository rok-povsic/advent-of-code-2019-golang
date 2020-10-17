package main

import (
	"strings"
	"testing"
)

func TestCompute(t *testing.T) {
	output := Compute(-1, "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")
	outputWithCommas := strings.ReplaceAll(strings.Trim(output, "\n"), "\n", ",")
	if outputWithCommas != "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99" {
		t.Errorf("Unexpected output %s", outputWithCommas)
	}

	output = Compute(-1, "1102,34915192,34915192,7,4,7,99,0")
	if len(strings.Trim(output, "\n")) != 16 {
		t.Errorf("Unexpected output %s", output)
	}

	output = Compute(-1, "104,1125899906842624,99")
	if output != "1125899906842624\n" {
		t.Errorf("Unexpected output %s", output)
	}
}
