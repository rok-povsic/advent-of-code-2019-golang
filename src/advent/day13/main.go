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

	programCode = "2" + programCode[1:] // Part 2

	actionChan := make(chan string)
	inputChan := make(chan int)
	outputChan := make(chan int)
	finishedChan := make(chan bool)
	go Compute(programCode, actionChan, inputChan, outputChan, finishedChan)

	maxX := 0
	maxY := 0

	grid := make(map[Coordinates]int)

	ballX, padX := -1, -1
	score := -1
	for true {
		action := <-actionChan
		if action == "NEED_INPUT" {
			ballX, padX = draw(grid, maxX, maxY)

			time.Sleep(10 * time.Millisecond)

			fmt.Printf("Score: %d\n", score)

			direction := 0
			if ballX < padX {
				direction = -1
			} else if ballX > padX {
				direction = 1
			}

			inputChan <- direction
			continue
		} else if action == "WILL_OUTPUT" {
		} else if action == "DONE" {
			fmt.Println("DONE")
			fmt.Printf("Final score: %d\n", score)
			return
		}

		x := <-outputChan
		if action == "WILL_OUTPUT" {
			_ = <-actionChan
		}
		y := <-outputChan
		if action == "WILL_OUTPUT" {
			_ = <-actionChan
		}
		tileId := <-outputChan

		if x == -1 && y == 0 {
			score = tileId
		}

		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		grid[Coordinates{x, y}] = tileId
	}
}

func draw(grid map[Coordinates]int, maxX int, maxY int) (int, int) {
	ballX := -1
	padX := -1

	for j := 0; j <= maxY; j++ {
		for i := 0; i <= maxX; i++ {
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
					padX = i
				case 4:
					fmt.Print("O")
					ballX = i
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	return ballX, padX
}

type Coordinates struct {
	x, y int
}
