package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func reverse(text string) (res string) {
	for _, v := range text {
		res = string(v) + res
	}
	return res
}

func findFirstInt(line string) string {
	re := regexp.MustCompile("\\d")
	digits := re.FindAllString(line, -1)
	first := digits[0]
	return first[0:1]
}

func findInt(line string) int {
	digit1 := findFirstInt(line)
	digit2 := findFirstInt(reverse(line))
	concatenateDigits := digit1 + digit2
	res, _ := strconv.Atoi(concatenateDigits)
	return res
}

func main() {

	file, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	res := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		res += findInt(line)

	}
	fmt.Println(res)
}
