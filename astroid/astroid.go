package astroid

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
			viewDeg := math.Atan2(float64(yDelta), float64(xDelta))*180/math.Pi + 90
			// adjust that zero degree is at 12 o'clock [0..360]
			if viewDeg < 0 {
				viewDeg = 360 + viewDeg
			}
			// round to remove float imprecision
			curRowAngles[col] = math.Round(viewDeg*100) / 100
		}
		viewAngles[row] = curRowAngles
	}
	return viewAngles
}

func (a Astroid) GetDistances() [][]float64 {
	distances := make([][]float64, a.Field.height())
	for row := 0; row < a.Field.height(); row++ {
		distanceRow := make([]float64, a.Field.width())
		for col := 0; col < a.Field.width(); col++ {
			xDelta := col - a.Coordinate.X
			yDelta := row - a.Coordinate.Y
			distance := math.Sqrt(math.Pow(float64(yDelta), 2) + math.Pow(float64(xDelta), 2))
			// round to remove float imprecision
			distanceRow[col] = distance
		}
		distances[row] = distanceRow
	}

	return distances
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

type BlastTarget struct {
	coord    Coord
	angle    float64
	distance float64
}

func (b BlastTarget) String() string {
	return fmt.Sprintf("Angle : %v", b.angle)
}

func buildTargets(station Astroid, others []Astroid) []BlastTarget {
	viewAngles := station.GetViewingAngles()
	distances := station.GetDistances()
	var targets []BlastTarget

	// now filter astroid positions
	for _, o := range others {
		blastIt := BlastTarget{
			o.Coordinate,
			viewAngles[o.Coordinate.Y][o.Coordinate.X],
			distances[o.Coordinate.Y][o.Coordinate.X],
		}
		targets = append(targets, blastIt)
	}
	return targets
}

func sortByViewAnglesAndDistance(targets []BlastTarget) {
	sort.Slice(targets, func(i, j int) bool {
		if targets[i].angle < targets[j].angle {
			return true
		}
		if targets[i].angle > targets[j].angle {
			return false
		}
		return targets[i].distance < targets[j].distance
	})
}

func recordBlast(targets []BlastTarget, blastAngle float64, blasted []BlastTarget) ([]BlastTarget, float64, []BlastTarget) {
	// first check if there are more than one target with same angle
	var targetsSameAngleIdxs []int
	for i, t := range targets {
		if t.angle == blastAngle {
			targetsSameAngleIdxs = append(targetsSameAngleIdxs, i)
		}
	}

	blastIdx := targetsSameAngleIdxs[0]
	blasted = append(blasted, targets[blastIdx])
	targets = append(targets[:blastIdx], targets[blastIdx+1:]...)

	// now progress to next angle (jumping over the ones with same angle)
	nextIdx := blastIdx + len(targetsSameAngleIdxs) - 1

	if nextIdx > len(targets)-1 {
		nextIdx -= len(targets)
	}

	if len(targets) == 1 {
		nextIdx = 0
	}

	return targets, targets[nextIdx].angle, blasted
}

func (a AsteroidBelt) BlastAll(station Astroid, others []Astroid) {
	targets := buildTargets(station, others)
	sortByViewAnglesAndDistance(targets)

	// find min viewing angle
	blastAngle := 360.
	for _, t := range targets {
		if t.angle >= 0 && t.angle < blastAngle {
			blastAngle = t.angle
		}
	}

	var blasted []BlastTarget

	iterCnt := 1
	for len(targets) > 1 {
		targets, blastAngle, blasted = recordBlast(targets, blastAngle, blasted)
		fmt.Println(iterCnt, "blasted", blasted[len(blasted)-1].coord, "at", blasted[len(blasted)-1].angle)
		iterCnt++
	}

	answer2 := blasted[199].coord.X*100 + blasted[199].coord.Y
	fmt.Println("Answer 2", answer2)
}
