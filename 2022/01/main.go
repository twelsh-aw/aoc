package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	weights []int
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	maxWeight := 0
	for _, w := range weights {
		if w > maxWeight {
			maxWeight = w
		}
	}

	fmt.Println(maxWeight)
}

func part2() {
	maxWeights := []int{0, 0, 0}
	for _, w := range weights {
		if w > maxWeights[2] {
			maxWeights = []int{
				maxWeights[0],
				maxWeights[1],
				w,
			}
			sort.Slice(maxWeights, func(i, j int) bool {
				return maxWeights[i] > maxWeights[j]
			})
		}
	}

	fmt.Println(maxWeights[0] + maxWeights[1] + maxWeights[2])
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	curWeight := 0
	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			weights = append(weights, curWeight)
			curWeight = 0
		} else {
			weight, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			curWeight += weight
		}
	}
}
