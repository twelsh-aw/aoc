package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var equations = []equation{}

type equation struct {
	result int
	nums   []int
	val    int
}

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := 0
	for _, eq := range equations {
		if testEquation(eq, false) {
			total += eq.result
		}
	}
	fmt.Printf("%v\n", total)
}

func part2() {
	total := 0
	for _, eq := range equations {
		if testEquation(eq, true) {
			total += eq.result
		}
	}
	fmt.Printf("%v\n", total)
}

func testEquation(eq equation, allowConcat bool) bool {
	tests := []equation{eq}
	for len(tests) > 0 {
		n := len(tests)
		for _, test := range tests {
			if len(test.nums) == 0 {
				if test.val == test.result {
					return true
				}
				continue
			}

			product := test.val * test.nums[0]
			sum := test.val + test.nums[0]
			concat, _ := strconv.Atoi(fmt.Sprintf("%v%v", test.val, test.nums[0]))

			if product != 0 && product <= test.result {
				tests = append(tests, equation{
					result: test.result,
					nums:   splice(test.nums, 0),
					val:    product,
				})
			}

			if sum != 0 && sum <= test.result {
				tests = append(tests, equation{
					result: test.result,
					nums:   splice(test.nums, 0),
					val:    sum,
				})
			}

			if allowConcat && concat != sum && concat <= test.result {
				tests = append(tests, equation{
					result: test.result,
					nums:   splice(test.nums, 0),
					val:    concat,
				})
			}
		}
		tests = tests[n:]
	}

	return false
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

		eq := equation{}
		eq.result, _ = strconv.Atoi(parts[0])

		nums := strings.Split(parts[1], " ")
		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			eq.nums = append(eq.nums, n)
		}

		equations = append(equations, eq)
	}
}

func splice[T any](a []T, i int) []T {
	if i >= len(a) || i < 0 {
		return a
	}
	ret := []T{}
	ret = append(ret, a[:i]...)
	ret = append(ret, a[i+1:]...)
	return ret
}
