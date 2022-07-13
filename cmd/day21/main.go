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

var diceFrequency map[int]int64 = map[int]int64{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

type PlayerState struct {
	Pos   int
	Score int
}

func (p PlayerState) IsWinner() bool {
	return p.Score >= 21
}

type DiracGame struct {
	Player1 PlayerState
	Player2 PlayerState
}

func (d DiracGame) Move(roll, playerID int) DiracGame {
	switch playerID {
	case 0:
		d.Player1.Pos = (d.Player1.Pos + roll) % 10
		d.Player1.Score += (d.Player1.Pos + 1)
	case 1:
		d.Player2.Pos = (d.Player2.Pos + roll) % 10
		d.Player2.Score += (d.Player2.Pos + 1)
	}
	return d
}

func (d DiracGame) String() string {
	return fmt.Sprintf("P1: %d@%d, P2: %d@%d", d.Player1.Pos, d.Player1.Score, d.Player2.Pos, d.Player2.Score)
}

func (d DiracGame) WinningPlayer() int {
	if d.Player1.IsWinner() {
		return 0
	}
	if d.Player2.IsWinner() {
		return 1
	}
	return -1
}

func NewGame(startP1, startP2 int) DiracGame {
	return DiracGame{
		Player1: PlayerState{Pos: startP1 - 1},
		Player2: PlayerState{Pos: startP2 - 1},
	}
}

type StateRecord struct {
	Game DiracGame
	Count int64
}

func part2(input []string) int64 {
	games := make([]StateRecord, 0)
	p1, p2 := parseStartingPositions(input)
	games[0] = StateRecord{
		Game: NewGame(p1, p2),
		Count: 1,
	}

	victories := make([]int64, 2)

	for playerID := 0; len(games) != 0; playerID = (playerID + 1) % 2 {
		nextGames, victoryCount := move(games, playerID)
		games = nextGames
		victories[playerID] += victoryCount
	}

	result := victories[0]
	if result < victories[1] {
		result = victories[1]
	}

	return result
}

func move(stateFrequencies []StateRecord, playerID int) ([], int64) {
	nextGameMap := make(map[DiracGame]int64)
	victories := int64(0)

	for _, record := range stateFrequencies {
		for roll, frequency := range diceFrequency {
			newgame := .Move(roll, playerID)
			newcount := count * frequency
			if newgame.WinningPlayer() == playerID {
				victories += newcount
				break
			}
			nextGameMap[newgame] += newcount
		}
	}
	return nextGameMap, victories
}
