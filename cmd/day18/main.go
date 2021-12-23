package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day18.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type node struct {
	value             int
	left, right, root *node
}

func NodeFromString(input string) *node {
	buf := strings.NewReader(input)
	n := nodeFromReader(buf)
	return n
}

func nodeFromReader(buf io.ByteReader) *node {
	result := &node{
		value: -1,
	}

	if c, _ := buf.ReadByte(); c == '[' {
		result.left = nodeFromReader(buf)
	} else if v := c - '0'; v <= 9 {
		result.value = int(v)
		return result
	}

	if c, _ := buf.ReadByte(); c != ',' {
		panic(errors.New("expected comma"))
	}

	result.right = nodeFromReader(buf)
	if c, _ := buf.ReadByte(); c != ']' {
		panic(errors.New("expected closing brace"))
	}
	return result
}

func (n *node) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *node) Explode() bool {

	type frame struct {
		n     *node
		level int
	}
	s := make([]frame, 0)

	s = append(s, frame{n.right, 1})
	s = append(s, frame{n.left, 1})

	for len(s) > 0 {
		p := s[len(s)-1]
		s = s[1:]

		if p.level > 4 && p.n.IsLeaf() {
			if t := p.n.root.root.left; t != nil && t.IsLeaf() {
				t.value += p.n.left.value
			}
			if t := p.n.root.root.right; t != nil && t.IsLeaf() {
				t.value += p.n.right.value
			}
			return true
		}

		if p.n.IsLeaf() {
			continue
		}

		s = append(s, frame{p.n.right, p.level + 1})
		s = append(s, frame{p.n.left, p.level + 1})
	}

	return false
}

func add(l, r *node) *node {
	result := &node{
		value: -1,
	}
	l.root, r.root = result, result
	result.left, result.right = l, r

	return result
}

func (n *node) String() string {
	if n.IsLeaf() {
		return strconv.Itoa(n.value)
	}
	return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
}

func part1(input []string) int {
	return 0
}

func part2(input []string) int {
	return 0
}
