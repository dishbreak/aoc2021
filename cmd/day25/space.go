package main

import (
	"fmt"
	"image"
	"strings"
)

type Space struct {
	maxDims   image.Point
	pts       map[image.Point]cukeState
	eastHerd  []image.Point
	southHerd []image.Point
}

type cukeState int

const (
	emptyCuke cukeState = iota
	southCuke
	eastCuke
)

func ToSpace(input []string) (s *Space) {
	s = &Space{
		maxDims:   image.Pt(len(input[0]), len(input)),
		pts:       make(map[image.Point]cukeState),
		eastHerd:  make([]image.Point, 0),
		southHerd: make([]image.Point, 0),
	}

	for y, line := range input {
		if line == "" {
			continue
		}
		for x, c := range line {
			pt := image.Pt(x, y)
			switch c {
			case '.':
				continue
			case '>':
				s.pts[pt] = eastCuke
				s.eastHerd = append(s.eastHerd, pt)
			case 'v':
				s.pts[pt] = southCuke
				s.southHerd = append(s.southHerd, pt)
			}
		}
	}

	return
}

func (s *Space) add(one, other image.Point) (result image.Point) {
	result = one.Add(other)
	result.X, result.Y = result.X%s.maxDims.X, result.Y%s.maxDims.Y
	return
}

func (s *Space) move(state cukeState, herd []image.Point, v image.Point) (newPts map[image.Point]cukeState, newHerd []image.Point, moves int) {
	newPts = make(map[image.Point]cukeState)
	newHerd = make([]image.Point, 0)

	for pt, val := range s.pts {
		if state == val {
			continue
		}
		newPts[pt] = val
	}

	for _, cuke := range herd {
		nextPt := s.add(cuke, v)
		if _, ok := s.pts[nextPt]; ok {
			newHerd = append(newHerd, cuke)
			newPts[cuke] = state
			continue
		}
		newPts[nextPt] = state
		newHerd = append(newHerd, nextPt)
		moves++
	}

	return
}

func (s *Space) Simulate() (rounds int) {
	moved := true
	for ; moved; rounds++ {
		moved = false

		nextPts, eastHerd, eMoves := s.move(eastCuke, s.eastHerd, image.Pt(1, 0))
		s.pts = nextPts
		s.eastHerd = eastHerd
		moved = moved || eMoves > 0

		nextPts, southHerd, sMoves := s.move(southCuke, s.southHerd, image.Pt(0, 1))
		s.pts = nextPts
		s.southHerd = southHerd
		moved = moved || sMoves > 0

		fmt.Printf("Round %3d: %4d east, %4d south\n", rounds+1, eMoves, sMoves)

	}
	return
}

func (s *Space) String() string {
	chars := make(map[image.Point]rune)
	for _, cuke := range s.eastHerd {
		chars[cuke] = '>'
	}
	for _, cuke := range s.southHerd {
		chars[cuke] = 'v'
	}

	var sb strings.Builder

	for y := 0; y < s.maxDims.Y; y++ {
		for x := 0; x < s.maxDims.X; x++ {
			if c, ok := chars[image.Pt(x, y)]; ok {
				sb.WriteRune(c)
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}
