package main

import (
	"strconv"
	"strings"
)

type indexPair struct {
	I, J int
}

type bingoBoard struct {
	lookup    map[int]indexPair
	rowCounts []int
	colCounts []int
	sum       int
}

func newBingoBoard(input []string) *bingoBoard {
	b := &bingoBoard{
		lookup:    make(map[int]indexPair),
		rowCounts: make([]int, 5),
		colCounts: make([]int, 5),
	}

	for i, line := range input {
		for j, val := range strings.Fields(line) {
			parsed, _ := strconv.Atoi(val)
			idx := indexPair{i, j}
			b.lookup[parsed] = idx
			b.sum += parsed
		}
	}

	return b
}

func (b *bingoBoard) playGame(rounds []int) (int, int) {
	for round, called := range rounds {
		idx, ok := b.lookup[called]
		if !ok {
			continue
		}

		b.sum -= called
		b.rowCounts[idx.I]++
		b.colCounts[idx.J]++

		if b.rowCounts[idx.I] == 5 || b.colCounts[idx.J] == 5 {
			return round, b.sum * called
		}
	}

	return len(rounds) + 1, -1
}
