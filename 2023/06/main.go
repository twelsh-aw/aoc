package main

import (
	"fmt"
)

type race struct {
	time     int
	distance int
}

var races []race

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	score := 1
	for _, race := range races {
		numWinningPaths := 0
		for i := 1; i <= race.time; i++ {
			if i*(race.time-i) > race.distance {
				numWinningPaths++
			}
		}
		score *= numWinningPaths
	}
	fmt.Printf("%v\n", score)
}

func part2() {
	r := race{
		time:     47986698,
		distance: 400121310111540,
	}
	numWinningPaths := 0
	for i := 1; i <= r.time; i++ {
		if i*(r.time-i) > r.distance {
			numWinningPaths++
		}
	}

	fmt.Printf("%v\n", numWinningPaths)
}

func readInput() {
	// example
	races = []race{
		{
			time:     7,
			distance: 9,
		},
		{
			time:     15,
			distance: 40,
		},
		{
			time:     30,
			distance: 200,
		},
	}

	// input
	races = []race{
		{
			time:     47,
			distance: 400,
		},
		{
			time:     98,
			distance: 1213,
		},
		{
			time:     66,
			distance: 1011,
		},
		{
			time:     98,
			distance: 1540,
		},
	}
}
