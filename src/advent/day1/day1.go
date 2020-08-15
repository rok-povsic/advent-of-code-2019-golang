package day1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Run() {
	masses := readData()

	additionalFuel := 0
	for _, mass := range masses {
		additionalFuel += RequiredFuelWithAdditionalFuel(mass)
	}

	fmt.Println("(Day 1 - Part 2) Required additional fuel: " + strconv.Itoa(additionalFuel))
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

func RequiredFuelWithAdditionalFuel(mass int) int {
	additionalFuel := requiredFuel(mass)
	totalFuel := additionalFuel
	for true {
		additionalFuel = requiredFuel(additionalFuel)
		if additionalFuel <= 0 {
			break
		}
		totalFuel += additionalFuel
	}

	return totalFuel
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}