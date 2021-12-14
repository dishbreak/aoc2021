package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = [][]string{
	{
		"NNCB",
	},
	{
		"CH -> B",
		"HH -> N",
		"CB -> H",
		"NH -> C",
		"HB -> C",
		"HC -> B",
		"HN -> C",
		"NN -> C",
		"BH -> H",
		"NC -> B",
		"NB -> B",
		"BN -> B",
		"BB -> N",
		"BC -> B",
		"CC -> N",
		"CN -> C",
	},
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 1588, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 2188189693529, part2(input))
}

func TestInsertionRule(t *testing.T) {
	ruleString := "FG -> H"
	expected := insertionRule{
		insertChar:   'H',
		matchingPair: "FG",
		newPairs:     []string{"FH", "HG"},
	}

	assert.Equal(t, expected, parseInsertionRule(ruleString))
}
