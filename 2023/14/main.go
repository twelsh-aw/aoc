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

var initialCoords = map[coord]string{}

var (
	numRows, numCols int
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	coords := clone(initialCoords)
	rollNorth(coords)
	total := getLoad(coords)
	fmt.Println(total)
}

func part2() {
	coordsByCycleNum := map[string]int{}
	coords := clone(initialCoords)
	n := 1000000000
	remainder := 0
	for i := 1; i <= n; i++ {
		// my rolling alg is slow so this takes a bit
		// loop occurs early enough stlll so not fixing slowness
		fmt.Println("step", i)
		rollNorth(coords)
		rollWest(coords)
		rollSouth(coords)
		rollEast(coords)
		orig, ok := coordsByCycleNum[toString(coords)]
		if !ok {
			coordsByCycleNum[toString(coords)] = i
			continue
		}
		cycleLength := i - orig
		numLoops := (n - orig) / cycleLength
		remainder = n - (numLoops * cycleLength) - orig
		fmt.Println("remainder is", remainder)
		break
	}

	for i := 0; i < remainder; i++ {
		rollNorth(coords)
		rollWest(coords)
		rollSouth(coords)
		rollEast(coords)
	}

	total := getLoad(coords)
	fmt.Println(total)
}

func rollNorth(coords map[coord]string) {
	for row := 1; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			c := coord{row, col}
			if coords[c] != "O" {
				continue
			}

			dest := c
			for up := row - 1; up >= 0; up-- {
				upCoord := coord{up, col}
				if coords[upCoord] == "." {
					dest.row = up
					continue
				}
				break
			}
			if dest != c {
				coords[dest] = "O"
				coords[c] = "."
			}
		}
	}
}

func rollWest(coords map[coord]string) {
	for col := 0; col < numCols; col++ {
		for row := 0; row < numRows; row++ {
			c := coord{row, col}
			if coords[c] != "O" {
				continue
			}

			dest := c
			for left := col - 1; left >= 0; left-- {
				leftCoord := coord{row, left}
				if coords[leftCoord] == "." {
					dest.col = left
					continue
				}
				break
			}
			if dest != c {
				coords[dest] = "O"
				coords[c] = "."
			}
		}
	}
}

func rollSouth(coords map[coord]string) {
	for row := numRows - 1; row >= 0; row-- {
		for col := 0; col < numCols; col++ {
			c := coord{row, col}
			if coords[c] != "O" {
				continue
			}

			dest := c
			for down := row + 1; down < numRows; down++ {
				downCoord := coord{down, col}
				if coords[downCoord] == "." {
					dest.row = down
					continue
				}
				break
			}
			if dest != c {
				coords[dest] = "O"
				coords[c] = "."
			}
		}
	}
}

func rollEast(coords map[coord]string) {
	for col := numCols - 1; col >= 0; col-- {
		for row := 0; row < numRows; row++ {
			c := coord{row, col}
			if coords[c] != "O" {
				continue
			}

			dest := c
			for right := col + 1; col < numCols; col++ {
				rightCoord := coord{row, right}
				if coords[rightCoord] == "." {
					dest.col = right
					continue
				}
				break
			}
			if dest != c {
				coords[dest] = "O"
				coords[c] = "."
			}
		}
	}
}

func getLoad(coords map[coord]string) int {
	total := 0
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			c := coord{row, col}
			if coords[c] == "O" {
				total += numRows - row
			}
		}
	}
	return total
}

func toString(coords map[coord]string) string {
	s := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			v := coords[coord{i, j}]
			s += v
		}
		s += "\n"
	}
	return s
}

func print(coords map[coord]string) {
	fmt.Println(toString(coords))
}

func clone(coords map[coord]string) map[coord]string {
	cloned := map[coord]string{}
	for k, v := range coords {
		cloned[k] = v
	}
	return cloned
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for i, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			continue
		}
		numRows++
		parts := strings.Split(line, "")
		numCols = len(parts)
		for j, v := range parts {
			c := coord{
				row: i,
				col: j,
			}
			initialCoords[c] = v
		}
	}
}
