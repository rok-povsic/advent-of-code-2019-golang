package main

import "testing"

func TestCompute(t *testing.T) {
	orbitalMap := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L"
	result := Compute(orbitalMap)
	if result != 42 {
		t.Errorf("Unexpected result: %d", result)
	}
}
