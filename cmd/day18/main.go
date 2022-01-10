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
		result.left.root = result
	} else if v := c - '0'; v <= 9 {
		result.value = int(v)
		return result
	}

	if c, _ := buf.ReadByte(); c != ',' {
		panic(errors.New("expected comma"))
	}

	result.right = nodeFromReader(buf)
	result.right.root = result
	if c, _ := buf.ReadByte(); c != ']' {
		panic(errors.New("expected closing brace"))
	}
	return result
}

func (n *node) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *node) Reduce() {
	for l := n.Explode(); len(l) > 0; l = n.Explode() {
		for _, t := range l {
			t.Split()
		}
	}
}

func (n *node) Magnitude() int {
	if n.IsLeaf() {
		return n.value
	}

	return 3*n.left.Magnitude() + 2*n.right.Magnitude()
}

func (n *node) explodeNode() []*node {
	result := make([]*node, 0)

	for p, r := n, n.root; r != nil; p, r = r, r.root {
		if r.right == p {
			r.left.value += p.right.value
			result = append(result, r.left)
		}
		if r.left == p {
			r.right.value += p.left.value
			result = append(result, r.right)
		}
		if len(result) == 2 {
			break
		}
	}
	return result
}

func (n *node) Explode() []*node {
	type frame struct {
		n     *node
		level int
	}
	s := make([]frame, 0)

	s = append(s, frame{n.right, 1})
	s = append(s, frame{n.left, 1})

	var result []*node

	for len(s) > 0 {
		p := s[len(s)-1]
		s = s[:len(s)-1]

		if !p.n.IsLeaf() {
			s = append(s, frame{p.n.right, p.level + 1})
			s = append(s, frame{p.n.left, p.level + 1})
			continue
		}

		if p.level >= 5 {
			pair := p.n.root
			result = pair.explodeNode()
			pair.value = 0
			pair.right, pair.left = nil, nil
			return result
		}

	}

	return result
}

func (n *node) Split() {
	if n.value < 10 {
		return
	}

	base, extra := n.value/2, n.value%2
	n.value = -1

	n.left = &node{value: base}
	n.right = &node{value: base + extra}
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
