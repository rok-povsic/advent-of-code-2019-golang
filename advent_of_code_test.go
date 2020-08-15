package main

import "testing"

func TestRequiredFuelWithAdditionalFuel(t *testing.T) {
	result := RequiredFuelWithAdditionalFuel(1969)
	if result != 966 {
		t.Errorf("Incorrect, got %d, expected %d", result, 966)
	}

	result = RequiredFuelWithAdditionalFuel(100756)
	if result != 50346 {
		t.Errorf("Incorrect, got %d, expected %d", result, 50346)
	}
}
