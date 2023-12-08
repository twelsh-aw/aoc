package main

import (
	"fmt"
	"os"
	"strings"
)

type node struct {
	left  string
	right string
}

var instructions []string
var locations map[string]node
var starts []string

func main() {
	readInput()
	fmt.Println(part1("AAA", "ZZZ"))
	part2()
}

func part1(cur string, end string) int {
	i := 0
	n := 0
	for {
		n++
		idx := i % len(instructions)
		dir := instructions[idx]
		if dir == "L" {
			cur = locations[cur].left
		} else if dir == "R" {
			cur = locations[cur].right
		}
		if end == "Z" && cur[2] == 'Z' {
			break
		} else if cur == end {
			break
		}
		i++
	}
	return n
}

func part2() {
	cur := append([]string{}, starts...)
	periods := []int{}
	for _, pos := range cur {
		period := part1(pos, "Z")
		periods = append(periods, period)
	}
	lcm := LCM(periods[0], periods[1], periods[2:]...)
	fmt.Printf("%v\n", lcm)
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	locations = map[string]node{}
	for i, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			continue
		}

		if i == 0 {
			instructions = strings.Split(line, "")
			continue
		}

		parts := strings.Split(line, " = ")
		trim := strings.TrimRight(strings.TrimLeft(parts[1], "("), ")")
		nodes := strings.Split(trim, ", ")
		if strings.HasSuffix(parts[0], "A") {
			starts = append(starts, parts[0])
		}
		locations[parts[0]] = node{
			left:  nodes[0],
			right: nodes[1],
		}
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
