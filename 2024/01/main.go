package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	list1 []int
	list2 []int
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	total := 0
	for i := 0; i < len(list1); i++ {
		total += absDiff(list1[i], list2[i])
	}

	fmt.Printf("%v\n", total)
}

func part2() {
	freqs := make(map[int]int)
	for i := 0; i < len(list1); i++ {
		freqs[list2[i]]++
	}

	total := 0
	for i := 0; i < len(list1); i++ {
		total += list1[i] * freqs[list1[i]]
	}

	fmt.Printf("%v\n", total)
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
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			panic(line)
		}

		p1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		p2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		list1 = append(list1, p1)
		list2 = append(list2, p2)
	}

	if len(list1) != len(list2) {
		panic("list1 and list2 not same length")
	}
}

func absDiff(x int, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}
