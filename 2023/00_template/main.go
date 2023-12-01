package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	fmt.Printf("%v\n", nil)
}

func part2() {
	fmt.Printf("%v\n", nil)
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		fmt.Println(line)
	}
}
