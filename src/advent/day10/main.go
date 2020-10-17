package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

type Location struct {
	x, y int
}

func BestAsteroidLocation(asteroidMap string) (int, int, int) {
	locations := parseLocations(asteroidMap)

	mostVisible := -1
	var mostVisibleLocation Location
	for _, location := range locations {
		fmt.Println(location)
		uniqueAngles := make(map[float64]bool)
		for _, otherLocation := range locations {
			if location == otherLocation {
				continue
			}
			fmt.Print("\t")
			fmt.Print(otherLocation)
			fmt.Print(" has angle of ")
			a := angle(location, otherLocation)
			fmt.Print(a)
			fmt.Println()
			uniqueAngles[a] = true
		}
		fmt.Printf("# of visible asteriods: %d\n", len(uniqueAngles))

		if mostVisible < len(uniqueAngles) {
			mostVisible = len(uniqueAngles)
			mostVisibleLocation = location
		}
	}

	return mostVisibleLocation.x, mostVisibleLocation.y, mostVisible
}

func parseLocations(asteroidMap string) []Location {
	var locations []Location

	curX := 0
	curY := 0
	for _, r := range asteroidMap {
		char := string(r)
		if char == "\n" {
			curY++
			curX = 0
			continue
		}

		if char == "#" {
			locations = append(locations, Location{curX, curY})
		}

		curX++
	}
	return locations
}

func angle(location Location, otherLocation Location) float64 {
	deltaX := float64(otherLocation.x - location.x)
	deltaY := float64(location.y - otherLocation.y)
	thetaRadians := math.Atan2(deltaY, deltaX)
	return thetaRadians
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("src/advent/day10/input.txt")
	check(err)
	asteroidMap := string(data)
	_, _, num := BestAsteroidLocation(asteroidMap)
	fmt.Println(num)
}
