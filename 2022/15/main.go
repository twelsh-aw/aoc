package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type sensor struct {
	pos           position
	nearestBeacon position
	distance      int
}

type position struct {
	x int
	y int
}

var (
	sensorBeaconCoords = make(map[position]string)
	sensors            []sensor
	minX               = math.MaxInt32
	maxX               int
	maxDistance        int
	posRegexp          = regexp.MustCompile("x=([-\\d]+), y=([-\\d]+)")
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	noBeaconCount := 0
	yPos := 2000000
	for xPos := minX - maxDistance; xPos <= maxX+maxDistance; xPos++ {
		p := position{xPos, yPos}
		if _, ok := sensorBeaconCoords[p]; ok {
			continue
		}

		for _, s := range sensors {
			if l1Distance(p, s.pos) <= s.distance {
				noBeaconCount++
				break
			}
		}
	}
	fmt.Printf("%v\n", noBeaconCount)
}

func part2() {
	var distressBeacon *position
	var pointsOnBoundary []position
	maxCoord := 4000000
	for _, s := range sensors {
		sensorBoundary := s.getBoundary(maxCoord)
		pointsOnBoundary = append(pointsOnBoundary, sensorBoundary...)
	}

	for _, p := range pointsOnBoundary {
		if _, ok := sensorBeaconCoords[p]; ok {
			continue
		}

		var isNonBeacon = false
		for _, s := range sensors {
			if l1Distance(p, s.pos) <= s.distance {
				isNonBeacon = true
				break
			}
		}

		if !isNonBeacon {
			distressBeacon = &p
			break
		}
	}

	if distressBeacon == nil {
		panic("no distress beacon found")
	}

	fmt.Printf("%v\n", (distressBeacon.x*4000000)+distressBeacon.y)
}

func (s *sensor) getBoundary(maxCoord int) []position {
	var positions []position
	distance := l1Distance(s.pos, s.nearestBeacon)
	for i := 0; i <= distance+1; i++ {
		xPos := s.pos.x + i
		xNeg := s.pos.x - i
		yPos := s.pos.y + (distance + 1 - i)
		yNeg := s.pos.y - (distance + 1 - i)
		if xPos >= 0 && xPos <= maxCoord {
			if yPos >= 0 && yPos <= maxCoord {
				positions = append(positions, position{xPos, yPos})
			}

			if yNeg != yPos && yNeg >= 0 && yNeg <= maxCoord {
				positions = append(positions, position{xPos, yNeg})
			}
		} else if xNeg != xPos && xNeg >= 0 && xNeg <= maxCoord {
			if yPos >= 0 && yPos <= maxCoord {
				positions = append(positions, position{xNeg, yPos})
			}

			if yNeg != yPos && yNeg >= 0 && yNeg <= maxCoord {
				positions = append(positions, position{xNeg, yNeg})
			}
		}
	}

	return positions
}

func l1Distance(p1, p2 position) int {
	return absDiff(p2.y, p1.y) + absDiff(p2.x, p1.x)
}

func absDiff(x, y int) int {
	if x-y < 0 {
		return y - x
	}

	return x - y
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			panic(line)
		}

		sensorPos := parsePosition(parts[0])
		beaconPos := parsePosition(parts[1])
		sensorBeaconCoords[sensorPos] = "S"
		sensorBeaconCoords[beaconPos] = "B"
		distance := l1Distance(sensorPos, beaconPos)

		s := sensor{
			pos:           sensorPos,
			nearestBeacon: beaconPos,
			distance:      distance,
		}
		sensors = append(sensors, s)

		if sensorPos.x < minX {
			minX = sensorPos.x
		}
		if beaconPos.x < minX {
			minX = beaconPos.x
		}
		if sensorPos.x > maxX {
			maxX = sensorPos.x
		}
		if beaconPos.x > maxX {
			maxX = beaconPos.x
		}
		if s.distance > maxDistance {
			maxDistance = s.distance
		}
	}
}

func parsePosition(part string) position {
	posParts := posRegexp.FindStringSubmatch(part)
	if posParts == nil || len(posParts) != 3 {
		panic(part)
	}

	pos := position{}
	var err error
	pos.x, err = strconv.Atoi(posParts[1])
	if err != nil {
		panic(err)
	}
	pos.y, err = strconv.Atoi(posParts[2])
	if err != nil {
		panic(err)
	}

	return pos
}
