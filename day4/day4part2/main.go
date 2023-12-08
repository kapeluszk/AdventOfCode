package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func prepareLines(wholeFile [][]string) (score int) {
	var length int = len(wholeFile)
	cardQuantity := make([]int, length)
	r, _ := regexp.Compile("(\\d+)")

	for i := range cardQuantity {
		cardQuantity[i] = 1
	}
	for i := range wholeFile {
		line := wholeFile[i]
		keys := r.FindAllString(line[0], -1)
		toBeChecked := r.FindAllString(line[1], -1)
		//
		var matchCount int = 0
		for _, key := range keys[1:] {
			for _, elem := range toBeChecked {
				if key == elem {
					matchCount++

				}
			}
		}
		if matchCount != 0 {
			startIndex := i + 1
			var index int = matchCount + i
			var endIndex int = min(length, index)
			currIndex := cardQuantity[i]
			for i := startIndex; i <= endIndex; i++ {
				cardQuantity[i] += currIndex
			}
		}
	}
	sum2 := 0
	for _, v := range cardQuantity {
		sum2 += v
	}
	fmt.Println("Part 2 Sum", sum2)
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
