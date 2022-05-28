package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

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

type result struct {
	Rounds int
	Score  int
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

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	problems := Generate(ctx, boards)
	solutions := Solver(ctx, problems, func(ctx context.Context, problem []string) result {
		r := result{}
		board := newBingoBoard(problem)
		r.Rounds, r.Score = board.playGame(rounds)
		return r
	})

	reports := 0
	for {
		select {
		case <-ctx.Done():
			if t, ok := ctx.Deadline(); ok && t.Before(time.Now()) {
				panic(errors.New("timed out"))
			} else if err := ctx.Err(); err != nil {
				panic(fmt.Errorf("error while executing: %w", err))
			}

		case r := <-solutions:
			reports++
			if minRounds > r.Rounds {
				winningScore = r.Score
				minRounds = r.Rounds
			}
			if reports == len(boards) {
				return winningScore
			}
		}
	}

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
