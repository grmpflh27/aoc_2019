package main

import (
	"fmt"
	"strconv"
	"strings"
)

func twoAdjacentTheSame(digits []int) bool {
	cur := digits[0]
	for _, d := range digits[1:] {
		if cur == d {
			return true
		}
		cur = d
	}
	return false
}

func isStrictlyIncreasing(digits []int) bool {
	cur := digits[0]
	for _, d := range digits[1:] {
		if cur > d {
			return false
		}
		cur = d
	}
	return true
}

func matchingDigitsNotGroupOfLarger(digits []int) bool {
	digitCnt := make(map[int]int)

	for _, d := range digits {
		if _, ok := digitCnt[d]; ok {
			digitCnt[d]++
		} else {
			digitCnt[d] = 1
		}
	}
	for _, v := range digitCnt {
		if v == 2 {
			return true
		}
	}

	return false
}

func filter(numberStr string, step2 bool) bool {
	digits := strings.Split(numberStr, "")

	intDigits := make([]int, len(digits))
	for i, d := range digits {
		intDigits[i], _ = strconv.Atoi(d)
	}

	//It is a six-digit number.
	if len(intDigits) != 6 {
		return false
	}
	//The value is within the range given in your puzzle input. --> implied

	//Two adjacent digits are the same (like 22 in 122345).
	if !twoAdjacentTheSame(intDigits) {
		return false
	}

	//Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
	if !isStrictlyIncreasing(intDigits) {
		return false
	}

	//two adjacent matching digits are not part of a larger group of matching digits
	if step2 {
		if !matchingDigitsNotGroupOfLarger(intDigits) {
			return false
		}
	}
	return true
}

func getFilteredCnt(startRange int, endRange int, step2 bool) int {
	cnt := 0
	cur := startRange
	for cur <= endRange {
		curStr := strconv.Itoa(cur)
		if filter(curStr, step2) {
			cnt++
		}
		cur++
	}
	return cnt
}

func main() {
	var day = 4
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")

	startRange, endRange := 265275, 781584

	// main loop
	cnt1 := getFilteredCnt(startRange, endRange, false)
	fmt.Printf("Answer 1: %v\n", cnt1)

	cnt2 := getFilteredCnt(startRange, endRange, true)
	fmt.Printf("Answer 2: %v\n", cnt2)

}
