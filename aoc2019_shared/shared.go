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

func Load(day int, sep string) []int {
	fileName := fmt.Sprintf("./input_%v.txt", day)
	lines := loadInput(fileName, sep)
	return convertToInt(lines)
}

func LoadStr(day int) [][]string {
	fileName := fmt.Sprintf("./input_%v.txt", day)
	lines := loadInput(fileName, "\n")

	fmt.Println("len lines", len(lines))
	final := make([][]string, 3)
	for i, line := range lines {
		words := strings.Split(line, ",")
		final[i] = words
	}
	return final
}

// TODO add oauth client
func loadInputViaGET(day int) string {
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
	return bodyString
}

func loadInput(fileName string, sep string) []string {
	fp, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Could not fetch from", fileName)
	}
	bodyBytes, err := ioutil.ReadAll(fp)

	if err != nil {
		log.Fatal("Could not read from", fileName)
	}
	bodyString := string(bodyBytes)
	lines := strings.Split(bodyString, sep)
	return lines
}

func convertToInt(lines []string) []int {
	var inputArray = []int{}

	for _, entry := range lines {
		intValue, err := strconv.Atoi(strings.TrimSpace(entry))
		if err != nil {
			panic(err)
		}
		inputArray = append(inputArray, intValue)
	}
	return inputArray
}
