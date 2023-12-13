package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var records = []record{}

type record struct {
	springs []string
	groups  []int
}

type recordAttempt struct {
	curSpringIndex int
	numGroupsSoFar int
	curGroupLength int
}

func main() {
	readInput()
	p1 := part1(records)
	fmt.Printf("%v\n", p1)
	p2 := part2()
	fmt.Printf("%v\n", p2)
}

func part1(records []record) int64 {
	total := int64(0)
	for _, r := range records {
		initial := recordAttempt{
			curSpringIndex: 0,
			curGroupLength: 0,
			numGroupsSoFar: 0,
		}
		seenMap := map[recordAttempt]int64{
			initial: 1,
		}
		r.springs = append(r.springs, ".")
		for springIndex := 0; springIndex < len(r.springs); springIndex++ {
			for groupsSoFar := 0; groupsSoFar <= len(r.groups); groupsSoFar++ {
				for groupLength := 0; groupLength <= len(r.springs); groupLength++ {
					attempt := recordAttempt{
						curSpringIndex: springIndex,
						curGroupLength: groupLength,
						numGroupsSoFar: groupsSoFar,
					}
					seen, ok := seenMap[attempt]
					if !ok {
						continue
					}
					char := r.springs[springIndex]
					if char == "." || char == "?" {
						if groupLength == 0 || (groupsSoFar > 0 && groupLength == r.groups[groupsSoFar-1]) {
							next := recordAttempt{
								curSpringIndex: springIndex + 1,
								curGroupLength: 0,
								numGroupsSoFar: groupsSoFar,
							}
							seenMap[next] += seen
						}
					}
					if char == "#" || char == "?" {
						nextGroupsSoFar := groupsSoFar
						if groupLength == 0 {
							nextGroupsSoFar++
						}
						next := recordAttempt{
							curSpringIndex: springIndex + 1,
							curGroupLength: groupLength + 1,
							numGroupsSoFar: nextGroupsSoFar,
						}
						seenMap[next] += seen
					}
				}
			}
		}

		final := recordAttempt{
			curSpringIndex: len(r.springs),
			numGroupsSoFar: len(r.groups),
			curGroupLength: 0,
		}
		total += seenMap[final]
	}

	return total
}

func part2() int64 {
	folded := []record{}
	for _, r := range records {
		next := record{}
		for i := 0; i < 5; i++ {
			next.springs = append(next.springs, r.springs...)
			next.groups = append(next.groups, r.groups...)
			if i != 4 {
				next.springs = append(next.springs, "?")
			}
		}
		folded = append(folded, next)
	}

	return part1(folded)
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
		line = strings.TrimSpace(line)
		r := record{}
		parts := strings.Split(line, " ")
		groupParts := strings.Split(parts[1], ",")
		for i := range groupParts {
			p, err := strconv.Atoi(groupParts[i])
			if err != nil {
				panic(err)
			}
			r.groups = append(r.groups, p)
		}

		springParts := strings.Split(parts[0], "")
		r.springs = append(r.springs, springParts...)
		records = append(records, r)
	}
}
