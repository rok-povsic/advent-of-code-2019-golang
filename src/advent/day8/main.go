package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func Run() {
	dat, err := ioutil.ReadFile("src/advent/day8/input.txt")
	check(err)

	text := string(dat)
	fmt.Println(Compute(text, 25, 6))
	Draw(text, 25, 6)
}

func Draw(text string, wide int, tall int) {
	digits := textToDigits(text)

	layerLength := wide * tall

	drawing := make([]int, 0)
	for i := 0; i < layerLength; i++ {
		drawing = append(drawing, 2)
	}

	numLayers := len(digits) / layerLength
	for layerIndex := numLayers - 1; layerIndex >= 0; layerIndex-- {
		from := layerIndex * layerLength
		to := (layerIndex + 1) * layerLength
		layer := digits[from:to]

		for i := 0; i < layerLength; i++ {
			if layer[i] != 2 {
				drawing[i] = layer[i]
			}
		}
	}

	for i := 0; i < layerLength; i++ {
		if drawing[i] == 0 {
			fmt.Print(" ")
		} else {
			fmt.Print(drawing[i])
		}

		if i != 0 && (i+1)%wide == 0 {
			fmt.Println()
		}
	}
}

func Compute(text string, wide int, tall int) int {
	digits := textToDigits(text)

	layerWithFewestZeros := make([]int, 0)
	curMinZeros := 9999999999999999

	layerLength := wide * tall
	numLayers := len(digits) / layerLength
	for layerIndex := 0; layerIndex < numLayers; layerIndex++ {
		from := layerIndex * layerLength
		to := (layerIndex + 1) * layerLength
		layer := digits[from:to]

		curZeros := countDigits(layer, 0)

		if curZeros < curMinZeros {
			curMinZeros = curZeros
			layerWithFewestZeros = layer
		}
	}

	return countDigits(layerWithFewestZeros, 1) * countDigits(layerWithFewestZeros, 2)
}

func countDigits(layer []int, digitToCount int) int {
	curZeros := 0
	for _, digit := range layer {
		if digit == digitToCount {
			curZeros++
		}
	}
	return curZeros
}

func textToDigits(text string) []int {
	result := make([]int, 0)
	for i := 0; i < len(text); i++ {
		char := string(text[i])
		digit, err := strconv.Atoi(char)
		check(err)
		result = append(result, digit)
	}
	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	Run()
}
