package day2

import "testing"

func TestIntcode(t *testing.T) {
	result := Compute("1,9,10,3,2,3,11,0,99,30,40,50")
	if result != "3500,9,10,70,2,3,11,0,99,30,40,50" {
		t.Errorf("Fail %s", result)
	}

	result = Compute("1,0,0,0,99")
	if result != "2,0,0,0,99" {
		t.Errorf("Fail %s", result)
	}

	result = Compute("2,3,0,3,99")
	if result != "2,3,0,6,99" {
		t.Errorf("Fail %s", result)
	}

	result = Compute("2,4,4,5,99,0")
	if result != "2,4,4,5,99,9801" {
		t.Errorf("Fail %s", result)
	}

	result = Compute("1,1,1,4,99,5,6,0,99")
	if result != "30,1,1,4,2,5,6,0,99" {
		t.Errorf("Fail %s", result)
	}
}