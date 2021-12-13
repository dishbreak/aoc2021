package main

import (
	"fmt"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInputAsSections("inputs/day13.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: \n%s\n", part2(input))
}

func part1(input [][]string) int {
	p := paperFromInput(input)
	p.fold(p.folds[0])

	return p.count()
}

func part2(input [][]string) string {
	p := paperFromInput(input)
	p.foldUp()
	return p.String()
}
