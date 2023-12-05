package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDestination(t *testing.T) {
	type testCase struct {
		source     int
		sourceMaps []sourceMap
		expected   int
	}

	seed2Soil := []sourceMap{
		{
			sourceStart: 98,
			destStart:   50,
			length:      2,
		},
		{
			sourceStart: 50,
			destStart:   52,
			length:      48,
		},
	}
	_ = fmt.Sprint(seed2Soil)

	soilToFert := []sourceMap{
		{
			sourceStart: 15,
			destStart:   0,
			length:      37,
		},
		{
			sourceStart: 52,
			destStart:   37,
			length:      2,
		},
		{
			sourceStart: 0,
			destStart:   39,
			length:      15,
		},
	}
	_ = fmt.Sprint(soilToFert)

	fertToWater := []sourceMap{
		{
			sourceStart: 53,
			destStart:   49,
			length:      8,
		},
		{
			sourceStart: 11,
			destStart:   0,
			length:      42,
		},
		{
			sourceStart: 0,
			destStart:   42,
			length:      7,
		},
		{
			sourceStart: 7,
			destStart:   57,
			length:      4,
		},
	}

	testCases := []testCase{
		{
			source:     79,
			sourceMaps: seed2Soil,
			expected:   81,
		},
		{
			source:     14,
			sourceMaps: seed2Soil,
			expected:   14,
		},
		{
			source:     55,
			sourceMaps: seed2Soil,
			expected:   57,
		},
		{
			source:     13,
			sourceMaps: seed2Soil,
			expected:   13,
		},
		{
			source:     81,
			sourceMaps: soilToFert,
			expected:   81,
		},
		{
			source:     14,
			sourceMaps: soilToFert,
			expected:   53,
		},
		{
			source:     53,
			sourceMaps: fertToWater,
			expected:   49,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			dest := getDestination(tc.source, tc.sourceMaps)
			require.Equal(t, tc.expected, dest)
		})
	}
}
