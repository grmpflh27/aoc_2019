package main

import (
	"bufio"
	"fmt"
	"os"
)

const SIZE = 5

const (
	BUG   = '#'
	EMPTY = '.'
)

type Field [5][5]rune

func (f Field) isBug(x, y int) bool {
	return f[y][x] == BUG
}

func (f Field) neighbouringBugs(x, y int) int {
	bugCnt := 0
	if x != 0 {
		if f.isBug(x-1, y) {
			bugCnt++
		}
	}
	if x != SIZE-1 {
		if f.isBug(x+1, y) {
			bugCnt++
		}
	}
	if y != 0 {
		if f.isBug(x, y-1) {
			bugCnt++
		}
	}
	if y != SIZE-1 {
		if f.isBug(x, y+1) {
			bugCnt++
		}
	}

	return bugCnt
}

func parseMap() Field {
	fp, _ := os.Open("./input_24.txt")
	scanner := bufio.NewScanner(fp)

	var field Field
	lineCnt := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, char := range line {
			field[lineCnt][i] = char
		}
		lineCnt++
	}
	return field
}

func gameOfLife(f Field) Field {
	var newField Field
	row := 0
	for row < SIZE {
		col := 0
		for col < SIZE {
			cur := f[row][col]
			adjBugs := f.neighbouringBugs(col, row)
			if cur == BUG {
				// dies
				if adjBugs != 1 {
					newField[row][col] = EMPTY
					// lives
				} else {
					newField[row][col] = BUG
				}
			} else {
				// infested
				if adjBugs == 1 || adjBugs == 2 {
					newField[row][col] = BUG
				} else {
					newField[row][col] = EMPTY
				}
			}
			col++
		}
		row++
	}
	return newField
}

func (f Field) biodiversityRating() int {
	rating := 0
	powerOfTwo := 0
	row := 0
	for row < SIZE {
		col := 0
		for col < SIZE {
			if f[row][col] == BUG {
				rating += 1 << powerOfTwo
			}
			powerOfTwo++
			col++
		}
		row++
	}
	return rating
}

func main() {

	var day = 24
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")

	recorder := make(map[Field]bool)

	field := parseMap()
	recorder[field] = true

	cnt := 0
	for {
		field = gameOfLife(field)
		if _, ok := recorder[field]; ok {
			fmt.Println("exists", field)
			break
		}
		recorder[field] = true
		cnt++
	}

	rating := field.biodiversityRating()
	fmt.Println(rating)
}
