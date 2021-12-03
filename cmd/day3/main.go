package main

import (
	"fmt"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day3.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	bits := len(input[0])
	counters := make([]int, bits)
	cutoff := len(input) / 2

	for _, reading := range input {
		if reading == "" {
			continue
		}
		for i, c := range reading {
			if c == '1' {
				counters[bits-1-i]++
			}
		}
	}

	epsilon, gamma := 0, 0
	for i, count := range counters {
		if count > cutoff {
			epsilon = epsilon | 1<<i
		}
		if count < cutoff {
			gamma = gamma | 1<<i
		}
	}

	return epsilon * gamma
}

type trieNode struct {
	data     string
	children map[rune]*trieNode
}

func insert(t *trieNode, s string) {
	for _, c := range s {
		if _, ok := t.children[c]; !ok {
			t.children[c] = &trieNode{
				children: make(map[rune]*trieNode),
			}
		}
		t = t.children[c]
	}
	t.data = s
}

func traverse(t *trieNode) []string {
	hits := make([]string, 0)
	q := make([]*trieNode, 1)
	q[0] = t

	for len(q) > 0 {
		head := q[0]
		q = q[1:]
		if head.data != "" {
			hits = append(hits, head.data)
			continue
		}
		for _, n := range head.children {
			q = append(q, n)
		}
	}
	return hits
}

func prefix(t *trieNode, s string) []string {
	for _, c := range s {
		if _, ok := t.children[c]; !ok {
			return nil
		}
		t = t.children[c]
	}

	return traverse(t)
}

func countBit(input []string, pos int) int {
	acc := 0
	for _, reading := range input {
		if reading[pos] == '1' {
			acc++
		}
	}
	return acc
}

type selector func(int, int) (rune, int)

func getOxyReading(input []string, t *trieNode) int {
	return getReading(input, t, func(hits, cutoff int) (rune, int) {
		if hits >= cutoff {
			return '1', 1
		}
		return '0', 0
	})
}

func getCo2Reading(input []string, t *trieNode) int {
	return getReading(input, t, func(hits, cutoff int) (rune, int) {
		if hits >= cutoff {
			return '0', 0
		}
		return '1', 1
	})
}

func toInt(binary string) int {
	val := 0
	for i, c := range binary {
		if c == '1' {
			val = val | 1<<(len(binary)-i-1)
		}
	}
	return val
}

func getReading(input []string, t *trieNode, nextBit selector) int {
	query := strings.Builder{}
	value := 0

	bits := len(input[0])

	for i := 0; i < bits; i++ {
		if len(input) == 1 {
			return toInt(input[0])
		}
		hits := countBit(input, i)
		nextRune, bitVal := nextBit(hits, len(input)/2+len(input)%2)
		query.WriteRune(nextRune)
		if bitVal == 1 {
			value = value | 1<<(bits-i-1)
		}
		input = prefix(t, query.String())
	}

	return value
}

func part2(input []string) int {
	input = input[:len(input)-1]
	t := &trieNode{
		children: make(map[rune]*trieNode),
	}

	for _, reading := range input {
		insert(t, reading)
	}

	return getCo2Reading(input, t) * getOxyReading(input, t)
}
