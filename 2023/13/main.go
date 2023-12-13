package main

import (
	"fmt"
	"os"
	"strings"
)

var grids []grid

type grid struct {
	coords  [][]string
	numRows int
	numCols int
}

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := 0
	for _, grid := range grids {
		total += grid.getReflectionScore(-1)
	}
	fmt.Printf("%v\n", total)
}

func part2() {
	total := 0
	for _, grid := range grids {
		origScore := grid.getReflectionScore(-1)
		newScore := 0
	smudge:
		for row := 0; row < grid.numRows; row++ {
			for col := 0; col < grid.numCols; col++ {
				smudged := grid.clone()
				smudged.coords[row][col] = smudge(smudged.coords[row][col])
				score := smudged.getReflectionScore(origScore)
				if score > 0 && score != origScore {
					newScore = score
					break smudge
				}
			}
		}
		if newScore == 0 {
			panic("no smudge")
		}
		total += newScore
	}
	fmt.Printf("%v\n", total)
}

func (g *grid) print() {
	for row := 0; row < g.numRows; row++ {
		fmt.Println(strings.Join(g.coords[row], " "))
	}
}

func (g *grid) getReflectionScore(continueOnScore int) int {
	rr := g.getRowReflectionScore(continueOnScore)
	if rr > 0 {
		return rr
	} else {
		transposed := g.transpose()
		return transposed.getRowReflectionScore(continueOnScore*100) / 100
	}
}

func (g *grid) getRowReflectionScore(continueOnScore int) int {
	for row := 1; row < g.numRows; row++ {
		rowsAbove := g.coords[0:row]
		rowsBelow := g.coords[row:]

		minLength := len(rowsAbove)
		if len(rowsBelow) < minLength {
			minLength = len(rowsBelow)
		}

		isMirror := true
	checkMirror:
		for i := 0; i < minLength; i++ {
			rowAbove := rowsAbove[len(rowsAbove)-i-1]
			rowBelow := rowsBelow[i]
			for j := 0; j < g.numCols; j++ {
				if rowAbove[j] != rowBelow[j] {
					isMirror = false
					break checkMirror
				}
			}
		}

		if isMirror {
			score := len(rowsAbove) * 100
			if score != continueOnScore {
				return score
			}
		}
	}

	return 0
}

func (g *grid) transpose() grid {
	transposed := grid{
		numRows: g.numCols,
		numCols: g.numRows,
	}
	coords := [][]string{}
	for j := 0; j < g.numCols; j++ {
		row := []string{}
		for i := 0; i < g.numRows; i++ {
			row = append(row, g.coords[i][j])
		}
		coords = append(coords, row)
	}
	transposed.coords = coords
	return transposed
}

func (g *grid) clone() grid {
	c := grid{
		numRows: g.numRows,
		numCols: g.numCols,
	}
	for _, row := range g.coords {
		next := append([]string{}, row...)
		c.coords = append(c.coords, next)
	}
	return c
}

func smudge(char string) string {
	if char == "." {
		return "#"
	} else if char == "#" {
		return "."
	} else {
		panic(char)
	}
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	var curGrid grid
	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			grids = append(grids, curGrid)
			curGrid = grid{}
			continue
		}
		curGrid.numRows++
		curGrid.numCols = len(line)
		curGrid.coords = append(curGrid.coords, strings.Split(line, ""))
	}
}
