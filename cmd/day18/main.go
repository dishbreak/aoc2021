package main

import (
	"fmt"
	"strconv"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day18.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type node struct {
	value       int
	left, right *node
}

func add(l, r *node) *node {
	result := &node{
		value: -1,
	}
	result.left = l
	result.right = r

	return result
}

func (n *node) String() string {
	if n.value != -1 {
		return strconv.Itoa(n.value)
	}
	return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
}

func part1(input []string) int {
	return 0
}

func part2(input []string) int {
	return 0
}
