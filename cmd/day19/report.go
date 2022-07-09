package main

import (
	"strconv"
	"strings"
)

type Report struct {
	ID         int
	Beacons    []Point3D
	position   Point3D
	normalized []Point3D
}

func NewReport(input []string) *Report {
	r := &Report{}
	parts := strings.Fields(input[0])
	parsed, _ := strconv.Atoi(parts[2])
	r.ID = parsed

	r.Beacons = make([]Point3D, len(input)-1)

	for i, line := range input[1:] {
		r.Beacons[i] = NewPoint3DFromString(line)
	}

	return r
}

func (r *Report) Rotate(t Transform) []Point3D {
	result := make([]Point3D, len(r.Beacons))

	for i, pt := range r.Beacons {
		result[i] = t(pt)
	}

	return result
}
