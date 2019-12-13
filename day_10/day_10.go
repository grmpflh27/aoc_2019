package main

import (
	"fmt"

	"github.com/grmpflh27/aoc_2019/astroid"
)

func setViewableAstroidCnt(this *astroid.Astroid, others []astroid.Astroid, viewAngles [][]float64) int {
	for _, o := range others {
		object := this.Field[o.Coordinate.Y][o.Coordinate.X]
		// first check if it is already SHADOWED
		if object == astroid.SHADOWED {
			//fmt.Printf("%v already shadowed ... bypassing\n", o.coord)
			continue
		}

		curViewAngle := viewAngles[o.Coordinate.Y][o.Coordinate.X]
		this.CanSeeCnt++

		// now mark everything with the same viewing angle as shadowed
		this.MarkSameViewingAngleAsShadowed(viewAngles, curViewAngle)
	}

	return this.CanSeeCnt
}

func getOthers(belt astroid.AsteroidBelt, thisIdx int) []astroid.Astroid {
	others := make([]astroid.Astroid, len(belt.Astroids)-1)
	otherCnt := 0
	for j, other := range belt.Astroids {
		if thisIdx != j {
			others[otherCnt] = other
			otherCnt++
		}
	}
	return others
}

func main() {
	var day = 10
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")

	belt := astroid.LoadAsteroidBelt()

	maxCnt := 0
	maxIdx := 0
	for i, this := range belt.Astroids {
		viewAngles := this.GetViewingAngles()

		others := getOthers(belt, i)
		curCnt := setViewableAstroidCnt(&this, others, viewAngles)

		if curCnt > maxCnt {
			maxCnt = curCnt
			maxIdx = i
		}
	}
	fmt.Println("Answer 1:", maxCnt, "from", belt.Astroids[maxIdx].Coordinate)

}
