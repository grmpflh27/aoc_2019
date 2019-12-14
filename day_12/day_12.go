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
	defer fp.Close()

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

func getStepsOfFullOrbit(moons [4]Moon) {

	xs := make(map[string]bool)
	ys := make(map[string]bool)
	zs := make(map[string]bool)

	returnX, returnY, returnZ := 0, 0, 0
	t := 0

	for {
		applyGravity(&moons)
		applyVelocity(&moons)

		if returnX == 0 {
			x := fmt.Sprintf("%v %v %v %v %v %v %v %v", moons[0].pos.X, moons[1].pos.X, moons[2].pos.X, moons[3].pos.X,
				moons[0].v.X, moons[1].v.X, moons[2].v.X, moons[3].v.X)
			if xs[x] {
				returnX = t
			}
			xs[x] = true
		}

		if returnY == 0 {
			y := fmt.Sprintf("%v %v %v %v %v %v %v %v", moons[0].pos.Y, moons[1].pos.Y, moons[2].pos.Y, moons[3].pos.Y,
				moons[0].v.Y, moons[1].v.Y, moons[2].v.Y, moons[3].v.Y)
			if ys[y] {
				returnY = t
			}
			ys[y] = true
		}

		if returnZ == 0 {
			z := fmt.Sprintf("%v %v %v %v %v %v %v %v", moons[0].pos.Z, moons[1].pos.Z, moons[2].pos.Z, moons[3].pos.Z,
				moons[0].v.Z, moons[1].v.Z, moons[2].v.Z, moons[3].v.Z)
			if zs[z] {
				returnZ = t
			}
			zs[z] = true
		}

		if returnX != 0 && returnY != 0 && returnZ != 0 {
			break
		}
		t++
	}
	fmt.Println("Answer 2: ", LCM(returnX, returnY, returnZ))
}

// 'borrowed' from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {

	var day = 12
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")

	// --- PART 1 ----
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

	// --- PART 2 ----

	initMoons := parse()
	getStepsOfFullOrbit(initMoons)

}
