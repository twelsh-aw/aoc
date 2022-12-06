package main

import (
	"fmt"
	"os"
	"strings"
)

var line []string

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	result := countUntilDistinct(4)
	fmt.Printf("%v\n", result)
}

func part2() {
	result := countUntilDistinct(14)
	fmt.Printf("%v\n", result)
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	line = strings.Split(string(b), "")
}

func countUntilDistinct(length int) int {
	result := -1
	for i := 0; i < len(line)-length; i++ {
		markers := line[i : i+length]
		hasDuplicates := false
		occurrences := make(map[string]bool)
		for _, m := range markers {
			if occurrences[m] {
				hasDuplicates = true
				break
			}
			occurrences[m] = true
		}

		if !hasDuplicates {
			result = i + length
			break
		}
	}

	return result
}
