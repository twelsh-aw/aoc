package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type card struct {
	number  int
	winning map[int]bool
	values  []int
}

var cards []card

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := 0
	for _, c := range cards {
		total += c.score()
	}
	fmt.Printf("%v\n", total)
}

func part2() {
	total := 0
	matchesByCardNumber := map[int]int{}
	cardNumbers := []int{}
	for _, c := range cards {
		matchesByCardNumber[c.number] = c.matches()
		cardNumbers = append(cardNumbers, c.number)
	}

	for len(cardNumbers) > 0 {
		nextNumbers := []int{}
		for _, c := range cardNumbers {
			matches := matchesByCardNumber[c]
			total++
			for i := 1; i <= matches; i++ {
				nextNumbers = append(nextNumbers, c+i)
			}
		}
		cardNumbers = append([]int{}, nextNumbers...)
	}

	fmt.Printf("%v\n", total)
}

func (c *card) score() int {
	score := 0
	for _, v := range c.values {
		if c.winning[v] {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}

	return score
}

func (c *card) matches() int {
	matches := 0
	for _, v := range c.values {
		if c.winning[v] {
			matches++
		}
	}

	return matches
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

		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			panic(line)
		}

		num, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(parts[0], "Card")))
		if err != nil {
			panic(parts[0])
		}

		c := card{
			number:  num,
			winning: map[int]bool{},
		}

		numbers := strings.Split(parts[1], " | ")
		if len(numbers) != 2 {
			panic(parts[1])
		}

		winning := strings.Split(numbers[0], " ")
		for _, w := range winning {
			if len(w) == 0 {
				continue
			}
			v, err := strconv.Atoi(w)
			if err != nil {
				panic(winning)
			}
			c.winning[v] = true
		}

		ticket := strings.Split(numbers[1], " ")
		for _, t := range ticket {
			if len(t) == 0 {
				continue
			}
			v, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			c.values = append(c.values, v)
		}

		cards = append(cards, c)
	}
}
