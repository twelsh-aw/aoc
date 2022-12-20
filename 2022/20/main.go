package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type number struct {
	value     int
	origIndex int
}

var (
	numbers    []number
	indexJumps = make(map[number]int)
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	mixed := make([]number, len(numbers))
	for i := range numbers {
		mixed[i] = numbers[i]
	}

	mixed = mixNumbers(numbers, mixed)
	total := getScore(mixed)

	fmt.Printf("%v\n", total)
}

func part2() {
	keyedNumbers := make([]number, len(numbers))
	mixed := make([]number, len(numbers))
	for i := range numbers {
		keyedNumbers[i] = numbers[i]
		keyedNumbers[i].value *= 811589153
		mixed[i] = keyedNumbers[i]
	}

	for i := 0; i < 10; i++ {
		mixed = mixNumbers(keyedNumbers, mixed)
	}

	total := getScore(mixed)
	fmt.Printf("%v\n", total)
}

func mixNumbers(original, mixed []number) []number {
	for _, num := range original {
		var next []number
		curIndex := -1
		for i := range mixed {
			if mixed[i] == num {
				curIndex = i
				continue
			}
			next = append(next, mixed[i])
		}

		if curIndex == -1 {
			panic("could not find num")
		}

		newIndex := getNewIndex(curIndex, num, len(mixed))
		mixed = append(append(append([]number{}, next[0:newIndex]...), num), next[newIndex:]...)
		if len(mixed) != len(numbers) {
			panic("length bad")
		}
	}

	return mixed
}

func getNewIndex(curIndex int, num number, length int) int {
	if num.value == 0 {
		return curIndex
	}

	if (num.value > 0 && curIndex+num.value < length) || (num.value < 0 && curIndex+num.value > 0) {
		return curIndex + num.value
	}

	next := (curIndex + num.value) % (length - 1)
	if next <= 0 {
		next += length - 1
	}
	return next
}

func getScore(mixed []number) int {
	total := 0
	curIndex := 0
	for i := range mixed {
		if mixed[i].value == 0 {
			curIndex = i
		}
	}

	for i := 1; i <= 3000; i++ {
		curIndex = (curIndex + 1) % len(mixed)
		if i%1000 == 0 {
			total += mixed[curIndex].value
		}
	}

	return total
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for i, line := range strings.Split(string(b), "\n") {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, number{
			value:     n,
			origIndex: i,
		})
	}
}
