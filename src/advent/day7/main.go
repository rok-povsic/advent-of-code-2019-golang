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

	highestThrust := 0
	for _, phases := range generatedAllPhases() {
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

func generatedAllPhases() [][]int {
	all := make([][]int, 0)

	values := []int{0, 1, 2, 3, 4}
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

func ComputeThrustWithRecursion(text string, phases []int) int {
	inputChan1 := make(chan int)
	inputChan2 := make(chan int)
	inputChan3 := make(chan int)
	inputChan4 := make(chan int)
	inputChan5 := make(chan int)

	finalOutputChan := make(chan int)

	// TODO
	exitChan1 := make(chan int)
	exitChan2 := make(chan int)
	exitChan3 := make(chan int)
	exitChan4 := make(chan int)
	finalExitChan := make(chan int)

	isFirstAmplifierDone := false

	go ComputeWithRecursion(text, inputChan1, inputChan2, exitChan1)
	go ComputeWithRecursion(text, inputChan2, inputChan3, exitChan2)
	go ComputeWithRecursion(text, inputChan3, inputChan4, exitChan3)
	go ComputeWithRecursion(text, inputChan4, inputChan5, exitChan4)
	go ComputeWithRecursion(text, inputChan5, finalOutputChan, finalExitChan)

	inputChan1 <- phases[0]
	inputChan2 <- phases[1]
	inputChan3 <- phases[2]
	inputChan4 <- phases[3]
	inputChan5 <- phases[4]
	inputChan1 <- 0

	output := -1
L:
	for true {
		select {
		case finalOutput := <-finalOutputChan:
			fmt.Println("Received final output: " + strconv.Itoa(finalOutput))
			if isFirstAmplifierDone {
				output = finalOutput
			} else {
				inputChan1 <- finalOutput
			}
			break
		case <-exitChan1:
			isFirstAmplifierDone = true
			break
		case <-exitChan2:
			break
		case <-exitChan3:
			break
		case <-exitChan4:
			break
		case <-finalExitChan:
			break L
		}
	}

	return output
}

func ComputeThrust(text string, phases []int) int {
	output := 0
	for amplifier := 0; amplifier < 5; amplifier++ {
		inputs := []int{phases[amplifier], output}
		response := Compute(inputs, text)
		output, _ = strconv.Atoi(strings.Trim(response, "\n"))
		fmt.Println("Output: " + strconv.Itoa(output))
	}
	return output
}

// TODO: The recursion and non-recursion variant should be combined.
func ComputeWithRecursion(programCode string, inputChan chan int, outputChan chan int, exitChan chan int) {
	program := program(programCode)

	curInput := 0

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
				program[addr] = <-inputChan
				curInput++
				pc += 2
			}
		case 4: // output
			{
				val := paramValue(modes, program, pc, 1)
				outputChan <- val
				fmt.Println("Output: " + strconv.Itoa(val))
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
				exitChan <- 0
				break L
			}
		default:
			panic("Unknown instruction " + strconv.Itoa(instruction))
		}
	}
}

func Compute(inputs []int, programCode string) string {
	fmt.Printf("Running program with inputs: %d %d\n", inputs[0], inputs[1])

	output := ""

	program := program(programCode)

	curInput := 0

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
				program[addr] = inputs[curInput]
				curInput++
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
