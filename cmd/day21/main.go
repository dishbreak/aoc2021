package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day21.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func parseStartingPositions(input []string) (int, int) {
	parts := strings.Fields(input[0])
	p1start, _ := strconv.Atoi(parts[len(parts)-1])

	parts = strings.Fields(input[1])
	p2start, _ := strconv.Atoi(parts[len(parts)-1])

	return p1start, p2start
}

type DiracDice struct {
	Val   int
	Rolls int
}

func (d *DiracDice) Roll() int {
	d.Val++
	if d.Val > 100 {
		d.Val = 1
	}
	d.Rolls++
	return d.Val
}

type Player struct {
	pos, score int
}

func (p *Player) Play(d *DiracDice) {
	p.pos = (p.pos + d.Roll() + d.Roll() + d.Roll()) % 10
	p.score += p.pos + 1
}

func (p *Player) IsWinner() bool {
	return p.score >= 1000
}

func part1(input []string) int {
	p1, p2 := parseStartingPositions(input)

	players := []*Player{
		{
			pos: p1 - 1,
		},
		{
			pos: p2 - 1,
		},
	}

	var loser *Player
	d := &DiracDice{}
	for i := 0; ; i++ {
		p := players[i%2]
		p.Play(d)
		if p.IsWinner() {
			loser = players[(i+1)%2]
			break
		}
	}

	return d.Rolls * loser.score
}

func part2(input []string) int {
	return 0
}
