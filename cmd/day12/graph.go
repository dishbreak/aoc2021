package main

import "strings"

type cave struct {
	ID        string
	IsBig     bool
	Neighbors []*cave
}

func (c *cave) String() string {
	return c.ID
}

type cavern struct {
	Caves map[string]*cave
}

func buildCavern(input []string) *cavern {
	c := &cavern{
		Caves: make(map[string]*cave),
	}

	for _, id := range []string{"start", "end"} {
		c.Caves[id] = &cave{
			ID:        id,
			Neighbors: make([]*cave, 0),
		}
	}

	edges := make([][]string, len(input))
	for i, line := range input {
		edges[i] = strings.Split(line, "-")
		for _, endpoint := range edges[i] {
			if _, ok := c.Caves[endpoint]; !ok {
				c.Caves[endpoint] = &cave{
					ID:        endpoint,
					IsBig:     strings.ToUpper(endpoint) == endpoint,
					Neighbors: make([]*cave, 0),
				}
			}
		}
	}

	for _, edge := range edges {
		one, other := edge[0], edge[1]
		c.Caves[one].Neighbors = append(c.Caves[one].Neighbors, c.Caves[other])
		c.Caves[other].Neighbors = append(c.Caves[other].Neighbors, c.Caves[one])
	}

	return c
}

type path struct {
	c        *cave
	visited  map[*cave]bool
	sequence []*cave
}

func (p *path) String() string {
	s := make([]string, len(p.sequence))
	for i, stop := range p.sequence {
		s[i] = stop.String()
	}

	return strings.Join(s, " -> ")
}

func (p *path) Branch(c *cave) *path {
	n := &path{
		c:        c,
		visited:  make(map[*cave]bool, len(p.visited)),
		sequence: make([]*cave, len(p.visited)+1),
	}

	for k, v := range p.visited {
		n.visited[k] = v
	}

	copy(n.sequence, p.sequence)
	n.sequence[len(n.sequence)-1] = c
	return n
}
