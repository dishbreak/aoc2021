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
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				panic(fmt.Errorf("failed to parse '%s': %s", input, err))
			} else {
				panic(err)
			}
		}
	}()

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
	} else if c == 'x' {
		acc := 0
		for d, _ := buf.ReadByte(); d != 'x'; d, _ = buf.ReadByte() {
			acc = (acc * 10) + int(d-'0')
		}
		result.value = acc
		return result
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
	reduced := false
	for !reduced {
		// don't short circuit, so evaluate and then set reduced
		exploded := n.Explode()

		if exploded {
			continue
		}

		split := n.Split()

		reduced = !exploded && !split
	}
}

func (n *node) Magnitude() int {
	if n.IsLeaf() {
		return n.value
	}

	return 3*n.left.Magnitude() + 2*n.right.Magnitude()
}

func (n *node) rightMost() *node {
	if n.IsLeaf() {
		return n
	}

	c := n
	for !c.IsLeaf() {
		c = c.right
	}
	return c
}

func (n *node) leftOf() *node {
	if n.root == nil {
		return nil
	}

	if n.root.right == n {
		return n.root.left.rightMost()
	}
	return n.root.leftOf()
}

func (n *node) leftMost() *node {
	if n.IsLeaf() {
		return n
	}

	c := n
	for !c.IsLeaf() {
		c = c.left
	}
	return c
}

func (n *node) rightOf() *node {
	if n.root == nil {
		return nil
	}

	if n.root.left == n {
		return n.root.right.leftMost()
	}

	return n.root.rightOf()
}

func (n *node) explodeNode() []*node {
	result := make([]*node, 0)

	if l := n.leftOf(); l != nil {
		l.value += n.left.value
		result = append(result, l)
	}

	if r := n.rightOf(); r != nil {
		r.value += n.right.value
		result = append(result, r)
	}

	return result
}

func (n *node) Explode() bool {
	return n.dfsSearch(func(f frame) bool {
		if f.level >= 5 {
			pair := f.n.root
			pair.explodeNode()
			pair.value = 0
			pair.right, pair.left = nil, nil
			return true
		}
		return false
	})
}

type frame struct {
	n     *node
	level int
}

func (n *node) dfsSearch(callback func(frame) bool) bool {
	s := make([]frame, 0)

	s = append(s, frame{n.right, 1})
	s = append(s, frame{n.left, 1})

	for len(s) > 0 {
		p := s[len(s)-1]
		s = s[:len(s)-1]

		if !p.n.IsLeaf() {
			s = append(s, frame{p.n.right, p.level + 1})
			s = append(s, frame{p.n.left, p.level + 1})
			continue
		}

		hit := callback(p)
		if hit {
			return true
		}

	}

	return false
}

func (n *node) Split() bool {
	return n.dfsSearch(func(f frame) bool {
		if f.n.value < 10 {
			return false
		}

		base, extra := f.n.value/2, f.n.value%2
		f.n.value = -1

		f.n.left = &node{value: base, root: f.n}
		f.n.right = &node{value: base + extra, root: f.n}
		return true
	})

}

func add(l, r *node) *node {
	result := &node{
		value: -1,
	}
	l.root, r.root = result, result
	result.left, result.right = l, r

	result.Reduce()

	return result
}

func (n *node) String() string {
	if n.IsLeaf() {
		return strconv.Itoa(n.value)
	}
	return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
}

func part1(input []string) int {
	operands := make([]*node, len(input)-1)
	for i, val := range input[:len(input)-1] {
		operands[i] = NodeFromString(val)
	}

	acc := operands[0]
	for _, other := range operands[1:] {
		if other == nil {
			continue
		}
		acc = add(acc, other)
	}

	return acc.Magnitude()
}

func addStrs(one, other string) int {
	n1, n2 := NodeFromString(one), NodeFromString(other)
	return add(n1, n2).Magnitude()
}

func part2(input []string) int {
	operands := input[:len(input)-1]

	max := -1
	for i := 0; i < len(operands); i++ {
		for j := 0; j < len(operands); j++ {
			if i == j {
				continue
			}
			val := addStrs(operands[i], operands[j])
			if val > max {
				max = val
			}
			val = addStrs(operands[j], operands[i])
			if val > max {
				max = val
			}
		}
	}

	return max
}
