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
	reports = processReports(reports)

	translated := make(map[Point3D]bool)
	for _, r := range reports {
		for _, p := range r.Beacons {
			translated[p] = true
		}
	}

	return len(translated)
}

func part2(input [][]string) int {
	reports := make([]*Report, len(input))

	for i, chunk := range input {
		reports[i] = NewReport(chunk)
	}
	reports = processReports(reports)

	max := -1
	for i, one := range reports {
		for j, other := range reports {
			if i == j {
				continue
			}
			if d := one.Position.Dist(other.Position); d > max {
				max = d
			}
		}
	}

	return max
}

func processReports(reports []*Report) []*Report {
	q := make([]*Report, 1)
	q[0] = reports[0]
	reports[0].RelativeTo = 0

	for len(q) > 0 {
		one := q[0]
		q = q[1:]

		for _, other := range reports {
			if one.ID == other.ID || other.RelativeTo != -1 {
				continue
			}

			if m := testRotation(one, other); m != nil {
				q = append(q, other)
				normalized := make([]Point3D, len(other.Beacons))
				for i, pt := range other.Beacons {
					normalized[i] = m.t(pt).Add(m.v)
				}
				other.Beacons = normalized
				other.Position = m.v
				other.RelativeTo = one.ID
			}
		}
	}
	return reports
}

type rotationMatch struct {
	one, other *Report
	t          Transform
	v          Point3D
}

func testRotation(one, other *Report) (r *rotationMatch) {
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
			r = &rotationMatch{
				one:   one,
				other: other,
				t:     t,
				v:     v,
			}
			break
		}
	}
	return
}
