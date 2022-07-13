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

type PlayerState struct {
	Score int
	Pos   int
}

type DiceGame struct {
	Player1, Player2 PlayerState
	RollsRemaining   int
}

func (d DiceGame) Clone(rollsRemaining int) DiceGame {
	d.RollsRemaining = rollsRemaining
	return d
}

type VictoryRecord struct {
	Player1, Player2 int64
}

func part2(input []string) int64 {

	memo := make(map[DiceGame]VictoryRecord)
	p1, p2 := parseStartingPositions(input)

	d := DiceGame{
		Player1: PlayerState{
			Pos: p1 - 1,
		},
		Player2: PlayerState{
			Pos: p2 - 1,
		},
		RollsRemaining: 3,
	}

	var p1Wins, p2Wins int64
	for diceValue := 1; diceValue <= 3; diceValue++ {
		a, b := playGame(memo, d, diceValue, true)
		p1Wins += a
		p2Wins += b
	}

	if p1Wins < p2Wins {
		return p2Wins
	}
	return p1Wins
}

func playGame(memo map[DiceGame]VictoryRecord, d DiceGame, roll int, isPlayer1 bool) (p1Wins, p2Wins int64) {
	if r, ok := memo[d]; ok {
		return r.Player1, r.Player2
	}

	if d.Player1.Score >= 21 {
		panic("whoops that shouldn't happen")
	}

	defer func() {
		memo[d] = VictoryRecord{p1Wins, p2Wins}
	}()

	if isPlayer1 {
		d.Player1.Pos = (d.Player1.Pos + roll) % 10
	} else {
		d.Player2.Pos = (d.Player2.Pos + roll) % 10
	}

	d.RollsRemaining -= 1

	if d.RollsRemaining == 0 {
		if isPlayer1 {
			d.Player1.Score += d.Player1.Pos + 1
			if d.Player1.Score >= 21 {
				p1Wins = 1
				return
			}
		} else {
			d.Player2.Score += d.Player2.Pos + 1
			if d.Player2.Score >= 21 {
				p2Wins = 1
				return
			}
		}
		isPlayer1 = !isPlayer1
		d.RollsRemaining = 3
	}

	for diceValue := 1; diceValue <= 3; diceValue++ {
		a, b := playGame(memo, d, diceValue, isPlayer1)
		p1Wins += a
		p2Wins += b
	}

	return
}
