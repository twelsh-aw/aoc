package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type rule struct {
	cond condition
	dest string
}

type condition struct {
	variable string
	operator string
	value    int64
}

type part struct {
	x int64
	m int64
	a int64
	s int64
}

var (
	workflows  = map[string][]rule{}
	alwaysTrue = condition{operator: "true"}
	parts      = []part{}
	partsRe    = regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := int64(0)
	for _, part := range parts {
		if isPartApproved(part) {
			total += part.x + part.m + part.a + part.s
		}
	}
	fmt.Printf("%v\n", total)
}

func part2() {
	// we don't need to test all 1 to 4000 in each dimensions
	// only need to test combinations of values "between splits" of conditions
	type splitPoint struct {
		val      int64
		op       string
		variable string
	}
	// testPoint represents an interval (inclusive) in a single variable in (x,m,a,s) where all evaluations of conditions are the same.
	// we track it's length and choose the start of the interval to be the representative point
	type testPoint struct {
		start  int64
		length int64
	}
	splits := map[string][]splitPoint{}
	testPoints := map[string][]testPoint{}
	seen := map[splitPoint]bool{}
	for _, wf := range workflows {
		for _, rule := range wf {
			if rule.cond == alwaysTrue {
				continue
			}
			variable := rule.cond.variable
			splitPoint := splitPoint{
				val:      rule.cond.value,
				op:       rule.cond.operator,
				variable: rule.cond.variable,
			}
			// ignore duplicate splits
			if !seen[splitPoint] {
				splits[variable] = append(splits[variable], splitPoint)
			}
			seen[splitPoint] = true
		}
	}
	for _, k := range []string{"x", "m", "a", "s"} {
		// we ensure ties "100 <" and "100 >" sort the "<" first
		// so that it makes intervals [X-99, 100-100, 101-Y]
		sort.Slice(splits[k], func(i, j int) bool {
			a := splits[k][i]
			b := splits[k][j]
			aVal, bVal := a.val, b.val
			if aVal < bVal {
				return true
			} else if aVal > bVal {
				return false
			} else if a.op == "<" && b.op == ">" {
				return true
			}

			return a.op == "<"
		})
		cur := int64(1)
		for _, s := range splits[k] {
			var next int64
			if s.op == "<" {
				// next and below is true
				// next+1 and above is false
				next = s.val - 1
			} else if s.op == ">" {
				// next and below is false
				// next + 1 and above is true
				next = s.val
			} else {
				panic(s)
			}
			tp := testPoint{
				start:  cur,
				length: next - cur + 1,
			}
			// splits like: ["880 >", "881 <"] become intervals -> [X-880, 881-880, 881-Y]
			// we can safely drop the middle invalid interval
			if next == cur-1 {
				continue
			} else if next < cur {
				panic("unexpected interval sequence")
			}
			testPoints[k] = append(testPoints[k], tp)
			cur = next + 1
		}
		testPoints[k] = append(testPoints[k], testPoint{cur, 4000 - cur + 1})
	}

	approved := &atomic.Int64{}
	progress := &atomic.Int64{}
	total := int64(len(testPoints["x"])) * int64(len(testPoints["m"])) * int64(len(testPoints["a"])) * int64(len(testPoints["s"]))
	wg := sync.WaitGroup{}
	for _, x := range testPoints["x"] {
		for _, m := range testPoints["m"] {
			wg.Add(1) // goroutines across two dimensions to speed up
			go func(x, m testPoint) {
				for _, a := range testPoints["a"] {
					for _, s := range testPoints["s"] {
						combo := part{
							x: x.start,
							m: m.start,
							a: a.start,
							s: s.start,
						}
						if isPartApproved(combo) {
							approved.Add(x.length * m.length * a.length * s.length)
						}
						if x := progress.Add(1); x%100000000 == 0 {
							fmt.Printf("%.2f%%\n", 100*float64(x)/float64(total))
						}
					}
				}
				wg.Done()
			}(x, m)
		}
	}
	wg.Wait()
	fmt.Printf("%v\n", approved.Load())
}

func isPartApproved(part part) bool {
	cur := "in"
	for {
		next := apply(part, cur)
		if next == "A" {
			return true
		} else if next == "R" {
			return false
		}
		cur = next
	}
}

func apply(part part, wid string) string {
	wf, ok := workflows[wid]
	if !ok {
		panic(wid)
	}
	for _, rule := range wf {
		if rule.cond == alwaysTrue {
			return rule.dest
		}
		var valFn func() int64
		switch rule.cond.variable {
		case "x":
			valFn = func() int64 { return part.x }
		case "m":
			valFn = func() int64 { return part.m }
		case "a":
			valFn = func() int64 { return part.a }
		case "s":
			valFn = func() int64 { return part.s }
		default:
			panic(rule.cond.variable)
		}
		v := valFn()
		switch rule.cond.operator {
		case "<":
			if v < rule.cond.value {
				return rule.dest
			}
		case ">":
			if v > rule.cond.value {
				return rule.dest
			}
		default:
			panic(rule.cond.operator)
		}
	}

	panic("no matching rule")
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	inWorkflows := true
	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			inWorkflows = false
			continue
		}

		if inWorkflows {
			lineParts := strings.Split(line, "{")
			if len(lineParts) != 2 {
				panic(line)
			}
			id := lineParts[0]
			workflows[id] = []rule{}
			rulesParts := strings.Split(strings.TrimSuffix(lineParts[1], "}"), ",")
			for _, p := range rulesParts {
				if !strings.Contains(p, ":") {
					workflows[id] = append(workflows[id], rule{
						cond: alwaysTrue,
						dest: p,
					})
					continue
				}

				rulePart := strings.Split(p, ":")
				if len(rulePart) != 2 {
					panic(p)
				}
				dest := rulePart[1]
				op := "<"
				condParts := strings.Split(rulePart[0], "<")
				if strings.Contains(p, ">") {
					op = ">"
					condParts = strings.Split(rulePart[0], ">")
				}
				if len(condParts) != 2 {
					panic(rulePart[0])
				}
				variable := condParts[0]
				val, err := strconv.Atoi(condParts[1])
				if err != nil {
					panic(err)
				}
				r := rule{
					cond: condition{
						variable: variable,
						operator: op,
						value:    int64(val),
					},
					dest: dest,
				}
				workflows[id] = append(workflows[id], r)
			}
		} else { // in parts
			matches := partsRe.FindStringSubmatch(line)
			if len(matches) != 5 {
				panic(line)
			}
			x, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			m, err := strconv.Atoi(matches[2])
			if err != nil {
				panic(err)
			}
			a, err := strconv.Atoi(matches[3])
			if err != nil {
				panic(err)
			}
			s, err := strconv.Atoi(matches[4])
			if err != nil {
				panic(err)
			}
			p := part{
				x: int64(x),
				m: int64(m),
				a: int64(a),
				s: int64(s),
			}
			parts = append(parts, p)
		}
	}
}
