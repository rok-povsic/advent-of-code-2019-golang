package main

import (
	"fmt"
	"io/ioutil"

	//"io/ioutil"
	"strconv"
	"strings"
)

func Run() {
	programCode := loadProgramCode()

	// Part 1
	//values := []int{0, 1, 2, 3, 4}

	// Part 2
	values := []int{5, 6, 7, 8, 9}

	highestThrust := 0
	for _, phases := range generatedAllPhases(values) {
		fmt.Println(
			"Running phase " +
				strconv.Itoa(phases[0]) + ", " +
				strconv.Itoa(phases[1]) + ", " +
				strconv.Itoa(phases[2]) + ", " +
				strconv.Itoa(phases[3]) + ", " +
				strconv.Itoa(phases[4]))

		currentThrust := ComputeThrust(programCode, phases)
		if highestThrust < currentThrust {
			highestThrust = currentThrust
		}
	}
	fmt.Printf("Highest thrust: %d\n", highestThrust)
}

func loadProgramCode() string {
	data, err := ioutil.ReadFile("src/advent/day7/input.txt")
	check(err)
	return string(data)
}

func generatedAllPhases(values []int) [][]int {
	all := make([][]int, 0)

	for _, value := range permutations(values) {
		// Probably there's a better way of doing this..
		num1, _ := strconv.Atoi(string(value[0]))
		num2, _ := strconv.Atoi(string(value[1]))
		num3, _ := strconv.Atoi(string(value[2]))
		num4, _ := strconv.Atoi(string(value[3]))
		num5, _ := strconv.Atoi(string(value[4]))
		arr := []int{num1, num2, num3, num4, num5}
		all = append(all, arr)
	}

	return all
}

func permutations(values []int) []string {
	result := make([]string, 0)

	if len(values) == 1 {
		result = append(result, strconv.Itoa(values[0]))
		return result
	}

	for i, value := range values {
		innerValues := permutations(valuesWithoutOne(i, values))
		for _, innerValue := range innerValues {
			result = append(result, strconv.Itoa(value)+innerValue)
		}
	}

	return result
}

/*
 * Probably there's a better way to do this in Go but I didn't find it after a quick search.
 */
func valuesWithoutOne(positionToIgnore int, values []int) []int {
	result := make([]int, 0)

	for i, value := range values {
		if i != positionToIgnore {
			result = append(result, value)
		}
	}

	return result
}

func ComputeThrust(programCode string, phases []int) int {
	hasFeedbackLoop := phases[0] >= 5

	program1 := program(programCode)
	computationStopped1 := ComputeWithRecursion(program1, -1, 0)
	computationStopped1 = ComputeWithRecursion(program1, phases[0], computationStopped1.pc)
	computationStopped1 = ComputeWithRecursion(program1, 0, computationStopped1.pc)

	program2 := program(programCode)
	computationStopped2 := ComputeWithRecursion(program2, -1, 0)
	computationStopped2 = ComputeWithRecursion(program2, phases[1], computationStopped2.pc)
	computationStopped2 = ComputeWithRecursion(program2, computationStopped1.output, computationStopped2.pc)

	program3 := program(programCode)
	computationStopped3 := ComputeWithRecursion(program3, -1, 0)
	computationStopped3 = ComputeWithRecursion(program3, phases[2], computationStopped3.pc)
	computationStopped3 = ComputeWithRecursion(program3, computationStopped2.output, computationStopped3.pc)

	program4 := program(programCode)
	computationStopped4 := ComputeWithRecursion(program4, -1, 0)
	computationStopped4 = ComputeWithRecursion(program4, phases[3], computationStopped4.pc)
	computationStopped4 = ComputeWithRecursion(program4, computationStopped3.output, computationStopped4.pc)

	program5 := program(programCode)
	computationStopped5 := ComputeWithRecursion(program5, -1, 0)
	computationStopped5 = ComputeWithRecursion(program5, phases[4], computationStopped5.pc)
	computationStopped5 = ComputeWithRecursion(program5, computationStopped4.output, computationStopped5.pc)

	if !hasFeedbackLoop {
		return computationStopped5.output
	}

	for true {
		if computationStopped1.mode == "EXITED" {
			return computationStopped5.output
		}

		computationStopped1 = ComputeWithRecursion(program1, computationStopped5.output, computationStopped1.pc)
		computationStopped2 = ComputeWithRecursion(program2, computationStopped1.output, computationStopped2.pc)
		computationStopped3 = ComputeWithRecursion(program3, computationStopped2.output, computationStopped3.pc)
		computationStopped4 = ComputeWithRecursion(program4, computationStopped3.output, computationStopped4.pc)
		computationStopped5 = ComputeWithRecursion(program5, computationStopped4.output, computationStopped5.pc)
	}

	return -1
}

type ComputationStopped struct {
	mode   string // Values: "EXITED", "WAITING_FOR_INPUT"
	output int    // -1 if no output
	pc     int
}

func ComputeWithRecursion(program []int, input int, pc int) ComputationStopped {
	output := -1

	for true {
		opcode := strconv.Itoa(program[pc])
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
				if input == -1 {
					return ComputationStopped{"WAITING_FOR_INPUT", output, pc}
				}

				addr := program[pc+1]
				program[addr] = input
				pc += 2

				input = -1
			}
		case 4: // output
			{
				val := paramValue(modes, program, pc, 1)
				output = val
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
				return ComputationStopped{"EXITED", output, pc}
			}
		default:
			panic("Unknown instruction " + strconv.Itoa(instruction))
		}
	}

	return ComputationStopped{"INVALID", -1, -1}
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
