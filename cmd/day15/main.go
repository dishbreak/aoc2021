package main

import (
	"container/heap"
	"errors"
	"fmt"
	"image"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day15.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type path struct {
	pt      image.Point
	risk    int
	visited map[image.Point]int
	steps   []image.Point
}

func (p *path) branch(o image.Point) *path {
	other := &path{
		pt:      o,
		risk:    p.risk,
		visited: make(map[image.Point]int, len(p.visited)),
		steps:   make([]image.Point, len(p.steps)+1),
	}

	for i, step := range p.steps {
		other.steps[i] = step
	}

	for k, v := range p.visited {
		other.visited[k] = v
	}

	other.steps[len(other.steps)-1] = o

	return other
}

type pointDist struct {
	p image.Point
	d int
}

type pointHeap []pointDist

func (h pointHeap) Len() int {
	return len(h)
}

func (h pointHeap) Less(i, j int) bool {
	return h[i].d < h[j].d
}

func (h pointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pointHeap) Push(x interface{}) {
	*h = append(*h, x.(pointDist))
}

func (h *pointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func dijkstra(space map[image.Point]int, end image.Point) int {
	points := (end.X + 1) * (end.Y + 1)
	start := image.Point{}
	infinity := points * 9
	dist := make(map[image.Point]int, points)
	visited := make(map[image.Point]int, points)
	minHeap := &pointHeap{}

	for k := range space {
		dist[k] = infinity
	}

	dist[start] = 0

	for current := start; ; {
		if current.Eq(end) {
			return dist[end]
		}
		for _, n := range []image.Point{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
			next := current.Add(n)
			if _, ok := space[next]; !ok {
				continue
			}
			if _, ok := visited[next]; ok {
				continue
			}

			tDist, ok := dist[next]
			if !ok {
				panic(errors.New("mismatch between space and dist maps"))
			}

			if nDist := dist[current] + space[next]; nDist < tDist {
				dist[next] = nDist
				heap.Push(minHeap, pointDist{next, nDist})
			}
		}
		visited[current]++
		current = heap.Pop(minHeap).(pointDist).p
	}
}

func part1(input []string) int {
	space := make(map[image.Point]int)

	for y, line := range input {
		if line == "" {
			continue
		}
		for x, col := range line {
			p := image.Point{x, y}
			space[p] = int(col - '0')
		}
	}

	end := image.Point{len(input[0]) - 1, len(input) - 2}

	return dijkstra(space, end)
}

func part2(input []string) int {
	space := make(map[image.Point]int)
	b := make(map[image.Point]int)

	tEnd := image.Point{len(input[0]) - 1, len(input) - 2}
	tEnd = tEnd.Add(image.Point{1, 1})
	end := tEnd.Mul(5).Add(image.Point{-1, -1})

	for y, line := range input {
		if line == "" {
			continue
		}
		for x, col := range line {
			p := image.Point{x, y}
			r := int(col - '0')
			space[p] = r
			b[p] = r
		}
	}

	base := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	lookups := make(map[int][][]int)
	for i := 1; i <= 9; i++ {
		progression := make([]int, 9)
		copy(progression, base)
		progression = append(progression[i-1:], progression[:i-1]...)
		lookups[i] = make([][]int, 5)
		for r := 0; r < 5; r++ {
			lookups[i][r] = make([]int, 5)
			if r != 0 {
				progression = append(progression[1:], progression[:1]...)
			}
			copy(lookups[i][r], progression)
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 0 && j == 0 {
				continue
			}
			for k, v := range b {
				translated := image.Point{k.X + (i * tEnd.X), k.Y + (j * tEnd.Y)}
				tVal := lookups[v][j][i]
				space[translated] = tVal
			}
		}
	}

	return dijkstra(space, end)
}
