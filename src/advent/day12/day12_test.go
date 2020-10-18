package main

import (
	"testing"
)

func TestSample1(t *testing.T) {
	moons := moons1()
	totalEnergy := Simulate(moons, 10)
	if totalEnergy != 179 {
		t.Errorf("Unexpected total energy: %d\n", totalEnergy)
	}
}

func TestRepeatStepsSample1(t *testing.T) {
	steps := RepeatSteps(moons1(), moons1())
	if steps != 2772 {
		t.Errorf("Unexpected steps: %d\n", steps)
	}
}

func TestSample2(t *testing.T) {
	moons := moons2()
	totalEnergy := Simulate(moons, 100)
	if totalEnergy != 1940 {
		t.Errorf("Unexpected total energy: %d\n", totalEnergy)
	}
}

func TestRepeatStepsSample2(t *testing.T) {
	steps := RepeatSteps(moons2(), moons2())
	if steps != 4686774924 {
		t.Errorf("Unexpected steps: %d\n", steps)
	}
}

func moons1() []Moon {
	moons := []Moon{
		MakeMoon(-1, 0, 2),
		MakeMoon(2, -10, -7),
		MakeMoon(4, -8, 8),
		MakeMoon(3, 5, -1),
	}
	return moons
}

func moons2() []Moon {
	moons := []Moon{
		MakeMoon(-8, -10, 0),
		MakeMoon(5, 5, 10),
		MakeMoon(2, -7, 3),
		MakeMoon(9, -8, -3),
	}
	return moons
}
