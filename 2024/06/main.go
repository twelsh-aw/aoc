package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/maps"
)

type (
	coord struct {
		row int
		col int
	}
	position struct {
		coord coord
		dir   string
	}
)

var (
	grid  = [][]string{}
	start = position{}
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	visited := getVisited()
	fmt.Printf("%v\n", len(visited))
}

func part2() {
	visited := getVisited()
	numLoops := 0
	for _, c := range visited {
		cur := start
		loopTracker := map[position]bool{}
		cloned := clone(grid)
		cloned[c.row][c.col] = "#"
		for {
			if loopTracker[cur] {
				numLoops++
				break
			} else {
				loopTracker[cur] = true
			}
			next := getNext(cloned, cur)
			if cur == next {
				break
			}
			cur = next
		}
	}
	fmt.Printf("%v\n", numLoops)
}

func getVisited() []coord {
	cur := start
	visited := map[coord]bool{}
	for {
		visited[cur.coord] = true
		next := getNext(grid, cur)
		if cur == next {
			break
		}
		cur = next
	}
	return maps.Keys(visited)
}

func getNext(grid [][]string, pos position) position {
	cur := pos.coord
	next := cur
	if pos.dir == "^" && cur.row != 0 {
		next = coord{cur.row - 1, cur.col}
	} else if pos.dir == "v" && cur.row != len(grid)-1 {
		next = coord{cur.row + 1, cur.col}
	} else if pos.dir == "<" && cur.col != 0 {
		next = coord{cur.row, cur.col - 1}
	} else if pos.dir == ">" && cur.col != len(grid[0])-1 {
		next = coord{cur.row, cur.col + 1}
	}

	if grid[next.row][next.col] == "#" {
		return position{
			coord: cur,
			dir:   rotate(pos.dir),
		}
	}
	return position{
		coord: next,
		dir:   pos.dir,
	}
}

func rotate(dir string) string {
	switch dir {
	case "^":
		return ">"
	case ">":
		return "v"
	case "v":
		return "<"
	case "<":
		return "^"
	default:
		panic("invalid direction")
	}
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			continue
		}
		row := []string{}
		for col, c := range strings.Split(line, "") {
			if c == "^" || c == "v" || c == "<" || c == ">" {
				start = position{
					coord: coord{len(grid), col},
					dir:   c,
				}
				row = append(row, ".")
				continue
			}
			row = append(row, c)
		}

		grid = append(grid, row)
	}
}

func clone(m [][]string) [][]string {
	c := make([][]string, len(m))
	for i := range m {
		c[i] = make([]string, len(m[i]))
		for j := range m[i] {
			v := m[i][j]
			c[i][j] = v
		}
	}

	return c
}
