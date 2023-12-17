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

type crucible struct {
	pos              coord
	dir              direction
	consecutiveSteps int
}

type direction string

const (
	dirRight = "r"
	dirLeft  = "l"
	dirUp    = "u"
	dirDown  = "d"
)

var (
	coords           = map[coord]int{}
	numRows, numCols int
)

func main() {
	readInput()
	part1(1, 3)
	part2()
}

func part1(minSteps, maxSteps int) {
	start := crucible{
		coord{0, 0},
		dirRight,
		0,
	}
	end := coord{numRows - 1, numCols - 1}
	dist := map[crucible]int{
		start: 0,
	}
	visited := map[crucible]bool{}
	possibleNodes := numRows * numCols * 4 * (maxSteps - minSteps + 1)
	for {
		var cur crucible
		minDist := math.MaxInt
		canContinue := false
		for k, v := range dist {
			if visited[k] {
				continue
			}
			if v < minDist {
				minDist = v
				cur = k
				canContinue = true
			}
		}
		if !canContinue {
			break
		}
		visited[cur] = true
		remainingNodes := possibleNodes - len(visited)
		if remainingNodes%10000 == 0 {
			fmt.Println("remaining nodes", remainingNodes)
		}
		if cur.pos == end { // stop at end
			break
		}
		adj := cur.getNextPositions(visited, minSteps, maxSteps)
		for _, n := range adj {
			prevCost, exists := dist[n]
			if !exists {
				prevCost = math.MaxInt
			}

			cost := dist[cur] + moveCost(cur.pos, n.pos)
			if cost < prevCost {
				dist[n] = cost
			}
		}
		delete(dist, cur)
	}
	minScore := math.MaxInt
	for k, v := range dist {
		if k.pos == end {
			if v < minScore {
				minScore = v
			}
		}
	}
	fmt.Printf("%v\n", minScore)
}

func part2() {
	part1(4, 10)
}

func (c *crucible) getNextPositions(visited map[crucible]bool, minSteps, maxSteps int) []crucible {
	next := []crucible{}
	if c.consecutiveSteps == 0 {
		forward, isValid := c.pos.move(c.dir, minSteps)
		ncf := crucible{forward, c.dir, minSteps}
		if isValid && !visited[ncf] {
			next = append(next, ncf)
		}
	}

	if c.consecutiveSteps >= minSteps && c.consecutiveSteps < maxSteps {
		forward, isValid := c.pos.move(c.dir, 1)
		ncf := crucible{forward, c.dir, c.consecutiveSteps + 1}
		if isValid && !visited[ncf] {
			next = append(next, ncf)
		}
	}

	var leftDir, rightDir direction
	switch c.dir {
	case dirUp:
		leftDir = dirLeft
		rightDir = dirRight
	case dirLeft:
		leftDir = dirDown
		rightDir = dirUp
	case dirDown:
		leftDir = dirRight
		rightDir = dirLeft
	case dirRight:
		leftDir = dirUp
		rightDir = dirDown
	default:
		panic(c.dir)
	}

	left, isValid := c.pos.move(leftDir, minSteps)
	ncl := crucible{left, leftDir, minSteps}
	if isValid && !visited[ncl] {
		next = append(next, ncl)
	}

	right, isValid := c.pos.move(rightDir, minSteps)
	ncr := crucible{right, rightDir, minSteps}
	if isValid && !visited[ncr] {
		next = append(next, ncr)
	}

	return next
}

func moveCost(from, to coord) int {
	if from == to {
		panic("bad move")
	}
	cost := 0
	if from.row == to.row {
		sign := 1
		if from.col > to.col {
			sign = -1
		}

		next := coord{from.row, from.col + sign}
		for {
			cost += coords[next]
			if next == to {
				break
			}
			next.col += sign
		}
	} else if from.col == to.col {
		sign := 1
		if from.row > to.row {
			sign = -1
		}

		next := coord{from.row + sign, from.col}
		for {
			cost += coords[next]
			if next == to {
				break
			}
			next.row += sign
		}
	} else {
		panic("bad move")
	}

	return cost
}

func (c *coord) move(dir direction, steps int) (next coord, valid bool) {
	next = *c
	switch dir {
	case dirRight:
		next.col += steps
	case dirLeft:
		next.col -= steps
	case dirDown:
		next.row += steps
	case dirUp:
		next.row -= steps
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
		for col, val := range parts {
			i, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			coords[coord{row, col}] = i
		}
	}
}
