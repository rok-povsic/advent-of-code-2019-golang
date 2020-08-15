package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	program := "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,19,1,19,9,23,1,23,6,27,1,9,27,31,1,31,10,35,2,13,35,39,1,39,10,43,1,43,9,47,1,47,13,51,1,51,13,55,2,55,6,59,1,59,5,63,2,10,63,67,1,67,9,71,1,71,13,75,1,6,75,79,1,10,79,83,2,9,83,87,1,87,5,91,2,91,9,95,1,6,95,99,1,99,5,103,2,103,10,107,1,107,6,111,2,9,111,115,2,9,115,119,2,13,119,123,1,123,9,127,1,5,127,131,1,131,2,135,1,135,6,0,99,2,0,14,0"
	L: for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			result := Compute(noun, verb, program)
			output := strings.Split(result, ",")[0]
			if output == "19690720" {
				fmt.Printf("noun = %d, verb = %d\n", noun, verb)
				fmt.Printf("100 * noun + verb = %d\n", 100 * noun + verb)
				break L
			}
		}
	}
}

func Compute(noun int, verb int, text string) string {
	instrs := instrs(text)

	instrs[1] = noun
	instrs[2] = verb

	pc := 0

	L: for true {
		instr := instrs[pc]

		switch instr {
		case 1: {
			aAddr := instrs[pc+1]
			bAddr := instrs[pc+2]
			cAddr := instrs[pc+3]
			a := instrs[aAddr]
			b := instrs[bAddr]
			c := a + b
			instrs[cAddr] = c
			pc += 4
		}
		case 2: {
			aAddr := instrs[pc+1]
			bAddr := instrs[pc+2]
			cAddr := instrs[pc+3]
			a := instrs[aAddr]
			b := instrs[bAddr]
			c := a * b
			instrs[cAddr] = c
			pc += 4
		}
		case 99: {
			break L
		}
		default:
		}
	}

	var instrsStr []string
	for _, instr := range instrs {
		instrsStr = append(instrsStr, strconv.Itoa(instr))
	}

	return strings.Join(instrsStr, ",")
}

func instrs(text string) []int {
	parts := strings.Split(text, ",")

	var instrs []int
	for _, part := range parts {
		instr, err := strconv.Atoi(part)
		check(err)
		instrs = append(instrs, instr)
	}
	return instrs
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}