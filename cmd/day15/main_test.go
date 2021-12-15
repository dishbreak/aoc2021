package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"1163751742",
	"1381373672",
	"2136511328",
	"3694931569",
	"7463417111",
	"1319128137",
	"1359912421",
	"3125421639",
	"1293138521",
	"2311944581",
	"",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 40, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 315, part2(input))
}
