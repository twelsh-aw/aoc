package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

type instruction struct {
	dir    direction
	num    int
	colour string
}

type edge struct {
	coords [2]coord
	ins    instruction
}

type direction string

const (
	dirRight = "R"
	dirLeft  = "L"
	dirUp    = "U"
	dirDown  = "D"
)

var (
	instructions = []instruction{}
)

func main() {
	readInput()
	part1(instructions)
	part2()
}

func part1(instructions []instruction) {
	edges := []edge{}
	cur, next := coord{}, coord{}
	for _, i := range instructions {
		switch i.dir {
		case dirRight:
			next.x += i.num
		case dirLeft:
			next.x -= i.num
		case dirDown:
			next.y -= i.num
		case dirUp:
			next.y += i.num
		}
		edges = append(edges, edge{[2]coord{cur, next}, i})
		cur = next
	}
	if cur.x != 0 || cur.y != 0 {
		panic("not a loop")
	}

	// https://en.wikipedia.org/wiki/Shoelace_formula#Trapezoid_formula
	// we measure the areas of trapezoids: (ax, 0),(bx, 0),(ax, ay),(bx, by) from points oriented counter-clockwise
	// in our case our trapezoids are just rectangles which makes this easy.
	// here we map 2-D coord (x,y) to represent the middle of the 3-D coord (x,y,z) (x, y, 0.5)
	interior := 1 // initial dig
	for i := len(edges) - 1; i >= 0; i-- {
		edge := edges[i]
		a := edge.coords[1]
		b := edge.coords[0]
		if edge.ins.dir == dirUp || edge.ins.dir == dirDown { // width 0 rectangle
			continue
		}
		height := a.y
		width := edge.ins.num
		area := height * width
		if a.x < b.x {
			interior -= area
		} else {
			interior += area
		}
	}

	// the above interior works by only counting "half" of the boundary (0.5m^2 per coord on boundary)
	boundary := 0
	for _, i := range instructions {
		boundary += i.num
	}
	fmt.Println(interior + boundary/2)
}

func part2() {
	newInstructions := []instruction{}
	for _, ins := range instructions {
		switch ins.colour[7] {
		case '0':
			ins.dir = dirRight
		case '1':
			ins.dir = dirDown
		case '2':
			ins.dir = dirLeft
		case '3':
			ins.dir = dirUp
		default:
			panic(string(ins.colour[7]))
		}

		n, err := strconv.ParseInt(ins.colour[2:7], 16, 0)
		if err != nil {
			panic(err)
		}
		ins.num = int(n)
		newInstructions = append(newInstructions, ins)
	}
	part1(newInstructions)
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
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			panic(line)
		}
		i := instruction{}
		i.dir = direction(parts[0])
		i.num, err = strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		i.colour = parts[2]
		instructions = append(instructions, i)
	}
}
