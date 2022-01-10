package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{}

func TestPart1(t *testing.T) {
	assert.Equal(t, 0, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, part2(input))
}

func newNode(l, r int) *node {
	return add(
		&node{value: l},
		&node{value: r},
	)
}

func TestAdd(t *testing.T) {
	type testCase struct {
		operands []*node
		expected string
	}

	testCases := []testCase{
		{
			operands: []*node{
				newNode(1, 1),
				newNode(2, 2),
				newNode(3, 3),
				newNode(4, 4),
			},
			expected: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			acc := tc.operands[0]
			for _, other := range tc.operands[1:] {
				acc = add(acc, other)
			}
			assert.Equal(t, tc.expected, acc.String())
		})
	}
}

func TestExplode(t *testing.T) {
	type testCase struct {
		input, expected string
	}

	testCases := []testCase{
		{
			input:    "[[[[[9,8],1],2],3],4]",
			expected: "[[[[0,9],2],3],4]",
		},
		{
			input:    "[7,[6,[5,[4,[3,2]]]]]",
			expected: "[7,[6,[5,[7,0]]]]",
		},
		{
			input:    "[[6,[5,[4,[3,2]]]],1]",
			expected: "[[6,[5,[7,0]]],3]",
		},
		{
			input:    "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			expected: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			input:    "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			expected: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			n := NodeFromString(tc.input)
			n.Explode()
			assert.Equal(t, tc.expected, n.String())
		})
	}
}

func TestMagnitude(t *testing.T) {
	type testCase struct {
		input     string
		magnitude int
	}

	testCases := []testCase{
		{input: "[[1,2],[[3,4],5]]", magnitude: 143.},
		{input: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", magnitude: 1384.},
		{input: "[[[[1,1],[2,2]],[3,3]],[4,4]]", magnitude: 445.},
		{input: "[[[[3,0],[5,3]],[4,4]],[5,5]]", magnitude: 791.},
		{input: "[[[[5,0],[7,4]],[5,5]],[6,6]]", magnitude: 1137.},
		{input: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", magnitude: 3488},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			n := NodeFromString(tc.input)
			assert.Equal(t, tc.magnitude, n.Magnitude())
		})
	}
}

func TestNodeFromString(t *testing.T) {
	tests := []string{
		"[1,2]",
		"[[1,2],3]",
		"[9,[8,7]]",
		"[[1,9],[8,5]]",
		"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
		"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]",
		"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			n := NodeFromString(test)
			assert.Equal(t, test, n.String())
		})
	}
}
