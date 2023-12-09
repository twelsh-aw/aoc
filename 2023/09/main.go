package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var histories [][]int

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := 0
	for _, history := range histories {
		cur := append([]int{}, history...)
		diffs := [][]int{cur}
		for {
			diff, allZero := getDiff(cur)
			if allZero {
				break
			}
			diffs = append(diffs, diff)
			cur = append([]int{}, diff...)
		}

		next := 0
		for i := len(diffs) - 1; i >= 0; i-- {
			next += diffs[i][len(diffs[i])-1]
		}

		total += next
	}
	fmt.Printf("%v\n", total)
}

func part2() {
	total := 0
	for _, history := range histories {
		cur := append([]int{}, history...)
		diffs := [][]int{cur}
		for {
			diff, allZero := getDiff(cur)
			if allZero {
				break
			}
			diffs = append(diffs, diff)
			cur = append([]int{}, diff...)
		}

		prev := 0
		for i := len(diffs) - 1; i >= 0; i-- {
			prev = diffs[i][0] - prev
		}

		total += prev
	}
	fmt.Printf("%v\n", total)
}

func getDiff(nums []int) (diff []int, allZero bool) {
	allZero = true
	for i := 0; i < len(nums)-1; i++ {
		d := nums[i+1] - nums[i]
		diff = append(diff, d)
		if d != 0 {
			allZero = false
		}
	}
	return
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
		history := []int{}
		for _, part := range parts {
			n, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			history = append(history, n)
		}
		histories = append(histories, history)
	}
}
