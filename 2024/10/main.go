package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	grid    = [][]int{}
	starts  []coord
	numRows int
	numCols int
)

type coord struct {
	row int
	col int
}

type path struct {
	start coord
	end   coord
}

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := 0
	pathsSeen := map[path]bool{}
	curPaths := []path{}
	for _, start := range starts {
		curPaths = append(curPaths, path{start, start})
	}

	for len(curPaths) > 0 {
		nextPaths := []path{}
		for _, path := range curPaths {
			if pathsSeen[path] {
				continue
			}
			pathsSeen[path] = true
			if grid[path.end.row][path.end.col] == 9 {
				total++
			}
			next := getNext(path)
			for _, n := range next {
				nextPaths = append(nextPaths, n)
			}
		}
		curPaths = nextPaths
	}

	fmt.Printf("%v\n", total)
}

func part2() {
	total := 0
	curPaths := []path{}
	for _, start := range starts {
		curPaths = append(curPaths, path{start, start})
	}

	for len(curPaths) > 0 {
		nextPaths := []path{}
		for _, path := range curPaths {
			if grid[path.end.row][path.end.col] == 9 {
				total++
			}
			next := getNext(path)
			for _, n := range next {
				nextPaths = append(nextPaths, n)
			}
		}
		curPaths = nextPaths
	}

	fmt.Printf("%v\n", total)
}

func getNext(p path) []path {
	pos := p.end
	value := grid[pos.row][pos.col]

	left := coord{pos.row, pos.col - 1}
	right := coord{pos.row, pos.col + 1}
	up := coord{pos.row - 1, pos.col}
	down := coord{pos.row + 1, pos.col}

	next := []path{}
	if left.col >= 0 && grid[left.row][left.col] == value+1 {
		next = append(next, path{p.start, left})
	}
	if right.col < numCols && grid[right.row][right.col] == value+1 {
		next = append(next, path{p.start, right})
	}
	if up.row >= 0 && grid[up.row][up.col] == value+1 {
		next = append(next, path{p.start, up})
	}
	if down.row < numRows && grid[down.row][down.col] == value+1 {
		next = append(next, path{p.start, down})
	}
	return next
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

		nums := strings.Split(line, "")
		grid = append(grid, make([]int, len(nums)))
		for col, num := range nums {
			n, _ := strconv.Atoi(num)
			if n == 0 {
				starts = append(starts, coord{row, col})
			}
			grid[row][col] = n
		}
	}

	numRows = len(grid)
	numCols = len(grid[0])
}
