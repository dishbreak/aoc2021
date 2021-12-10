package main

import (
	"fmt"
	"math/bits"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day8.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type noteEntry struct {
	signalPatterns []signalReading
	displays       []signalReading
}

type signalReading struct {
	flags    uint
	segments int
}

func encodeSignal(input string) signalReading {
	segments := len(input)
	val := uint(0)
	for _, c := range input {
		offset := c - 'a'
		val = val | 1<<offset
	}
	return signalReading{
		flags:    val,
		segments: segments,
	}
}

func (n noteEntry) descramble() int {
	legend := make(map[uint]int, 10)
	numbers := make([]uint, 10)

	// first pass: lock in the unique segments
	for _, signal := range n.signalPatterns {
		switch signal.segments {
		case 2:
			legend[signal.flags] = 1
			numbers[1] = signal.flags
		case 3:
			legend[signal.flags] = 7
			numbers[7] = signal.flags
		case 4:
			legend[signal.flags] = 4
			numbers[4] = signal.flags
		case 7:
			legend[signal.flags] = 8
			numbers[8] = signal.flags
		}
	}

	// second pass: use bitmasking to deduce the rest
	// "mystery" bitvector given as bv
	// 5-segment signals:
	// 3 -> bv & bitvector[1] == bitvector[1]
	// 2 -> bv &^ bitvector[4] has 3 bits set
	// 5 -> not 3 or 2

	// 6-segment signals:
	// 9 -> bv & bitvector[4] == bitvector[4]
	// 0 -> bv & bitvector[1] == bitvector[1]
	// 6 -> not 9 or 0
	for _, signal := range n.signalPatterns {
		if _, ok := legend[signal.flags]; ok {
			continue
		}
		switch signal.segments {
		case 5:
			if signal.flags&numbers[1] == numbers[1] {
				legend[signal.flags] = 3
			} else if masked := signal.flags &^ numbers[4]; bits.OnesCount(masked) == 3 {
				legend[signal.flags] = 2
			} else {
				legend[signal.flags] = 5
			}
		case 6:
			if signal.flags&numbers[4] == numbers[4] {
				legend[signal.flags] = 9
			} else if signal.flags&numbers[1] == numbers[1] {
				legend[signal.flags] = 0
			} else {
				legend[signal.flags] = 6
			}
		}
	}

	return legend[n.displays[0].flags]*1000 +
		legend[n.displays[1].flags]*100 +
		legend[n.displays[2].flags]*10 +
		legend[n.displays[3].flags]
}

func parseLine(input string) noteEntry {
	parts := strings.Split(input, " | ")

	n := noteEntry{
		signalPatterns: make([]signalReading, 10),
		displays:       make([]signalReading, 4),
	}

	for i, pattern := range strings.Fields(parts[0]) {
		n.signalPatterns[i] = encodeSignal(pattern)
	}

	for i, display := range strings.Fields(parts[1]) {
		n.displays[i] = encodeSignal(display)
	}

	return n
}

func part1(input []string) int {
	acc := 0
	for _, s := range input {
		if s == "" {
			continue
		}
		n := parseLine(s)
		for _, display := range n.displays {
			if (display.segments >= 2 && display.segments <= 4) || display.segments == 7 {
				acc++
			}
		}
	}
	return acc
}

func part2(input []string) int {
	acc := 0
	for _, s := range input {
		if s == "" {
			continue
		}
		n := parseLine(s)
		acc += n.descramble()
	}
	return acc
}
