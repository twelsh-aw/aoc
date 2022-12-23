package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type coordRange struct {
	minRow int
	maxRow int
	minCol int
	maxCol int
}

type coord struct {
	row int
	col int
}

type direction string

const (
	north     direction = "N"
	northEast direction = "NE"
	northWest direction = "NW"
	south     direction = "S"
	southEast direction = "SE"
	southWest direction = "SW"
	west      direction = "W"
	east      direction = "E"
)

var (
	originalElves = make(map[coord]string)
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	elves := clone(originalElves)
	directions := []direction{north, south, west, east}
	for round := 1; round <= 10; round++ {
		elves, directions, _ = moveElves(elves, directions)
	}

	rng := getCoordRange(elves)
	rectangleArea := rng.length() * rng.width()
	numEmpty := rectangleArea - len(elves)
	fmt.Printf("%v\n", numEmpty)
}

func part2() {
	elves := clone(originalElves)
	directions := []direction{north, south, west, east}
	roundNumber := 0
	numMoves := -1
	for numMoves != 0 {
		roundNumber++
		elves, directions, numMoves = moveElves(elves, directions)
	}

	fmt.Printf("%v\n", roundNumber)
}

func moveElves(elves map[coord]string, directions []direction) (map[coord]string, []direction, int) {
	nextCoords := make(map[coord]coord)
	nextCounts := make(map[coord]int)
	for elfCoord := range elves {
		next := getNextCoord(elfCoord, elves, directions)
		nextCoords[elfCoord] = next
		nextCounts[next]++
	}

	numMoves := 0
	for cur, next := range nextCoords {
		if nextCounts[next] > 1 {
			continue
		}

		delete(elves, cur)
		elves[next] = "#"
		if next != cur {
			numMoves++
		}
	}

	directions = append(append([]direction{}, directions[1:]...), directions[0])
	return elves, directions, numMoves
}

func getCoordRange(elves map[coord]string) coordRange {
	r := coordRange{
		minRow: math.MaxInt,
		maxRow: 0,
		minCol: math.MaxInt,
		maxCol: 0,
	}
	for c := range elves {
		if c.row > r.maxRow {
			r.maxRow = c.row
		}
		if c.row < r.minRow {
			r.minRow = c.row
		}
		if c.col > r.maxCol {
			r.maxCol = c.col
		}
		if c.col < r.minCol {
			r.minCol = c.col
		}
	}

	return r
}

func (r coordRange) length() int {
	return r.maxRow - r.minRow + 1
}

func (r coordRange) width() int {
	return r.maxCol - r.minCol + 1
}

func getNextCoord(cur coord, all map[coord]string, directions []direction) coord {
	allDirs := []direction{north, northEast, northWest, south, southEast, southWest, east, west}
	coordsByDirection := make(map[direction]coord)
	canMove := false
	for _, d := range allDirs {
		coordsByDirection[d] = getCoordInDirection(cur, d)
		if all[coordsByDirection[d]] == "#" {
			canMove = true
		}
	}
	if !canMove {
		return cur
	}

	for _, dir := range directions {
		var toCheck []direction
		switch dir {
		case north:
			toCheck = []direction{north, northEast, northWest}
		case south:
			toCheck = []direction{south, southEast, southWest}
		case east:
			toCheck = []direction{east, northEast, southEast}
		case west:
			toCheck = []direction{west, southWest, northWest}
		default:
			panic(dir)
		}

		canMove = true
		for _, d := range toCheck {
			if all[coordsByDirection[d]] == "#" {
				canMove = false
			}
		}
		if canMove {
			return coordsByDirection[dir]
		}
	}

	return cur
}

func getCoordInDirection(cur coord, dir direction) coord {
	switch dir {
	case north:
		return coord{cur.row - 1, cur.col}
	case northEast:
		return coord{cur.row - 1, cur.col + 1}
	case northWest:
		return coord{cur.row - 1, cur.col - 1}
	case south:
		return coord{cur.row + 1, cur.col}
	case southEast:
		return coord{cur.row + 1, cur.col + 1}
	case southWest:
		return coord{cur.row + 1, cur.col - 1}
	case east:
		return coord{cur.row, cur.col + 1}
	case west:
		return coord{cur.row, cur.col - 1}
	default:
		panic(dir)
	}
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for row, line := range strings.Split(string(b), "\n") {
		for col, val := range strings.Split(line, "") {
			if val == "#" {
				c := coord{row: row, col: col}
				originalElves[c] = val
			}
		}
	}
}

func clone[K comparable, V comparable](m map[K]V) map[K]V {
	c := make(map[K]V)
	for k, v := range m {
		c[k] = v
	}

	return c
}

func print(elves map[coord]string) {
	rng := getCoordRange(elves)
	for row := rng.minRow; row <= rng.maxRow; row++ {
		var rowString []string
		for col := rng.minCol; col <= rng.maxCol; col++ {
			if elves[coord{row, col}] == "#" {
				rowString = append(rowString, "#")
			} else {
				rowString = append(rowString, ".")
			}
		}
		fmt.Println(strings.Join(rowString, " "))
	}
}
