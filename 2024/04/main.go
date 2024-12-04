package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	input  [][]string
	numRow int
	numCol int
)

type coord struct {
	row int
	col int
}

type word struct {
	text   string
	coords []coord
}

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := 0
	for row := 0; row < numRow; row++ {
		for col := 0; col < numCol; col++ {
			if input[row][col] == "X" {
				words := getWords(row, col, 4)
				for _, word := range words {
					if word == "XMAS" {
						total++
					}
				}
			}
		}
	}
	fmt.Printf("%v\n", total)
}

func part2() {
	centerMAS := map[coord]int{}
	for row := 0; row < numRow; row++ {
		for col := 0; col < numCol; col++ {
			if input[row][col] == "M" {
				words := getDiagWords(row, col, 3)
				for _, word := range words {
					if word.text == "MAS" {
						middle := word.coords[1]
						centerMAS[middle]++
					}
				}
			}
		}
	}

	total := 0
	for _, count := range centerMAS {
		if count == 2 {
			total++
		}
		if count > 2 {
			panic("too many middle coords counted")
		}
	}
	fmt.Printf("%v\n", total)
}

func getWords(row, col int, length int) []string {
	words := []string{}
	if col >= length-1 {
		left := ""
		for j := 0; j < length; j++ {
			left += input[row][col-j]
		}
		words = append(words, left)
	}
	if col < numCol-(length-1) {
		right := ""
		for j := 0; j < length; j++ {
			right += input[row][col+j]
		}
		words = append(words, right)
	}
	if row >= length-1 {
		up := ""
		for j := 0; j < length; j++ {
			up += input[row-j][col]
		}
		words = append(words, up)
	}
	if row < numRow-(length-1) {
		down := ""
		for j := 0; j < length; j++ {
			down += input[row+j][col]
		}
		words = append(words, down)
	}
	diagWords := getDiagWords(row, col, length)
	for _, d := range diagWords {
		words = append(words, d.text)
	}

	return words
}

func getDiagWords(row, col int, length int) []word {
	words := []word{}
	if col >= length-1 && row >= length-1 {
		upLeft := ""
		coords := []coord{}
		for j := 0; j < length; j++ {
			upLeft += input[row-j][col-j]
			coords = append(coords, coord{row - j, col - j})
		}
		words = append(words, word{upLeft, coords})
	}
	if col < numCol-(length-1) && row >= (length-1) {
		upRight := ""
		coords := []coord{}
		for j := 0; j < length; j++ {
			upRight += input[row-j][col+j]
			coords = append(coords, coord{row - j, col + j})
		}
		words = append(words, word{upRight, coords})
	}
	if col >= length-1 && row < numRow-(length-1) {
		downLeft := ""
		coords := []coord{}
		for j := 0; j < length; j++ {
			downLeft += input[row+j][col-j]
			coords = append(coords, coord{row + j, col - j})
		}
		words = append(words, word{downLeft, coords})
	}
	if col < numCol-(length-1) && row < numRow-(length-1) {
		downRight := ""
		coords := []coord{}
		for j := 0; j < length; j++ {
			downRight += input[row+j][col+j]
			coords = append(coords, coord{row + j, col + j})
		}
		words = append(words, word{downRight, coords})
	}

	return words
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
		row := strings.Split(line, "")
		input = append(input, row)
	}

	numRow = len(input)
	numCol = len(input[0])
}
