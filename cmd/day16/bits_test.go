package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopBits(t *testing.T) {
	type popCall struct {
		bits   int
		result uint
	}
	type testCase struct {
		input []byte
		pops  []popCall
	}

	testCases := []testCase{
		{
			input: []byte{
				0b11010011,
			},
			pops: []popCall{
				{
					bits:   3,
					result: uint(6),
				},
				{
					bits:   1,
					result: uint(1),
				},
				{
					bits:   4,
					result: uint(3),
				},
			},
		},
		{
			input: []byte{
				0b11010011, 0b00101001, 0b00110011, 0b10101010,
			},
			pops: []popCall{
				{
					bits:   3,
					result: uint(6),
				},
				{
					bits:   5,
					result: uint(0b10011),
				},
				{
					bits:   17,
					result: uint(0b00101001001100111),
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			b := &BitBuffer{
				data:  tc.input,
				limit: len(tc.input) * 8,
			}

			for _, call := range tc.pops {
				actual, err := b.PopBits(call.bits)
				assert.Nil(t, err)
				assert.Equal(t, call.result, actual)
			}
		})
	}
}
