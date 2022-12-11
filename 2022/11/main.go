package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	monkey struct {
		ID            int
		StartItems    []uint64
		CurItems      []uint64
		Op            operation
		TestDivisible uint64
		TestTossTrue  int
		TestTossFalse int
	}

	operation interface {
		Inspect(uint64) uint64
	}

	addOp struct {
		Double bool
		Val    uint64
	}

	multOp struct {
		Square bool
		Val    uint64
	}

	scaleFn func(uint642 uint64) uint64
)

var (
	monkeys []monkey
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	for i := range monkeys {
		monkeys[i].CurItems = append([]uint64{}, monkeys[i].StartItems...)
	}

	divideBy3 := func(in uint64) uint64 {
		return in / 3
	}

	score := getMonkeyBusiness(20, divideBy3)
	fmt.Printf("%v\n", score)
}

func part2() {
	coprimeDivisors := uint64(1)
	for i := range monkeys {
		monkeys[i].CurItems = append([]uint64{}, monkeys[i].StartItems...)
		coprimeDivisors *= monkeys[i].TestDivisible
	}

	remainder := func(in uint64) uint64 {
		return in % coprimeDivisors
	}

	score := getMonkeyBusiness(10000, remainder)
	fmt.Printf("%v\n", score)
}

func getMonkeyBusiness(numRounds int, scale scaleFn) uint64 {
	counts := make(map[int]uint64)
	for round := 1; round <= numRounds; round++ {
		for i, m := range monkeys {
			for _, item := range m.CurItems {
				counts[m.ID]++
				newVal := scale(m.Op.Inspect(item))
				if newVal%m.TestDivisible == 0 {
					monkeys[m.TestTossTrue].CurItems = append(monkeys[m.TestTossTrue].CurItems, newVal)
				} else {
					monkeys[m.TestTossFalse].CurItems = append(monkeys[m.TestTossFalse].CurItems, newVal)
				}
			}
			monkeys[i].CurItems = []uint64{}
		}
	}

	top := uint64(0)
	second := uint64(0)
	for _, v := range counts {
		if v >= top {
			second = top
			top = v
		} else if v > second {
			second = v
		}
	}

	return top * second
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	curMonkey := monkey{}
	for _, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			monkeys = append(monkeys, curMonkey)
			curMonkey = monkey{}
		}

		if strings.HasPrefix(line, "Monkey") {
			curMonkey.ID = len(monkeys)
		}

		if strings.HasPrefix(line, "Operation") {
			if add := strings.Split(line, "+"); len(add) == 2 {
				toAdd := strings.TrimSpace(add[1])
				op := addOp{}
				if toAdd == "old" {
					op.Double = true
				} else {
					val, err := strconv.Atoi(toAdd)
					if err != nil {
						panic(err)
					}
					op.Val = uint64(val)
				}
				curMonkey.Op = op
			}

			if mult := strings.Split(line, "*"); len(mult) == 2 {
				toMult := strings.TrimSpace(mult[1])
				op := multOp{}
				if toMult == "old" {
					op.Square = true
				} else {
					val, err := strconv.Atoi(toMult)
					if err != nil {
						panic(err)
					}
					op.Val = uint64(val)
				}
				curMonkey.Op = op
			}
		}

		if strings.HasPrefix(line, "Starting items:") {
			items := strings.Split(line, "Starting items:")
			if len(items) != 2 {
				panic(line)
			}

			values := strings.Split(strings.TrimSpace(items[1]), ",")
			for _, val := range values {
				v, err := strconv.Atoi(strings.TrimSpace(val))
				if err != nil {
					panic(err)
				}

				curMonkey.StartItems = append(curMonkey.StartItems, uint64(v))
			}
		}

		if strings.HasPrefix(line, "Test") {
			test := strings.Split(line, "divisible by")
			if len(test) != 2 {
				panic(line)
			}

			val, err := strconv.Atoi(strings.TrimSpace(test[1]))
			if err != nil {
				panic(err)
			}

			curMonkey.TestDivisible = uint64(val)
		}

		if strings.HasPrefix(line, "If true:") {
			throw := strings.Split(line, "to monkey")
			if len(throw) != 2 {
				panic(line)
			}

			val, err := strconv.Atoi(strings.TrimSpace(throw[1]))
			if err != nil {
				panic(err)
			}

			curMonkey.TestTossTrue = val
		}

		if strings.HasPrefix(line, "If false:") {
			throw := strings.Split(line, "to monkey")
			if len(throw) != 2 {
				panic(line)
			}

			val, err := strconv.Atoi(strings.TrimSpace(throw[1]))
			if err != nil {
				panic(err)
			}

			curMonkey.TestTossFalse = val
		}
	}
}

func (op addOp) Inspect(in uint64) uint64 {
	if op.Double {
		return 2 * in
	} else {
		return in + op.Val
	}
}

func (op multOp) Inspect(in uint64) uint64 {
	if op.Square {
		return in * in
	} else {
		return in * op.Val
	}
}
