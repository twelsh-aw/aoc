package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type (
	moduleType  string
	moduleState string
	pulseType   string
)

const (
	flipFlop    moduleType  = "%"
	conjunction moduleType  = "&"
	broadcaster moduleType  = "broadcaster"
	stateOn     moduleState = "on"
	stateOff    moduleState = "off"
	pulseHigh   pulseType   = "high"
	pulseLow    pulseType   = "low"
)

type module struct {
	id           string
	moduleType   moduleType
	destinations []string
	sources      []string
	state        moduleState
	lastPulse    pulseType
}

type pulse struct {
	source    string
	dest      string
	pulseType pulseType
}

type pulseLoop struct {
	id        string
	pulseType pulseType
	length    int
}

var (
	modules  = map[string]*module{}
	deadEnd  = &module{}
	finalEnd = &module{id: "rx"}
)

func main() {
	readInput()
	part1(1000, nil)
	part2()
}

func part1(maxNumPresses int, pulseLoop *pulseLoop) {
	cycleMin, cycleMax := 0, 0
	numLowPulses := 0
	numHighPulses := 0
	for i := 0; i < maxNumPresses; i++ {
		// always start with button press which sends low to broadcaster
		toProcess := []pulse{
			{
				source:    "button",
				dest:      "broadcaster",
				pulseType: pulseLow,
			},
		}

		// continue as long as pulses to process
		for len(toProcess) > 0 {
			// for part2; finding loops
			if pulseLoop != nil {
				if modules[pulseLoop.id].lastPulse == pulseLoop.pulseType {
					if cycleMin == 0 {
						cycleMin = i + 1
					} else if cycleMax == 0 && cycleMin != i+1 {
						cycleMax = i + 1
						pulseLoop.length = cycleMax - cycleMin
						return
					}
				}
			}
			next := []pulse{}
			// process pulses in order
			for _, p := range toProcess {
				if p.pulseType == pulseLow {
					numLowPulses++
				} else if p.pulseType == pulseHigh {
					numHighPulses++
				}
				mod, ok := modules[p.dest]
				if !ok {
					panic(p)
				}
				if mod == deadEnd {
					continue
				}
				if mod == finalEnd {
					if p.pulseType == pulseLow {
						return
					}
					continue
				}

				switch mod.moduleType {
				case broadcaster:
					next = append(next, mod.broadcast(p.pulseType)...)
				case flipFlop:
					next = append(next, mod.flipFlop(p.pulseType)...)
				case conjunction:
					next = append(next, mod.conjuct(p.pulseType)...)
				default:
					panic(mod.moduleType)
				}
			}
			toProcess = append([]pulse{}, next...)
		}
	}
	fmt.Printf("%v\n", numLowPulses*numHighPulses)
}

func part2() {
	// reset state from part1
	modules = make(map[string]*module)
	readInput()
	modules["rx"] = finalEnd

	// rx will get low pulse from dx iff all of (qt, qb, ng, mp) all last sent high to dx iff all of (ck, cs, jh, dx) all sent lows.
	// after this point, it gets more complicated but there appears to be a pattern there as well
	// we monitor for loops in (ck, cs, jh, dx) to find iterations where lows are sent
	ckLoop := &pulseLoop{id: "ck", pulseType: pulseLow}
	csLoop := &pulseLoop{id: "cs", pulseType: pulseLow}
	jhLoop := &pulseLoop{id: "jh", pulseType: pulseLow}
	dxLoop := &pulseLoop{id: "dx", pulseType: pulseLow}
	part1(math.MaxInt, ckLoop)
	part1(math.MaxInt, csLoop)
	part1(math.MaxInt, jhLoop)
	part1(math.MaxInt, dxLoop)
	steps := LCM(ckLoop.length, csLoop.length, jhLoop.length, dxLoop.length)
	fmt.Printf("%v\n", steps)
}

func (m *module) broadcast(p pulseType) []pulse {
	m.lastPulse = p
	next := []pulse{}
	for _, d := range m.destinations {
		next = append(next, pulse{source: m.id, dest: d, pulseType: p})
	}
	return next
}

func (m *module) flipFlop(p pulseType) []pulse {
	if p == pulseHigh {
		return []pulse{}
	}

	var nextPulse pulseType
	if m.state == stateOff {
		nextPulse = pulseHigh
		m.state = stateOn
	} else if m.state == stateOn {
		nextPulse = pulseLow
		m.state = stateOff
	} else {
		panic(m.state)
	}
	m.lastPulse = nextPulse
	next := []pulse{}
	for _, d := range m.destinations {
		next = append(next, pulse{source: m.id, dest: d, pulseType: nextPulse})
	}
	return next
}

func (m *module) conjuct(p pulseType) []pulse {
	var nextPulse pulseType = pulseLow
	for _, s := range m.sources {
		if modules[s].lastPulse != pulseHigh {
			nextPulse = pulseHigh
			break
		}
	}

	m.lastPulse = nextPulse
	next := []pulse{}
	for _, d := range m.destinations {
		next = append(next, pulse{source: m.id, dest: d, pulseType: nextPulse})
	}
	return next
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
		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			panic(line)
		}

		m := &module{state: stateOff, lastPulse: ""}
		mod := parts[0]
		if mod == "broadcaster" {
			m.id = mod
			m.moduleType = broadcaster
		} else {
			m.id = mod[1:]
			m.moduleType = moduleType(mod[0:1])
			if m.moduleType != flipFlop && m.moduleType != conjunction {
				panic(m.moduleType)
			}
		}

		destParts := strings.Split(parts[1], ", ")
		m.destinations = append(m.destinations, destParts...)
		modules[m.id] = m
	}

	// invert destinations to sources
	for _, m := range modules {
		for _, d := range m.destinations {
			md, ok := modules[d]
			if !ok { // dead end module
				modules[d] = deadEnd
				continue
			}
			md.sources = append(md.sources, m.id)
		}
	}
}

func (p pulse) String() string {
	return fmt.Sprintf("%s -%s-> %s", p.source, p.pulseType, p.dest)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
