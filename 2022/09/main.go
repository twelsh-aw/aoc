package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type direction string

const (
	up    direction = "U"
	down  direction = "D"
	left  direction = "L"
	right direction = "R"
)

type move struct {
	dir direction
	num int
}

type position struct {
	x int
	y int
}

var (
	moves []move
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	headPos := position{0, 0}
	tails := []*position{
		{0, 0},
	}

	tailVisited := make(map[position]bool)
	for _, m := range moves {
		for i := 0; i < m.num; i++ {
			moveDirection(&headPos, tails, m.dir)
			tailVisited[*tails[0]] = true
		}
	}
	fmt.Printf("%+v\n", len(tailVisited))
}

func part2() {
	headPos := position{0, 0}
	var tails []*position
	for i := 0; i < 9; i++ {
		tails = append(tails, &position{0, 0})
	}

	tailVisited := make(map[position]bool)
	for _, m := range moves {
		for i := 0; i < m.num; i++ {
			moveDirection(&headPos, tails, m.dir)
			tailVisited[*tails[8]] = true
		}
	}

	fmt.Printf("%+v\n", len(tailVisited))
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			panic(fmt.Sprintf("unexpected num parts: %s", line))
		}

		dir := direction(parts[0])
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		m := move{dir, num}
		moves = append(moves, m)
	}
}

func moveDirection(headPos *position, tails []*position, dir direction) {
	switch dir {
	case up:
		headPos.y++
	case down:
		headPos.y--
	case left:
		headPos.x--
	case right:
		headPos.x++
	}

	for i := 0; i < len(tails); i++ {
		if i == 0 {
			follow(tails[i], headPos)
			continue
		}

		follow(tails[i], tails[i-1])
	}
}

func follow(tailPos *position, aheadPos *position) {
	if tailPos.x == aheadPos.x { // left or right by more than 1 space
		if tailPos.y+1 < aheadPos.y {
			tailPos.y++
		} else if tailPos.y-1 > aheadPos.y {
			tailPos.y--
		}
	} else if tailPos.y == aheadPos.y { // up or down by more than one space
		if tailPos.x+1 < aheadPos.x {
			tailPos.x++
		} else if tailPos.x-1 > aheadPos.x {
			tailPos.x--
		}
	} else if tailPos.x != aheadPos.x && tailPos.y != aheadPos.y { // diagonal by more than one space in at least one direction
		if tailPos.y+1 < aheadPos.y {
			tailPos.y++
			if tailPos.x > aheadPos.x {
				tailPos.x--
			} else {
				tailPos.x++
			}
		} else if tailPos.y-1 > aheadPos.y {
			tailPos.y--
			if tailPos.x > aheadPos.x {
				tailPos.x--
			} else {
				tailPos.x++
			}
		} else if tailPos.x+1 < aheadPos.x {
			tailPos.x++
			if tailPos.y > aheadPos.y {
				tailPos.y--
			} else {
				tailPos.y++
			}
		} else if tailPos.x-1 > aheadPos.x {
			tailPos.x--
			if tailPos.y > aheadPos.y {
				tailPos.y--
			} else {
				tailPos.y++
			}
		}
	}
}
