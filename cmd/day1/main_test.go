package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"199",
	"200",
	"208",
	"210",
	"200",
	"207",
	"240",
	"269",
	"260",
	"263",
	"",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 7, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 5, part2(input))
}
