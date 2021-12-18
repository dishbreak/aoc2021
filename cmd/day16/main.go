package main

import (
	"fmt"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day16.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

/*
0011 1000 0000 0000 0110 1111 0100 0101 0010 1001 0001 0010 0000 0000
VVVT TTIL LLLL LLLL LLLL LLAA AAAA AAAA ABBB BBBB BBBB BBBB B---

1110 1110 0000 0000 1101 0100 0000 1100 1000 0010 0011 0000 0110 0000
VVVT TTIL LLLL LLLL LLAA AAAA AAAA ABBB BBBB BBBB CCCC CCCC CCC- ----
                      VV VTTT S
*/

func parseLiteralPacket(b *BitBuffer) int {
	val := 0
	for i := b.PopBits(1); i != 0; i = b.PopBits(1) {
		val = (val << 4) | b.PopBits(4)
	}
	val = (val << 4) | b.PopBits(4)
	return val
}

func parseVersionSums(b *BitBuffer) int {
	bits := b.Pos()
	offset := bits % 8
	byteNum := bits / 8
	fmt.Printf("bit %d (word %d offset %d)\n", bits, byteNum, offset)

	version := b.PopBits(3)
	acc := version
	typeId := b.PopBits(3)

	if typeId == 4 {
		parseLiteralPacket(b) // to seek to the end of the packet.
		return acc
	}

	lengthId := b.PopBits(1)

	if lengthId == 1 {
		subpackets := b.PopBits(11)
		for i := 0; i < subpackets; i++ {
			acc += parseVersionSums(b)
		}
		return acc
	}

	bitCount := b.PopBits(15)
	limit := b.Pos() + bitCount
	for i := b.Pos(); i < limit; i = b.Pos() {
		acc += parseVersionSums(b)
	}

	return acc
}

func part1(input []string) int {
	acc := 0
	for _, line := range input {
		b, err := NewBitBuffer(line)
		if err != nil {
			panic(err)
		}
		acc += parseVersionSums(b)
	}
	return acc
}

func part2(input []string) int {
	return 0
}
