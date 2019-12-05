package aoc2019_shared

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// TODO add oauth client
func loadInputViaGET(day int) {
	url := fmt.Sprintf("https://adventofcode.com/2019/day/%v/input", day)
	fmt.Printf("fetching %s ...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Could not fetch from", url)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Could not read from", url)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func Load(day int) []int {
	fileName := fmt.Sprintf("./input_%v.txt", day)
	lines := loadInput(fileName)
	return convertToInt(lines)
}

func loadInput(filename string) []string {
	fp, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Could not fetch from", fileName)
	}
	bodyBytes, err := ioutil.ReadAll(fp)

	if err != nil {
		log.Fatal("Could not read from", fileName)
	}
	bodyString := string(bodyBytes)
	lines := strings.Split(bodyString, "\n")
	return lines
}

func convertToInt(lines []string) []int {
	var inputArray = []int{}

	for _, entry := range lines {
		fmt.Println(entry)
		intValue, err := strconv.Atoi(strings.TrimSpace(entry))
		if err != nil {
			panic(err)
		}
		inputArray = append(inputArray, intValue)
	}
	fmt.Println(inputArray)
	return inputArray
}
