package main

import (
	"fmt"

	"github.com/grmpflh27/aoc_2019/aoc2019_shared"
)

func main() {
	var day = 2
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")

	input := aoc2019_shared.Load(day, ",")

	//1
	input[1] = 12
	input[2] = 2

	answer1 := runIntcode(input)
	fmt.Printf("Answer 1: %v\n", answer1)

	// 2
	targetOutput := 19690720
	answer2 := runIntCodeWithTargetValue(input, targetOutput)
	fmt.Printf("Answer 2: %v\n", answer2)

}

func runIntcode(input []int) int {
	var halt = 99
	var add = 1
	var mul = 2

	tmp := make([]int, len(input))
	copy(tmp, input)

	curId := 0
opCodeLoop:
	for {
		opcode := tmp[curId]
		switch opcode {
		case halt:
			break opCodeLoop
		case add, mul:
			target, op1, op2 := tmp[curId+3], tmp[curId+1], tmp[curId+2]
			switch opcode {
			case add:
				tmp[target] = tmp[op1] + tmp[op2]
			case mul:
				tmp[target] = tmp[op1] * tmp[op2]
			}
		}
		curId += 4
	}
	return tmp[0]

}

func runIntCodeWithTargetValue(input []int, targetOutput int) int {

	// brute force
	noun, verb := 99, 99

	for {
		input[1] = noun
		input[2] = verb

		curAnswer := runIntcode(input)

		if curAnswer == targetOutput {
			//fmt.Printf("%v & %v\n", noun, verb)
			break
		}

		verb -= 1
		if verb < 0 {
			verb = 99
			noun -= 1
		}

		if noun < 0 {
			panic("Did not find solution")
		}

	}

	return 100*noun + verb
}
