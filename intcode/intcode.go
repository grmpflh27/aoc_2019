package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

type Mode int

type SupportedOpCodes int

const (
	HALT          SupportedOpCodes = 99
	ADD           SupportedOpCodes = 1
	MUL           SupportedOpCodes = 2
	INPUT         SupportedOpCodes = 3
	OUTPUT        SupportedOpCodes = 4
	JUMP_IF_TRUE  SupportedOpCodes = 5
	JUMP_IF_FALSE SupportedOpCodes = 6
	LESS_THAN     SupportedOpCodes = 7
	EQUALS        SupportedOpCodes = 8
)

const (
	Position  Mode = 0
	Immediate Mode = 1
)

type Instruction struct {
	idx        int
	opscode    SupportedOpCodes
	firstMode  Mode
	secondMode Mode
	thirdMode  Mode
}

func ParseInstruction(idx int, compountOpscode int) Instruction {
	digits := splitIntoDigitsAndPad(compountOpscode)

	opsCode := digits[3]*10 + digits[4]
	instr := Instruction{
		idx,
		SupportedOpCodes(opsCode),
		Mode(digits[2]),
		Mode(digits[1]),
		Mode(digits[0]),
	}
	return instr
}

func splitIntoDigitsAndPad(opscode int) []int {
	digitsStr := strings.Split(strconv.Itoa(opscode), "")
	digits := make([]int, 5)

	startAt := 5 - len(digitsStr)

	for cnt, d := range digitsStr {
		digits[startAt+cnt], _ = strconv.Atoi(d)
	}

	return digits
}

func getParam(mode Mode, idx int, input []int) int {
	if mode == Immediate {
		ret := input[idx]
		return ret
	} else {
		return input[input[idx]]
	}
}

func setTarget(value int, mode Mode, idx int, input []int) {
	if mode == Immediate {
		input[idx] = value
	} else {
		input[input[idx]] = value
	}
}

func (instr *Instruction) Process(input []int) int {
	nextIdx := 0
	switch instr.opscode {
	case ADD, MUL:
		firstParam := getParam(instr.firstMode, instr.idx+1, input)
		secondParam := getParam(instr.secondMode, instr.idx+2, input)
		nextIdx = instr.idx + 4
		switch instr.opscode {
		case ADD:
			value := firstParam + secondParam
			setTarget(value, instr.thirdMode, instr.idx+3, input)
		case MUL:
			value := firstParam * secondParam
			setTarget(value, instr.thirdMode, instr.idx+3, input)
		}
	case INPUT:
		fmt.Println("Please provide number")
		var value int
		fmt.Scanf("%d", &value)
		setTarget(value, instr.firstMode, instr.idx+1, input)
		nextIdx = instr.idx + 2
	case OUTPUT:
		outputValue := getParam(instr.firstMode, instr.idx+1, input)
		fmt.Printf("Output at %v = %v\n", instr.idx+1, outputValue)
		nextIdx = instr.idx + 2
	// comment those 3 cases out for solution 1
	case JUMP_IF_TRUE, JUMP_IF_FALSE:
		firstParam := getParam(instr.firstMode, instr.idx+1, input)
		jumpValue := getParam(instr.secondMode, instr.idx+2, input)
		nextIdx = instr.idx + 3
		switch instr.opscode {
		case JUMP_IF_TRUE:
			if firstParam != 0 {
				nextIdx = jumpValue
			}
		case JUMP_IF_FALSE:
			if firstParam == 0 {
				nextIdx = jumpValue
			}
		}
	case LESS_THAN, EQUALS:
		firstParam := getParam(instr.firstMode, instr.idx+1, input)
		secondParam := getParam(instr.secondMode, instr.idx+2, input)
		nextIdx = instr.idx + 4

		switch instr.opscode {
		case LESS_THAN:
			if firstParam < secondParam {
				setTarget(1, instr.thirdMode, instr.idx+3, input)
			} else {
				setTarget(0, instr.thirdMode, instr.idx+3, input)
			}
		case EQUALS:
			if firstParam == secondParam {
				setTarget(1, instr.thirdMode, instr.idx+3, input)
			} else {
				setTarget(0, instr.thirdMode, instr.idx+3, input)
			}
		}
	}
	return nextIdx
}
