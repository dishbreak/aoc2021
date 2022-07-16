package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Point3D struct {
	X, Y, Z int
}

type Cuboid struct {
	Min      Point3D
	Max      Point3D
	On       bool
	nonEmpty bool
}

func (c Cuboid) Volume() int64 {
	mult := -1
	if c.On {
		mult = 1
	}
	return int64(mult * (c.Max.X - c.Min.X + 1) * (c.Max.Y - c.Min.Y + 1) * (c.Max.Z - c.Min.Z + 1))
}

func (c Cuboid) Empty() bool {
	return !c.nonEmpty
}

func getMinMax(input string) (min, max int) {
	parts := strings.Split(input, "..")
	min, _ = strconv.Atoi(parts[0])
	max, _ = strconv.Atoi(parts[1])
	return
}

func getDims(input string) (pMin, pMax Point3D) {
	parts := strings.Split(input, ",")
	pMin.X, pMax.X = getMinMax(parts[0][2:])
	pMin.Y, pMax.Y = getMinMax(parts[1][2:])
	pMin.Z, pMax.Z = getMinMax(parts[2][2:])
	return
}

func ToCuboid(input string) (c Cuboid) {
	parts := strings.Fields(input)
	if parts[0] == "on" {
		c.On = true
	}
	c.nonEmpty = true
	c.Min, c.Max = getDims(parts[1])
	return
}

func min(one, other int) int {
	if one < other {
		return one
	}
	return other
}

func max(one, other int) int {
	if one > other {
		return one
	}
	return other
}

func (c Point3D) LessThan(other Point3D) bool {
	if c.X > other.X || c.Y > other.Y || c.Z > other.Z {
		return false
	}
	return true
}

func Intersection(one, other Cuboid) (c Cuboid) {
	c.Min = Point3D{
		X: max(one.Min.X, other.Min.X),
		Y: max(one.Min.Y, other.Min.Y),
		Z: max(one.Min.Z, other.Min.Z),
	}
	c.Max = Point3D{
		X: min(one.Max.X, other.Max.X),
		Y: min(one.Max.Y, other.Max.Y),
		Z: min(one.Max.Z, other.Max.Z),
	}

	if !c.Min.LessThan(c.Max) {
		c = Cuboid{}
		return
	}

	c.On = !one.On
	c.nonEmpty = true
	return
}

func (c Cuboid) String() string {
	if c.Empty() {
		return "(empty)"
	}
	state := "off"
	if c.On {
		state = "on"
	}
	return fmt.Sprintf("%s x=%d..%d,y=%d..%d,z=%d..%d",
		state, c.Min.X, c.Max.X, c.Min.Y, c.Max.Y, c.Min.Z, c.Max.Z,
	)
}
