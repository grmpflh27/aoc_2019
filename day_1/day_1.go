package main

import (
	"fmt"

	"github.com/grmpflh27/aoc_2019/aoc2019_shared"
)

func transform(in int) int {
	return (in / 3) - 2
}

func main() {
	var day = 1
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")
	masses := aoc2019_shared.Load(day)

	//1
	
	sumModules := 0
	for _, mass := range masses {
		cur := transform(mass)
		sumModules += cur
	}
	fmt.Printf("Answer 1: %v\n", sumModules)

	// 2

	sumFuels := 0
	for _, mass := range masses {
		for {
			mass = transform(mass)
			if mass <= 0 {
				break
			}
			sumFuels += mass
		}
	}
	fmt.Printf("Answer 2: %v\n", sumModules)

}
