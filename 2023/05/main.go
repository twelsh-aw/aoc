package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	seeds                 = []int{}
	seedToSoil            = []sourceMap{}
	soilToFertilizer      = []sourceMap{}
	fertilizerToWater     = []sourceMap{}
	waterToLight          = []sourceMap{}
	lightToTemperature    = []sourceMap{}
	temperatureToHumidity = []sourceMap{}
	humidityToLocation    = []sourceMap{}
)

type sourceMap struct {
	sourceStart int
	destStart   int
	length      int
}

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	minScore := math.MaxInt
	for _, seed := range seeds {
		soil := getDestination(seed, seedToSoil)
		fert := getDestination(soil, soilToFertilizer)
		water := getDestination(fert, fertilizerToWater)
		light := getDestination(water, waterToLight)
		temp := getDestination(light, lightToTemperature)
		hum := getDestination(temp, temperatureToHumidity)
		loc := getDestination(hum, humidityToLocation)
		if loc < minScore {
			minScore = loc
		}
	}
	fmt.Printf("%v\n", minScore)
}

func part2() {
	minScore := math.MaxInt
	pairs := getSeedPairs()
	for _, pair := range pairs {
		// fmt.Println("testing pair", pair)
		for i := 0; i < pair[1]; i++ {
			seed := pair[0] + i
			soil := getDestination(seed, seedToSoil)
			fert := getDestination(soil, soilToFertilizer)
			water := getDestination(fert, fertilizerToWater)
			light := getDestination(water, waterToLight)
			temp := getDestination(light, lightToTemperature)
			hum := getDestination(temp, temperatureToHumidity)
			loc := getDestination(hum, humidityToLocation)
			if loc < minScore {
				minScore = loc
			}
		}
	}
	fmt.Printf("%v\n", minScore)
}

func getSeedPairs() [][2]int {
	pairs := [][2]int{}
	curPair := [2]int{}
	for i, seed := range seeds {
		if i%2 == 0 {
			if i != 0 {
				pairs = append(pairs, curPair)
			}
			curPair = [2]int{}
			curPair[0] = seed
		} else {
			curPair[1] = seed
		}
	}
	pairs = append(pairs, curPair)
	return pairs
}

func getDestination(source int, sourceMaps []sourceMap) int {
	for _, sm := range sourceMaps {
		if sm.sourceStart <= source && source < sm.sourceStart+sm.length {
			return sm.destStart + source - sm.sourceStart
		}
	}
	return source
}

func readInput() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	var curMap *[]sourceMap
	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			continue
		}

		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			switch parts[0] {
			case "seeds":
				nums := strings.Split(parts[1], " ")
				for _, num := range nums {
					if num != "" {
						seed, err := strconv.Atoi(num)
						if err != nil {
							panic(err)
						}
						seeds = append(seeds, seed)
					}
				}
			case "seed-to-soil map":
				curMap = &seedToSoil
			case "soil-to-fertilizer map":
				curMap = &soilToFertilizer
			case "fertilizer-to-water map":
				curMap = &fertilizerToWater
			case "water-to-light map":
				curMap = &waterToLight
			case "light-to-temperature map":
				curMap = &lightToTemperature
			case "temperature-to-humidity map":
				curMap = &temperatureToHumidity
			case "humidity-to-location map":
				curMap = &humidityToLocation
			default:
				panic(line)
			}
		} else {
			nums := strings.Split(line, " ")
			if len(nums) != 3 {
				panic(line)
			}
			ss, err := strconv.Atoi(nums[1])
			if err != nil {
				panic(err)
			}
			ds, err := strconv.Atoi(nums[0])
			if err != nil {
				panic(err)
			}
			l, err := strconv.Atoi(nums[2])
			if err != nil {
				panic(err)
			}
			lineMap := sourceMap{
				sourceStart: ss,
				destStart:   ds,
				length:      l,
			}
			*curMap = append(*curMap, lineMap)
		}
	}
}
