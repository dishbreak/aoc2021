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

func TestIntersection(t *testing.T) {
	type testCase struct {
		one, other, expected string
	}

	testCases := map[string]testCase{
		"overlapping": {
			one:      "on x=10..12,y=9..11,z=9..17",
			other:    "on x=11..12,y=10..12,z=7..14",
			expected: "off x=11..12,y=10..11,z=9..14",
		},
		"non overlapping": {
			one:      "on x=10..12,y=9..11,z=9..17",
			other:    "on x=11..12,y=14..19,z=7..14",
			expected: "(empty)",
		},
		"both off": {
			one:      "off x=10..12,y=9..11,z=9..17",
			other:    "off x=11..12,y=10..12,z=7..14",
			expected: "on x=11..12,y=10..11,z=9..14",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			o1, o2 := ToCuboid(tc.one), ToCuboid(tc.other)
			assert.Equal(t, tc.expected, Intersection(o1, o2).String())
		})
	}
}
