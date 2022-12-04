package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	first  assignment
	second assignment
}

type assignment struct {
	start int
	end   int
}

var pairs []pair

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	numFullOverlaps := 0
	for _, p := range pairs {
		if p.first.start <= p.second.start && p.first.end >= p.second.end {
			numFullOverlaps++
		} else if p.second.start <= p.first.start && p.second.end >= p.first.end {
			numFullOverlaps++
		}
	}
	fmt.Printf("%v\n", numFullOverlaps)
}

func part2() {
	numOverlaps := 0
	for _, p := range pairs {
		if p.first.end < p.second.start {
			continue
		}
		if p.first.start > p.second.end {
			continue
		}
		numOverlaps++
	}
	fmt.Printf("%v\n", numOverlaps)
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			panic("not enough parts")
		}

		var p pair
		for i, part := range parts {
			zones := strings.Split(part, "-")
			if len(zones) != 2 {
				panic("not enough zones")
			}

			start, err := strconv.Atoi(zones[0])
			if err != nil {
				panic(fmt.Sprintf("zone: %s, %s", zones[0], err))
			}

			end, err := strconv.Atoi(zones[1])
			if err != nil {
				panic(fmt.Sprintf("zone: %s, %s", zones[1], err))
			}
			a := assignment{
				start: start,
				end:   end,
			}

			if i == 0 {
				p.first = a
			} else {
				p.second = a
			}
		}

		pairs = append(pairs, p)
	}
}
