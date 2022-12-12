package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct {
	row int
	col int
}

type path struct {
	pos        position
	stepsSoFar int
}

var (
	coords [][]int
	start  position
	end    position
	nRow   int
	nCol   int
)

var letters = map[string]int{
	"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26,
	"S": 1, "E": 26,
}

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	toVisit := []path{{start, 0}}
	visited := make(map[position]bool)
	visited[start] = true

	ms := getMinStepsUntilEnd(toVisit, visited)
	fmt.Println(ms)
}

func part2() {
	var toVisit []path
	visited := make(map[position]bool)
	for row := range coords {
		for col := range coords[row] {
			if coords[row][col] == 1 {
				pos := position{row, col}
				toVisit = append(toVisit, path{pos: pos, stepsSoFar: 0})
				visited[pos] = true
			}
		}
	}

	ms := getMinStepsUntilEnd(toVisit, visited)
	fmt.Println(ms)
}

func getMinStepsUntilEnd(toVisit []path, visited map[position]bool) int {
	for {
		orig := len(toVisit)
		for _, p := range toVisit {
			adj := getAdjacent(p.pos)
			for _, next := range adj {
				if next == end {
					return 1 + p.stepsSoFar
				}

				if visited[next] {
					continue
				}

				visited[next] = true
				toVisit = append(toVisit, path{next, 1 + p.stepsSoFar})
			}
		}

		after := len(toVisit)
		if orig == after {
			fmt.Println("no such paths")
			break
		}

		toVisit = toVisit[orig:]
	}

	return -1
}

func getAdjacent(cur position) []position {
	var positions []position
	elevation := coords[cur.row][cur.col]
	if cur.row > 0 {
		up := position{row: cur.row - 1, col: cur.col}
		if coords[up.row][up.col] <= elevation+1 {
			positions = append(positions, up)
		}
	}

	if cur.row < nRow-1 {
		down := position{row: cur.row + 1, col: cur.col}
		if coords[down.row][down.col] <= elevation+1 {
			positions = append(positions, down)
		}
	}

	if cur.col > 0 {
		left := position{row: cur.row, col: cur.col - 1}
		if coords[left.row][left.col] <= elevation+1 {
			positions = append(positions, left)
		}
	}

	if cur.col < nCol-1 {
		right := position{row: cur.row, col: cur.col + 1}
		if coords[right.row][right.col] <= elevation+1 {
			positions = append(positions, right)
		}
	}

	return positions
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		if i := strings.Index(line, "S"); i >= 0 {
			start.row = len(coords)
			start.col = i
		}

		if i := strings.Index(line, "E"); i >= 0 {
			end.row = len(coords)
			end.col = i
		}

		var row []int
		for _, letter := range strings.Split(line, "") {
			row = append(row, letters[letter])
		}
		coords = append(coords, row)
	}

	nRow = len(coords)
	nCol = len(coords[0])
}
