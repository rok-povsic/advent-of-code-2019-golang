package main

import (
	"fmt"
	"io/ioutil"
)

type Coordinates struct {
	x, y int
}

func main() {
	data, err := ioutil.ReadFile("src/advent/day11/input.txt")
	check(err)
	text := string(data)

	inputChan := make(chan int)
	outputChan := make(chan int)
	finishedChan := make(chan bool)
	go Compute(text, inputChan, outputChan, finishedChan)

	// dX is 1 when turned right, -1 when turned left.
	// dY is 1 when turned up, -1 when turned down.
	// Exactly one of dX and dY is always zero.
	dX := 0
	dY := 1

	// Current coordinates.
	currentCoordinates := Coordinates{0, 0}
	colors := make(map[Coordinates]int) // coordinate to color (0 is black, 1 is white)

	for true {
		fmt.Printf("Current coordinates: x=%d, y=%d\n", currentCoordinates.x, currentCoordinates.y)
		fmt.Printf("# of painted grids so far: %d\n", len(colors))

		curColor, exists := colors[currentCoordinates]
		if !exists {
			curColor = 0
		}
		inputChan <- curColor
		newColor := <-outputChan
		rotation := <-outputChan
		fmt.Printf("newColor: %d\n", newColor)
		fmt.Printf("rotation: %d\n", rotation)

		colors[currentCoordinates] = newColor

		if rotation == 0 { // Turn left
			if dY == 1 {
				dX = -1
				dY = 0
			} else if dY == -1 {
				dX = 1
				dY = 0
			} else if dX == -1 {
				dX = 0
				dY = -1
			} else if dX == 1 {
				dX = 0
				dY = 1
			}
		} else if rotation == 1 { // Turn right
			if dY == 1 {
				dX = 1
				dY = 0
			} else if dY == -1 {
				dX = -1
				dY = 0
			} else if dX == -1 {
				dX = 0
				dY = 1
			} else if dX == 1 {
				dX = 0
				dY = -1
			}
		}

		currentCoordinates = Coordinates{currentCoordinates.x + dX, currentCoordinates.y + dY}
	}
}
