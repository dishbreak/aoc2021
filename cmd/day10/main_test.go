package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
	"",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 26397, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 288957, part2(input))
}
