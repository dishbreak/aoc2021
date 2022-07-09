package main

import (
	"strconv"
	"strings"
)

type Point3D struct {
	X, Y, Z int
}

func NewPoint3DFromString(input string) Point3D {
	parts := strings.Split(input, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])

	return Point3D{x, y, z}
}

type Transform func(Point3D) Point3D

// Taken from Šimon Tóth's C++ solution
// https://itnext.io/modern-c-in-advent-of-code-day19-ff9525afb2ee
var Rotations []Transform = []Transform{
	//  {{{X, POS}, {Y, POS}, {Z, POS}}},
	func(p Point3D) Point3D {
		return Point3D{p.X, p.Y, p.Z}
	},
	//	{{{X, POS}, {Z, POS}, {Y, NEG}}},
	func(p Point3D) Point3D {
		return Point3D{p.X, p.Z, -p.Y}
	},
	//  {{{X, POS}, {Y, NEG}, {Z, NEG}}},
	func(p Point3D) Point3D {
		return Point3D{p.X, -p.Y, -p.Z}
	},
	//	{{{X, POS}, {Z, NEG}, {Y, POS}}},
	func(p Point3D) Point3D {
		return Point3D{p.X, -p.Z, p.Y}
	},
	//  {{{X, NEG}, {Y, POS}, {Z, NEG}}},
	func(p Point3D) Point3D {
		return Point3D{-p.X, p.Y, -p.Z}
	},
	//	{{{X, NEG}, {Z, NEG}, {Y, NEG}}},
	func(p Point3D) Point3D {
		return Point3D{-p.X, -p.Z, -p.Y}
	},
	//  {{{X, NEG}, {Y, NEG}, {Z, POS}}},
	func(p Point3D) Point3D {
		return Point3D{-p.X, -p.Y, p.Z}
	},
	//	{{{X, NEG}, {Z, POS}, {Y, POS}}},
	func(p Point3D) Point3D {
		return Point3D{-p.X, p.Z, p.Y}
	},
	//  {{{Y, POS}, {X, POS}, {Z, NEG}}},
	func(p Point3D) Point3D {
		return Point3D{p.Y, p.X, -p.Z}
	},
	//	{{{Y, POS}, {Z, NEG}, {X, NEG}}},
	func(p Point3D) Point3D {
		return Point3D{p.Y, -p.Z, -p.X}
	},
	//  {{{Y, POS}, {X, NEG}, {Z, POS}}},
	func(p Point3D) Point3D {
		return Point3D{p.Y, -p.X, p.Z}
	},
	//	{{{Y, POS}, {Z, POS}, {X, POS}}},
	func(p Point3D) Point3D {
		return Point3D{p.Y, p.Z, p.X}
	},
	//  {{{Y, NEG}, {X, POS}, {Z, POS}}},
	func(p Point3D) Point3D {
		return Point3D{-p.Y, p.X, p.Z}
	},
	//	{{{Y, NEG}, {Z, POS}, {X, NEG}}},
	func(p Point3D) Point3D {
		return Point3D{-p.Y, p.Z, -p.X}
	},
	//  {{{Y, NEG}, {X, NEG}, {Z, NEG}}},
	func(p Point3D) Point3D {
		return Point3D{-p.Y, -p.X, -p.Z}
	},
	//	{{{Y, NEG}, {Z, NEG}, {X, POS}}},
	func(p Point3D) Point3D {
		return Point3D{-p.Y, -p.Z, p.X}
	},
	//  {{{Z, POS}, {X, POS}, {Y, POS}}},
	func(p Point3D) Point3D {
		return Point3D{p.Z, p.X, p.Y}
	},
	//	{{{Z, POS}, {Y, POS}, {X, NEG}}},
	func(p Point3D) Point3D {
		return Point3D{p.Z, p.Y, -p.X}
	},
	//  {{{Z, POS}, {X, NEG}, {Y, NEG}}},
	func(p Point3D) Point3D {
		return Point3D{p.Z, -p.X, -p.Y}
	},
	//	{{{Z, POS}, {Y, NEG}, {X, POS}}},
	func(p Point3D) Point3D {
		return Point3D{p.Z, -p.Y, p.X}
	},
	//  {{{Z, NEG}, {X, POS}, {Y, NEG}}},
	func(p Point3D) Point3D {
		return Point3D{-p.Z, p.X, -p.Y}
	},
	//	{{{Z, NEG}, {Y, NEG}, {X, NEG}}},
	func(p Point3D) Point3D {
		return Point3D{-p.Z, -p.Y, -p.X}
	},
	//  {{{Z, NEG}, {X, NEG}, {Y, POS}}},
	func(p Point3D) Point3D {
		return Point3D{-p.Z, -p.X, p.Y}
	},
	//	{{{Z, NEG}, {Y, POS}, {X, POS}}},
	func(p Point3D) Point3D {
		return Point3D{-p.Z, p.Y, p.X}
	},
}

func (p Point3D) Sub(o Point3D) Point3D {
	return Point3D{
		X: p.X - o.X,
		Y: p.Y - o.Y,
		Z: p.Z - o.Z,
	}
}

func (p Point3D) Add(o Point3D) Point3D {
	return Point3D{
		X: p.X + o.X,
		Y: p.Y + o.Y,
		Z: p.Z + o.Z,
	}
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func (p Point3D) Dist(o Point3D) int {
	d := p.Sub(o)
	return abs(d.X) + abs(d.Y) + abs(d.Z)
}
