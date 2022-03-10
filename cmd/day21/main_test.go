package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"Player 1 starting position: 4",
	"Player 2 starting position: 8",
}

func TestParseStartingPos(t *testing.T) {
	p1, p2 := parseStartingPositions(input)
	assert.Equal(t, 4, p1)
	assert.Equal(t, 8, p2)
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 739785, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, part2(input))
}
