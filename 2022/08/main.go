package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	direction string
	position  struct {
		i int
		j int
	}
)

var (
	trees [][]int
	nRow  int
	nCol  int
)

const (
	up    direction = "up"
	down  direction = "down"
	right direction = "right"
	left  direction = "left"
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	numSeen := 0
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if canPositionBeSeen(position{i, j}) {
				numSeen++
			}
		}
	}

	fmt.Printf("%v\n", numSeen)
}

func part2() {
	maxScore := 0
	for i := 1; i < nRow-1; i++ {
		for j := 1; j < nCol-1; j++ {
			s := getScore(position{i, j})
			if s > maxScore {
				maxScore = s
			}
		}
	}

	fmt.Printf("%v\n", maxScore)
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for i, line := range strings.Split(string(b), "\n") {
		trees = append(trees, []int{})
		for _, v := range strings.Split(line, "") {
			s, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			trees[i] = append(trees[i], s)
		}
	}

	nRow = len(trees)
	nCol = len(trees[0])
}

func canPositionBeSeen(pos position) bool {
	height := trees[pos.i][pos.j]
	points := getHeightsFromDirection(pos, up)
	if canHeightBeSeenFrom(height, points) {
		return true
	}

	points = getHeightsFromDirection(pos, down)
	if canHeightBeSeenFrom(height, points) {
		return true
	}

	points = getHeightsFromDirection(pos, left)
	if canHeightBeSeenFrom(height, points) {
		return true
	}

	points = getHeightsFromDirection(pos, right)
	if canHeightBeSeenFrom(height, points) {
		return true
	}

	return false
}

func getScore(pos position) int {
	height := trees[pos.i][pos.j]
	score := 1

	points := getHeightsFromDirection(pos, up)
	score *= countTreesCanSee(height, points)

	points = getHeightsFromDirection(pos, down)
	score *= countTreesCanSee(height, points)

	points = getHeightsFromDirection(pos, left)
	score *= countTreesCanSee(height, points)

	points = getHeightsFromDirection(pos, right)
	score *= countTreesCanSee(height, points)

	return score
}

func getHeightsFromDirection(pos position, dir direction) []int {
	var points []int
	switch dir {
	case up:
		for i := pos.i - 1; i >= 0; i-- {
			points = append(points, trees[i][pos.j])
		}
		break
	case down:
		for i := pos.i + 1; i < nRow; i++ {
			points = append(points, trees[i][pos.j])
		}
		break
	case left:
		for j := pos.j - 1; j >= 0; j-- {
			points = append(points, trees[pos.i][j])
		}
		break
	case right:
		for j := pos.j + 1; j < nCol; j++ {
			points = append(points, trees[pos.i][j])
		}
		break
	}

	return points
}

func canHeightBeSeenFrom(height int, others []int) bool {
	for _, o := range others {
		if o >= height {
			return false
		}
	}
	return true
}

func countTreesCanSee(height int, others []int) int {
	numSeen := 0
	for _, o := range others {
		numSeen++
		if o >= height {
			break
		}
	}
	return numSeen
}
