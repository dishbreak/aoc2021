package main

import (
	"fmt"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day10.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

var scores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func evaluateLine(line string) int {
	s := make([]rune, 0)
	for _, c := range line {
		switch c {
		case '(':
			s = append(s, ')')
		case '[':
			s = append(s, ']')
		case '{':
			s = append(s, '}')
		case '<':
			s = append(s, '>')
		default:
			if s[len(s)-1] != c {
				return scores[c]
			}
			s = s[:len(s)-1]
		}
	}
	return 0
}

func part1(input []string) int {
	acc := 0
	for _, line := range input {
		if line == "" {
			continue
		}
		acc += evaluateLine(line)
	}
	return acc
}

func part2(input []string) int {
	return 0
}
