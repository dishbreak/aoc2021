package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
	"",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 150, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 900, part2(input))
}
