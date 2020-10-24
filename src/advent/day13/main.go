package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	data, err := ioutil.ReadFile("src/advent/day13/input.txt")
	check(err)
	programCode := string(data)

	inputChan := make(chan int)
	outputChan := make(chan int)
	finishedChan := make(chan bool)
	go Compute(programCode, inputChan, outputChan, finishedChan)

	maxX := 0
	maxY := 0

	grid := make(map[Coordinates]int)

	for true {
		x := <-outputChan
		y := <-outputChan
		tileId := <-outputChan

		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		grid[Coordinates{x, y}] = tileId

		fmt.Printf("%d %d %d\n", x, y, tileId)

		draw(grid, maxX, maxY)

		time.Sleep(10 * time.Millisecond)
	}
}

func draw(grid map[Coordinates]int, maxX int, maxY int) {
	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			tileId, exists := grid[Coordinates{i, j}]
			if exists {
				switch tileId {
				case 0:
					fmt.Print(" ")
				case 1:
					fmt.Print("#")
				case 2:
					fmt.Print("x")
				case 3:
					fmt.Print("-")
				case 4:
					fmt.Print("O")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println(" &")
	}
	fmt.Println("---------")
}

type Coordinates struct {
	x, y int
}
