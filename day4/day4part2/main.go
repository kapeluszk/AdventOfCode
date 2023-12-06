package main

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strings"
)

func prepareLines(wholeFile [][]string) (score int) {
	lenght := len(wholeFile)
	cardQuantity := make([]int, lenght)
	r, _ := regexp.Compile("(\\d+)")

	for i := range cardQuantity {
		cardQuantity[i] = 1
	}
	for i, _ := range wholeFile {
		line := wholeFile[i]
		keys := r.FindAllString(line[0], -1)
		toBeChecked := r.FindAllString(line[1], -1)
		//
		matchCount := 0
		for _, key := range keys[1:] {
			for _, elem := range toBeChecked {
				if key == elem {
					matchCount++

				}
			}
		}
		if matchCount != 0 {
			startIndex := i + 1
			endIndex := math.MinInt(lenght, i+matchCount)
		}
	}
	return score
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var wholeFile [][]string
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), "|")
		wholeFile = append(wholeFile, splitted)
	}

	prepareLines(wholeFile)
}
