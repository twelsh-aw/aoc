package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type blueprint struct {
	ID                     int
	OreRobotOreCost        int
	ClayRobotOreCost       int
	ObsidianRobotOreCost   int
	ObsidianRobotClayCost  int
	GeodeRobotOreCost      int
	GeodeRobotObsidianCost int
	visited                map[string]bool
	TotalOreCost           int
}

type inventory struct {
	roboInventory
	oreInventory
}

type roboInventory struct {
	NumOreRobots      int
	NumClayRobots     int
	NumObsidianRobots int
	NumGeodeRobots    int
	TimeLeft          int
}

type oreInventory struct {
	NumOre      int
	NumClay     int
	NumObsidian int
	NumGeodes   int
}

var (
	blueprints []blueprint
	logTimes   = true
	logCounts  = false
)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	total := 0
	for _, bp := range blueprints {
		start := time.Now()
		maxGeodes := bp.GetMaxGeodes(24)
		qualityLevel := bp.ID * maxGeodes
		total += qualityLevel
		if logTimes {
			fmt.Println("part1 id:", bp.ID)
			fmt.Println("max:", maxGeodes)
			fmt.Println("time:", time.Since(start))
		}
	}
	fmt.Printf("%v\n", total)
}

func part2() {
	numBlueprints := 3
	if len(blueprints) < numBlueprints {
		numBlueprints = len(blueprints)
	}

	total := 1
	for _, bp := range blueprints[0:numBlueprints] {
		start := time.Now()
		maxGeodes := bp.GetMaxGeodes(32)
		total *= maxGeodes
		if logTimes {
			fmt.Println("part2 id:", bp.ID)
			fmt.Println("max:", maxGeodes)
			fmt.Println("time:", time.Since(start))
		}
	}
	fmt.Printf("%v\n", nil)
}

func (bp *blueprint) GetMaxGeodes(time int) int {
	toCheck := []inventory{
		{
			roboInventory: roboInventory{
				NumOreRobots: 1,
				TimeLeft:     time,
			},
		},
	}

	bp.visited = make(map[string]bool)
	maxByRobots := make(map[roboInventory]oreInventory)
	maxByRobots[toCheck[0].roboInventory] = toCheck[0].oreInventory
	for {
		numOrig := len(toCheck)
		numFiltered := 0
		for _, state := range toCheck {
			if state.TimeLeft <= 0 {
				continue
			}

			if bp.visited[state.key()] {
				numFiltered++
				continue
			}

			bp.visited[state.key()] = true

			if max, ok := maxByRobots[state.roboInventory]; ok {
				if state.oreInventory != max {
					if state.NumOre <= max.NumOre &&
						state.NumClay <= max.NumClay &&
						state.NumObsidian <= max.NumObsidian &&
						state.NumGeodes <= max.NumGeodes {
						numFiltered++
						continue
					}

					if state.NumGeodes+1 < max.NumGeodes {
						numFiltered++
						continue
					}

					if state.NumObsidian+bp.GeodeRobotObsidianCost < max.NumObsidian {
						numFiltered++
						continue
					}

					if state.NumObsidian > max.NumObsidian+bp.GeodeRobotObsidianCost {
						numFiltered++
						continue
					}

					if state.NumClay+bp.ObsidianRobotClayCost < max.NumClay {
						numFiltered++
						continue
					}

					if state.NumClay > max.NumClay+bp.ObsidianRobotClayCost {
						numFiltered++
						continue
					}

					if state.NumOre+bp.TotalOreCost < max.NumOre {
						numFiltered++
						continue
					}

					if state.NumOre > max.NumOre+bp.TotalOreCost {
						numFiltered++
						continue
					}
				}
			} else {
				panic("not in maxByRobots map")
			}

			nextStates := bp.getNextStates(state)
			for _, next := range nextStates {
				//fmt.Println(next)
				toCheck = append(toCheck, next)
				key := next.roboInventory
				max := maxByRobots[key]

				if next.NumGeodes < max.NumGeodes {
					continue
				}

				if next.NumGeodes > max.NumGeodes {
					maxByRobots[key] = next.oreInventory
					continue
				}

				if next.NumObsidian < max.NumObsidian {
					continue
				}

				if next.NumObsidian > max.NumObsidian {
					maxByRobots[key] = next.oreInventory
					continue
				}

				if next.NumClay < max.NumClay {
					continue
				}

				if next.NumClay > max.NumClay {
					maxByRobots[key] = next.oreInventory
					continue
				}

				if next.NumOre < max.NumOre {
					continue
				}

				if next.NumOre > max.NumOre {
					maxByRobots[key] = next.oreInventory
					continue
				}
			}
		}

		numAdded := len(toCheck) - numOrig
		if numAdded == 0 {
			break
		}

		if logCounts {
			fmt.Println("id", bp.ID, "minute", time-toCheck[0].TimeLeft+1, "added", numAdded, "filtered", numFiltered)
		}
		toCheck = toCheck[numOrig:]
	}

	bp.visited = nil
	maxGeodes := 0
	for _, v := range maxByRobots {
		if v.NumGeodes > maxGeodes {
			maxGeodes = v.NumGeodes
		}
	}
	return maxGeodes
}

