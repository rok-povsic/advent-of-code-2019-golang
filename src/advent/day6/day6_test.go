package main

import "testing"

func TestCompute(t *testing.T) {
	orbitalMap := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN"
	allEdges, nearestPath := Compute(orbitalMap)
	if allEdges != 54 {
		t.Errorf("Unexpected result: %d", allEdges)
	}
	if nearestPath != 4 {
		t.Errorf("Unexpected result: %d", nearestPath)
	}
}
