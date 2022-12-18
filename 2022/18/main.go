package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
	z int
}

var (
	input            []position
	minX, minY, minZ = math.MaxInt, math.MaxInt, math.MaxInt
	maxX, maxY, maxZ = 0, 0, 0
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	droplet := make(map[position]bool)
	for _, pos := range input {
		droplet[pos] = true
	}

	surfaceArea := 0
	for pos := range droplet {
		faces := getAdjacent(pos)
		for _, face := range faces {
			if !droplet[face] {
				surfaceArea++
			}
		}
	}
	fmt.Printf("%v\n", surfaceArea)
}

func part2() {
	droplet := make(map[position]bool)
	for _, pos := range input {
		droplet[pos] = true
	}

	surfaceArea := 0
	for pos := range droplet {
		faces := getAdjacent(pos)
		for _, face := range faces {
			if !droplet[face] {
				if isExterior(face, droplet) {
					surfaceArea++
				}
			}
		}
	}
	fmt.Printf("%v\n", surfaceArea)
}

func getAdjacent(pos position) [6]position {
	return [6]position{
		{
			pos.x + 1,
			pos.y,
			pos.z,
		},
		{
			pos.x - 1,
			pos.y,
			pos.z,
		},
		{
			pos.x,
			pos.y + 1,
			pos.z,
		},
		{
			pos.x,
			pos.y - 1,
			pos.z,
		},
		{
			pos.x,
			pos.y,
			pos.z + 1,
		},
		{
			pos.x,
			pos.y,
			pos.z - 1,
		},
	}
}

func isExterior(face position, droplet map[position]bool) bool {
	toCheck := []position{face}
	exteriorSoFar := make(map[position]bool)
	for {
		numOriginal := len(toCheck)
		for _, pos := range toCheck {
			adj := getAdjacent(pos)
			for _, next := range adj {
				if droplet[next] {
					continue
				}
				if exteriorSoFar[next] {
					continue
				}
				exteriorSoFar[next] = true
				if next.x > maxX || next.x < minX || next.y > maxY || next.y < minY || next.z > maxZ || next.z < minZ {
					return true
				}
				toCheck = append(toCheck, next)
			}
		}
		numAdded := len(toCheck) - numOriginal
		if numAdded == 0 {
			return false
		}
		toCheck = toCheck[numOriginal:]
	}
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		coords := strings.Split(line, ",")
		if len(coords) != 3 {
			panic(line)
		}

		pos := position{}
		pos.x, err = strconv.Atoi(coords[0])
		if err != nil {
			panic(err)
		}

		pos.y, err = strconv.Atoi(coords[1])
		if err != nil {
			panic(err)
		}

		pos.z, err = strconv.Atoi(coords[2])
		if err != nil {
			panic(err)
		}

		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.x < minX {
			minX = pos.x
		}
		if pos.y > maxY {
			maxY = pos.y
		}
		if pos.y < minY {
			minY = pos.y
		}
		if pos.z > maxZ {
			maxZ = pos.z
		}
		if pos.z < minZ {
			minZ = pos.z
		}

		input = append(input, pos)
	}
}
