package main

import (
	"fmt"
	"sort"

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

func evaluateLine(line string) (int, []rune) {
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
				return scores[c], s
			}
			s = s[:len(s)-1]
		}
	}
	return 0, s
}

func part1(input []string) int {
	acc := 0
	for _, line := range input {
		if line == "" {
			continue
		}
		score, _ := evaluateLine(line)
		acc += score
	}
	return acc
}

var completeScores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func autoComplete(seq []rune) int {
	score := 0
	for i := len(seq) - 1; i >= 0; i-- {
		score *= 5
		score += completeScores[seq[i]]
	}
	return score
}

func part2(input []string) int {
	scores := make([]int, 0)

	for _, line := range input {
		if line == "" {
			continue
		}
		errScore, remaining := evaluateLine(line)
		if errScore != 0 {
			continue
		}

		scores = append(scores, autoComplete(remaining))

	}

	sort.IntSlice(scores).Sort()
	return scores[len(scores)/2]
}
