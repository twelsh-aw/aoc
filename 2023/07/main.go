package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type handType int

const (
	fiveKind  = 6
	fourKind  = 5
	fullHouse = 4
	threeKind = 3
	twoPair   = 2
	onePair   = 1
	highCard  = 0
)

var cardRanks = map[string]int{
	"2":     1,
	"3":     2,
	"4":     3,
	"5":     4,
	"6":     5,
	"7":     6,
	"8":     7,
	"9":     8,
	"T":     9,
	"J":     10,
	"Q":     11,
	"K":     12,
	"A":     13,
	"Joker": 0,
}

type hand struct {
	cards        [5]string
	countsByCard map[string]int
	handType     handType
	bid          int
}

var hands []hand

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	sortHands()
	total := 0
	for i, hand := range hands {
		rank := i + 1
		total += rank * hand.bid
	}
	fmt.Printf("%v\n", total)
}

func part2() {
	for i := range hands {
		hands[i].applyJokers()
	}

	sortHands()
	total := 0
	for i, hand := range hands {
		rank := i + 1
		total += rank * hand.bid
	}
	fmt.Printf("%v\n", total)
}

func sortHands() {
	sort.Slice(hands, func(i, j int) bool {
		handA := hands[i]
		handB := hands[j]
		if handA.handType < handB.handType {
			return true
		} else if handB.handType < handA.handType {
			return false
		}

		// equal hands
		for n := range handA.cards {
			cardA := cardRanks[handA.cards[n]]
			cardB := cardRanks[handB.cards[n]]
			if cardA < cardB {
				return true
			} else if cardB < cardA {
				return false
			}
		}
		return true
	})
}

func (h *hand) applyJokers() {
	numJokers := h.countsByCard["J"]
	switch numJokers {
	case 0:
		return
	case 1:
		switch h.handType {
		case fourKind:
			h.handType = fiveKind
		case threeKind:
			h.handType = fourKind
		case twoPair:
			h.handType = fullHouse
		case onePair:
			h.handType = threeKind
		case highCard:
			h.handType = onePair
		}
	case 2:
		switch h.handType {
		case threeKind:
			h.handType = fiveKind
		case fullHouse:
			h.handType = fiveKind
		case twoPair:
			h.handType = fourKind
		case onePair:
			h.handType = threeKind
		}
	case 3:
		switch h.handType {
		case fullHouse:
			h.handType = fiveKind
		case threeKind:
			h.handType = fourKind
		}
	case 4:
		switch h.handType {
		case fourKind:
			h.handType = fiveKind
		}
	}

	for i := range h.cards {
		if h.cards[i] == "J" {
			h.cards[i] = "Joker"
		}
	}
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

		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			panic(line)
		}
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		cards := strings.Split(parts[0], "")
		if len(cards) != 5 {
			panic(cards)
		}
		h := hand{
			cards:        [5]string{},
			countsByCard: map[string]int{},
			bid:          bid,
		}
		for i, card := range cards {
			h.cards[i] = card
			h.countsByCard[h.cards[i]]++
		}

		h.handType = getHandType(h.countsByCard)
		hands = append(hands, h)
	}
}

func getHandType(countsByCard map[string]int) handType {
	hasSet := false
	numPairs := 0
	for _, count := range countsByCard {
		if count == 5 {
			return fiveKind
		}
		if count == 4 {
			return fourKind
		}
		if count == 3 {
			hasSet = true
		}
		if count == 2 {
			numPairs++
		}
	}
	if hasSet && numPairs == 1 {
		return fullHouse
	} else if hasSet {
		return threeKind
	}
	if numPairs == 2 {
		return twoPair
	} else if numPairs == 1 {
		return onePair
	}

	return highCard
}
