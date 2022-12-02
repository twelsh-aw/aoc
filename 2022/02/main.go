package main

import (
	"fmt"
	"os"
	"strings"
)

type move int

const (
	Rock     move = 1
	Paper    move = 2
	Scissors move = 3
)

type result int

const (
	Lose = 0
	Draw = 3
	Win  = 6
)

type round struct {
	Opponent move
	Player   move
}

type roundResult struct {
	Opponent     move
	PlayerResult result
}

var rounds []round
var roundResults []roundResult

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	score := 0
	for _, r := range rounds {
		score += int(r.Player)
		if r.Player == r.Opponent {
			score += Draw
		} else if (r.Player == Rock && r.Opponent == Scissors) || (r.Player == Paper && r.Opponent == Rock) || (r.Player == Scissors && r.Opponent == Paper) {
			score += Win
		}
	}
	fmt.Printf("%v\n", score)
}

func part2() {
	score := 0
	for _, r := range roundResults {
		score += int(r.PlayerResult)
		if r.PlayerResult == Draw {
			score += int(r.Opponent)
		} else if r.PlayerResult == Lose {
			switch r.Opponent {
			case Rock:
				score += int(Scissors)
			case Paper:
				score += int(Rock)
			case Scissors:
				score += int(Paper)
			}
		} else if r.PlayerResult == Win {
			switch r.Opponent {
			case Rock:
				score += int(Paper)
			case Paper:
				score += int(Scissors)
			case Scissors:
				score += int(Rock)
			}
		}
	}
	fmt.Printf("%v\n", score)
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		moves := strings.Split(line, " ")
		if len(moves) != 2 {
			panic(fmt.Sprintf("invalid line: %s", line))
		}

		r := round{}
		if moves[0] == "A" {
			r.Opponent = Rock
		} else if moves[0] == "B" {
			r.Opponent = Paper
		} else if moves[0] == "C" {
			r.Opponent = Scissors
		} else {
			panic("unknown opponent move")
		}

		rr := roundResult{}
		rr.Opponent = r.Opponent

		if moves[1] == "X" {
			r.Player = Rock
			rr.PlayerResult = Lose
		} else if moves[1] == "Y" {
			r.Player = Paper
			rr.PlayerResult = Draw
		} else if moves[1] == "Z" {
			r.Player = Scissors
			rr.PlayerResult = Win
		} else {
			panic("unknown player move")
		}

		rounds = append(rounds, r)
		roundResults = append(roundResults, rr)
	}
}
