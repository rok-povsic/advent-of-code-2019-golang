package main

import (
	"fmt"
	"strconv"
)

func Run() {
	from := 240920
	to := 789857

	countPart1 := 0
	countPart2 := 0
	for num := from; num <= to; num++ {
		if isValidForPart1(num) {
			countPart1++
		}
		if isValidForPart2(num) {
			countPart2++
		}
	}
	fmt.Printf("Part 1: %d\n", countPart1)
	fmt.Printf("Part 2: %d\n", countPart2)
}

func isValidForPart1(num int) bool {
	str := strconv.Itoa(num)
	isDuplicated := false
	for i := 0; i < 5; i++ {
		cur, err := strconv.Atoi(str[i : i+1])
		check(err)
		next, err := strconv.Atoi(str[i+1 : i+2])
		check(err)
		if cur > next {
			return false
		}
		if cur == next {
			isDuplicated = true
		}
	}
	return isDuplicated
}

func isValidForPart2(num int) bool {
	str := strconv.Itoa(num)
	for i := 0; i < 5; i++ {
		cur, err := strconv.Atoi(str[i : i+1])
		check(err)
		next, err := strconv.Atoi(str[i+1 : i+2])
		check(err)
		if cur > next {
			return false
		}
	}

	var m = make(map[int]int)
	for i := 0; i < 6; i++ {
		cur, err := strconv.Atoi(str[i : i+1])
		check(err)
		m[cur]++
	}
	for _, v := range m {
		if v == 2 {
			return true
		}
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	Run()
}
