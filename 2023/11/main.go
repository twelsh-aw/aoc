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
	coords           [][]string
	combos           [][2]int
	numRows, numCols int
)

var galaxies []coord

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	expanded := getExpandedGalaxies(2)
	total := 0
	for _, combo := range combos {
		g1 := expanded[combo[0]]
		g2 := expanded[combo[1]]
		dist := absDiff(g1.row, g2.row) + absDiff(g1.col, g2.col)
		total += dist
	}
	fmt.Printf("%v\n", total)
}

func part2() {
	expanded := getExpandedGalaxies(1000000)
	total := 0
	for _, combo := range combos {
		g1 := expanded[combo[0]]
		g2 := expanded[combo[1]]
		dist := absDiff(g1.row, g2.row) + absDiff(g1.col, g2.col)
		total += dist
	}
	fmt.Printf("%v\n", total)
}

func absDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff *= -1
	}
	return diff
}

func getExpandedGalaxies(factor int) []coord {
	expandRows := []int{}
	for i := 0; i < numRows; i++ {
		hasGalaxy := false
		for _, g := range galaxies {
			if g.row == i {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			expandRows = append(expandRows, i)
		}
	}

	expandCols := []int{}
	for j := 0; j < numCols; j++ {
		hasGalaxy := false
		for _, g := range galaxies {
			if g.col == j {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			expandCols = append(expandCols, j)
		}
	}

	next := append([]coord{}, galaxies...)
	for i := range expandRows {
		row := expandRows[i] + (i * (factor - 1))
		for k, g := range next {
			if g.row > row {
				next[k] = coord{g.row + factor - 1, g.col}
			}
		}
	}

	for j := range expandCols {
		col := expandCols[j] + (j * (factor - 1))
		for k, g := range next {
			if g.col > col {
				next[k] = coord{g.row, g.col + factor - 1}
			}
		}
	}

	return next
}

// func printCoords() {
// 	for i := 0; i < numRows; i++ {
// 		for j := 0; j < numCols; j++ {
// 			fmt.Printf("%s ", coords[i][j])
// 		}
// 		fmt.Print("\n")
// 	}
// }

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		numRows++
		parts := strings.Split(line, "")
		numCols = len(parts)
		coords = append(coords, make([]string, numCols))
		for j, v := range parts {
			coords[i][j] = v
			if v == "#" {
				galaxies = append(galaxies, coord{i, j})
			}
		}
	}

	for i := range galaxies {
		for j := range galaxies {
			if i < j {
				combos = append(combos, [2]int{i, j})
			}
		}
	}
}
