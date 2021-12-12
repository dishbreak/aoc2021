package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{}

func TestPart1(t *testing.T) {
	type testCase struct {
		input    []string
		expected int
	}

	testCases := []testCase{
		{
			input: []string{
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d",
				"A-end",
				"b-end",
				"",
			},
			expected: 10,
		},
		{
			input: []string{
				"dc-end",
				"HN-start",
				"start-kj",
				"dc-start",
				"dc-HN",
				"LN-dc",
				"HN-end",
				"kj-sa",
				"kj-HN",
				"kj-dc",
				"",
			},
			expected: 19,
		},
		{
			input: []string{
				"fs-end",
				"he-DX",
				"fs-he",
				"start-DX",
				"pj-DX",
				"end-zg",
				"zg-sl",
				"zg-pj",
				"pj-he",
				"RW-he",
				"fs-DX",
				"pj-RW",
				"zg-RW",
				"start-pj",
				"he-WI",
				"zg-he",
				"pj-fs",
				"start-RW",
				"",
			},
			expected: 226,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			actual := part1(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, part2(input))
}
