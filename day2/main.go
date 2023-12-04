package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var cubeLimit = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func processGames(color, line string) bool {
	r, _ := regexp.Compile("(\\d+) " + color)
	matches := r.FindAllStringSubmatch(line, -1)
	validColor := true

	for _, submatch := range matches {
		colorCount, _ := strconv.Atoi(submatch[1])

		if validColor && (colorCount > cubeLimit[color]) {
			validColor = false
		}
	}
	return validColor
}

func separateNumberFromText(text string) (int, error) {
	r, _ := regexp.Compile("(\\d+)")
	res := r.FindAllString(text, -1)

	return strconv.Atoi(res[0])
}

func findGameIndex(line string) (index int) {
	r, _ := regexp.Compile("Game (\\d+)")
	gameMatch := r.FindAllString(line, -1)
	index, _ = separateNumberFromText(gameMatch[0])

	return index
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	res := 0
	for scanner.Scan() {
		line := scanner.Text()

		gameIndex := findGameIndex(line)
		validRed := processGames("red", line)
		validBlue := processGames("blue", line)
		validGreen := processGames("green", line)

		if validRed && validGreen && validBlue {
			res += gameIndex
		}
	}
	fmt.Println(res)
}
