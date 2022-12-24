package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type position struct {
	row int
	col int
}

type blizzardPath struct {
	cur   position
	steps int
	route int
	dest  position
}

type direction string

const (
	left  direction = "<"
	right direction = ">"
	down  direction = "v"
	up    direction = "^"
)

var (
	startPos, destPos                    position
	originalBlizz                        = make(map[int]position)
	blizzardDirections                   = make(map[int]direction)
	walls                                = make(map[position]bool)
	numBlizzards, maxWallCol, maxWallRow int
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	minSteps := navigate(1)
	fmt.Printf("%v\n", minSteps)
}

func part2() {
	minSteps := navigate(3)
	fmt.Printf("%v\n", minSteps)
}

func navigate(numRoutes int) int {
	pathsToCheck := []blizzardPath{
		{
			cur:   startPos,
			steps: 0,
			route: 1,
			dest:  destPos,
		},
	}

	minSteps := 0
	closestDistanceSoFarByRoute := make(map[int]int)
	for i := 1; i <= numRoutes; i++ {
		closestDistanceSoFarByRoute[i] = math.MaxInt
	}

	blizzard := clone(originalBlizz)
	occupied := make(map[position]bool)
	visited := make(map[string]bool)

	for {
		numOrigPaths := len(pathsToCheck)
		blizzard, occupied = moveBlizard(blizzard)
		blizzKey := getKey(blizzard)
		for _, path := range pathsToCheck {
			curDistance := l1Distance(path.cur, path.dest)
			if curDistance == 0 {
				if path.route == numRoutes {
					minSteps = path.steps
					break
				}

				path.route++
				if path.route%2 == 0 {
					path.dest = startPos
				} else {
					path.dest = destPos
				}

				curDistance = l1Distance(path.cur, path.dest)
			}

			key := fmt.Sprintf("%v_%v;%v_%v", path.route, path.cur.row, path.cur.col, blizzKey)
			if visited[key] {
				continue
			}

			visited[key] = true

			closestDistanceSoFar := closestDistanceSoFarByRoute[path.route]
			if curDistance < closestDistanceSoFar {
				closestDistanceSoFarByRoute[path.route] = curDistance
			} else if curDistance > 3*closestDistanceSoFar { // get greedy to keep visited map size
				continue
			}

			next := getNextPositions(path.cur, occupied)
			for _, n := range next {
				nextPath := blizzardPath{
					cur:   n,
					steps: path.steps + 1,
					dest:  path.dest,
					route: path.route,
				}
				pathsToCheck = append(pathsToCheck, nextPath)
			}
		}

		if minSteps > 0 {
			break
		}

		numAddedPaths := len(pathsToCheck) - numOrigPaths
		if numAddedPaths == 0 {
			panic("no paths found")
		}

		pathsToCheck = pathsToCheck[numOrigPaths:]
	}

	return minSteps
}

func getNextPositions(cur position, occupied map[position]bool) []position {
	var next []position
	if !occupied[cur] {
		next = append(next, cur)
	}

	pLeft := position{cur.row, cur.col - 1}
	if !walls[pLeft] && !occupied[pLeft] {
		next = append(next, pLeft)
	}

	pRight := position{cur.row, cur.col + 1}
	if !walls[pRight] && !occupied[pRight] {
		next = append(next, pRight)
	}

	pUp := position{cur.row - 1, cur.col}
	if cur != startPos && !walls[pUp] && !occupied[pUp] {
		next = append(next, pUp)
	}

	pDown := position{cur.row + 1, cur.col}
	if cur != destPos && !walls[pDown] && !occupied[pDown] {
		next = append(next, pDown)
	}

	return next
}

func moveBlizard(blizzard map[int]position) (map[int]position, map[position]bool) {
	occupied := make(map[position]bool)
	nextBlizzard := make(map[int]position, len(blizzard))
	for blizzID, p := range blizzard {
		d, ok := blizzardDirections[blizzID]
		if !ok {
			panic(blizzID)
		}

		next := position{p.row, p.col}
		switch d {
		case left:
			next.col--
			if walls[next] {
				next.col = maxWallCol - 1
			}
		case right:
			next.col++
			if walls[next] {
				next.col = 1
			}
		case down:
			next.row++
			if walls[next] {
				if p.col == startPos.col {
					next.row = startPos.row
				} else {
					next.row = 1
				}
			}
		case up:
			next.row--
			if walls[next] {
				if p.col == destPos.col {
					next.row = destPos.row
				} else {
					next.row = maxWallRow - 1
				}
			}
		default:
			panic(d)
		}

		if walls[next] {
			panic(next)
		}

		nextBlizzard[blizzID] = next
		occupied[next] = true
	}

	return nextBlizzard, occupied
}

func getKey(blizz map[int]position) string {
	var parts []string
	for id := 1; id <= numBlizzards; id++ {
		parts = append(parts, fmt.Sprintf("%v;%v", blizz[id].row, blizz[id].col))
	}

	return strings.Join(parts, "_")
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	for row, line := range lines {
		for col, v := range strings.Split(line, "") {
			p := position{row, col}
			if v == "#" {
				walls[p] = true
				if col > maxWallCol {
					maxWallCol = col
				}
				if row > maxWallRow {
					maxWallRow = row
				}
				continue
			}
			if v == "." {
				if row == 0 {
					startPos = p
				}
				if row == len(lines)-1 {
					destPos = p
				}
				continue
			}
			d := direction(v)
			if d != left && d != right && d != up && d != down {
				panic(d)
			}
			blizzID := numBlizzards + 1
			originalBlizz[blizzID] = p
			blizzardDirections[blizzID] = d
			numBlizzards++
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

func l1Distance(p1, p2 position) int {
	return absDiff(p2.row, p1.row) + absDiff(p2.col, p1.col)
}

func absDiff(x, y int) int {
	if x-y < 0 {
		return y - x
	}

	return x - y
}
