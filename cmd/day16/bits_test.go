package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMask(t *testing.T) {
	type testCase struct {
		bits, offset int
		result       byte
	}

	testCases := []testCase{
		{
			bits:   3,
			offset: 3,
			result: byte(0b00111000),
		},
		{
			bits:   1,
			offset: 5,
			result: byte(0b00100000),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.result, getMask(tc.bits, tc.offset))
		})
	}
}

func TestPopBits(t *testing.T) {
	type popCall struct {
		bits   int
		result int
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
					result: 6,
				},
				{
					bits:   1,
					result: 1,
				},
				{
					bits:   4,
					result: 3,
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
					result: 6,
				},
				{
					bits:   8,
					result: 0b10011001,
				},
				{
					bits:   17,
					result: 0b01001001100111010,
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			b := &BitBuffer{
				data: tc.input,
			}

			for _, call := range tc.pops {
				assert.Equal(t, call.result, b.PopBits(call.bits))
			}
		})
	}
}
