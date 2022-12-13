package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type pair struct {
	first  string
	second string
}

var pairs []pair

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	score := 0
	for idx := 1; idx <= len(pairs); idx++ {
		if inCorrectOrder(pairs[idx-1]) {
			score += idx
		}
	}

	fmt.Printf("%v\n", score)
}

func part2() {
	dividers := []string{"[[2]]", "[[6]]"}
	sorted := append([]string{}, dividers...)
	for _, p := range pairs {
		sorted = append(sorted, p.first)
		sorted = append(sorted, p.second)
	}

	sort.Slice(sorted, func(i, j int) bool {
		return inCorrectOrder(pair{sorted[i], sorted[j]})
	})

	score := 1
	for idx := 1; idx <= len(sorted); idx++ {
		if sorted[idx-1] == dividers[0] || sorted[idx-1] == dividers[1] {
			score *= idx
		}
	}
	fmt.Printf("%v\n", score)
}

func inCorrectOrder(p pair) bool {
	p1 := getNextPair(p.first)
	p2 := getNextPair(p.second)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from pair: %v (%v, %v) with error: %s\n", p, p1, p2, r)
			panic(r)
		}
	}()

	if p1.first == "" && p2.first == "" {
		if p2.second == p2.second {
			return true
		}
		return inCorrectOrder(pair{p1.second, p2.second})
	} else if p1.first == "" {
		return true
	} else if p2.first == "" {
		return false
	}

	if p1.first[0:1] == "[" && p2.first[0:1] == "[" {
		if p1.first == p2.first {
			return inCorrectOrder(pair{p1.second, p2.second})
		}
		return inCorrectOrder(pair{p1.first, p2.first})
	} else if p1.first[0:1] == "[" {
		if p1.first == "["+p2.first+"]" {
			return inCorrectOrder(pair{p1.second, p2.second})
		}

		return inCorrectOrder(pair{p1.first, "[" + p2.first + "]"})
	} else if p2.first[0:1] == "[" {
		if "["+p1.first+"]" == p2.first {
			return inCorrectOrder(pair{p1.second, p2.second})
		}

		return inCorrectOrder(pair{"[" + p1.first + "]", p2.first})
	}

	v1, err := strconv.Atoi(p1.first)
	if err != nil {
		panic(err)
	}

	v2, err := strconv.Atoi(p2.first)
	if err != nil {
		panic(err)
	}

	if v1 == v2 {
		return inCorrectOrder(pair{p1.second, p2.second})
	}

	return v1 < v2
}

func getNextPair(v string) pair {
	if v == "" {
		return pair{}
	}

	if v[0:1] != "[" {
		panic(v)
	}

	if v[1:2] == "]" {
		first := ""
		second := v[2:]
		return pair{first, second}
	}

	if v[1:2] != "[" {
		comma := strings.Index(v[1:], ",")
		if comma == -1 {
			comma = len(v) - 2
		}
		first := v[1 : comma+1]
		second := ""
		if len(v) > comma+2 {
			second = "[" + v[comma+2:]
		}
		return pair{first, second}
	}

	// nested array case: find where to split things
	numLeft := 0
	numRight := 0
	for i := 1; i < len(v); i++ {
		if v[i:i+1] == "[" {
			numLeft++
		} else if v[i:i+1] == "]" {
			numRight++
		}

		if numLeft == numRight {
			first := v[1 : i+1]
			second := ""
			if len(v) > i+2 {
				second = "[" + v[i+2:]
			}

			return pair{first, second}
		}
	}

	panic("no pairs made")
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	curPair := pair{}
	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			pairs = append(pairs, curPair)
			curPair = pair{}
		}

		if curPair.first == "" {
			curPair.first = line
		} else {
			curPair.second = line
		}
	}
}
