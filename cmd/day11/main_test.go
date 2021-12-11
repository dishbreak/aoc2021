package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"5483143223",
	"2745854711",
	"5264556173",
	"6141336146",
	"6357385478",
	"4167524645",
	"2176841721",
	"6882881134",
	"4846848554",
	"5283751526",
	"",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 1656, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, part2(input))
}
