package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type stacks map[int][]string

type move struct {
	Num  int
	From int
	To   int
}

var (
	re               = regexp.MustCompile("move ([0-9]+) from ([1-9]) to ([1-9])")
	curStacks stacks = make(map[int][]string)
	moves     []move
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	rearranged := clone(curStacks)
	for _, m := range moves {
		toMove := append([]string{}, rearranged[m.From][0:m.Num]...)
		toMove = reverse(toMove)
		newDest := append(toMove, rearranged[m.To]...)
		newSrc := append([]string{}, rearranged[m.From][m.Num:]...)
		rearranged[m.To] = append([]string{}, newDest...)
		rearranged[m.From] = append([]string{}, newSrc...)
	}

	output := ""
	for i := 1; i <= 9; i++ {
		output += rearranged[i][0]
	}

	fmt.Printf("%+v\n", output)
}

func part2() {
	rearranged := clone(curStacks)
	for _, m := range moves {
		toMove := append([]string{}, rearranged[m.From][0:m.Num]...)
		newDest := append(toMove, rearranged[m.To]...)
		newSrc := append([]string{}, rearranged[m.From][m.Num:]...)
		rearranged[m.To] = append([]string{}, newDest...)
		rearranged[m.From] = append([]string{}, newSrc...)
	}

	output := ""
	for i := 1; i <= 9; i++ {
		output += rearranged[i][0]
	}

	fmt.Printf("%+v\n", output)
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	parseMoves := false
	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			parseMoves = true
			continue
		}

		if parseMoves {
			matches := re.FindStringSubmatch(line)
			if matches == nil || len(matches) != 4 {
				panic(fmt.Sprintf("line %s did not match regexp", line))
			}

			m := move{}
			m.Num, err = strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}

			m.From, err = strconv.Atoi(matches[2])
			if err != nil {
				panic(err)
			}

			m.To, err = strconv.Atoi(matches[3])
			if err != nil {
				panic(err)
			}

			moves = append(moves, m)
		} else {
			for i := 1; i <= 9; i++ {
				// 1: 1-2
				// 2: 5-6
				// 3: 9-10
				// ...
				// 9:
				start := 4*i - 3
				obj := ""
				if len(line) > start+1 {
					obj = line[start : start+1]
				}

				if obj == fmt.Sprintf("%v", i) {
					continue
				}

				if len(strings.TrimSpace(obj)) > 0 {
					curStacks[i] = append(curStacks[i], obj)
				}
			}
		}
	}
}

func clone(toClone stacks) stacks {
	var cloned stacks = make(map[int][]string)
	for k, v := range toClone {
		cloned[k] = []string{}
		for _, obj := range v {
			cloned[k] = append(cloned[k], obj)
		}
	}
	return cloned
}

func reverse(s []string) []string {
	out := make([]string, len(s), len(s))
	n := len(s)
	for i := 0; i < n; i++ {
		out[i] = s[n-(i+1)]
	}

	return out
}

func compareLetters(s1, s2 stacks) {
	counts1 := make(map[string]int)
	for _, v := range s1 {
		for _, o := range v {
			counts1[o]++
		}
	}

	counts2 := make(map[string]int)
	for _, v := range s2 {
		for _, o := range v {
			counts2[o]++
		}
	}

	for k, v := range counts1 {
		if counts2[k] != v {
			panic("mismatched after move")
		}
	}
}
