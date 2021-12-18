package main

import (
	"encoding/hex"
)

type sliceBitBuffer struct {
	data []byte
	pos  int
}

type 

func NewBitBuffer(input string) (*BitBuffer, error) {
	b := &BitBuffer{}
	data, err := hex.DecodeString(input)
	b.data = data
	return b, err
}

func getMask(bits, offset int) byte {
	return byte(0) | byte(((1<<bits)-1)<<(offset%8))
}

func (b *BitBuffer) Pos() int {
	return b.pos
}

func (b *BitBuffer) PopBits(bits int) int {

	if bits+(b.pos%8) <= 8 {
		b.pos += bits
		byteNum, offset := (b.pos-1)/8, b.pos%8
		m := getMask(bits, 8-offset)
		maskedOut := b.data[byteNum] & m
		shiftAmt := (8 - offset) % 8
		shifted := maskedOut >> byte(shiftAmt)
		return int(shifted)
	}

	byteNum, offset := (b.pos-1)/8, b.pos%8
	headMask := getMask(8-offset, 0)
	bits -= 8 - offset
	val := int(b.data[byteNum] & headMask)
	b.pos += 8 - offset

	for bits > 8 {
		b.pos += 8
		byteNum = (b.pos - 1) / 8
		val = (val << 8) | int(b.data[byteNum])
		bits -= 8
	}

	tailMask := getMask(bits, 8-bits)

	b.pos += bits
	byteNum = (b.pos - 1) / 8
	val = (val << bits) | int((b.data[byteNum]&tailMask)>>(8-bits))

	return val
}
