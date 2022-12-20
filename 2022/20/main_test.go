package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNewIndex(t *testing.T) {
	const length = 3
	var next int

	// no change
	for i := 0; i < length; i++ {
		next := getNewIndex(i, number{0, 0}, length)
		require.Equal(t, i, next)
	}

	next = getNewIndex(0, number{1, 0}, length)
	require.Equal(t, 1, next)

	next = getNewIndex(0, number{2, 0}, length)
	require.Equal(t, 2, next)

	next = getNewIndex(0, number{3, 0}, length)
	require.Equal(t, 1, next)

	next = getNewIndex(0, number{4, 0}, length)
	require.Equal(t, 2, next)

	next = getNewIndex(0, number{5, 0}, length)
	require.Equal(t, 1, next)

	next = getNewIndex(0, number{6, 0}, length)
	require.Equal(t, 2, next)

	next = getNewIndex(0, number{7, 0}, length)
	require.Equal(t, 1, next)

	next = getNewIndex(0, number{8, 0}, length)
	require.Equal(t, 2, next)

	next = getNewIndex(0, number{9, 0}, length)
	require.Equal(t, 1, next)

	next = getNewIndex(0, number{-1, 0}, length)
	require.Equal(t, 1, next)

	next = getNewIndex(0, number{-2, 0}, length)
	require.Equal(t, 2, next)

	next = getNewIndex(0, number{-3, 0}, length)
	require.Equal(t, 1, next)

	next = getNewIndex(0, number{-4, 0}, length)
	require.Equal(t, 2, next)

	next = getNewIndex(0, number{-5, 0}, length)
	require.Equal(t, 1, next)

	next = getNewIndex(0, number{-6, 0}, length)
	require.Equal(t, 2, next)

	next = getNewIndex(0, number{-7, 0}, length)
	require.Equal(t, 1, next)

	next = getNewIndex(0, number{-8, 0}, length)
	require.Equal(t, 2, next)

	next = getNewIndex(0, number{-9, 0}, length)
	require.Equal(t, 1, next)
}
