package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	id    int
	draws []draw
}

type draw struct {
	red   int
	blue  int
	green int
}

type colour string

const (
	red   colour = "red"
	blue  colour = "blue"
	green colour = "green"
)

var games []game

var extractDigitRE = regexp.MustCompile(`([0-9]+) .*`)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := 0
	for _, g := range games {
		valid := true
		for _, d := range g.draws {
			if d.blue > 14 || d.green > 13 || d.red > 12 {
				valid = false
				break
			}
		}
		if valid {
			total += g.id
		}
	}
	fmt.Printf("%v\n", total)
}

func part2() {
	total := 0
	for _, g := range games {
		maxRed := 0
		maxBlue := 0
		maxGreen := 0
		for _, d := range g.draws {
			if d.blue > maxBlue {
				maxBlue = d.blue
			}
			if d.green > maxGreen {
				maxGreen = d.green
			}
			if d.red > maxRed {
				maxRed = d.red
			}
		}
		product := maxBlue * maxGreen * maxRed
		total += product
	}
	fmt.Printf("%v\n", total)
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

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			panic("bad input")
		}

		sid := parts[0][5:]
		id, err := strconv.Atoi(sid)
		if err != nil {
			panic(err)
		}

		g := game{
			id: id,
		}

		draws := strings.Split(parts[1], ";")
		for _, ds := range draws {
			d := draw{}
			colours := strings.Split(ds, ", ")
			for _, c := range colours {
				switch {
				case strings.Contains(c, string(green)):
					d.green = parseColour(c)
				case strings.Contains(c, string(blue)):
					d.blue = parseColour(c)
				case strings.Contains(c, string(red)):
					d.red = parseColour(c)
				}
			}
			g.draws = append(g.draws, d)
		}

		games = append(games, g)
	}
}

func parseColour(c string) int {
	matches := extractDigitRE.FindStringSubmatch(c)
	if len(matches) != 2 {
		panic(c)
	}
	s, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	return s
}
