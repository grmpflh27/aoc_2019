package main

import (
	"fmt"

	"github.com/grmpflh27/aoc_2019/aoc2019_shared"
	"github.com/grmpflh27/aoc_2019/intcode"
)

func main() {
	var day = 9
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")

	input := aoc2019_shared.Load(day, ",")

	curIdx := 0
	endIdx := len(input) - 1

	padded := make([]int, 34463340)
	for i, instr := range input {
		padded[i] = instr
	}

	for curIdx < endIdx {
		opsCode := padded[curIdx]
		instr := intcode.ParseInstruction(curIdx, opsCode)

		if curIdx == 49 {
			fmt.Println(opsCode, " =!= ", instr)
		}
		nextIdx, _ := instr.Process(padded)
		curIdx = nextIdx

		if curIdx == 49 {
			fmt.Println(padded[:len(input)])
		}
	}

}
