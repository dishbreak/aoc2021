package main

import (
	"fmt"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day12.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func pathFind(c *cavern) []*path {
	start := c.Caves["start"]
	q := make([]*path, 1)
	ways := make([]*path, 0)
	q[0] = &path{
		sequence: []*cave{start},
		c:        start,
		visited:  map[*cave]bool{},
	}

	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		if n.c.ID == "end" {
			ways = append(ways, n)
			continue
		}

		n.visited[n.c] = true

		for _, next := range n.c.Neighbors {
			_, alreadyVisited := n.visited[next]
			if next.IsBig || !alreadyVisited {
				q = append(q, n.Branch(next))
			}
		}

	}

	return ways
}

func part1(input []string) int {
	c := buildCavern(input[:len(input)-1])
	return len(pathFind(c))
}

func part2(input []string) int {
	return 0
}
