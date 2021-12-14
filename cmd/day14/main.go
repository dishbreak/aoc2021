package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInputAsSections("inputs/day14.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type polymerChain struct {
	pairs map[string]int
	chars map[byte]int
}

type insertionRule struct {
	insertChar   byte
	matchingPair string
	newPairs     []string
}

func parsePolymerChain(input string) polymerChain {
	p := polymerChain{
		pairs: make(map[string]int),
		chars: make(map[byte]int),
	}

	for i := 0; i < len(input)-1; i++ {
		p.pairs[input[i:i+2]]++
		p.chars[input[i]]++
	}
	p.chars[input[len(input)-1]]++

	return p
}

func parseInsertionRule(input string) insertionRule {
	parts := strings.Fields(input)

	r := insertionRule{
		insertChar:   parts[2][0],
		matchingPair: parts[0],
	}

	first, second := r.matchingPair[0], r.matchingPair[1]
	r.newPairs = []string{
		string([]byte{first, r.insertChar}),
		string([]byte{r.insertChar, second}),
	}

	return r
}

func simulate(input [][]string, steps int) int {
	ruleSet := make(map[string]insertionRule, len(input[1]))
	p := parsePolymerChain(input[0][0])

	for _, line := range input[1] {
		r := parseInsertionRule(line)
		ruleSet[r.matchingPair] = r
	}

	for i := 0; i < steps; i++ {
		q := make(map[string]insertionRule)

		for pair := range p.pairs {
			if rule, ok := ruleSet[pair]; !ok {
				continue
			} else {
				q[pair] = rule
			}
		}

		nextChain := make(map[string]int)

		for pair, val := range p.pairs {
			if _, ok := q[pair]; ok {
				continue
			}
			nextChain[pair] = val
		}

		for pair, rule := range q {
			val := p.pairs[pair]

			p.chars[rule.insertChar] += val
			nextChain[rule.newPairs[0]] += val
			nextChain[rule.newPairs[1]] += val
		}
		p.pairs = nextChain
	}

	counts := make([]int, 0)

	for _, val := range p.chars {
		counts = append(counts, val)
	}

	sort.IntSlice(counts).Sort()
	return counts[len(counts)-1] - counts[0]
}

func part1(input [][]string) int {
	return (simulate(input, 10))
}

func part2(input [][]string) int {
	return simulate(input, 40)
}
