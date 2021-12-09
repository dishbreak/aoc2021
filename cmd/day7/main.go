package main

import (
	"fmt"
	"math/bits"
	"sort"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day7.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func toIntSlice(input []string) []int {
	result := make([]int, len(input))
	for idx, val := range input {
		result[idx], _ = strconv.Atoi(val)
	}
	return result
}

func abs(o int) int {
	if o >= 0 {
		return o
	}
	return -1 * o
}

func part1(input []string) int {
	positions := toIntSlice(strings.Split(input[0], ","))

	values := make(map[int]bool)

	minFuelUsed := 0
	for _, position := range positions {
		values[position] = true
		minFuelUsed += position
	}

	for value := range values {
		totalFuel := 0
		for _, position := range positions {
			totalFuel += abs(position - value)
		}
		if totalFuel < minFuelUsed {
			minFuelUsed = totalFuel
		}
	}

	return minFuelUsed

}

func part2(input []string) int {
	positions := toIntSlice(strings.Split(input[0], ","))

	values := make(map[int]bool)

	sort.IntSlice(positions).Sort()

	maxDiff := positions[len(positions)-1] - positions[0]
	lookupTable := make([]int, maxDiff+1)

	for i := 1; i < len(lookupTable); i++ {
		lookupTable[i] = i + lookupTable[i-1]
	}

	minFuelUsed := 1<<(bits.UintSize-1) - 1
	for _, position := range positions {
		values[position] = true
	}

	for value := 0; value <= positions[len(positions)-1]; value++ {
		totalFuel := 0
		for _, position := range positions {
			dist := abs(position - value)
			totalFuel += lookupTable[dist]
		}
		if totalFuel < minFuelUsed {
			minFuelUsed = totalFuel
		}
	}

	return minFuelUsed
}
