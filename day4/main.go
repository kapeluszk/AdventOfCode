package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func matchStrings(splitted []string) int {
	r, _ := regexp.Compile("(\\d+)")
	keys := r.FindAllString(splitted[0], -1)
	toBeChecked := r.FindAllString(splitted[1], -1)

	//part 1
	counterPart1 := 0
	for _, key := range keys[1:] {
		for _, elem := range toBeChecked {
			if key == elem {
				if counterPart1 == 0 {
					counterPart1 = 1
				} else {
					counterPart1 *= 2
				}

			}
		}

	}
	return counterPart1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	score := 0
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), "|")
		score += matchStrings(splitted)
	}
	fmt.Printf("Part 1 result is: %d", score)
}
