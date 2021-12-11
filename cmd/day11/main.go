package main

import (
	"fmt"
	"image"

	"github.com/dishbreak/aoc2020/lib"
	lib21 "github.com/dishbreak/aoc2021/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day11.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

var neighbors = []image.Point{
	{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
}

func part1(input []string) int {
	space := lib21.ToSpace(input[:len(input)-1])
	acc := 0

	for day := 0; day < 100; day++ {
		flashes := 0
		q := make([]image.Point, 0)
		for p := range space {
			space[p]++
			if space[p] > 9 {
				flashes++
				for _, n := range neighbors {
					q = append(q, p.Add(n))
				}
			}
		}

		for len(q) > 0 {
			p := q[0]
			q = q[1:]

			if _, ok := space[p]; !ok {
				continue
			}

			if space[p] >= 10 {
				continue
			}

			space[p]++
			if space[p] < 10 {
				continue
			}

			flashes++
			for _, n := range neighbors {
				q = append(q, p.Add(n))
			}

		}

		for p, v := range space {
			if v > 9 {
				space[p] = 0
			}
		}
		acc += flashes
	}
	return acc
}

func part2(input []string) int {
	space := lib21.ToSpace(input[:len(input)-1])

	for day := 0; day > -1; day++ {
		flashes := 0
		q := make([]image.Point, 0)
		for p := range space {
			space[p]++
			if space[p] > 9 {
				flashes++
				for _, n := range neighbors {
					q = append(q, p.Add(n))
				}
			}
		}

		for len(q) > 0 {
			p := q[0]
			q = q[1:]

			if _, ok := space[p]; !ok {
				continue
			}

			if space[p] >= 10 {
				continue
			}

			space[p]++
			if space[p] < 10 {
				continue
			}

			flashes++
			for _, n := range neighbors {
				q = append(q, p.Add(n))
			}

		}

		for p, v := range space {
			if v > 9 {
				space[p] = 0
			}
		}
		if flashes == 100 {
			return day + 1
		}
	}
	return -1
}
