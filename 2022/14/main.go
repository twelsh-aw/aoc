package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

var (
	startCave = make(map[position]string)
	minX      = math.MaxInt32
	maxX      int
	maxY      int
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	cave := clone(startCave)
	hitEdge := false
	numSand := 0
	for !hitEdge {
		cur := position{500, 0}
		var next position
		sandSettled := false
		for !sandSettled {
			next, sandSettled = getNextPosition(cur, cave)
			if !sandSettled && (next.x < minX || next.x > maxX || next.y > maxY) {
				hitEdge = true
				break
			}

			if !sandSettled {
				cur = next
			}
		}

		cave[cur] = "o"
		if !hitEdge {
			numSand++
		}
	}

	fmt.Printf("%v\n", numSand)
}

func part2() {
	cave := clone(startCave)
	startBlocked := false
	floor := maxY + 2
	numSand := 0
	for !startBlocked {
		cur := position{500, 0}
		var next position
		sandSettled := false
		for !sandSettled {
			if cur.y == floor-1 {
				sandSettled = true
				break
			}

			next, sandSettled = getNextPosition(cur, cave)
			if !sandSettled {
				cur = next
			}
		}

		cave[cur] = "o"
		numSand++
		startBlocked = cur.x == 500 && cur.y == 0
	}

	fmt.Printf("%v\n", numSand)
}

func getNextPosition(cur position, cave map[position]string) (position, bool) {
	down := position{cur.x, cur.y + 1}
	if _, ok := cave[down]; !ok {
		return down, false
	}

	downLeft := position{cur.x - 1, cur.y + 1}
	if _, ok := cave[downLeft]; !ok {
		return downLeft, false
	}

	downRight := position{cur.x + 1, cur.y + 1}
	if _, ok := cave[downRight]; !ok {
		return downRight, false
	}

	return position{}, true
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		parts := strings.Split(line, " -> ")
		curPos := position{}
		for _, part := range parts {
			nextPos := position{}
			p := strings.Split(part, ",")
			if len(p) != 2 {
				panic(part)
			}

			nextPos.x, err = strconv.Atoi(p[0])
			if err != nil {
				panic(err)
			}

			nextPos.y, err = strconv.Atoi(p[1])
			if err != nil {
				panic(err)
			}

			if nextPos.x < minX {
				minX = nextPos.x
			}

			if nextPos.x > maxX {
				maxX = nextPos.x
			}

			if nextPos.y > maxY {
				maxY = nextPos.y
			}

			startCave[nextPos] = "$"
			if curPos.x == 0 && curPos.y == 0 {
				curPos = nextPos
				continue
			}

			xDiff := nextPos.x - curPos.x
			yDiff := nextPos.y - curPos.y
			if xDiff != 0 && yDiff != 0 {
				panic("unexpected diffs")
			}

			for i := curPos.x; i != nextPos.x; i += sign(xDiff) {
				midPos := position{i, curPos.y}
				startCave[midPos] = "$"
			}

			for i := curPos.y; i != nextPos.y; i += sign(yDiff) {
				midPos := position{curPos.x, i}
				startCave[midPos] = "$"
			}

			curPos = nextPos
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

func sign(i int) int {
	if i > 0 {
		return 1
	}

	if i < 0 {
		return -1
	}

	return 0
}
