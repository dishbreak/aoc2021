package main

import (
	"testing"

	"github.com/dishbreak/aoc2020/lib"
	"github.com/stretchr/testify/assert"
)

var input = []string{
	"target area: x=20..30, y=-10..-5",
	"",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 45, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 112, part2(input))
}

func TestParseTargetArea(t *testing.T) {
	expected := targetArea{
		X: lib.Range{
			Min: 20,
			Max: 30,
		},
		Y: lib.Range{
			Min: -10,
			Max: -5,
		},
	}
	actual := parseTarget(input[0])
	assert.Equal(t, expected, actual)
}
