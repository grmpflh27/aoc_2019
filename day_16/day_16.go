package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func parse() []int {
	var inputSignal []int
	fp, _ := os.Open("./input_16.txt")
	bytes, _ := ioutil.ReadAll(fp)
	strInput := string(bytes)

	for _, b := range strInput {
		conv, _ := strconv.Atoi(string(b))
		inputSignal = append(inputSignal, conv)
	}
	return inputSignal
}

func getPhaseCoeffs(basePattern []int, length int) [][]int {
	var phaseCoeffs [][]int
	for offset := 0; offset < length; offset++ {
		coeffs := padOffsetted(basePattern, offset+1, length)
		phaseCoeffs = append(phaseCoeffs, coeffs)
	}

	return phaseCoeffs
}

func padOffsetted(basePattern []int, repeat int, size int) []int {
	padded := make([]int, size)

	inPtr := 0
	// start at offset (== repeat-1)
	outPtr := repeat - 1
	for outPtr < size {
		for k := 0; k < repeat; k++ {
			if outPtr >= size {
				break
			}
			padded[outPtr] = basePattern[inPtr%len(basePattern)]
			outPtr++
		}
		inPtr++

	}
	return padded
}

func applyPhase(inputSignal *[]int, phaseCoeffs [][]int) {
	for i, phase := range phaseCoeffs {
		scal := scalarProduct(*inputSignal, phase)
		(*inputSignal)[i] = abs(scal % 10)
	}
}

func scalarProduct(vec1 []int, vec2 []int) int {
	cum := 0
	for i := 0; i < len(vec1); i++ {
		cum += vec1[i] * vec2[i]
	}
	return cum
}

func abs(in int) int {
	if in < 0 {
		return -1 * in
	}
	return in
}

func main() {

	var day = 16
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")

	inputSignal := parse()
	fmt.Println(inputSignal)

	basePattern := []int{1, 0, -1, 0}
	// --- PART 1 ----
	phaseCoeffs := getPhaseCoeffs(basePattern, len(inputSignal))

	outputSignal := make([]int, len(inputSignal))
	copy(outputSignal, inputSignal)

	iterations := 100
	cnt := 0

	for cnt < iterations {
		applyPhase(&outputSignal, phaseCoeffs)
		cnt++
	}
	fmt.Println("Answer 1", arrayToString(outputSignal, "")[:8])

	// --- PART 2 ----
	// start anew
	messagesOffset := inputSignal[:7]
	totalOffset, _ := strconv.Atoi(arrayToString(messagesOffset, ""))
	fmt.Println(totalOffset)

	var realSignal []int
	for i := 0; i < 10000; i++ {
		for _, value := range inputSignal {
			realSignal = append(realSignal, value)
		}
	}

	realSignal = realSignal[totalOffset:]
	phaseCoeffs = getPhaseCoeffs(basePattern, len(realSignal))

	cnt = 0
	iterations = 100

	for phase := 0; phase < 100; phase++ {
		applyPhase(&realSignal, phaseCoeffs)
		cnt++
		fmt.Println(cnt)
	}

	fmt.Println("Answer 2:", realSignal[:8])
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
