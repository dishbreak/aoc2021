package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCuboid(t *testing.T) {
	input := "on x=10..12,y=9..13,z=11..17"
	expected := Cuboid{
		Min: Point3D{10, 9, 11},
		Max: Point3D{12, 13, 17},
		On:  true,
	}
	assert.Equal(t, expected, ToCuboid(input))
}
