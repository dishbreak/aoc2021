package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"2199943210",
	"3987894921",
	"9856789892",
	"8767896789",
	"9899965678",
	"",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 15, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 1134, part2(input))
}
