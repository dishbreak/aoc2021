package main

import (
	"fmt"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInputAsSections("inputs/day19.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

const threshold = 12

func part1(input [][]string) int {
	reports := make([]*Report, len(input))

	for i, chunk := range input {
		reports[i] = NewReport(chunk)
	}

	translated := make(map[Point3D]bool)
	for _, p := range reports[0].Beacons {
		translated[p] = true
	}

	matches := make([][]*rotationMatch, len(reports))
	for i := range matches {
		matches[i] = make([]*rotationMatch, 0)
	}
	fixed := make(map[int]bool)
	for i, one := range reports {
		for j, other := range reports {
			if i == j {
				continue
			}
			if fixed[j] {
				continue
			}
			if m := testRotation(one, other); m != nil {
				matches[one.ID] = append(matches[one.ID], m)
				fixed[i] = true
			}
		}
	}
	return len(translated)
}

type rotationMatch struct {
	one, other *Report
	t          Transform
	v          Point3D
}

func testRotation(one, other *Report) *rotationMatch {
	v := Point3D{}
	for _, t := range Rotations {
		pts := other.Rotate(t)

		vectors := make(map[Point3D]int)
		for _, b0 := range one.Beacons {
			for _, c0 := range pts {
				vectors[b0.Sub(c0)]++
			}
		}

		max := -1

		for k, hits := range vectors {
			if hits > max {
				max = hits
				v = k
			}
		}

		if max >= threshold {
			return &rotationMatch{
				one:   one,
				other: other,
				t:     t,
				v:     v,
			}
		}
	}
	return nil
}

func part2(input [][]string) int {
	return 0
}
