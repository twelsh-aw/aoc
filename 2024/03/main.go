package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input string

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := eval(input)
	fmt.Printf("%v\n", total)
}

func part2() {
	total := 0
	enabled := true
	left := input
	for len(left) > 0 {
		if enabled {
			nextDont := strings.Index(left, "don't()")
			if nextDont == -1 {
				total += eval(left)
				break
			}
			total += eval(left[:nextDont])
			left = left[nextDont:]
			enabled = false
		} else {
			nextDo := strings.Index(left, "do()")
			if nextDo == -1 {
				break
			}
			left = left[nextDo:]
			enabled = true
		}
	}
	fmt.Printf("%v\n", total)
}

func eval(input string) int {
	total := 0
	numIdx := 0
	nums := [2]string{}
	lastValid := rune(0)
	for _, char := range input {
		if (char == 'm' && lastValid == rune(0)) ||
			(char == 'u' && lastValid == 'm') ||
			(char == 'l' && lastValid == 'u') ||
			(char == '(' && lastValid == 'l') ||
			(char == ',' && isDigit(lastValid)) {
			lastValid = char
		} else if isDigit(char) && lastValid == '(' {
			numIdx = 0
			nums[numIdx] += string(char)
			lastValid = char
		} else if isDigit(char) && lastValid == ',' {
			numIdx = 1
			nums[numIdx] += string(char)
			lastValid = char
		} else if isDigit(char) && isDigit(lastValid) {
			nums[numIdx] += string(char)
			lastValid = char
		} else if char == ')' && isDigit(lastValid) {
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			total += num1 * num2

			lastValid = rune(0)
			numIdx = 0
			nums = [2]string{}
		} else {
			lastValid = rune(0)
			numIdx = 0
			nums = [2]string{}
		}
	}
	return total
}

func isDigit(char rune) bool {
	return char == '0' || char == '1' || char == '2' || char == '3' || char == '4' || char == '5' || char == '6' || char == '7' || char == '8' || char == '9'
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			continue
		}
		input += line
	}
}
