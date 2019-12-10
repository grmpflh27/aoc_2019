package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	WIDTH       = 25
	HEIGHT      = 6
	LAYERLENGTH = WIDTH * HEIGHT
)

func load() string {
	fp, _ := os.Open("./input_8.txt")
	fpBytes, _ := ioutil.ReadAll(fp)
	return string(fpBytes)
}

func countRune(chunk string, searchRune rune) int {
	var zeroCnt int
	for i := range chunk {
		if rune(chunk[i]) == searchRune {
			zeroCnt++
		}
	}
	return zeroCnt
}

func main() {
	var day = 8
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")

	oned := load()

	numberOfLayers := len(oned) / LAYERLENGTH
	fmt.Printf("# layers: %v\n", numberOfLayers)

	// chunk into layers
	byLayer := make([]string, numberOfLayers)
	for cnt := 0; cnt < numberOfLayers; cnt++ {
		byLayer[cnt] = oned[(cnt)*LAYERLENGTH : (cnt+1)*LAYERLENGTH]
	}

	// count zeros
	var minZeroCntLayerIndex int = 0
	var minZeroCnt int = LAYERLENGTH

	for i := range byLayer {
		curZeroCnt := countRune(byLayer[i], '0')
		if curZeroCnt < minZeroCnt {
			minZeroCnt = curZeroCnt
			minZeroCntLayerIndex = i
		}
	}

	OneCnt := countRune(byLayer[minZeroCntLayerIndex], '1')
	TwoCnt := countRune(byLayer[minZeroCntLayerIndex], '2')

	fmt.Println("Answer 1:", OneCnt*TwoCnt)
}
