package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type brick struct {
	start coord
	end   coord
}

type coord struct {
	x int
	y int
	z int
}

var (
	bricks = []brick{}
	filled = map[coord]bool{}
)

func main() {
	part1()
	part2()
}

func part1() {
	settleBricks()
	total := 0
	for _, b := range bricks {
		// disintegrate brick
		cloneFilled := cloneFilled()
		for _, c := range b.getCoords() {
			cloneFilled[c] = false
		}
		isSettled := true
		for i := range bricks {
			_, canMove := move(bricks[i], cloneFilled)
			if canMove {
				isSettled = false
				break
			}
		}
		if isSettled { // can disintegrate
			total++
		}
	}
	fmt.Printf("%v\n", total)
}

func part2() {
	settleBricks()
	numMoved := 0
	for _, b := range bricks {
		bricksMoved := map[int]bool{}
		// disintegrate brick
		cloneFilled := cloneFilled()
		for _, c := range b.getCoords() {
			cloneFilled[c] = false
		}
		cloneBricks := []brick{}
		for _, cb := range bricks {
			if b != cb {
				cloneBricks = append(cloneBricks, cb)
			}
		}
		for {
			isSettled := true
			for i := range cloneBricks {
				next, canMove := move(cloneBricks[i], cloneFilled)
				if canMove {
					if !bricksMoved[i] {
						numMoved++
					}
					bricksMoved[i] = true
					isSettled = false
					for _, c := range cloneBricks[i].getCoords() {
						delete(cloneFilled, c)
					}
					for _, c := range next.getCoords() {
						cloneFilled[c] = true
					}
					cloneBricks[i] = next
				}
			}
			if isSettled {
				break
			}
		}
	}

	fmt.Printf("%v\n", numMoved)
}

func settleBricks() {
	readInput()
	for {
		isSettled := true
		for i := range bricks {
			next, canMove := move(bricks[i], filled)
			if canMove {
				isSettled = false
				for _, c := range bricks[i].getCoords() {
					delete(filled, c)
				}
				for _, c := range next.getCoords() {
					filled[c] = true
				}
				bricks[i] = next
			}
		}
		if isSettled {
			break
		}
	}
}

func cloneFilled() map[coord]bool {
	cloned := map[coord]bool{}
	for k, v := range filled {
		cloned[k] = v
	}
	return cloned
}

func move(cur brick, filled map[coord]bool) (brick, bool) {
	if cur.end.z == 1 {
		return cur, false
	}
	curCoordsMap := map[coord]bool{}
	for _, c := range cur.getCoords() {
		curCoordsMap[c] = true
	}
	canMove := true
	moved := brick{
		start: cur.start,
		end:   cur.end,
	}
	moved.start.z--
	moved.end.z--
	for _, c := range moved.getCoords() {
		if filled[c] && !curCoordsMap[c] {
			canMove = false
			break
		}
	}

	return moved, canMove
}

func readInput() {
	bricks = []brick{}
	filled = map[coord]bool{}
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, "~")
		if len(parts) != 2 {
			panic(line)
		}
		b := brick{
			start: makeCoord(parts[0]),
			end:   makeCoord(parts[1]),
		}
		// ensure end has lower z
		if b.start.z < b.end.z {
			s := b.start
			b.start = b.end
			b.end = s
		}
		bricks = append(bricks, b)
	}

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].end.z < bricks[j].end.z
	})

	for _, b := range bricks {
		for _, c := range b.getCoords() {
			filled[c] = true
		}
	}
}

func makeCoord(csv string) coord {
	coord := coord{}
	for i, c := range strings.Split(csv, ",") {
		v, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		switch i {
		case 0:
			coord.x = v
		case 1:
			coord.y = v
		case 2:
			coord.z = v
		default:
			panic(csv)
		}
	}
	return coord
}

func sign(start, end coord) coord {
	c := coord{}
	if end.x-start.x >= 0 {
		c.x = 1
	} else {
		c.x = -1
	}
	if end.y-start.y >= 0 {
		c.y = 1
	} else {
		c.y = -1
	}
	if end.z-start.z >= 0 {
		c.z = 1
	} else {
		c.z = -1
	}
	return c
}

func (b *brick) getCoords() []coord {
	coords := []coord{}
	sign := sign(b.start, b.end)
	for x := b.start.x; true; x += sign.x {
		for y := b.start.y; true; y += sign.y {
			for z := b.start.z; true; z += sign.z {
				c := coord{x, y, z}
				coords = append(coords, c)
				if z == b.end.z {
					break
				}
			}
			if y == b.end.y {
				break
			}
		}
		if x == b.end.x {
			break
		}
	}
	return coords
}
