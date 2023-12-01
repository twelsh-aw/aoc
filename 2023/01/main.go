package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var (
	inputs []string
	digits = map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	words = map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := 0
	for _, i := range inputs {
		total += extractDigits(i)
	}

	fmt.Printf("%v\n", total)
}

func part2() {
	total := 0
	for _, i := range inputs {
		total += extractDigitsUsingWords(i)
	}

	fmt.Printf("%v\n", total)
}

func extractDigits(s string) int {
	var first, last *int
	for _, c := range s {
		if casted, ok := digits[c]; ok {
			last = &casted
			if first == nil {
				first = &casted
			}
		}
	}

	if first == nil || last == nil {
		panic("no digits in string")
	}

	return *first*10 + *last
}

func extractDigitsUsingWords(s string) int {
	var first, last *int
	firstIndex := len(s) - 1
	lastIndex := 0
	for idx, c := range s {
		if casted, ok := digits[c]; ok {
			last = &casted
			lastIndex = idx
			if first == nil {
				firstIndex = idx
				first = &casted
			}
		}
	}

	firstWord, hasFirstWord := getFirstNumberWord(s[:firstIndex])
	if hasFirstWord {
		first = &firstWord
	}
	lastWord, hasLastWord := getLastNumberWord(s[lastIndex:])
	if hasLastWord {
		last = &lastWord
	}

	if first == nil || last == nil {
		panic("no digits in string")
	}

	return *first*10 + *last
}

func getFirstNumberWord(s string) (int, bool) {
	firstIdx := math.MaxInt
	firstWord := ""
	found := false
	for word := range words {
		idx := strings.Index(s, word)
		if idx != -1 && idx < firstIdx {
			found = true
			firstIdx = idx
			firstWord = word
		}
	}

	return words[firstWord], found
}

func getLastNumberWord(s string) (int, bool) {
	lastIdx := -1
	lastWord := ""
	found := false
	for word := range words {
		idx := strings.LastIndex(s, word)
		if idx != -1 && idx > lastIdx {
			found = true
			lastIdx = idx
			lastWord = word
		}
	}

	return words[lastWord], found
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		if len(line) > 0 {
			inputs = append(inputs, line)
		}
	}
}
