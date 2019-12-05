package main

import (
	"fmt"
	"strconv"

	"github.com/grmpflh27/aoc_2019/aoc2019_shared"
)

type Coord struct {
	X int
	Y int
}

func parseMoves(strMoves []string) []Coord {
	moves := make([]Coord, len(strMoves))

	for i, cur := range strMoves {
		dir := string(cur[0])
		num, _ := strconv.Atoi(cur[1:])
		x, y := 0, 0
		switch dir {
		case "L":
			x = -num
		case "R":
			x = num
		case "U":
			y = num
		case "D":
			y = -num
		}
		moves[i] = Coord{x, y}
	}
	return moves
}

func walk(moves []Coord) []Coord {
	visitedCoords := make([]Coord, len(moves)+1)
	start := Coord{0, 0}

	visitedCoords[0] = start

	for i, _ := range moves {
		visitedCoords[i+1] = Coord{visitedCoords[i].X + moves[i].X, visitedCoords[i].Y + moves[i].Y}
	}
	return visitedCoords
}

func main() {
	var day = 3
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")
	words := aoc2019_shared.LoadStr(day)
	//1
	coords1 := walk(parseMoves(words[0]))
	fmt.Printf("FIRST GUY WALKS %v\n", coords1)

	coords2 := walk(parseMoves(words[1]))
	fmt.Printf("SECOND GUY WALKS: %t %v\n", coords2, coords2)

	fmt.Printf("Answer 1: %v\n")

	// create rightly sized board
	stitched := append(coords1, coords2...)

	minX, maxX, minY, maxY := 0, 0, 0, 0
	for _, coord := range stitched {
		if minX > coord.X {
			minX = coord.X
		}
		if maxX < coord.X {
			maxX = coord.X
		}
		if minY > coord.Y {
			minY = coord.Y
		}
		if maxY < coord.Y {
			maxY = coord.Y
		}
	}

	sizeX := maxX - minX + 1
	sizeY := maxY - minY + 1
	fmt.Printf("Building board %v x %v", sizeX, sizeY)

	// 2

	fmt.Printf("Answer 2: %v\n")

}
