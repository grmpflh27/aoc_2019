package astroid

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	EMPTY    = '.'
	ASTEROID = '#'
	SHADOWED = 'X'
)

func newAstroid(x int, y int, field AsteriodField) Astroid {
	aster := Astroid{}
	aster.Coordinate = Coord{x, y}
	aster.CanSeeCnt = 0

	tmp := make(AsteriodField, len(field))
	copy(tmp, field)
	aster.Field = tmp
	return aster
}

type Coord struct {
	X int
	Y int
}

type Astroid struct {
	Field      AsteriodField
	Coordinate Coord
	CanSeeCnt  int
}

func (a Astroid) String() string {
	return fmt.Sprintf("Asteroid at (%v,%v)", a.Coordinate.X, a.Coordinate.Y)
}

type AsteriodField []string

func (a AsteriodField) String() string {
	return fmt.Sprintf(strings.Join(a, "\n"))
}

func (a AsteriodField) width() int {
	return len(a[0])
}

func (a AsteriodField) height() int {
	return len(a)
}

type AsteroidBelt struct {
	Field    AsteriodField
	Astroids []Astroid
	width    int
	height   int
}

func (a AsteroidBelt) String() string {
	return a.Field.String()
}

func LoadAsteroidBelt() AsteroidBelt {
	var asteroids []Astroid
	var field AsteriodField
	fp, err := os.Open("./input_10.txt")
	if err != nil {
		panic("couldnt load")
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, line)
	}

	lineCnt := 0
	for _, line := range field {
		for colCnt, char := range line {
			if char == '#' {
				asteroids = append(asteroids, newAstroid(colCnt, lineCnt, field))
			}
		}
		lineCnt++
	}

	return AsteroidBelt{
		field,
		asteroids,
		len(field[0]),
		len(field),
	}
}

func (a AsteriodField) mark(shadowed Coord, symbol rune) {
	row := a[shadowed.Y]
	// strings are immutable --> convert to rune
	runeRow := []rune(row)
	runeRow[shadowed.X] = symbol
	a[shadowed.Y] = string(runeRow)
}

func (a Astroid) GetViewingAngles() [][]float64 {
	viewAngles := make([][]float64, a.Field.height())
	for row := 0; row < a.Field.height(); row++ {
		curRowAngles := make([]float64, a.Field.width())
		for col := 0; col < a.Field.width(); col++ {
			xDelta := col - a.Coordinate.X
			yDelta := row - a.Coordinate.Y
			viewDeg := math.Atan2(float64(yDelta), float64(xDelta)) * 180 / math.Pi
			// round to remove float imprecision
			curRowAngles[col] = math.Round(viewDeg*100) / 100
		}
		viewAngles[row] = curRowAngles
	}

	return viewAngles
}

func (a Astroid) MarkSameViewingAngleAsShadowed(viewAngles [][]float64, curViewAngle float64) {
	for row := 0; row < len(viewAngles); row++ {
		for col := 0; col < len(viewAngles[0]); col++ {
			if viewAngles[row][col] == curViewAngle {
				a.Field.mark(Coord{col, row}, SHADOWED)
			}
		}
	}
}
