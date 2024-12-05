package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	rules   = map[int]map[int]bool{}
	updates = [][]int{}
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := 0
	for _, update := range updates {
		if isGoodUpdate(update) {
			middle := update[len(update)/2]
			total += middle
		}
	}
	fmt.Printf("%v\n", total)
}

func part2() {
	total := 0
	for _, update := range updates {
		if isGoodUpdate(update) {
			continue
		}

		sort.Slice(update, func(i, j int) bool {
			if rules[update[i]][update[j]] {
				return true
			} else if rules[update[j]][update[i]] {
				return false
			} else {
				panic(fmt.Sprintf("unable to sort: %v. failed comparision: %v, %v", update, i, j))
			}
		})

		middle := update[len(update)/2]
		total += middle
	}
	fmt.Printf("%v\n", total)
}

func isGoodUpdate(update []int) bool {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			if !rules[update[i]][update[j]] {
				return false
			}
		}
	}
	return true
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	isRuleParsing := true
	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			isRuleParsing = false
			continue
		}
		if isRuleParsing {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				panic(parts)
			}
			num1, _ := strconv.Atoi(parts[0])
			num2, _ := strconv.Atoi(parts[1])
			if _, ok := rules[num1]; !ok {
				rules[num1] = map[int]bool{}
			}
			rules[num1][num2] = true
		} else {
			parts := strings.Split(line, ",")
			nums := []int{}
			for _, num := range parts {
				n, _ := strconv.Atoi(num)
				nums = append(nums, n)
			}
			updates = append(updates, nums)
		}
	}
}
