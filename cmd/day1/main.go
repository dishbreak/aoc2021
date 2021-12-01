package main

import (
	"fmt"
	"strconv"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func toIntSlice(input []string) []int {
	r := make([]int, len(input)-1)
	for i, val := range input[:len(input)-1] {
		parsed, _ := strconv.Atoi(val)
		r[i] = parsed
	}
	return r
}

func part1(input []string) int {
	acc := 0
	vals := toIntSlice(input)

	for i := 1; i < len(vals); i++ {
		if vals[i] > vals[i-1] {
			acc++
		}
	}
	return acc
}

func part2(input []string) int {
	vals := toIntSlice(input)

	state := make([]int, len(vals))
	copy(state, vals)

	// using a state vector, we can build up the windowed sums
	// by making multiple passes.
	for window := 2; window <= 3; window++ {
		for i := window - 1; i < len(state); i++ {
			state[i] += vals[i-window+1]
		}
	}

	acc := 0
	for i := 3; i < len(state); i++ {
		if state[i] > state[i-1] {
			acc++
		}
	}

	return acc
}
