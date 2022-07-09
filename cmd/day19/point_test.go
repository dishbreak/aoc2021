package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDist(t *testing.T) {
	one := Point3D{1105, -1205, 1229}
	other := Point3D{-92, -2380, -20}

	assert.Equal(t, 3621, one.Dist(other))
}
