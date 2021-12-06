package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	raw, err := lib.GetInput("inputs/day6.txt")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(raw[0], ",")
	input := make([]int, len(parts))
	for i, part := range parts {
		input[i], _ = strconv.Atoi(part)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(fishes []int) int {
	for day := 0; day < 80; day++ {
		newFish := 0
		for i, fish := range fishes {
			if fish == 0 {
				fishes[i] = 6
				newFish++
				continue
			}
			fishes[i]--
		}
		for i := 0; i < newFish; i++ {
			fishes = append(fishes, 8)
		}
	}
	return len(fishes)
}

func part2(fishes []int) int64 {
	return simulate(fishes, 256)
}

func simulate(fishes []int, days int) int64 {
	hist := make([]int, 9)

	for _, fish := range fishes {
		hist[fish]++
	}

	for day := 0; day < days; day++ {
		zero := hist[0]
		hist = append(hist[1:], zero)
		hist[6] += zero
		fmt.Println(day, hist)
	}

	sum := int64(0)
	for _, count := range hist {
		sum += int64(count)
	}

	return sum
}
