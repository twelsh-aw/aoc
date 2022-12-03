package main

import (
	"fmt"
	"os"
	"strings"
)

type item string

type rucksack struct {
	ComponentA map[item]bool
	ComponentB map[item]bool
	Common     item
	AllItems   map[item]bool
}

var scores = map[item]int{
	"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26,
	"A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52,
}

var rucksacks []rucksack

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	score := 0
	for _, rs := range rucksacks {
		score += scores[rs.Common]
	}
	fmt.Printf("%v\n", score)
}

func part2() {
	score := 0
	for i := 0; i < len(rucksacks)/3; i++ {
		var sacks [3]rucksack
		sacks[0] = rucksacks[3*i]
		sacks[1] = rucksacks[3*i+1]
		sacks[2] = rucksacks[3*i+2]
		common := getCommon(sacks)
		score += scores[common]
	}

	fmt.Printf("%v\n", score)
}

func getCommon(sacks [3]rucksack) item {
	for i := range sacks[0].AllItems {
		v := sacks[0].AllItems[i] && sacks[1].AllItems[i] && sacks[2].AllItems[i]
		if v {
			return i
		}
	}

	fmt.Println(sacks[0].AllItems)
	fmt.Println(sacks[1].AllItems)
	fmt.Println(sacks[2].AllItems)
	panic("nothing in common")
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		split := strings.Split(line, "")
		n := len(split)
		rs := rucksack{
			ComponentA: make(map[item]bool),
			ComponentB: make(map[item]bool),
			AllItems:   make(map[item]bool),
		}
		for i, v := range split {
			rs.AllItems[item(v)] = true
			if i < n/2 {
				rs.ComponentA[item(v)] = true
			} else {
				rs.ComponentB[item(v)] = true
			}

			_, okA := rs.ComponentA[item(v)]
			_, okB := rs.ComponentB[item(v)]
			if okA && okB {
				rs.Common = item(v)
			}
		}

		rucksacks = append(rucksacks, rs)
	}
}
