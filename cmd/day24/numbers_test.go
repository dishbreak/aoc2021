package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasZeroDigits(t *testing.T) {
	testCases := map[int]bool{
		1:     false,
		9:     false,
		0:     true,
		40123: true,
		54678: false,
	}

	for k, expected := range testCases {
		t.Run(fmt.Sprintf("test case %d", k), func(t *testing.T) {
			assert.Equal(t, expected, HasZeroDigits(k))
		})
	}
}

func TestAsDigitSlice(t *testing.T) {
	testCases := map[int][]int{
		1010:  {0, 1, 0, 1},
		54032: {2, 3, 0, 4, 5},
		1:     {1},
		0:     {0},
	}

	for k, expected := range testCases {
		t.Run(fmt.Sprint("test case", k), func(t *testing.T) {
			assert.Equal(t, expected, AsDigitSlice(k))
		})
	}
}
