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

func Compute(orbitalMap string) (int, int) {
	allObjects := make(map[string]bool)
	directOrbits := make(map[string]string)
	neighbors := make(map[string][]string)

	lines := strings.Split(orbitalMap, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ")")
		orbited := parts[0]
		orbitor := parts[1]

		allObjects[orbited] = true
		allObjects[orbitor] = true
		directOrbits[orbitor] = orbited

		orbitorNeighbors, contains := neighbors[orbitor]
		if !contains {
			orbitorNeighbors = make([]string, 0)
			neighbors[orbitor] = orbitorNeighbors
		}
		orbitedNeighbors, contains := neighbors[orbited]
		if !contains {
			orbitedNeighbors = make([]string, 0)
			neighbors[orbited] = orbitedNeighbors
		}

		neighbors[orbitor] = append(neighbors[orbitor], orbited)
		neighbors[orbited] = append(neighbors[orbited], orbitor)
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

	return count, bfs(neighbors) - 2
}

func bfs(neighbors map[string][]string) int {
	alreadyVisited := make(map[string]bool)

	nextList := make([]string, 0)
	nextList = append(nextList, "YOU")

	steps := 0
L:
	for len(nextList) > 0 {
		addLater := make([]string, 0)
		for _, next := range nextList {
			if next == "SAN" {
				break L
			}
			alreadyVisited[next] = true

			for _, neighbor := range neighbors[next] {
				_, contains := alreadyVisited[neighbor]
				if !contains {
					addLater = append(addLater, neighbor)
				}
			}
		}

		steps++

		nextList = make([]string, 0)
		for _, neighbor := range addLater {
			nextList = append(nextList, neighbor)
		}
	}

	return steps
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	Run()
}
