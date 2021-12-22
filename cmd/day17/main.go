package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day17.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type targetArea struct {
	X, Y lib.Range
}

func parseTarget(line string) targetArea {
	targetRegex := regexp.MustCompile("x=(-?[0-9]+)..(-?[0-9]+), y=(-?[0-9]+)..(-?[0-9]+)")
	parts := targetRegex.FindStringSubmatch(line)
	t := targetArea{}
	t.X.Min, _ = strconv.Atoi(parts[1])
	t.X.Max, _ = strconv.Atoi(parts[2])
	t.Y.Min, _ = strconv.Atoi(parts[3])
	t.Y.Max, _ = strconv.Atoi(parts[4])
	return t
}

func part1(input []string) int {
	t := parseTarget(input[0])
	result := 0
	// only consider Y velocity for this problem -- X and Y velocities are independent!

	// don't bother with negative velocities because we're interested in maximum height.
	for vY0 := 0; vY0 <= -1*t.Y.Min; vY0++ {
		// fun fact, the trajectory from initial height to maximum height is an arithmetic series.
		// https://mathworld.wolfram.com/ArithmeticSeries.html
		maxHt := (vY0 * (vY0 + 1)) / 2

		// we iteratively model the freefall back to the target area.
		for y, vY := maxHt, 0; ; {
			y += vY
			vY -= 1
			if t.Y.Contains(y) {
				if result <= maxHt {
					result = maxHt
				}
				break
			}
			if y < t.Y.Min {
				break
			}
		}
	}
	return result
}

func part2(input []string) int {
	target := parseTarget(input[0])

	acc := 0
	for vY0 := target.Y.Min; vY0 <= -1*target.Y.Min; vY0++ {
		for vX0 := 1; vX0 <= target.X.Max; vX0++ {
			for x, y, vX, vY := 0, 0, vX0, vY0; ; {
				x += vX
				y += vY
				if vX != 0 {
					vX -= 1
				}
				vY -= 1
				if target.X.Contains(x) && target.Y.Contains(y) {
					acc++
					break
				}
				if x > target.X.Max || y < target.Y.Min {
					break
				}
			}
		}
	}

	return acc
}
