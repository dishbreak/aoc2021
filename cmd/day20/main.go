package main

import (
	"fmt"
	"image"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInputAsSections("inputs/day20.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

var neighbors = []image.Point{
	{1, 1},
	{0, 1},
	{-1, 1},
	{1, 0},
	{0, 0},
	{-1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}

func part1(input [][]string) int {
	return enhance(input, 2)
}

func enhance(input [][]string, rounds int) int {
	algo := make([]bool, 512)
	for i, c := range input[0][0] {
		if c == '#' {
			algo[i] = true
		}
	}

	space := make(map[image.Point]bool)
	minDim := image.Point{}
	maxDim := image.Point{len(input[1][0]), len(input[1])}

	for y, line := range input[1] {
		for x, col := range line {
			space[image.Point{x, y}] = col == '#'
		}
	}

	for round := 0; round < rounds; round++ {
		minDim = minDim.Sub(image.Point{2, 2})
		maxDim = maxDim.Add(image.Point{2, 2})
		s2 := make(map[image.Point]bool)
		for y := minDim.Y; y < maxDim.Y; y++ {
			for x := minDim.X; x < maxDim.X; x++ {
				p := image.Point{x, y}
				val := 0
				for i, n := range neighbors {
					n2 := p.Add(n)
					if v, ok := space[n2]; (v && ok) || (algo[0] && !ok && round%2 == 1) {
						val = val | (1 << i)
					}
				}
				s2[p] = algo[val]
			}
		}
		space = s2
	}

	acc := 0
	for _, v := range space {
		if v {
			acc++
		}
	}
	return acc
}

func part2(input [][]string) int {
	return enhance(input, 50)
}
