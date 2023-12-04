package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	row int
	col int
}

type number struct {
	value     int
	positions []position
}

type symbol struct {
	value    string
	position position
}

var (
	numbers []number
	symbols []symbol
	maxRows int
	maxCols int
)

var digits = map[rune]bool{
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
}

var syms = map[rune]bool{
	'+': true,
	'*': true,
	'#': true,
	'$': true,
	'/': true,
	'-': true,
	'%': true,
	'=': true,
	'&': true,
	'@': true,
}

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := 0
	adjacentSymbolPositions := getAdjacentPositions(symbols)
	for _, num := range numbers {
		if _, isAdjacent := getAdjacent(num.positions, adjacentSymbolPositions); isAdjacent {
			total += num.value
		}
	}
	fmt.Printf("%v\n", total)
}

func part2() {
	gears := []symbol{}
	for _, sym := range symbols {
		if sym.value == "*" {
			gears = append(gears, sym)
		}
	}
	adjacentGearPositions := getAdjacentPositions(gears)

	gearProducts := map[symbol][]number{}
	for _, num := range numbers {
		if sym, isAdjacent := getAdjacent(num.positions, adjacentGearPositions); isAdjacent {
			gearProducts[sym] = append(gearProducts[sym], num)
		}
	}
	total := 0
	for _, numbers := range gearProducts {
		if len(numbers) != 2 {
			continue
		}
		total += numbers[0].value * numbers[1].value
	}
	fmt.Printf("%v\n", total)
}

func getAdjacentPositions(symbols []symbol) map[position]symbol {
	adjacentSymbols := map[position]symbol{}
	for _, sym := range symbols {
		adjacent := sym.position.getAdjacent()
		for _, adj := range adjacent {
			adjacentSymbols[adj] = sym
		}
	}
	return adjacentSymbols
}

func getAdjacent(positions []position, adjacent map[position]symbol) (symbol, bool) {
	for _, pos := range positions {
		if sym, ok := adjacent[pos]; ok {
			return sym, true
		}
	}
	return symbol{}, false
}

func (p *position) getAdjacent() []position {
	adj := []position{
		{
			p.row - 1,
			p.col - 1,
		},
		{
			p.row - 1,
			p.col,
		},
		{
			p.row - 1,
			p.col + 1,
		},
		{
			p.row,
			p.col - 1,
		},
		{
			p.row,
			p.col + 1,
		},
		{
			p.row + 1,
			p.col - 1,
		},
		{
			p.row + 1,
			p.col,
		},
		{
			p.row + 1,
			p.col + 1,
		},
	}
	filtered := []position{}
	for _, pos := range adj {
		if pos.row < 0 || pos.col < 0 || pos.row > maxRows || pos.col > maxCols {
			continue
		}
		filtered = append(filtered, pos)
	}
	return filtered
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(b), "\n")
	for row, line := range rows {
		if len(line) == 0 {
			continue
		}
		if row > maxRows {
			maxRows = row
		}
		num := number{}
		val := ""
		for col, char := range line {
			if col > maxCols {
				maxCols = col
			}
			pos := position{
				row: row,
				col: col,
			}
			if digits[char] {
				val += string(char)
				num.positions = append(num.positions, pos)
			} else {
				if len(val) > 0 {
					num.value, err = strconv.Atoi(val)
					if err != nil {
						panic(err)
					}
					numbers = append(numbers, num)
					num = number{}
					val = ""
				}
			}

			if _, ok := syms[char]; ok {
				sym := symbol{
					value:    string(char),
					position: pos,
				}
				symbols = append(symbols, sym)
			} else if char != '.' && !digits[char] {
				panic(char)
			}
		}
		if len(val) > 0 {
			num.value, err = strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, num)
			num = number{}
			val = ""
		}
	}
}
