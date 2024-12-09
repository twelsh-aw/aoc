package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input string

type blockRange struct {
	start int
	end   int
}

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	rangesIndexedById, emptyRangesOrderedByStart, filledSize := parseRanges()

	curFileId := 0
	curBackFileId := len(rangesIndexedById) - 1
	curEmptyIdx := 0

	inFront := false
	inBack := false
	curPos := 0
	total := 0

	for curPos < filledSize {
		curFile := rangesIndexedById[curFileId]
		emptyFile := emptyRangesOrderedByStart[curEmptyIdx]
		curBackFile := rangesIndexedById[curBackFileId]

		if curPos >= curFile.start && curPos <= curFile.end {
			inFront = true
			total += curFileId * curPos
			curPos++
		} else if inFront {
			inFront = false
			curFileId++
		}

		if curPos >= emptyFile.start && curPos <= emptyFile.end {
			inBack = true
			if curBackFile.end >= curBackFile.start {
				total += curBackFileId * curPos
				curBackFile.end--
				curPos++
			} else {
				curBackFileId--
			}
		} else if inBack {
			inBack = false
			curEmptyIdx++
		}
	}

	fmt.Println(total)
}

func part2() {
	rangesIndexedById, emptyRangesOrderedByStart, _ := parseRanges()

	for curBackFileId := len(rangesIndexedById) - 1; curBackFileId >= 0; curBackFileId-- {
		curBackFile := rangesIndexedById[curBackFileId]
		for emptyFileId, emptyFile := range emptyRangesOrderedByStart {
			if emptyFile.start > curBackFile.start {
				break
			}

			curFileLength := (curBackFile.end - curBackFile.start) + 1
			curEmptySpace := (emptyFile.end - emptyFile.start) + 1
			if curFileLength <= curEmptySpace {
				curBackFile.start = emptyFile.start
				curBackFile.end = emptyFile.start + curFileLength - 1

				emptyFile.start += curFileLength
				if emptyFile.start > emptyFile.end { // invalidate empty space
					emptyRangesOrderedByStart = splice(emptyRangesOrderedByStart, emptyFileId)
				}
				break
			}
		}
	}

	total := 0
	for curFileId, curFile := range rangesIndexedById {
		for i := curFile.start; i <= curFile.end; i++ {
			total += curFileId * i
		}
	}

	fmt.Println(total)
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

		input = line
	}
}

func parseRanges() ([]*blockRange, []*blockRange, int) {
	rangesIndexedById := []*blockRange{}
	emptyRangesOrderedByStart := []*blockRange{}
	id := -1
	idx := 0
	filledSize := 0
	for i, char := range input {
		if i%2 == 0 {
			id++
			blockSize := parseDigit(char)
			if blockSize == 0 {
				panic("file block size is 0")
			}
			r := blockRange{
				start: idx,
				end:   idx + blockSize - 1,
			}
			idx += blockSize
			rangesIndexedById = append(rangesIndexedById, &r)
			filledSize += blockSize
		} else {
			emptySize := parseDigit(char)
			if emptySize > 0 {
				r := blockRange{
					start: idx,
					end:   idx + emptySize - 1,
				}
				idx += emptySize
				emptyRangesOrderedByStart = append(emptyRangesOrderedByStart, &r)
			}
		}
	}

	return rangesIndexedById, emptyRangesOrderedByStart, filledSize
}

func parseDigit(char rune) int {
	v, _ := strconv.Atoi(string(char))
	return v
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
