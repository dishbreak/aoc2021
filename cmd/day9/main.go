package main

import (
	"fmt"
	"image"
	"sort"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day9.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func toSpace(input []string) map[image.Point]int {
	space := make(map[image.Point]int, len(input)*len(input[0]))
	for x, s := range input {
		for y, c := range s {
			space[image.Point{x, y}] = int(c - '0')
		}
	}
	return space
}

var neighbors []image.Point = []image.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func isLocalMin(space map[image.Point]int, p image.Point) bool {
	v, ok := space[p]
	if !ok {
		return false
	}

	for _, n := range neighbors {
		val, ok := space[p.Add(n)]
		if !ok {
			continue
		}
		if val <= v { // tricksy...
			return false
		}
	}
	return true
}

func part1(input []string) int {
	space := toSpace(input[:len(input)-1])

	acc := 0
	for p, v := range space {
		if isLocalMin(space, p) {
			acc += (v + 1)
		}
	}
	return acc
}

func part2(input []string) int {
	space := toSpace(input[:len(input)-1])

	lowPoints := make([]image.Point, 0)

	for p := range space {
		if isLocalMin(space, p) {
			lowPoints = append(lowPoints, p)
		}
	}

	basins := make([]int, len(lowPoints))

	for i, lp := range lowPoints {
		acc := 0
		seen := make(map[image.Point]bool, len(space))
		seen[lp] = true
		q := make([]image.Point, 1)
		q[0] = lp

		for len(q) > 0 {
			p := q[0]
			q = q[1:]

			val, ok := space[p]
			if !ok || val == 9 {
				continue
			}

			acc++
			for _, n := range neighbors {
				next := p.Add(n)
				if _, ok := seen[next]; ok {
					continue
				}
				seen[next] = true
				q = append(q, p.Add(n))
			}
		}
		basins[i] = acc
	}
	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	return basins[0] * basins[1] * basins[2]
}
