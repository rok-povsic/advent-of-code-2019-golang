package main

import "testing"

func TestNearestManhattanDistance(t *testing.T) {
	distance := NearestManhattanDistance("R8,U5,L5,D3\nU7,R6,D4,L4")
	if distance != 6 {
		t.Errorf("Fail: %d", distance)
	}

	distance = NearestManhattanDistance("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83")
	if distance != 159 {
		t.Errorf("Fail: %d", distance)
	}

	distance = NearestManhattanDistance("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	if distance != 135 {
		t.Errorf("Fail: %d", distance)
	}
}

func TestNearestSignalDistance(t *testing.T) {
	distance := NearestSignalDistance("R8,U5,L5,D3\nU7,R6,D4,L4")
	if distance != 30 {
		t.Errorf("Fail: %d", distance)
	}

	distance = NearestSignalDistance("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83")
	if distance != 610 {
		t.Errorf("Fail: %d", distance)
	}

	distance = NearestSignalDistance("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	if distance != 410 {
		t.Errorf("Fail: %d", distance)
	}
}
