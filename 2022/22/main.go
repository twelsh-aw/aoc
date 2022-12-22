package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	row int
	col int
}

type instruction struct {
	steps     int
	direction direction
}

type coordRange struct {
	min int
	max int
}

type direction string
type rotation string

const (
	right direction = "R"
	down  direction = "D"
	left  direction = "L"
	up    direction = "U"

	clockwise        rotation = "R"
	counterClockwise rotation = "L"
)

var (
	coords       = make(map[coord]string)
	instructions []instruction
	rowRanges    = make(map[int]coordRange)
	colRanges    = make(map[int]coordRange)
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	curPosition := coord{1, rowRanges[1].min}
	curDirection := right
	canMove := true
	for _, ins := range instructions {
		curDirection = ins.direction
		for i := 0; i < ins.steps; i++ {
			curPosition, canMove = move(curPosition, curDirection)
			if !canMove {
				break
			}
		}
	}

	password := (curPosition.row * 1000) + (curPosition.col * 4) + curDirection.asValue()
	fmt.Printf("%v\n", password)
}

func move(pos coord, dir direction) (coord, bool) {
	next := pos
	switch dir {
	case right:
		next.col++
		if next.col > rowRanges[pos.row].max {
			next.col = rowRanges[pos.row].min
		}
	case left:
		next.col--
		if next.col < rowRanges[pos.row].min {
			next.col = rowRanges[pos.row].max
		}
	case down:
		next.row++
		if next.row > colRanges[pos.col].max {
			next.row = colRanges[pos.col].min
		}
	case up:
		next.row--
		if next.row < colRanges[pos.col].min {
			next.row = colRanges[pos.col].max
		}
	}

	tile := coords[next]
	if tile == "." {
		return next, true
	} else if tile == "#" {
		return pos, false
	} else {
		panic(pos)
	}
}

func part2() {
	fmt.Printf("%v\n", nil)
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	parseInstructions := false
	curDirection := right
	curDigits := ""
	for i, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			parseInstructions = true
			continue
		}

		if !parseInstructions {
			row := i + 1
			if _, ok := rowRanges[row]; !ok {
				rowRanges[row] = coordRange{min: math.MaxInt, max: 0}
			}
			parts := strings.Split(line, "")
			for j, part := range parts {
				col := j + 1
				if _, ok := colRanges[col]; !ok {
					colRanges[col] = coordRange{min: math.MaxInt, max: 0}
				}

				if part == " " {
					continue
				} else if part == "#" || part == "." {
					c := coord{row, col}
					coords[c] = part

					rowRange := rowRanges[row]
					colRange := colRanges[col]
					if col > rowRange.max {
						rowRange.max = col
						rowRanges[row] = rowRange
					}
					if col < rowRange.min {
						rowRange.min = col
						rowRanges[row] = rowRange
					}
					if row > colRange.max {
						colRange.max = row
						colRanges[col] = colRange
					}
					if row < colRange.min {
						colRange.min = row
						colRanges[col] = colRange
					}
				} else {
					panic(part)
				}
			}
		}

		if parseInstructions {
			parts := strings.Split(line, "")
			for _, part := range parts {
				if part == string(clockwise) || part == string(counterClockwise) {
					ins := instruction{direction: curDirection}
					ins.steps, err = strconv.Atoi(curDigits)
					if err != nil {
						panic(err)
					}
					instructions = append(instructions, ins)
					curDirection = getNextDirection(curDirection, rotation(part))
					curDigits = ""
				} else {
					curDigits += part
				}
			}

			ins := instruction{direction: curDirection}
			ins.steps, err = strconv.Atoi(curDigits)
			if err != nil {
				panic(err)
			}
			instructions = append(instructions, ins)
		}
	}
}

func getNextDirection(d direction, r rotation) direction {
	switch d {
	case right:
		if r == clockwise {
			return down
		} else if r == counterClockwise {
			return up
		} else {
			panic(r)
		}
	case down:
		if r == clockwise {
			return left
		} else if r == counterClockwise {
			return right
		} else {
			panic(r)
		}
	case left:
		if r == clockwise {
			return up
		} else if r == counterClockwise {
			return down
		} else {
			panic(r)
		}
	case up:
		if r == clockwise {
			return right
		} else if r == counterClockwise {
			return left
		} else {
			panic(r)
		}
	default:
		panic(d)
	}
}

func (d direction) asValue() int {
	switch d {
	case right:
		return 0
	case down:
		return 1
	case left:
		return 2
	case up:
		return 3
	default:
		return -1
	}
}
