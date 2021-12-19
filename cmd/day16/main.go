package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"

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

type packet struct {
	typeId, version uint
	subpackets      []*packet
	value           uint
}

func parseLiteralValue(b *BitBuffer) (uint, error) {
	val := uint(0)
	for contBit, err := b.PopBits(1); contBit == 1; contBit, err = b.PopBits(1) {
		if err != nil {
			return 0, errors.Wrap(err, "failed to parse continue value")
		}
		if nybble, err := b.PopBits(4); err != nil {
			return 0, errors.Wrap(err, "failes to parse nybble")
		} else {
			val = (val << 4) | nybble
		}
	}
	if nybble, err := b.PopBits(4); err != nil {
		return 0, errors.Wrap(err, "failes to parse nybble")
	} else {
		val = (val << 4) | nybble
	}
	return val, nil
}

func parseSubpacketPktCount(b *BitBuffer) ([]*packet, error) {
	pktCount, err := b.PopBits(11)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get packet count")
	}

	result := make([]*packet, pktCount)

	for i := 0; i < int(pktCount); i++ {
		result[i], err = parsePacket(b)
		if err != nil {
			return result, errors.Wrap(err, "failed to parse subpacket")
		}
	}
	return result, nil
}

func parseSubpacketBitCount(b *BitBuffer) ([]*packet, error) {
	pktBitCount, err := b.PopBits(15)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get bit length")
	}
	stopPoint := b.pos + int(pktBitCount)
	if stopPoint > b.limit {
		return nil, errors.New("not enough bits left to parse")
	}

	result := make([]*packet, 0)
	for b.pos < stopPoint {
		p, err := parsePacket(b)
		result = append(result, p)
		if err != nil {
			return result, errors.Wrap(err, "failed to parse subpacket")
		}
	}

	return result, nil
}

type subPacketParser func(b *BitBuffer) ([]*packet, error)

func parsePacket(b *BitBuffer) (*packet, error) {
	p := &packet{
		subpackets: make([]*packet, 0),
	}

	if version, err := b.PopBits(3); err != nil {
		return p, errors.Wrap(err, "failed to parse version")
	} else {
		p.version = version
	}

	if typeId, err := b.PopBits(3); err != nil {
		return p, errors.Wrap(err, "failed to parse type")
	} else {
		p.typeId = typeId
	}

	if p.typeId == 4 {
		value, err := parseLiteralValue(b)
		if err != nil {
			return p, errors.Wrap(err, "failed to get packet literal value")
		}
		p.value = value
		return p, nil
	}

	parser := parseSubpacketBitCount
	lenId, err := b.PopBits(1)
	if err != nil {
		return p, errors.Wrap(err, "failed to get length type")
	}
	if lenId == 1 {
		parser = parseSubpacketPktCount
	}

	subPackets, err := parser(b)
	if err != nil {
		return p, errors.Wrap(err, "failed to get subpackets")
	}

	p.subpackets = subPackets
	return p, nil
}

func parseVersionSums(p *packet) int {
	acc := 0
	acc += int(p.version)

	for _, subP := range p.subpackets {
		acc += parseVersionSums(subP)
	}
	return acc
}

func part1(input []string) int {
	acc := 0
	for _, line := range input {
		if line == "" {
			continue
		}
		b, _ := NewBitBuffer(line)
		p, err := parsePacket(b)
		if err != nil {
			spew.Println(p)
			spew.Println(b)
			panic(err)
		}
		acc += parseVersionSums(p)
	}
	return acc
}

func part2(input []string) int {
	return 0
}
