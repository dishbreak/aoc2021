package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAluExec(t *testing.T) {
	type testCase struct {
		program      []string
		input        []int
		expectedVars map[string]int64
	}

	testCases := map[string]testCase{
		"negate value": {
			program: []string{
				"inp x",
				"mul x -1",
			},
			input: []int{34},
			expectedVars: map[string]int64{
				"w": int64(0),
				"x": int64(-34),
				"y": int64(0),
				"z": int64(0),
			},
		},
		"3x happy path": {
			program: []string{
				"inp z",
				"inp x",
				"mul z 3",
				"eql z x",
			},
			input: []int{3, 9},
			expectedVars: map[string]int64{
				"w": 0,
				"x": 9,
				"y": 0,
				"z": 1,
			},
		},
		"3x sad path": {
			program: []string{
				"inp z",
				"inp x",
				"mul z 3",
				"eql z x",
			},
			input: []int{3, 8},
			expectedVars: map[string]int64{
				"w": 0,
				"x": 8,
				"y": 0,
				"z": 0,
			},
		},
		"binary representation": {
			program: []string{
				"inp w",
				"add z w",
				"mod z 2",
				"div w 2",
				"add y w",
				"mod y 2",
				"div w 2",
				"add x w",
				"mod x 2",
				"div w 2",
				"mod w 2",
			},
			input: []int{10},
			expectedVars: map[string]int64{
				"w": 1,
				"x": 0,
				"y": 1,
				"z": 0,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			alu := NewAlu()
			alu.LoadProgram(tc.program)
			alu.LoadInput(tc.input)
			alu.Execute()
			assert.Equal(t, tc.expectedVars, alu.GetOutput())
		})
	}
}
