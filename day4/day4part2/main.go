package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func prepareLines(splitted []string, i, int) int {
	r, _ := regexp.Compile("(\\d+)")
	keys := r.FindAllString(splitted[0], -1)
	toBeChecked := r.FindAllString(splitted[1], -1)

	matchCount := 0
	for _, key := range keys[1:] {
		for _, elem := range toBeChecked {
			if key == elem {
				matchCount++

			}
		}

	}
	if matchCount != 0 {

	}
	return matchCount
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

	for i, _ := range wholeFile {
		startIndex := i + 1
		endIndex :=
	}

}
