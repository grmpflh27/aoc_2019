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

type Move struct {
	from      Coord
	to        Coord
	direction string
}

func (m Move) distance() int {
	return Abs(m.to.Y-m.from.Y) + Abs(m.to.X-m.from.X)
}

func parseMoves(strMoves []string) []Move {
	moves := make([]Move, len(strMoves))

	curCoord := Coord{0, 0}
	for i := range strMoves {
		nextCoord, direction := parseCoord(strMoves[i], curCoord)
		moves[i] = Move{
			curCoord,
			nextCoord,
			direction,
		}
		curCoord = nextCoord

	}
	return moves
}

func parseCoord(strMove string, coord Coord) (Coord, string) {
	dir := string(strMove[0])
	num, _ := strconv.Atoi(strMove[1:])
	x, y := 0, 0
	switch dir {
	case "L":
		x = coord.X - num
		y = coord.Y
	case "R":
		x = coord.X + num
		y = coord.Y
	case "U":
		y = coord.Y + num
		x = coord.X
	case "D":
		y = coord.Y - num
		x = coord.X
	}
	return Coord{x, y}, dir
}

type Board [][]int

func initBoard(stitched []Move) (Board, Coord) {
	// create rightly sized board
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for _, move := range stitched {
		coord := move.to
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

	originOffset := Coord{-1 * minX, -1 * minY}
	fmt.Printf("Origin offset: %v\n", originOffset)

	sizeX := maxX - minX + 1
	sizeY := maxY - minY + 1
	fmt.Printf("Building board %v x %v\n", sizeX, sizeY)

	board := make(Board, sizeY)
	for i := range board {
		board[i] = make([]int, sizeX)
	}

	return board, originOffset
}

func transposeMoves(moves []Move, originOffset Coord) {
	idx, end := 0, len(moves)
	for idx < end {
		m := &moves[idx]
		m.from.X = m.from.X + originOffset.X
		m.from.Y = m.from.Y + originOffset.Y
		m.to.X = m.to.X + originOffset.X
		m.to.Y = m.to.Y + originOffset.Y
		idx++
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func placeFootsteps(moves []Move, board Board) {

	for _, move := range moves {
		switch move.direction {
		case "L", "R":
			start := Min(move.from.X, move.to.X)
			end := Max(move.from.X, move.to.X)
			for i := start; i <= end; i++ {
				if board[move.from.Y][i] != 1 {
					board[move.from.Y][i] = 1
				}
			}
		case "U", "D":
			start := Min(move.from.Y, move.to.Y)
			end := Max(move.from.Y, move.to.Y)
			for i := start; i <= end; i++ {
				if board[i][move.from.X] != 1 {
					board[i][move.from.X] = 1
				}
			}
		}
	}
}

func findIntersections(board1 Board, board2 Board, originOffset Coord) []Coord {
	intersections := []Coord{}
	for rowIdx := range board1 {
		for colIdx := range board1[rowIdx] {
			if board1[rowIdx][colIdx] == 1 && board2[rowIdx][colIdx] == 1 {
				cur := Coord{colIdx, rowIdx}
				if cur != originOffset {
					intersections = append(intersections, cur)
				}
			}
		}
	}
	return intersections
}

func manhattanDistance(from Coord, to Coord) int {
	return Abs(to.X-from.X) + Abs(to.Y-from.Y)
}

func minManhattanDistance(intersections []Coord, origin Coord) (int, Coord) {
	manh := 100000
	minCoord := origin

	for _, inter := range intersections {
		if inter != origin {
			curDistance := manhattanDistance(origin, inter)
			if manh > curDistance {
				manh = curDistance
				minCoord = inter
			}
		}
	}
	return manh, minCoord
}

func countStepsToMinIntersection(moves []Move, intersections []Coord) []int {
	stepCounts := make([]int, len(intersections))

	for i, inter := range intersections {
		stepCounts[i] = countStepsTo(moves, inter)
	}
	return stepCounts
}

func countStepsTo(moves []Move, intersection Coord) int {
	var steps int = 0
	for _, move := range moves {
		switch move.direction {
		case "R":
			start := Min(move.from.X, move.to.X)
			end := Max(move.from.X, move.to.X)
			for i := start + 1; i <= end; i++ {
				if move.to.Y == intersection.Y && i == intersection.X {
					steps++
					return steps
				}
				steps++
			}
		case "L":
			start := Min(move.from.X, move.to.X) - 1
			end := Max(move.from.X, move.to.X) - 1
			for i := end; i > start; i-- {
				if move.to.Y == intersection.Y && i == intersection.X {
					steps++
					return steps
				}
				steps++
			}
		case "U":
			start := Min(move.from.Y, move.to.Y)
			end := Max(move.from.Y, move.to.Y)
			for i := start + 1; i <= end; i++ {
				if i == intersection.Y && move.to.X == intersection.X {
					steps++
					return steps
				}
				steps++
			}
		case "D":
			start := Min(move.from.Y, move.to.Y) - 1
			end := Max(move.from.Y, move.to.Y) - 1
			for i := end; i > start; i-- {
				if i == intersection.Y && move.to.X == intersection.X {
					steps++
					return steps
				}
				steps++
			}
		}
	}
	return steps
}

func main() {
	var day = 3
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")
	words := aoc2019_shared.LoadStr(day, ",")
	//1
	moves1 := parseMoves(words[0])
	moves2 := parseMoves(words[1])

	stitched := append(moves1, moves2...)
	board1, _ := initBoard(stitched)
	board2, originOffset := initBoard(stitched)

	transposeMoves(moves1, originOffset)
	transposeMoves(moves2, originOffset)

	placeFootsteps(moves1, board1)
	placeFootsteps(moves2, board2)

	intersections := findIntersections(board1, board2, originOffset)
	minDistance, _ := minManhattanDistance(intersections, originOffset)
	fmt.Printf("Answer 1: %v\n", minDistance)

	// 2
	cnt1 := countStepsToMinIntersection(moves1, intersections)
	cnt2 := countStepsToMinIntersection(moves2, intersections)

	sums := make([]int, len(cnt1))
	for i := range cnt1 {
		sums[i] = cnt1[i] + cnt2[i]
	}

	minSum := sums[0]
	for _, curSum := range sums[1:] {
		if curSum < minSum {
			minSum = curSum
		}
	}
	fmt.Printf("Answer 2: %v\n", minSum)

}
