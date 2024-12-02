package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reports [][]int

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := 0
	for _, report := range reports {
		if ok, _ := isSafe(report); ok {
			total++
		}
	}

	fmt.Printf("%v\n", total)
}

func part2() {
	total := 0
	for _, report := range reports {
		ok, i := isSafe(report)
		if ok {
			total++
			continue
		}

		reportA := splice(report, i)
		reportB := splice(report, i-1)
		reportC := splice(report, i-2)
		if ok, _ := isSafe(reportA); ok {
			total++
		} else if ok, _ := isSafe(reportB); ok {
			total++
		} else if ok, _ := isSafe(reportC); ok {
			total++
		}
	}

	fmt.Printf("%v\n", total)
}

// isSafe returns true if the given report is safe, and false otherwise.
// It also returns the index of first violating index
func isSafe(report []int) (bool, int) {
	if len(report) <= 1 {
		return true, -1
	}
	diff := report[1] - report[0]
	if abs(diff) < 1 || abs(diff) > 3 {
		return false, 1
	}
	isIncrease := diff > 0

	for i := 2; i < len(report); i++ {
		diff := report[i] - report[i-1]
		safeDirection := (isIncrease && diff > 0) || (!isIncrease && diff < 0)
		safeMagnitude := abs(diff) >= 1 && abs(diff) <= 3
		if safeDirection && safeMagnitude {
			continue
		}
		return false, i
	}

	return true, -1
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

		parts := strings.Split(line, " ")
		levels := []int{}
		for _, part := range parts {
			level, err := strconv.Atoi(part)
			if err != nil {
				panic(part)
			}
			levels = append(levels, level)
		}

		reports = append(reports, levels)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func splice(a []int, i int) []int {
	if i >= len(a) || i < 0 {
		return a
	}
	ret := []int{}
	ret = append(ret, a[:i]...)
	ret = append(ret, a[i+1:]...)
	return ret
}
