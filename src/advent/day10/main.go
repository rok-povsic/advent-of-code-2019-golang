package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
)

type Location struct {
	x, y int
}

func TwohundrethVaporizedAsteroid(asteroidMap string, deployedX int, deployedY int) (int, int) {
	asteroidLocations := parseLocations(asteroidMap)
	gunLocation := Location{deployedX, deployedY}

	// This is problematic due to floating point arithmetic but it turns out to work in this example, so no need
	// to spend more time on it.
	angleToLocations := make(map[float64][]Location)
	for _, asteroidLocation := range asteroidLocations {
		if gunLocation == asteroidLocation {
			continue
		}
		a := angle(gunLocation, asteroidLocation)
		_, exists := angleToLocations[a]
		if !exists {
			angleToLocations[a] = make([]Location, 0)
		}
		angleToLocations[a] = append(angleToLocations[a], asteroidLocation)
	}
	for _, locations := range angleToLocations {
		sort.Slice(locations, func(i, j int) bool {
			distanceToFirst := (locations[i].x - gunLocation.x) + (locations[i].y - gunLocation.y)
			distanceToSecond := (locations[j].x - gunLocation.x) + (locations[j].y - gunLocation.y)
			return distanceToFirst > distanceToSecond
		})
	}

	countAsteroids := 0
	for true {
		ascendingAngles := make([]float64, len(angleToLocations))
		i := 0
		for angle := range angleToLocations {
			ascendingAngles[i] = angle
			i++
		}
		sort.Float64s(ascendingAngles)
		for _, angle := range ascendingAngles {
			fmt.Println(angle)
			locations := angleToLocations[angle]
			vaporizedLocation := locations[0]
			fmt.Println(vaporizedLocation)
			locations = locations[1:]
			countAsteroids++

			if len(locations) == 0 {
				delete(angleToLocations, angle)
			}

			if countAsteroids == 200 {
				return vaporizedLocation.x, vaporizedLocation.y
			}
		}
	}
	return -1, -1
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
	thetaRadians := math.Atan2(deltaX, deltaY)
	// Rotate the coordinate system to start at the top.
	if thetaRadians < 0 {
		thetaRadians += 2 * math.Pi
	}
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
	x, y, num := BestAsteroidLocation(asteroidMap)
	fmt.Printf("gun location, x=%d, y=%d\n", x, y)
	fmt.Println(num)

	twohundrethX, twohundrethY := TwohundrethVaporizedAsteroid(asteroidMap, x, y)
	fmt.Println(twohundrethX*100 + twohundrethY)
}
