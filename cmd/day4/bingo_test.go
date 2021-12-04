package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBingoGame(t *testing.T) {
	input := []string{
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		" 2  0 12  3  7",
	}

	b := newBingoBoard(input)

	rounds := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	roundCount, score := b.playGame(rounds)
	assert.Equal(t, 4512, score)
	assert.Equal(t, 11, roundCount)
}
