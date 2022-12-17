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

type position struct {
	x int
	y int
}

type rock struct {
	coords [][]int
}

type chamber struct {
	coords           map[position]string
	maxHeignt        int
	curRockPositions []position
}

type coordinateMove func(pos position) position

type cycle struct {
	firstOccurrence int
	lastOccurrence  int
	length          int
}

const (
	chamberWidth      = 7
	rockStartRowAbove = 3
	rockStartColumn   = 2
)

var (
	movements []string
	rocks     = []rock{
		{
			[][]int{
				{0, 1, 2, 3}, // bottom row
			},
		},
		{
			[][]int{
				{1}, // bottom row
				{0, 1, 2},
				{1}, // top row
			},
		},
		{
			[][]int{
				{0, 1, 2}, // bottom row
				{2},
				{2}, // top row
			},
		},
		{
			[][]int{
				{0}, // bottom row
				{0},
				{0},
				{0}, // top row
			},
		},
		{
			[][]int{
				{0, 1}, // bottom row
				{0, 1}, // top row
			},
		},
	}

	moveLeft coordinateMove = func(p position) position {
		return position{p.x - 1, p.y}
	}
	moveRight coordinateMove = func(p position) position {
		return position{p.x + 1, p.y}
	}
	moveDown coordinateMove = func(p position) position {
		return position{p.x, p.y - 1}
	}
)

func part1() {
	height := simulate(2022)
	fmt.Println(height)
}

func part2() {
	cyc := detectCycle()
	heightUntilCycle := simulate(cyc.firstOccurrence)
	heightAfterCycle := simulate(cyc.lastOccurrence)
	heightAddedByCycle := heightAfterCycle - heightUntilCycle
	numCycles := (1000000000000 - cyc.firstOccurrence) / cyc.length
	remainderCycles := (1000000000000 - cyc.firstOccurrence) % cyc.length
	heightAfterRemainder := simulate(cyc.firstOccurrence + remainderCycles)
	heightAddedByRemainder := heightAfterRemainder - heightUntilCycle
	totalHeight := heightUntilCycle + heightAddedByRemainder + (heightAddedByCycle * numCycles)
	fmt.Println(totalHeight)
}

func simulate(numRocks int) int {
	ch := chamber{
		coords:    make(map[position]string),
		maxHeignt: -1,
	}

	moveNum := 0
	for rockNum := 0; rockNum < numRocks; rockNum++ {
		rck := rocks[rockNum%len(rocks)]
		ch.InitRock(rck)

		rockSettled := false
		moveByJets := true
		for !rockSettled {
			direction := "v"
			if moveByJets {
				direction = movements[moveNum%len(movements)]
				moveNum++
			}

			didMove := ch.MoveRock(direction)
			if !moveByJets && !didMove {
				ch.SettleCurrentRock()
				rockSettled = true
			}

			moveByJets = !moveByJets
		}
	}

	return ch.maxHeignt + 1
}

func detectCycle() cycle {
	ch := chamber{
		coords:    make(map[position]string),
		maxHeignt: -1,
	}

	moveNum := 0
	cycles := make(map[string]cycle)
	for rockNum := 0; true; rockNum++ {
		if key, filled := ch.areTopRowsFilled(); filled {
			rockIndex := rockNum % len(rocks)
			moveIndex := moveNum % len(movements)
			key += fmt.Sprintf("_%s_%s", rockIndex, moveIndex)
			if v, ok := cycles[key]; ok {
				v.lastOccurrence = rockNum
				v.length = v.lastOccurrence - v.firstOccurrence
				return v
			} else {
				cycles[key] = cycle{
					firstOccurrence: rockNum,
				}
			}
		}

		rck := rocks[rockNum%len(rocks)]
		ch.InitRock(rck)
		rockSettled := false
		moveByJets := true
		for !rockSettled {
			direction := "v"
			if moveByJets {
				direction = movements[moveNum%len(movements)]
				moveNum++
			}

			didMove := ch.MoveRock(direction)
			if !moveByJets && !didMove {
				ch.SettleCurrentRock()
				rockSettled = true
			}

			moveByJets = !moveByJets
		}
	}

	panic("did not find cycle")
}

func (c *chamber) InitRock(rck rock) {
	for i := 0; i < chamberWidth; i++ {
		for j := 1; j <= rockStartRowAbove; j++ {
			pos := position{i, c.maxHeignt + j}
			c.coords[pos] = "."
		}
	}

	c.curRockPositions = nil
	for rowIdx, rockRow := range rck.coords {
		for _, rockCol := range rockRow {
			pos := position{
				x: rockStartColumn + rockCol,
				y: c.maxHeignt + rockStartRowAbove + rowIdx + 1,
			}
			c.coords[pos] = "@"
			c.curRockPositions = append(c.curRockPositions, pos)
		}
	}
}

func (c *chamber) MoveRock(direction string) bool {
	//fmt.Println("Moving", direction)
	switch direction {
	case "<":
		return c.moveRock(moveLeft)
	case ">":
		return c.moveRock(moveRight)
	case "v":
		return c.moveRock(moveDown)
	default:
		panic(direction)
	}
}

func (c *chamber) moveRock(m coordinateMove) bool {
	var newPositions []position
	for _, pos := range c.curRockPositions {
		next := m(pos)
		if next.x < 0 || next.x >= chamberWidth || next.y < 0 {
			return false
		}

		if c.coords[next] == "#" {
			return false
		}

		newPositions = append(newPositions, next)
	}

	for _, pos := range c.curRockPositions {
		c.coords[pos] = "."
	}
	for _, pos := range newPositions {
		c.coords[pos] = "@"
	}
	c.curRockPositions = newPositions

	return true
}

func (c *chamber) SettleCurrentRock() {
	for _, pos := range c.curRockPositions {
		if pos.y > c.maxHeignt {
			c.maxHeignt = pos.y
		}

		c.coords[pos] = "#"
	}
	c.curRockPositions = nil
}

func (c *chamber) Print() {
	for row := c.maxHeignt + rockStartRowAbove + 4; row >= 0; row-- {
		var rowCoords []string
		for col := 0; col < chamberWidth; col++ {
			val := c.coords[position{col, row}]
			if len(val) == 0 {
				val = "."
			}

			rowCoords = append(rowCoords, val)
		}
		fmt.Println(strings.Join(rowCoords, " "))
	}
	fmt.Println("= = = = = = =")
}

func (c *chamber) areTopRowsFilled() (string, bool) {
	key := ""
	for i := 0; i < chamberWidth; i++ {
		isColFilled := false
		for j := 0; j < 4; j++ {
			val := c.coords[position{i, c.maxHeignt - j}]
			key += val
			if c.coords[position{i, c.maxHeignt - j}] == "#" {
				isColFilled = true
			}
		}

		if !isColFilled {
			return "", false
		}
	}

	return key, true
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	movements = strings.Split(string(b), "")
}
