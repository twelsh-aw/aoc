package main

import (
	"fmt"
	"os"
	"strings"
)

type coord struct {
	row int
	col int
}

var (
	coords           = map[coord]string{}
	start            = coord{}
	startBoard       = board{0, 0}
	numRows, numCols int
)

type move struct {
	coord coord
	dir   direction
}

type direction string

const (
	north direction = "N"
	south direction = "S"
	east  direction = "E"
	west  direction = "W"
)

type board struct {
	row int
	col int
}

type boardResult struct {
	numFilledByParity map[int]int
}

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	n := 64
	br := solveBoard(n, false)
	total := br.numFilledByParity[n%2]
	fmt.Printf("%v\n", total)
}

func solveBoard(n int, allowInfinite bool) boardResult {
	first := start
	seen := map[coord]bool{
		first: true,
	}
	cur := []move{
		{first, ""},
	}
	numFilledByParity := map[int]int{
		0: 1,
	}
	br := boardResult{
		numFilledByParity: numFilledByParity,
	}
	for i := 1; i <= n; i++ {
		next := []move{}
		for _, m := range cur {
			for _, nm := range m.getNextValidMoves(seen, allowInfinite) {
				if nm.coord.row >= 0 && nm.coord.col >= 0 && nm.coord.row < numRows && nm.coord.col < numCols {
					next = append(next, nm)
				}
			}
		}
		if len(next) == 0 {
			return br
		}
		cur = append([]move{}, next...)
		numFilledByParity[i%2] += len(next)
	}
	return br
}

func part2() {
	// n := 26501365
	//
}

func (m *move) getNextValidMoves(seen map[coord]bool, allowInfinite bool) []move {
	moves := []move{}
	dirs := []direction{north, south, east, west}
	for _, dir := range dirs {
		// if we previously got to coord by making move in dir, we don't need to go backwards
		if m.dir == dir.opposite() {
			continue
		}
		next := m.coord
		switch dir {
		case north:
			next.row--
		case south:
			next.row++
		case east:
			next.col++
		case west:
			next.col--
		}
		val := coords[next]
		if val == "" && allowInfinite {
			val = coords[next.getMod()]
		} else if val == "" {
			continue
		}

		if seen[next] || val == "#" {
			continue
		}
		moves = append(moves, move{next, dir})
		seen[next] = true
	}
	return moves
}

func (c *coord) getMod() coord {
	mod := coord{c.row % numRows, c.col % numCols}
	if mod.row < 0 {
		mod.row += numRows
	}
	if mod.col < 0 {
		mod.col += numCols
	}
	return mod
}

func (d direction) opposite() direction {
	switch d {
	case north:
		return south
	case east:
		return west
	case west:
		return east
	case south:
		return north
	default:
		panic(d)
	}
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for row, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			continue
		}
		numRows++
		parts := strings.Split(line, "")
		numCols = len(parts)
		for col, v := range parts {
			co := coord{row, col}
			coords[co] = v
			if v == "S" {
				start = co
			}
		}
	}
}
