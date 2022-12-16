package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type tunnel struct {
	name    string
	flow    int
	leadsTo []string
}

type path struct {
	prevTunnel         string
	curTunnel          string
	prevElephantTunnel string
	curElephantTunnel  string
	pressureSoFar      int
	timeLeft           int
	valvesOpened       map[string]bool
}

var (
	tunnels             = make(map[string]tunnel)
	maxFlow, secondFlow int
	tunnelRegexp        = regexp.MustCompile("Valve ([A-Z]+) has flow rate=(\\d+); tunnel[s]? lead[s]? to valve[s]? ([A-Z ,]+)")
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	maxPressure := getMaxPressure(30, false)
	fmt.Printf("%v\n", maxPressure)
}

func part2() {
	maxPressure := getMaxPressure(26, true)
	fmt.Printf("%v\n", maxPressure)
}

func getMaxPressure(startTime int, includeElephant bool) int {
	valvesOpened := make(map[string]bool)
	for _, tun := range tunnels {
		valvesOpened[tun.name] = tun.flow == 0
	}

	maxPressure := 0
	startPath := path{
		curTunnel:     "AA",
		pressureSoFar: 0,
		timeLeft:      startTime,
		valvesOpened:  valvesOpened,
	}

	if includeElephant {
		startPath.curElephantTunnel = "AA"
	}

	pathsToCheck := []path{startPath}

	for {
		for _, p := range pathsToCheck {
			if p.pressureSoFar > maxPressure {
				maxPressure = p.pressureSoFar
			}
		}

		origNumPaths := len(pathsToCheck)
		//fmt.Println(startTime, origNumPaths)
		//startTime--
		for _, curPath := range pathsToCheck {
			nextPaths := getNextPaths(curPath, includeElephant)
			for _, n := range nextPaths {
				// want easy to calculate bound to filter things by
				maxPotential := n.getMaxPotentialFlowLeft(includeElephant)
				if maxPotential+n.pressureSoFar < maxPressure { // can't catch up
					continue
				}

				pathsToCheck = append(pathsToCheck, n)
			}
		}

		newNumPaths := len(pathsToCheck) - origNumPaths
		if newNumPaths <= 0 {
			break
		}
		pathsToCheck = append([]path{}, pathsToCheck[origNumPaths:]...)
	}

	return maxPressure
}

func (p *path) getMaxPotentialFlowLeft(includeElephant bool) int {
	if includeElephant {
		return (maxFlow + secondFlow) * (p.timeLeft - 1)
	}
	return (maxFlow) * (p.timeLeft - 1)
}

func getNextPaths(p path, includeElephant bool) []path {
	if p.timeLeft <= 1 {
		return []path{}
	}

	var paths []path
	curTunnel := tunnels[p.curTunnel]
	curElephantTunnel := tunnels[p.curElephantTunnel]
	canHumanOpen := !p.valvesOpened[curTunnel.name]
	canElephantOpen := includeElephant && curElephantTunnel.name != curTunnel.name && !p.valvesOpened[curElephantTunnel.name]

	if canHumanOpen && canElephantOpen {
		next := path{
			curTunnel:         p.curTunnel,
			curElephantTunnel: p.curElephantTunnel,
			pressureSoFar:     p.pressureSoFar + ((curTunnel.flow + curElephantTunnel.flow) * (p.timeLeft - 1)),
			timeLeft:          p.timeLeft - 1,
			valvesOpened:      clone(p.valvesOpened),
		}
		next.valvesOpened[curTunnel.name] = true
		next.valvesOpened[curElephantTunnel.name] = true
		paths = append(paths, next)
	}

	if canHumanOpen {
		next := path{
			curTunnel:     p.curTunnel,
			pressureSoFar: p.pressureSoFar + (curTunnel.flow * (p.timeLeft - 1)),
			timeLeft:      p.timeLeft - 1,
			valvesOpened:  clone(p.valvesOpened),
		}
		next.valvesOpened[curTunnel.name] = true
		paths = append(paths, next)

		for _, tun := range curElephantTunnel.leadsTo {
			if tun == p.prevTunnel || tun == p.prevElephantTunnel {
				continue
			}

			next.curElephantTunnel = tun
			next.prevElephantTunnel = p.curElephantTunnel
			paths = append(paths, next)
		}
	}

	if canElephantOpen {
		next := path{
			curElephantTunnel: p.curElephantTunnel,
			pressureSoFar:     p.pressureSoFar + (curElephantTunnel.flow * (p.timeLeft - 1)),
			timeLeft:          p.timeLeft - 1,
			valvesOpened:      clone(p.valvesOpened),
		}
		next.valvesOpened[curElephantTunnel.name] = true
		paths = append(paths, next)

		for _, tun := range curTunnel.leadsTo {
			if tun == p.prevTunnel || tun == p.prevElephantTunnel {
				continue
			}

			next.curTunnel = tun
			next.prevTunnel = p.curTunnel
			paths = append(paths, next)
		}
	}

	for _, tun := range curTunnel.leadsTo {
		if tun == p.prevTunnel || tun == p.prevElephantTunnel {
			continue
		}

		next := path{
			prevTunnel:    p.curTunnel,
			curTunnel:     tun,
			pressureSoFar: p.pressureSoFar,
			timeLeft:      p.timeLeft - 1,
			valvesOpened:  clone(p.valvesOpened),
		}
		paths = append(paths, next)

		for _, etun := range curElephantTunnel.leadsTo {
			if etun == p.prevTunnel || etun == p.prevElephantTunnel {
				continue
			}

			next.curElephantTunnel = etun
			next.prevElephantTunnel = p.curElephantTunnel
			paths = append(paths, next)
		}
	}

	return paths
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		matches := tunnelRegexp.FindStringSubmatch(line)
		if len(matches) != 4 {
			panic(line)
		}

		name := matches[1]
		flow, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}

		if flow >= maxFlow {
			secondFlow = maxFlow
			maxFlow = flow
		} else if flow > secondFlow {
			secondFlow = flow
		}

		connected := strings.Split(matches[3], ", ")
		tunnels[name] = tunnel{
			name:    name,
			flow:    flow,
			leadsTo: connected,
		}
	}
}

func clone[K comparable, V comparable](m map[K]V) map[K]V {
	c := make(map[K]V)
	for k, v := range m {
		c[k] = v
	}

	return c
}
