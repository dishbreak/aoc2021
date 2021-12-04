package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInputAsSections("inputs/day4.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input [][]string) int {
	parts := strings.Split(input[0][0], ",")
	rounds := make([]int, len(parts))
	for i, part := range parts {
		rounds[i], _ = strconv.Atoi(part)
	}

	boards := input[1:]

	minRounds := len(rounds) + 1
	winningScore := -1

	for _, board := range boards {
		b := newBingoBoard(board)
		count, score := b.playGame(rounds)
		if count < minRounds {
			winningScore = score
			minRounds = count
		}
	}

	return winningScore
}

func part2(input [][]string) int {
	parts := strings.Split(input[0][0], ",")
	rounds := make([]int, len(parts))
	for i, part := range parts {
		rounds[i], _ = strconv.Atoi(part)
	}

	boards := input[1:]

	maxRounds := -1
	winningScore := -1

	for _, board := range boards {
		b := newBingoBoard(board)
		count, score := b.playGame(rounds)
		if count > maxRounds && score > 0 {
			winningScore = score
			maxRounds = count
		}
	}

	return winningScore
}
