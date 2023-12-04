package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func convertToInt(s string) int {
	stringToInt := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"enin":  9,
		"thgie": 8,
		"neves": 7,
		"xis":   6,
		"evif":  5,
		"ruof":  4,
		"eerht": 3,
		"owt":   2,
		"eno":   1,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	if stringToInt[s] != 0 || s == "0" {
		return stringToInt[s]
	} else if stringToInt[s[0:1]] != 0 || s[0:1] == "0" {
		return stringToInt[s[0:1]]
	}
	res, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return res

}

func reverse(text string) (res string) {
	for _, v := range text {
		res = string(v) + res
	}
	return res
}

func findFirstInt(line string) (res [2]int) {
	re := regexp.MustCompile("(\\d|one|two|three|four|five|six|seven|eight|nine|enin|thgie|neves|xis|evif|ruof|eerht|owt|eno)")
	res[0] = convertToInt(re.FindString(line))
	res[1] = convertToInt(re.FindString(reverse(line)))
	return res
}

func findInt(line string) int {
	digits := findFirstInt(line)
	res := digits[0]*10 + digits[1]
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
