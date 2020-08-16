package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Run() {
	dat, err := ioutil.ReadFile("src/advent/day6/input.txt")
	check(err)

	text := string(dat)
	fmt.Println(Compute(text))
}

func Compute(orbitalMap string) int {
	allObjects := make(map[string]bool)
	directOrbits := make(map[string]string)

	lines := strings.Split(orbitalMap, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ")")
		orbited := parts[0]
		orbitor := parts[1]

		allObjects[orbited] = true
		allObjects[orbitor] = true
		directOrbits[orbitor] = orbited
	}

	count := 0
	for object, _ := range allObjects {
		orbited := object
		for true {
			orbitor, hasOrbitor := directOrbits[orbited]
			if !hasOrbitor {
				break
			}
			count++
			orbited = orbitor
		}
	}
	return count
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	Run()
}
