package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []int{3, 4, 3, 1, 2}

func TestPart1(t *testing.T) {
	assert.Equal(t, 5934, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, int64(5934), simulate(input, 80))
	assert.Equal(t, int64(26984457539), part2(input))
}
