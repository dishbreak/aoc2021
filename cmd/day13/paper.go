package main

import (
	"image"
	"strconv"
	"strings"
)

type paper struct {
	dots  [][]int
	max   image.Point
	folds []fold
}

type Axis int

const (
	AxisX Axis = 0
	AxisY Axis = 1
)

type fold struct {
	direction Axis
	value     int
}

func paperFromInput(input [][]string) *paper {
	dotPairs := input[0]
	foldSteps := input[1]

	p := &paper{
		max:   image.Point{-1, -1},
		folds: make([]fold, len(foldSteps)),
	}

	points := make([]image.Point, len(dotPairs))
	for i, dotPair := range dotPairs {
		pt := image.Point{}
		parts := strings.Split(dotPair, ",")
		pt.X, _ = strconv.Atoi(parts[0])
		pt.Y, _ = strconv.Atoi(parts[1])
		if pt.X > p.max.X {
			p.max.X = pt.X
		}
		if pt.Y > p.max.Y {
			p.max.Y = pt.Y
		}
		points[i] = pt
	}

	p.dots = make([][]int, p.max.Y+1)
	for i := range p.dots {
		p.dots[i] = make([]int, p.max.X+1)
	}

	for _, pt := range points {
		p.dots[pt.Y][pt.X] = 1
	}

	for i, foldStep := range foldSteps {
		words := strings.Fields(foldStep)
		parts := strings.Split(words[2], "=")

		f := fold{}
		if parts[0] == "y" {
			f.direction = AxisY
		}
		f.value, _ = strconv.Atoi(parts[1])

		p.folds[i] = f
	}

	return p
}

func (p *paper) foldUp() {
	for _, f := range p.folds {
		p.fold(f)
	}
}

func (p *paper) fold(f fold) {
	switch f.direction {
	case AxisX:
		p.foldX(f.value)
	case AxisY:
		p.foldY(f.value)
	}
}

func (p *paper) foldX(value int) {
	n := make([][]int, len(p.dots))
	for i := range n {
		n[i] = make([]int, value)
		copy(n[i], p.dots[i][:value])
	}

	for y := 0; y < len(n); y++ {
		for x := 0; x < value; x++ {
			translated := (2 * value) - x
			n[y][x] = max(p.dots[y][translated], p.dots[y][x])
		}
	}
	p.dots = n
}

func max(one, other int) int {
	if one > other {
		return one
	}
	return other
}

func (p *paper) foldY(value int) {
	n := make([][]int, value)
	copy(n, p.dots[:value])

	for y := p.max.Y; y > value; y-- {
		for x := 0; x < len(p.dots[y]); x++ {
			translated := y - (2 * y) + (2 * value)
			n[translated][x] = max(n[translated][x], p.dots[y][x])
		}
	}
	p.max.Y = value - 1
	p.dots = n
}

func (p *paper) count() int {
	acc := 0

	for y := 0; y < len(p.dots); y++ {
		for x := 0; x < len(p.dots[y]); x++ {
			acc += p.dots[y][x]
		}
	}
	return acc
}

func (p *paper) String() string {
	b := strings.Builder{}
	for y := 0; y < len(p.dots); y++ {
		for x := 0; x < len(p.dots[y]); x++ {
			if p.dots[y][x] == 1 {
				b.WriteByte('#')
				continue
			}
			b.WriteByte(' ')
		}
		b.WriteByte('\n')
	}
	return b.String()
}
