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
