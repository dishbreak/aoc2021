package main

import (
	"fmt"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day25.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	s := ToSpace(input[:len(input)-1])
	return s.Simulate()
}
