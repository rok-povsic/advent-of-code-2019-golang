package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	masses := readData()

	additionalFuel := 0
	for _, mass := range masses {
		additionalFuel += requiredFuel(mass)
	}

	fmt.Println("(Day 1) Required additional fuel: " + strconv.Itoa(additionalFuel))
}

func readData() []int {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	text := string(dat)
	lines := strings.Split(text, "\n")

	var masses []int
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		check(err)

		masses = append(masses, mass)
	}
	return masses
}

func requiredFuel(mass int) int {
	fuel := mass/3 - 2
	return fuel
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}