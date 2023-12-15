package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputs = []string{}

type box struct {
	lenses []lens
	labels map[string]bool
}

type lens struct {
	label  string
	length int
}

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := 0
	for _, s := range inputs {
		score := hash(s)
		total += score
	}
	fmt.Printf("%v\n", total)
}

func part2() {
	boxes := [256]box{}
	for i := range boxes {
		boxes[i].labels = map[string]bool{}
	}
	for _, s := range inputs {
		if strings.HasSuffix(s, "-") {
			label := s[:len(s)-1]
			boxNum := hash(label)
			if !boxes[boxNum].labels[label] {
				continue
			}
			next := []lens{}
			for _, lens := range boxes[boxNum].lenses {
				if lens.label != label {
					next = append(next, lens)
				}
			}
			boxes[boxNum].lenses = next
			delete(boxes[boxNum].labels, label)
			continue
		}

		parts := strings.Split(s, "=")
		if len(parts) != 2 {
			panic(parts)
		}
		label := parts[0]
		boxNum := hash(label)
		length, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		lens := lens{
			label:  label,
			length: length,
		}
		if !boxes[boxNum].labels[label] {
			boxes[boxNum].lenses = append(boxes[boxNum].lenses, lens)
			boxes[boxNum].labels[label] = true
			continue
		}

		for i, lens := range boxes[boxNum].lenses {
			if lens.label == label {
				boxes[boxNum].lenses[i].length = length
				break
			}
		}
	}

	total := 0
	for boxNum, box := range boxes {
		for lensNum, lens := range box.lenses {
			total += (boxNum + 1) * (lensNum + 1) * lens.length
		}
	}
	fmt.Printf("%v\n", total)
}

func hash(s string) int {
	score := 0
	for _, char := range s {
		score += int(char)
		score *= 17
		score %= 256
	}
	return score
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
		inputs = strings.Split(line, ",")
	}
}