func (bp *blueprint) buildNewOreRobot(i inventory) inventory {
	i.NumOre -= bp.OreRobotOreCost
	i.NumOreRobots++
	return i
}

func (bp *blueprint) buildNewClayRobot(i inventory) inventory {
	i.NumOre -= bp.ClayRobotOreCost
	i.NumClayRobots++
	return i
}

func (bp *blueprint) buildNewObsidianRobot(i inventory) inventory {
	i.NumOre -= bp.ObsidianRobotOreCost
	i.NumClay -= bp.ObsidianRobotClayCost
	i.NumObsidianRobots++
	return i
}

func (bp *blueprint) buildNewGeodeRobot(i inventory) inventory {
	i.NumOre -= bp.GeodeRobotOreCost
	i.NumObsidian -= bp.GeodeRobotObsidianCost
	i.NumGeodeRobots++
	return i
}

func (bp *blueprint) collect(i inventory) inventory {
	i.TimeLeft--
	i.NumOre += i.NumOreRobots
	i.NumClay += i.NumClayRobots
	i.NumObsidian += i.NumObsidianRobots
	i.NumGeodes += i.NumGeodeRobots
	return i
}

func (bp *blueprint) getNextStates(state inventory) []inventory {
	var nextStates []inventory

	collected := bp.collect(state)
	if true { // build nothing
		nextStates = append(nextStates, collected)
	}

	if state.NumOre >= bp.GeodeRobotOreCost && state.NumObsidian >= bp.GeodeRobotObsidianCost {
		next := bp.buildNewGeodeRobot(collected)
		nextStates = append(nextStates, next) //bp.getNextStates(next)...)
	}

	if state.NumObsidianRobots < bp.GeodeRobotObsidianCost {
		if state.NumOre >= bp.ObsidianRobotOreCost && state.NumClay >= bp.ObsidianRobotClayCost {
			next := bp.buildNewObsidianRobot(collected)
			nextStates = append(nextStates, next) //bp.getNextStates(next)...)
		}
	}

	if state.NumClayRobots < bp.ObsidianRobotClayCost {
		if state.NumOre >= bp.ClayRobotOreCost {
			next := bp.buildNewClayRobot(collected)
			nextStates = append(nextStates, next) //bp.getNextStates(next)...)
		}
	}

	if state.NumOreRobots < bp.TotalOreCost {
		if state.NumOre >= bp.OreRobotOreCost {
			next := bp.buildNewOreRobot(collected)
			nextStates = append(nextStates, next) //bp.getNextStates(next)...)
		}
	}

	return nextStates
}

func (i *inventory) key() string {
	return strings.Join([]string{
		// fmt.Sprint(i.TimeLeft),
		fmt.Sprint(i.NumOre),
		fmt.Sprint(i.NumClay),
		fmt.Sprint(i.NumObsidian),
		fmt.Sprint(i.NumGeodes),
		fmt.Sprint(i.NumOreRobots),
		fmt.Sprint(i.NumClayRobots),
		fmt.Sprint(i.NumObsidianRobots),
		fmt.Sprint(i.NumGeodeRobots),
	}, "_")
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	blueprintRe := regexp.MustCompile("Blueprint (\\d+): Each ore robot costs (\\d) ore\\. Each clay robot costs (\\d) ore\\. Each obsidian robot costs (\\d) ore and (\\d+) clay\\. Each geode robot costs (\\d) ore and (\\d+) obsidian\\.")
	for _, line := range strings.Split(string(b), "\n") {
		matches := blueprintRe.FindStringSubmatch(line)
		if len(matches) != 8 {
			panic(line)
		}

		bp := blueprint{}
		bp.ID, err = strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		bp.OreRobotOreCost, err = strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		bp.ClayRobotOreCost, err = strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}
		bp.ObsidianRobotOreCost, err = strconv.Atoi(matches[4])
		if err != nil {
			panic(err)
		}
		bp.ObsidianRobotClayCost, err = strconv.Atoi(matches[5])
		if err != nil {
			panic(err)
		}
		bp.GeodeRobotOreCost, err = strconv.Atoi(matches[6])
		if err != nil {
			panic(err)
		}
		bp.GeodeRobotObsidianCost, err = strconv.Atoi(matches[7])
		if err != nil {
			panic(err)
		}
		bp.TotalOreCost = bp.ClayRobotOreCost + bp.ObsidianRobotOreCost + bp.GeodeRobotOreCost
		blueprints = append(blueprints, bp)
	}

	////example cases from puzzle:
	//blueprints = []blueprint{
	//	{
	//		1, 4, 2, 3, 14, 2, 7, nil, 7,
	//	},
	//	{
	//		2, 2, 3, 3, 8, 3, 12, nil, 9,
	//	},
	//}
}
