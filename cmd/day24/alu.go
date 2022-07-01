package main

import (
	"strconv"
	"strings"
)

type AluInstruction func(string, string)

type Alu struct {
	vars    map[string]int64
	input   []int
	program []string
}

func NewAlu() *Alu {
	a := &Alu{
		vars: make(map[string]int64, 4),
	}

	for _, val := range []string{"w", "x", "y", "z"} {
		a.vars[val] = 0
	}

	return a
}

func (a *Alu) LoadProgram(program []string) {
	a.program = make([]string, len(program))
	copy(a.program, program)
}

func (a *Alu) getVal(s string) int64 {
	v, ok := a.vars[s]
	if ok {
		return v
	}
	parsed, _ := strconv.Atoi(s)
	return int64(parsed)
}

func (a *Alu) LoadInput(input []int) {
	a.input = make([]int, len(input))
	copy(a.input, input)
}

func (a *Alu) Execute() {
	for _, inst := range a.program {
		a.runInstruction(inst)
	}
}

func (a *Alu) runInstruction(inst string) {
	parts := strings.Fields(inst)
	switch parts[0] {
	case "inp":
		a.vars[parts[1]] = int64(a.input[0])
		a.input = a.input[1:]
	case "add":
		one, other := a.getVal(parts[1]), a.getVal(parts[2])
		a.vars[parts[1]] = one + other
	case "mul":
		one, other := a.getVal(parts[1]), a.getVal(parts[2])
		a.vars[parts[1]] = one * other
	case "div":
		one, other := a.getVal(parts[1]), a.getVal(parts[2])
		a.vars[parts[1]] = one / other
	case "mod":
		one, other := a.getVal(parts[1]), a.getVal(parts[2])
		a.vars[parts[1]] = one % other
	case "eql":
		one, other := a.getVal(parts[1]), a.getVal(parts[2])
		a.vars[parts[1]] = 0
		if one == other {
			a.vars[parts[1]] = 1
		}
	}
}

func (a *Alu) GetOutput() map[string]int64 {
	result := make(map[string]int64, len(a.vars))

	for k, v := range a.vars {
		result[k] = v
	}

	return result
}
