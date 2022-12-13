package main

import (
	"fmt"
	"testing"
)

func TestInCorrectOrder(t *testing.T) {
	testCases := []pair{
		{
			"[1,1,3,1,1]",
			"[1,1,5,1,1]",
		},
		{
			"[[1],[2,3,4]]",
			"[[1],4]",
		},
		{
			"[9]",
			"[[8,7,6]]",
		},
		{
			"[[4,4],4,4]",
			"[[4,4],4,4,4]",
		},
		{
			"[7,7,7,7]",
			"[7,7,7]",
		},
		{
			"[]",
			"[3]",
		},
		{
			"[[[]]]",
			"[[]]",
		},
		{
			"[1,[2,[3,[4,[5,6,7]]]],8,9]",
			"[1,[2,[3,[4,[5,6,0]]]],8,9]",
		},
		{
			"[[10,0,2]]",
			"[[[[4,3],5,[2,2,6,9],[],[8,4,4,7,2]],[3],1],[[6,5,[8],[10,9,9],6],[10],[3,[0,4,6,0],[3,4,8],9],0,[3,1,[1,10]]],[8],[[[1,0,6,1]],[[7,2,1,3,6],[7,8,0],9],[[9,7],7,4]]]",
		},
	}

	expectedCases := []bool{
		true,
		true,
		false,
		true,
		false,
		true,
		false,
		false,
		false,
	}

	for i, tc := range testCases {
		actual := inCorrectOrder(tc)
		expected := expectedCases[i]
		if actual != expected {
			panic(fmt.Sprintf("expected %v, got %v", expected, actual))
		}
	}
}

func TestGetNextPair(t *testing.T) {
	testCases := []string{
		"[1,1,3,1,1]",
		"[[1],[2,3,4]]",
		"[[1],4]",
		"[9]",
		"[[8,7,6]]",
		"[[4,4],4,4]",
		"[]",
		"[[[]]]",
		"[1,[2,[3,[4,[5,6,7]]]],8,9]",
		"[[],[8,4,4,7,2]]",
		"[10,0,2]",
		"[10]",
	}
	expectedCases := []pair{
		{
			"1",
			"[1,3,1,1]",
		},
		{
			"[1]",
			"[[2,3,4]]",
		},
		{
			"[1]",
			"[4]",
		},
		{
			"9",
			"",
		},
		{
			"[8,7,6]",
			"",
		},
		{
			"[4,4]",
			"[4,4]",
		},
		{
			"",
			"",
		},
		{
			"[[]]",
			"",
		},
		{
			"1",
			"[[2,[3,[4,[5,6,7]]]],8,9]",
		},
		{
			"[]",
			"[[8,4,4,7,2]]",
		},
		{
			"10",
			"[0,2]",
		},
		{
			"10",
			"",
		},
	}
	for i, tc := range testCases {
		actual := getNextPair(tc)
		expected := expectedCases[i]
		if actual != expected {
			panic(fmt.Sprintf("expected %v, got %v", expected, actual))
		}
	}
}
