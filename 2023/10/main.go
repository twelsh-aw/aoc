package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type direction string

const (
	north = "N"
	south = "S"
	east  = "E"
	west  = "W"
)

type coord struct {
	row int
	col int
}

type pipeBoundary struct {
	coord      coord
	adjFaceDir direction
}

var coords = map[coord]string{}

var pathLoop = map[coord]int{}

var start coord

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	initialPipes := []string{
		"-", "|", "L", "J", "7", "F",
	}
	for _, val := range initialPipes {
		curCoords := []coord{start}
		curDist := 0
		coords[start] = val
		path := map[coord]int{
			start: curDist,
		}
		pathValid := true
	outer:
		for {
			if len(curCoords) == 0 {
				break
			}
			curDist++
			nextCoords := []coord{}
			for _, coord := range curCoords {
				adj := getAdjacentCoordsForPipe(coord, coords[coord])
				for _, next := range adj {
					if !hasConnectingPipe(next) {
						pathValid = false
						break outer
					}
					if _, alreadySeen := path[next.coord]; !alreadySeen {
						nextCoords = append(nextCoords, next.coord)
						path[next.coord] = curDist
					}
				}
			}

			curCoords = append([]coord{}, nextCoords...)
		}

		if pathValid {
			fmt.Println(curDist - 1)
			pathLoop = path
			return
		}
	}
}

func part2() {
	numInterior := 0
	for _, c := range orderedCoords() {
		if _, isOnLoop := pathLoop[c]; isOnLoop {
			continue
		}
		loopIntersectionsEast := 0
		curCoord := c
		// JCT: rays from points in interior of boundary intersect boundary an odd number of times
		// fix path going east from "south part" of coord
		for {
			curCoord = getNextCoord(curCoord, east)
			pipe, inGrid := coords[curCoord]
			if !inGrid {
				break
			}
			_, inLoop := pathLoop[curCoord]
			if inLoop {
				if pipe == "|" || pipe == "7" || pipe == "F" {
					// only pipes that run south will result in an intersection for the chosen ray point
					loopIntersectionsEast++
				}
			}
		}

		isInterior := loopIntersectionsEast%2 != 0
		if isInterior {
			numInterior++
		}
	}
	fmt.Printf("%v\n", numInterior)
}

func orderedCoords() []coord {
	ordered := []coord{}
	for k := range coords {
		ordered = append(ordered, k)
	}

	sort.Slice(ordered, func(i, j int) bool {
		if ordered[i].row < ordered[j].row {
			return true
		} else if ordered[i].row > ordered[j].row {
			return false
		}

		return ordered[i].col < ordered[j].col
	})

	return ordered
}

func getAdjacentCoordsForPipe(c coord, val string) [2]pipeBoundary {
	// 	| is a vertical pipe connecting north and south.
	// - is a horizontal pipe connecting east and west.
	// L is a 90-degree bend connecting north and east.
	// J is a 90-degree bend connecting north and west.
	// 7 is a 90-degree bend connecting south and west.
	// F is a 90-degree bend connecting south and east.
	switch val {
	case "-":
		return [2]pipeBoundary{
			{
				// go west
				coord{c.row, c.col - 1},
				east,
			},
			{
				// go east
				coord{c.row, c.col + 1},
				west,
			},
		}
	case "|":
		return [2]pipeBoundary{
			{
				// go north
				coord{c.row - 1, c.col},
				south,
			},
			{
				// go south
				coord{c.row + 1, c.col},
				north,
			},
		}
	case "L":
		return [2]pipeBoundary{
			{
				// go east
				coord{c.row, c.col + 1},
				west,
			},
			{
				// go north
				coord{c.row - 1, c.col},
				south,
			},
		}
	case "J":
		return [2]pipeBoundary{
			{
				// go west
				coord{c.row, c.col - 1},
				east,
			},
			{
				// go north
				coord{c.row - 1, c.col},
				south,
			},
		}
	case "7":
		return [2]pipeBoundary{
			{
				// go west
				coord{c.row, c.col - 1},
				east,
			},
			{
				// go south
				coord{c.row + 1, c.col},
				north,
			},
		}
	case "F":
		return [2]pipeBoundary{
			{
				// go east
				coord{c.row, c.col + 1},
				west,
			},
			{
				// go south
				coord{c.row + 1, c.col},
				north,
			},
		}
	default:
		panic(val)
	}
}

func hasConnectingPipe(adj pipeBoundary) bool {
	nextVal, ok := coords[adj.coord]
	if !ok {
		return false
	}
	if nextVal == "." {
		return false
	}
	if nextVal == "S" {
		return true
	}
	switch adj.adjFaceDir {
	case north:
		return nextVal == "L" || nextVal == "J" || nextVal == "|"
	case south:
		return nextVal == "7" || nextVal == "F" || nextVal == "|"
	case west:
		return nextVal == "J" || nextVal == "7" || nextVal == "-"
	case east:
		return nextVal == "L" || nextVal == "F" || nextVal == "-"
	default:
		panic(adj.adjFaceDir)
	}
}

func getNextCoord(c coord, d direction) coord {
	switch d {
	case north:
		return coord{c.row - 1, c.col}
	case south:
		return coord{c.row + 1, c.col}
	case east:
		return coord{c.row, c.col + 1}
	case west:
		return coord{c.row, c.col - 1}
	default:
		panic(d)
	}
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for i, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			continue
		}
		for j, val := range strings.Split(line, "") {
			c := coord{i, j}
			coords[c] = val
			if val == "S" {
				start = c
			}
		}
	}
}
