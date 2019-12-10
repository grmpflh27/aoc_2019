package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/grmpflh27/aoc_2019/aoc2019_shared"
	"github.com/grmpflh27/aoc_2019/intcode"
)

func setupinputBuffer(phaseSequence []int) {
	// clear input buffer
	intcode.InputBuffer = intcode.InputBuffer[:0]
	intcode.InputBuffer = append(intcode.InputBuffer, phaseSequence[0])
	intcode.InputBuffer = append(intcode.InputBuffer, 0)
	for _, p := range phaseSequence[1:] {
		intcode.InputBuffer = append(intcode.InputBuffer, p)
	}
}

func runProgram(input []int, phaseSequence []int) int {

	// copy input to workingCopy

	setupinputBuffer(phaseSequence)
	curIdx := 0
	endIdx := len(input) - 1

	for curIdx < endIdx {
		opsCode := input[curIdx]
		instr := intcode.ParseInstruction(curIdx, opsCode)
		nextIdx, err := instr.Process(input)
		if err != nil {
			break
		}
		curIdx = nextIdx
	}
	return intcode.OutputBuffer
}

func setupinputBufferForFeedback(phaseSequence []int) {
	// clear input buffer
	intcode.InputBuffer = intcode.InputBuffer[:0]
	intcode.InputBuffer = append(intcode.InputBuffer, 0)
	intcode.InputBuffer = append(intcode.InputBuffer, phaseSequence[0])
	for _, p := range phaseSequence[1:] {
		intcode.InputBuffer = append(intcode.InputBuffer, p)
	}
}

func runProgramWithFeedback(input []int, phaseSequence []int) int {

	setupinputBufferForFeedback(phaseSequence)
	curIdx := 0
	endIdx := len(input) - 1
	fmt.Println(intcode.InputBuffer)
	for curIdx < endIdx {
		opsCode := input[curIdx]
		instr := intcode.ParseInstruction(curIdx, opsCode)
		nextIdx, _ := instr.Process(input)
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

func parsePhaseDigits(phaseSequence string) []int {
	digitStr := strings.Split(phaseSequence, "")
	intPhaseSequence := make([]int, len(digitStr))

	for i, d := range digitStr {
		intPhaseSequence[i], _ = strconv.Atoi(d)
	}
	return intPhaseSequence
}

func main() {
	var day = 7
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")

	permStr := flag.String("perm", "01234", "5 digit permutation string, e.g. 01234")
	feedback := flag.Bool("feedback", false, "set to start program with feedback loop")
	flag.Parse()

	input := aoc2019_shared.Load(day, ",")
	// create permutations
	var permutations []string
	Perm([]rune(*permStr), func(a []rune) {
		permutations = append(permutations, string(a))
	})

	maxPhaseOutput := 0
	bestPerm := ""
	curPhaseOutput := 0
	for _, perm := range permutations {
		phaseSequence := parsePhaseDigits(perm)
		workingCopy := make([]int, len(input))
		copy(workingCopy, input)
		if *feedback {
			curPhaseOutput = runProgramWithFeedback(workingCopy, phaseSequence)
		} else {
			curPhaseOutput = runProgram(workingCopy, phaseSequence)
		}

		if maxPhaseOutput < curPhaseOutput {
			bestPerm = perm
			maxPhaseOutput = curPhaseOutput
		}
	}

	withOrWithout := "without"
	if *feedback {
		withOrWithout = "with"
	}
	fmt.Println("Best output to thrusters", maxPhaseOutput, " by ", bestPerm, withOrWithout, "feedback")

}
