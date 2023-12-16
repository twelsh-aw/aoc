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

type direction string

const (
	dirRight = "r"
	dirLeft  = "l"
	dirUp    = "u"
	dirDown  = "d"
)

type beam struct {
	c coord
	d direction
}

var coords = map[coord]string{}
var numRows, numCols int

func main() {
	readInput()
	initial := beam{
		coord{0, 0},
		dirRight,
	}
	fmt.Println(part1(initial))
	fmt.Println(part2())
}

func part1(initial beam) int {
	beams := []beam{initial}
	seen := map[beam]bool{}
	energized := map[coord]bool{}
	seen[initial] = true
	energized[initial.c] = true
	for len(beams) > 0 {
		nextBeams := []beam{}
		for _, beam := range beams {
			for _, next := range beam.getNext() {
				if seen[next] {
					continue
				}
				nextBeams = append(nextBeams, next)
				seen[next] = true
				energized[next.c] = true
			}
		}
		beams = append([]beam{}, nextBeams...)
	}

	return len(energized)
}

func part2() int {
	initials := []beam{}
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if row == 0 {
				initials = append(initials, beam{
					coord{row, col},
					dirDown,
				})
			}
			if col == 0 {
				initials = append(initials, beam{
					coord{row, col},
					dirRight,
				})
			}
			if row == numRows-1 {
				initials = append(initials, beam{
					coord{row, col},
					dirUp,
				})
			}
			if col == numCols-1 {
				initials = append(initials, beam{
					coord{row, col},
					dirLeft,
				})
			}
		}
	}

	max := 0
	for _, initial := range initials {
		score := part1(initial)
		if score > max {
			max = score
		}
	}
	return max
}

func (b *beam) getNext() []beam {
	next := []beam{}
	val := coords[b.c]
	nextDirections := b.getNextDirections(val)
	for _, dir := range nextDirections {
		nextCoord, ok := b.c.move(dir)
		if ok {
			next = append(next, beam{nextCoord, dir})
		}
	}
	return next
}

func (b *beam) getNextDirections(val string) []direction {
	switch val {
	case ".":
		return []direction{b.d}
	case "-":
		if b.d == dirRight || b.d == dirLeft {
			return []direction{b.d}
		} else {
			return []direction{dirLeft, dirRight}
		}
	case "|":
		if b.d == dirUp || b.d == dirDown {
			return []direction{b.d}
		} else {
			return []direction{dirUp, dirDown}
		}
	case "/":
		if b.d == dirRight {
			return []direction{dirUp}
		}
		if b.d == dirLeft {
			return []direction{dirDown}
		}
		if b.d == dirUp {
			return []direction{dirRight}
		}
		if b.d == dirDown {
			return []direction{dirLeft}
		}
	case "\\":
		if b.d == dirRight {
			return []direction{dirDown}
		}
		if b.d == dirLeft {
			return []direction{dirUp}
		}
		if b.d == dirUp {
			return []direction{dirLeft}
		}
		if b.d == dirDown {
			return []direction{dirRight}
		}
	default:
		panic(val)
	}
	panic("switch fall")
}

func (c *coord) move(dir direction) (next coord, valid bool) {
	next = *c
	switch dir {
	case dirRight:
		next.col++
	case dirLeft:
		next.col--
	case dirDown:
		next.row++
	case dirUp:
		next.row--
	default:
		panic(dir)
	}
	valid = next.row >= 0 && next.row < numRows && next.col >= 0 && next.col < numCols
	return
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
			c := coord{row, col}
			coords[c] = v
		}
	}
}
