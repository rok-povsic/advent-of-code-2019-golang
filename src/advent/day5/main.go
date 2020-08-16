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
	Compute(5, text)
}

func Compute(input int, text string) string {
	fmt.Printf("Running program with input: %d\n", input)

	output := ""

	program := program(text)

	pc := 0
L:
	for true {
		opcode := strconv.Itoa(program[pc])
		//fmt.Printf("pc: %d, opcode: %s\n", pc, opcode)
		instruction, err := strconv.Atoi(opcode[max(0, len(opcode)-2):])
		check(err)

		modes := opcode[:max(0, len(opcode)-2)]

		switch instruction {
		case 1: // sum
			{
				a := paramValue(modes, program, pc, 1)
				b := paramValue(modes, program, pc, 2)

				cAddr := program[pc+3]
				c := a + b
				program[cAddr] = c
				pc += 4
			}
		case 2: // multiply
			{
				a := paramValue(modes, program, pc, 1)
				b := paramValue(modes, program, pc, 2)

				cAddr := program[pc+3]
				c := a * b
				program[cAddr] = c
				pc += 4
			}
		case 3: // input
			{
				addr := program[pc+1]
				program[addr] = input
				pc += 2
			}
		case 4: // output
			{
				val := paramValue(modes, program, pc, 1)
				text := strconv.Itoa(val)
				fmt.Println("Output: " + text)
				output += text + "\n"
				pc += 2
			}
		case 5: // jump-if-true
			{
				condition := paramValue(modes, program, pc, 1)
				jumpAddr := paramValue(modes, program, pc, 2)
				if condition != 0 {
					pc = jumpAddr
				} else {
					pc += 3
				}
			}
		case 6: // jump-if-false
			{
				condition := paramValue(modes, program, pc, 1)
				jumpAddr := paramValue(modes, program, pc, 2)
				if condition == 0 {
					pc = jumpAddr
				} else {
					pc += 3
				}
			}
		case 7: // less than
			{
				a := paramValue(modes, program, pc, 1)
				b := paramValue(modes, program, pc, 2)
				var result int
				if a < b {
					result = 1
				} else {
					result = 0
				}
				resultAddrAddr := program[pc+3]
				program[resultAddrAddr] = result
				pc += 4
			}
		case 8: // equals
			{
				a := paramValue(modes, program, pc, 1)
				b := paramValue(modes, program, pc, 2)
				var result int
				if a == b {
					result = 1
				} else {
					result = 0
				}
				resultAddrAddr := program[pc+3]
				program[resultAddrAddr] = result
				pc += 4
			}
		case 99: // exit
			{
				fmt.Print("EXIT\n\n")
				break L
			}
		default:
			panic("Unknown instruction " + strconv.Itoa(instruction))
		}
	}

	return output
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
