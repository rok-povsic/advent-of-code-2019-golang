package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Compute(programCode string, inputChan chan int, outputChan chan int, finishedChan chan bool) {
	fmt.Println("Running program")

	program := program(programCode)

	pc := 0
	relativeBase := 0
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
				a := paramValue(modes, program, pc, 1, relativeBase)
				b := paramValue(modes, program, pc, 2, relativeBase)

				cAddr := paramAddr(modes, program, pc, 3, relativeBase)
				c := a + b
				program = writeDataToProgram(program, cAddr, c)
				pc += 4
			}
		case 2: // multiply
			{
				a := paramValue(modes, program, pc, 1, relativeBase)
				b := paramValue(modes, program, pc, 2, relativeBase)

				cAddr := paramAddr(modes, program, pc, 3, relativeBase)
				c := a * b
				program = writeDataToProgram(program, cAddr, c)
				pc += 4
			}
		case 3: // input
			{
				fmt.Println("Waiting for input")
				input := <-inputChan
				fmt.Printf("Got input: %d\n", input)

				addr := paramAddr(modes, program, pc, 1, relativeBase)
				program = writeDataToProgram(program, addr, input)
				pc += 2
			}
		case 4: // output
			{
				val := paramValue(modes, program, pc, 1, relativeBase)
				fmt.Printf("Output: %d\n", val)
				outputChan <- val
				pc += 2
			}
		case 5: // jump-if-true
			{
				condition := paramValue(modes, program, pc, 1, relativeBase)
				jumpAddr := paramValue(modes, program, pc, 2, relativeBase)
				if condition != 0 {
					pc = jumpAddr
				} else {
					pc += 3
				}
			}
		case 6: // jump-if-false
			{
				condition := paramValue(modes, program, pc, 1, relativeBase)
				jumpAddr := paramValue(modes, program, pc, 2, relativeBase)
				if condition == 0 {
					pc = jumpAddr
				} else {
					pc += 3
				}
			}
		case 7: // less than
			{
				a := paramValue(modes, program, pc, 1, relativeBase)
				b := paramValue(modes, program, pc, 2, relativeBase)
				var result int
				if a < b {
					result = 1
				} else {
					result = 0
				}
				resultAddrAddr := paramAddr(modes, program, pc, 3, relativeBase)
				program = writeDataToProgram(program, resultAddrAddr, result)
				pc += 4
			}
		case 8: // equals
			{
				a := paramValue(modes, program, pc, 1, relativeBase)
				b := paramValue(modes, program, pc, 2, relativeBase)
				var result int
				if a == b {
					result = 1
				} else {
					result = 0
				}
				resultAddrAddr := paramAddr(modes, program, pc, 3, relativeBase)
				program = writeDataToProgram(program, resultAddrAddr, result)
				pc += 4
			}
		case 9: // adjust relative base
			{
				val := paramValue(modes, program, pc, 1, relativeBase)
				relativeBase += val
				pc += 2
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

	finishedChan <- true
}

func writeDataToProgram(program []int, addr int, value int) []int {
	if addr >= len(program) {
		newProgram := make([]int, addr+1)
		copy(newProgram, program)
		program = newProgram
	}
	program[addr] = value
	return program
}

func paramValue(modes string, program []int, pc int, pos int, relativeBase int) int {
	addr := paramAddr(modes, program, pc, pos, relativeBase)
	if addr < len(program) {
		return program[addr]
	} else {
		return 0
	}
}

func paramAddr(modes string, program []int, pc int, pos int, relativeBase int) int {
	var addr int
	if isImmediate(modes, pos) {
		addr = pc + pos
	} else if isRelative(modes, pos) {
		addr = relativeBase + program[pc+pos]
	} else { // position mode
		addr = program[pc+pos]
	}
	return addr
}

func isImmediate(modes string, pos int) bool {
	idx := len(modes) - pos
	if idx < 0 {
		return false
	}
	return modes[idx:idx+1] == "1"
}

func isRelative(modes string, pos int) bool {
	idx := len(modes) - pos
	if idx < 0 {
		return false
	}
	return modes[idx:idx+1] == "2"
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
