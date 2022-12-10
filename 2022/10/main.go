package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	signals []int
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	curRegister := 1
	signalStrengths := make(map[int]int)
	adds := make(map[int]int)
	for cycle := 1; cycle <= len(signals); cycle++ {
		signalStrengths[cycle] = curRegister * cycle
		toAdd := signals[cycle-1]
		adds[cycle+1] = toAdd
		if addNow, ok := adds[cycle]; ok {
			curRegister += addNow
		}
	}

	strength := 0
	for _, cycle := range []int{20, 60, 100, 140, 180, 220} {
		strength += signalStrengths[cycle]
	}

	fmt.Printf("%v\n", strength)
}

func part2() {
	curRegister := 1 // controls middle position of sprite
	adds := make(map[int]int)

	var crt [][]string
	crt = make([][]string, 6)
	for i := range crt {
		crt[i] = make([]string, 40)
	}

	for cycle := 1; cycle <= len(signals); cycle++ {
		curIndex := cycle - 1
		toAdd := signals[curIndex]
		adds[cycle+1] = toAdd

		row := curIndex / 40
		col := curIndex % 40
		if cycle > 240 {
			continue
		}

		if curRegister-1 <= col && col <= curRegister+1 {
			crt[row][col] = "#"
		} else {
			crt[row][col] = "."
		}

		if addNow, ok := adds[cycle]; ok {
			curRegister += addNow
		}
	}

	for _, data := range crt {
		fmt.Println(strings.Join(data, " "))
	}
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		if strings.HasPrefix(line, "addx") {
			s, err := strconv.Atoi(strings.Split(line, " ")[1])
			if err != nil {
				panic(err)
			}

			signals = append(signals, s)
		}

		signals = append(signals, 0)
	}
}
