package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
	"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
	"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
	"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
	"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
	"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
	"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
	"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
	"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
	"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
	"",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 26, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 61229, part2(input))
}

func TestEncodeSignal(t *testing.T) {
	type testCase struct {
		input  string
		result uint
	}

	testCases := []testCase{
		{"be", uint(18)},
		{"cfbegad", uint(127)},
		{"cbdgef", uint(126)},
		{"fgaecd", uint(125)},
		{"cgeb", uint(86)},
		{"fdcge", uint(124)},
	}

	for i, tc := range testCases {
		assert.Equal(t, tc.result, encodeSignal(tc.input).flags, "test case %d failed: %s", i, tc.input)
	}
}
