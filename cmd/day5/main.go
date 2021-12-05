package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day5.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type direction int

const (
	NorthSouth direction = 1
	EastWest   direction = 2
	Other      direction = 0
)

type point struct {
	I, J int
}

func (p point) Sub(o point) point {
	i := p.I - o.I
	j := p.J - o.J
	return point{i, j}
}

func (p point) Add(o point) point {
	return point{
		p.I + o.I,
		p.J + o.J,
	}
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func (p point) Norm() point {
	var n int
	if p.I == 0 {
		n = abs(p.J)
	} else {
		n = abs(p.I)
	}
	return point{
		p.I / n,
		p.J / n,
	}
}

func (p point) Equal(o point) bool {
	return p.I == o.I && p.J == o.J
}

func (l line) Trace() []point {
	result := make([]point, 0)
	for i := l.A; !i.Equal(l.B); i = i.Add(l.V) {
		result = append(result, i)
	}
	result = append(result, l.B)
	return result
}

type line struct {
	A, B point
	D    direction
	V    point
}

func parsePoint(input string) point {
	p := point{}
	parts := strings.Split(input, ",")
	p.I, _ = strconv.Atoi(parts[0])
	p.J, _ = strconv.Atoi(parts[1])
	return p
}

func parseLine(input string) line {
	parts := strings.Fields(input)
	l := line{
		A: parsePoint(parts[0]),
		B: parsePoint(parts[2]),
		D: Other,
	}
	if l.A.I == l.B.I {
		l.D = EastWest
	}
	if l.A.J == l.B.J {
		l.D = NorthSouth
	}
	l.V = l.B.Sub(l.A).Norm()
	return l
}

func part1(input []string) int {
	hits := make(map[point]int)
	overlaps := make(map[point]int)
	for _, v := range input {
		if v == "" {
			continue
		}
		l := parseLine(v)
		if l.D == Other {
			continue
		}
		for _, p := range l.Trace() {
			hits[p]++
			if hits[p] > 1 {
				overlaps[p]++
			}
		}
	}
	return len(overlaps)
}

func part2(input []string) int {
	hits := make(map[point]int)
	overlaps := make(map[point]int)
	for _, v := range input {
		if v == "" {
			continue
		}
		l := parseLine(v)
		for _, p := range l.Trace() {
			hits[p]++
			if hits[p] > 1 {
				overlaps[p]++
			}
		}
	}
	return len(overlaps)
}
