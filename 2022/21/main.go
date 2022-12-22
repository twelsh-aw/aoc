package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operation struct {
	monkeyA string
	monkeyB string
	operand operand
}

type value struct {
	int
	coefficient float64
	unknown     bool
}

type result struct {
	value *value
	op    *operation
}

type operand string

const (
	plus   operand = "+"
	minus  operand = "-"
	divide operand = "/"
	mult   operand = "*"
	equals operand = "="
)

var (
	origMonkeys = make(map[string]*result)
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	monkeys := clone(origMonkeys)
	resloveRootMonkey(monkeys)
	fmt.Printf("%v\n", monkeys["root"].value.int)
}

func part2() {
	monkeys := clone(origMonkeys)
	monkeys["root"].op.operand = equals
	monkeys["humn"].value.unknown = true
	monkeys["humn"].value.coefficient = 1
	monkeys["humn"].value.int = 0
	resloveRootMonkey(monkeys)

	num := float64(monkeys["root"].value.int) / monkeys["root"].value.coefficient

	fmt.Printf("%v\n", int(num))
}

func resloveRootMonkey(monkeys map[string]*result) {
	monkeysToCheck := []string{"root"}
	for monkeys["root"].value == nil {
		numOrig := len(monkeysToCheck)
		alreadyAdded := make(map[string]bool)
		for _, name := range monkeysToCheck {
			monkey := monkeys[name]
			if monkey.value != nil {
				continue
			}
			if monkey.op == nil {
				panic("no op and no value")
			}

			monkeyA := monkeys[monkey.op.monkeyA]
			monkeyB := monkeys[monkey.op.monkeyB]
			if monkeyA.value == nil && !alreadyAdded[monkey.op.monkeyA] {
				alreadyAdded[monkey.op.monkeyA] = true
				monkeysToCheck = append(monkeysToCheck, monkey.op.monkeyA)
			}

			if monkeyB.value == nil && !alreadyAdded[monkey.op.monkeyB] {
				alreadyAdded[monkey.op.monkeyB] = true
				monkeysToCheck = append(monkeysToCheck, monkey.op.monkeyB)
			}

			if monkeyA.value != nil && monkeyB.value != nil {
				r := calculate(monkeyA.value, monkeyB.value, monkey.op.operand)
				monkey.value = &r
			} else if !alreadyAdded[name] {
				alreadyAdded[name] = true
				monkeysToCheck = append(monkeysToCheck, name)
			}
		}

		numAdded := len(monkeysToCheck) - numOrig
		if numAdded == 0 {
			break
		}

		monkeysToCheck = monkeysToCheck[numOrig:]
	}
}

func calculate(valueA, valueB *value, operand operand) value {
	if valueA.unknown && valueB.unknown {
		panic("didn't account for this")
	}

	r := 0
	c := float64(1)
	unknown := valueA.unknown || valueB.unknown
	switch operand {
	case plus:
		r = valueA.int + valueB.int
		if valueA.unknown {
			c = valueA.coefficient
		} else if valueB.unknown {
			c = valueB.coefficient
		}

	case minus:
		r = valueA.int - valueB.int
		if valueA.unknown {
			c = valueA.coefficient
		} else if valueB.unknown {
			c = -1 * valueB.coefficient
		}
	case mult:
		r = valueA.int * valueB.int
		if valueA.unknown {
			c = valueA.coefficient * float64(valueB.int)
		} else if valueB.unknown {
			c = valueB.coefficient * float64(valueA.int)
		}
	case divide:
		r = valueA.int / valueB.int
		if valueA.unknown {
			c = valueA.coefficient / float64(valueB.int)
		} else if valueB.unknown {
			c = valueB.coefficient / float64(valueA.int)
		}
	case equals:
		if !valueA.unknown && !valueB.unknown {
			panic("no unknowns to solve")
		}

		if valueA.unknown {
			r = valueB.int - valueA.int
			c = valueA.coefficient
		} else if valueB.unknown {
			r = valueA.int - valueB.int
			c = valueB.coefficient
		}
	default:
		panic(operand)
	}

	return value{int: r, coefficient: c, unknown: unknown}
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			panic(line)
		}

		name := parts[0]
		rawResult := parts[1]
		res := result{}
		if opParts := strings.Split(rawResult, " + "); len(opParts) == 2 {
			op := &operation{
				monkeyA: opParts[0],
				monkeyB: opParts[1],
				operand: plus,
			}
			res.op = op
			origMonkeys[name] = &res
			continue
		}

		if opParts := strings.Split(rawResult, " - "); len(opParts) == 2 {
			op := &operation{
				monkeyA: opParts[0],
				monkeyB: opParts[1],
				operand: minus,
			}
			res.op = op
			origMonkeys[name] = &res
			continue
		}

		if opParts := strings.Split(rawResult, " / "); len(opParts) == 2 {
			op := &operation{
				monkeyA: opParts[0],
				monkeyB: opParts[1],
				operand: divide,
			}
			res.op = op
			origMonkeys[name] = &res
			continue
		}

		if opParts := strings.Split(rawResult, " * "); len(opParts) == 2 {
			op := &operation{
				monkeyA: opParts[0],
				monkeyB: opParts[1],
				operand: mult,
			}
			res.op = op
			origMonkeys[name] = &res
			continue
		}

		number, err := strconv.Atoi(rawResult)
		if err != nil {
			panic(err)
		}

		res.value = &value{int: number}
		origMonkeys[name] = &res
	}
}

func clone(m map[string]*result) map[string]*result {
	c := make(map[string]*result)
	for k, v := range m {
		c[k] = v.clone()
	}

	return c
}

func (r *result) clone() *result {
	cloned := &result{}
	if r.value != nil {
		v := *r.value
		cloned.value = &v
	}

	if r.op != nil {
		v := *r.op
		cloned.op = &v
	}

	return cloned
}
