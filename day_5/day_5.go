package main

import (
	"fmt"

	"github.com/grmpflh27/aoc_2019/aoc2019_shared"
	"github.com/grmpflh27/aoc_2019/intcode"
)

func main() {
	var day = 5
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")

	input := aoc2019_shared.Load(day, ",")

	curIdx := 0
	endIdx := len(input) - 1

	for curIdx < endIdx {
		opsCode := input[curIdx]
		instr := intcode.ParseInstruction(curIdx, opsCode)
		nextIdx, _ := instr.Process(input)
		curIdx = nextIdx
	}

}
