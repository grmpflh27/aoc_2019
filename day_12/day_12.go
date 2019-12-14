package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Position struct {
	X int
	Y int
	Z int
}

type Velocity struct {
	X int
	Y int
	Z int
}

type Moon struct {
	pos Position
	v   Velocity
}

func (m Moon) String() string {
	return fmt.Sprintf("pos=<x=%v, y=%v, z=%v>, vel=<x=%v, y=%v, z=%v>", m.pos.X, m.pos.Y, m.pos.Z, m.v.X, m.v.Y, m.v.Z)
}

func parse() [4]Moon {
	re := regexp.MustCompile(`<x=(-?\d+),\sy=(-?\d+),\sz=(-?\d+)>`)
	fp, _ := os.Open("./input_12.txt")

	scanner := bufio.NewScanner(fp)

	moons := [4]Moon{}
	cnt := 0
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		z, _ := strconv.Atoi(match[3])
		moons[cnt] = Moon{
			Position{x, y, z},
			Velocity{},
		}
		cnt++
	}
	return moons
}

func _get(point1, point2 int) int {
	if point1 == point2 {
		return 0
	}
	if point1 < point2 {
		return 1
	}
	return -1
}

func applyGravity(moons *[4]Moon) {
	for i := range moons {
		for j := range moons {
			if i != j {
				moons[i].v.X += _get(moons[i].pos.X, moons[j].pos.X)
				moons[i].v.Y += _get(moons[i].pos.Y, moons[j].pos.Y)
				moons[i].v.Z += _get(moons[i].pos.Z, moons[j].pos.Z)
			}
		}
	}
}

func applyVelocity(moons *[4]Moon) {
	for i := range moons {
		moons[i].pos.X += moons[i].v.X
		moons[i].pos.Y += moons[i].v.Y
		moons[i].pos.Z += moons[i].v.Z
	}
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (m Moon) calcPotentialEnergy() int {
	pot := Abs(m.pos.X) + Abs(m.pos.Y) + Abs(m.pos.Z)
	kin := Abs(m.v.X) + Abs(m.v.Y) + Abs(m.v.Z)
	return pot * kin
}

func main() {

	var day = 12
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")

	totalTime := flag.Int("time", 10, "simulation time")
	flag.Parse()
	moons := parse()
	time := 0

	for time < *totalTime {
		applyGravity(&moons)
		applyVelocity(&moons)
		time++
	}

	totalEnergy := 0

	for _, m := range moons {
		totalEnergy += m.calcPotentialEnergy()
	}
	fmt.Println("TOTAL ENERGY", totalEnergy)
}
