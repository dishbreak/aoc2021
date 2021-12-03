package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
	"",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 198, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 230, part2(input))
}
