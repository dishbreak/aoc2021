package main

import (
	"encoding/hex"

	"github.com/pkg/errors"
)

type BitBuffer struct {
	data  []byte
	pos   int
	limit int
}

func NewBitBuffer(input string) (*BitBuffer, error) {
	b := &BitBuffer{}
	if buf, err := hex.DecodeString(input); err != nil {
		return b, errors.Wrap(err, "failed to read hex string")
	} else {
		b.data = buf
		b.limit = len(buf) * 8
	}

	return b, nil
}

func (b *BitBuffer) Pos() int {
	return b.pos
}

func (b *BitBuffer) PopBits(bits int) (uint, error) {
	if b.pos+bits > b.limit {
		return 0, errors.New("end of buffer")
	}

	val := uint(0)
	for i := 0; i < bits; i++ {
		byteNum, offset := b.pos/8, b.pos%8
		bit := uint(0)
		if uint(b.data[byteNum]&(128>>offset)) > 0 {
			bit = uint(1)
		}
		val = (val << 1) | bit
		b.pos++
	}

	return val, nil
}
