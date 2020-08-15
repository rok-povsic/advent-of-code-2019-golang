package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Run() {
	data, err := ioutil.ReadFile("src/advent/day5/input.txt")
	check(err)
	text := string(data)
	Compute(text)
}

func Compute(text string) string {
	program := program(text)

	pc := 0

	//fmt.Println(program)
L:
	for true {
		opcode := strconv.Itoa(program[pc])
		//fmt.Printf("pc: %d, opcode: %s\n", pc, opcode)
		instruction, err := strconv.Atoi(opcode[max(0, len(opcode)-2):])
		check(err)

		modes := opcode[:max(0, len(opcode)-2)]

		switch instruction {
		case 1:
			{
				a := paramValue(modes, program, pc, 1)
				b := paramValue(modes, program, pc, 2)

				cAddr := program[pc+3]
				c := a + b
				program[cAddr] = c
				pc += 4
			}
		case 2:
			{
				a := paramValue(modes, program, pc, 1)
				b := paramValue(modes, program, pc, 2)

				cAddr := program[pc+3]
				c := a * b
				program[cAddr] = c
				pc += 4
			}
		case 3:
			{
				val := input()
				addr := program[pc+1]
				program[addr] = val
				pc += 2
			}
		case 4:
			{
				addr := program[pc+1]
				output(program[addr])
				pc += 2
			}
		case 99:
			{
				break L
			}
		default:
			panic("Unknown instruction " + strconv.Itoa(instruction))
		}
	}

	var instrsStr []string
	for _, instr := range program {
		instrsStr = append(instrsStr, strconv.Itoa(instr))
	}

	return strings.Join(instrsStr, ",")
}

func paramValue(modes string, program []int, pc int, pos int) int {
	if isImmediate(modes, pos) {
		return program[pc+pos]
	} else {
		addr := program[pc+pos]
		return program[addr]
	}
}

func isImmediate(modes string, pos int) bool {
	idx := len(modes) - pos
	if idx < 0 {
		return false
	}
	return modes[idx:idx+1] == "1"
}

func program(text string) []int {
	parts := strings.Split(text, ",")

	var program []int
	for _, part := range parts {
		instr, err := strconv.Atoi(part)
		check(err)
		program = append(program, instr)
	}
	return program
}

func input() int {
	num := 1
	fmt.Printf("> %d\n", num)
	return num
}

func output(num int) {
	fmt.Println(num)
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	Run()
}
