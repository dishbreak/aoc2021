package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"v...>>.vv>",
	".vv>>.vv..",
	">>.>v>...v",
	">>v>>.>.v.",
	"v>v.vv.v..",
	">.>>..v...",
	".vv..>.>v.",
	"v.v..>>v.v",
	"....v..v.>",
}

// var input = []string{
// 	"...>...",
// 	".......",
// 	"......>",
// 	"v.....>",
// 	"......>",
// 	".......",
// 	"..vvv..",
// }

func TestPart1(t *testing.T) {
	assert.Equal(t, 58, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, part2(input))
}
