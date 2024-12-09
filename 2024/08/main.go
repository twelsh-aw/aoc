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
	antennas = map[string][]coord{}
	numRow   int
	numCol   int
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	antinodes := map[coord]bool{}
	for _, coords := range antennas {
		for i := 0; i < len(coords); i++ {
			for j := 0; j < len(coords); j++ {
				if i == j {
					continue
				}
				vertical := coords[j].row - coords[i].row
				horizontal := coords[j].col - coords[i].col
				next := coord{
					row: coords[j].row + vertical,
					col: coords[j].col + horizontal,
				}

				if next.row >= 0 && next.row < numRow && next.col >= 0 && next.col < numCol {
					antinodes[next] = true
				}
			}
		}
	}
	fmt.Printf("%v\n", len(antinodes))
}

func part2() {
	antinodes := map[coord]bool{}
	for _, coords := range antennas {
		for i := 0; i < len(coords); i++ {
			for j := 0; j < len(coords); j++ {
				if i == j {
					continue
				}
				k := -1
				for {
					k++
					vertical := coords[j].row - coords[i].row
					horizontal := coords[j].col - coords[i].col
					next := coord{
						row: coords[j].row + (k * vertical),
						col: coords[j].col + (k * horizontal),
					}

					if next.row >= 0 && next.row < numRow && next.col >= 0 && next.col < numCol {
						antinodes[next] = true
					} else {
						break
					}
				}
			}
		}
	}
	fmt.Printf("%v\n", len(antinodes))
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
		numRow++
		for col, char := range strings.Split(line, "") {
			if row == 0 {
				numCol++
			}
			if char == "." {
				continue
			}
			antennas[char] = append(antennas[char], coord{row, col})
		}
	}
}
