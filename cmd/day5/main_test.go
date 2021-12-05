package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"0,9 -> 5,9",
	"8,0 -> 0,8",
	"9,4 -> 3,4",
	"2,2 -> 2,1",
	"7,0 -> 7,4",
	"6,4 -> 2,0",
	"0,9 -> 2,9",
	"3,4 -> 1,4",
	"0,0 -> 8,8",
	"5,5 -> 8,2",
	"",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 5, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 12, part2(input))
}
