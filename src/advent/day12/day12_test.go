package main

import (
	"testing"
)

func TestSample1(t *testing.T) {
	moons := []Moon{
		MakeMoon(-1, 0, 2),
		MakeMoon(2, -10, -7),
		MakeMoon(4, -8, 8),
		MakeMoon(3, 5, -1),
	}
	totalEnergy := Simulate(moons, 10)
	if totalEnergy != 179 {
		t.Errorf("Unexpected total energy: %d\n", totalEnergy)
	}
}

func TestSample2(t *testing.T) {
	moons := []Moon{
		MakeMoon(-8, -10, 0),
		MakeMoon(5, 5, 10),
		MakeMoon(2, -7, 3),
		MakeMoon(9, -8, -3),
	}
	totalEnergy := Simulate(moons, 100)
	if totalEnergy != 1940 {
		t.Errorf("Unexpected total energy: %d\n", totalEnergy)
	}
}
