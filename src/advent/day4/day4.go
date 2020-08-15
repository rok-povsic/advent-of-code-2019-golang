package day4

import (
	"fmt"
	"strconv"
)

func Run() {
	from := 240920
	to := 789857

	count := 0
	for num := from; num <= to; num++ {
		if isValid(num) {
			count++
		}
	}
	fmt.Println(count)
}

func isValid(num int) bool {
	str := strconv.Itoa(num)
	isDuplicated := false
	for i := 0; i < 5; i++ {
		cur, err := strconv.Atoi(str[i:i+1])
		check(err)
		next, err := strconv.Atoi(str[i+1:i+2])
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
