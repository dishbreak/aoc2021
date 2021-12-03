package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	depth, dist := 0, 0

	for _, inst := range input {
		if inst == "" {
			continue
		}
		parts := strings.Fields(inst)
		dir := parts[0]
		value, _ := strconv.Atoi(parts[1])
		switch dir {
		case "up":
			depth -= value
		case "down":
			depth += value
		case "forward":
			dist += value
		}
	}
	return depth * dist
}

func part2(input []string) int {
	depth, dist, aim := 0, 0, 0

	for _, inst := range input {
		if inst == "" {
			continue
		}
		parts := strings.Fields(inst)
		dir := parts[0]
		value, _ := strconv.Atoi(parts[1])
		switch dir {
		case "up":
			aim -= value
		case "down":
			aim += value
		case "forward":
			dist += value
			depth += value * aim
		}
	}
	return depth * dist
}
