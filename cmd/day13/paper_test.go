package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputFoldX = [][]string{
	{
		"0,0",
		"2,0",
		"3,0",
		"6,0",
		"9,0",
		"0,1",
		"4,1",
		"6,2",
		"10,2",
		"0,3",
		"4,3",
		"1,4",
		"3,4",
		"6,4",
		"8,4",
		"9,4",
		"10,4",
	},
	{
		"fold along x=5",
	},
}

var expected = `#####
#...#
#...#
#...#
#####
.....
.....
`

func TestFoldX(t *testing.T) {
	p := paperFromInput(inputFoldX)
	p.fold(p.folds[0])

	assert.Equal(t, 16, p.count())
}

func TestFoldUp(t *testing.T) {
	p := paperFromInput(input)
	p.foldUp()
	assert.Equal(t, expected, p.String())
}
