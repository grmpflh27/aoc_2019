package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/grmpflh27/aoc_2019/aoc2019_shared"
	"github.com/grmpflh27/aoc_2019/intcode"
)

func runProgram(input []int, phaseSequence string) int {
	workingCopy := make([]int, len(input))
	copy(workingCopy, input)
	digitStr := strings.Split(phaseSequence, "")
	intPhaseSequence := make([]int, len(digitStr))

	for i, d := range digitStr {
		intPhaseSequence[i], _ = strconv.Atoi(d)
	}
	fmt.Printf("INIT input: %v\n", input)
	curIdx := 0
	endIdx := len(input) - 1

	// init code
	intcode.InputBuffer = append(intcode.InputBuffer, intPhaseSequence[0])
	intcode.InputBuffer = append(intcode.InputBuffer, 0)

	for _, p := range intPhaseSequence[1:] {
		intcode.InputBuffer = append(intcode.InputBuffer, p)
	}

	for curIdx < endIdx {
		opsCode := workingCopy[curIdx]
		instr := intcode.ParseInstruction(curIdx, opsCode)
		nextIdx, err := instr.Process(workingCopy)
		if err != nil {
			break
		}
		curIdx = nextIdx
	}
	return intcode.OutputBuffer
}

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	var day = 7
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")

	input := aoc2019_shared.Load(day, ",")
	// create permutations
	permutations := make([]string, 100)
	Perm([]rune("01234"), func(a []rune) {
		permutations = append(permutations, string(a))
	})

	maxPhaseOutput := 0
	bestPerm := ""

	for _, perm := range permutations {
		curPhaseOutput := runProgram(input, perm)

		if maxPhaseOutput < curPhaseOutput {
			bestPerm = perm
			maxPhaseOutput = curPhaseOutput
		}
	}

	fmt.Println("Best output to thrusters", maxPhaseOutput, " by ", bestPerm)
}
