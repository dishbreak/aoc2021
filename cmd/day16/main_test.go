package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, part2(input))
}

func TestParseVersionSums(t *testing.T) {
	type testCase struct {
		hexDump string
		sum     int
	}

	testCases := []testCase{
		{
			hexDump: "8A004A801A8002F478",
			sum:     16,
		},
		{
			hexDump: "620080001611562C8802118E34",
			sum:     12,
		},
		{
			hexDump: "C0015000016115A2E0802F182340",
			sum:     23,
		},
		{
			hexDump: "A0016C880162017C3686B18A3D4780",
			sum:     31,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			b, err := NewBitBuffer(tc.hexDump)
			assert.Nil(t, err)
			p, err := parsePacket(b)
			assert.Nil(t, err)
			assert.Equal(t, tc.sum, parseVersionSums(p))
		})
	}
}

func TestParsePacketValues(t *testing.T) {
	type testCase struct {
		hexDump string
		sum     int
	}

	testCases := []testCase{
		{
			hexDump: "C200B40A82",
			sum:     3,
		},
		{
			hexDump: "04005AC33890",
			sum:     54,
		},
		{
			hexDump: "880086C3E88112",
			sum:     7,
		},
		{
			hexDump: "CE00C43D881120",
			sum:     9,
		},
		{
			hexDump: "9C0141080250320F1802104A08",
			sum:     1,
		},
		{
			hexDump: "D8005AC2A8F0",
			sum:     1,
		},
		{
			hexDump: "F600BC2D8F",
			sum:     0,
		},
		{
			hexDump: "9C005AC2F8F0",
			sum:     0,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.sum, parseToValue(tc.hexDump))
		})
	}
}
