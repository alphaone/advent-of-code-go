package vectors_test

import (
	"math"
	"slices"
	"testing"

	"github.com/alphaone/advent/y2019/day10/vectors"
	"github.com/stretchr/testify/assert"
)

func TestProportional(t *testing.T) {
	assert.True(t, vectors.New(1, 2).ProportionalTo(vectors.New(2, 4)))
	assert.True(t, vectors.New(3, 0).ProportionalTo(vectors.New(4, 0)))
	assert.True(t, vectors.New(0, 2).ProportionalTo(vectors.New(0, 3)))
}

func TestReduce(t *testing.T) {
	assert.Equal(t, vectors.New(0, 1), vectors.New(0, 4).Reduce())
	assert.Equal(t, vectors.New(1, 2), vectors.New(2, 4).Reduce())
	assert.Equal(t, vectors.New(5, 2), vectors.New(25, 10).Reduce())

	assert.Equal(t, vectors.New(-1, 0), vectors.New(-4, 0).Reduce())
}

func TestOrdering(t *testing.T) {
	expected := []vectors.Vector{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}
	actual := []vectors.Vector{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{-1, 1},
		{1, -1},
		{-1, -1},
	}
	slices.SortFunc(actual, vectors.SortClockwiseFunc)
	assert.Equal(t, expected, actual)
}

func TestAngle(t *testing.T) {
	assert.InDelta(t, 0.0, vectors.New(1, 0).Angle(vectors.New(1, 0)), 0.0001)
	assert.InDelta(t, 0.0, vectors.New(1, 0).Angle(vectors.New(2, 0)), 0.0001)
	assert.InDelta(t, math.Pi/2, vectors.New(1, 0).Angle(vectors.New(0, 1)), 0.0001)
	assert.InDelta(t, math.Pi, vectors.New(1, 0).Angle(vectors.New(-1, 0)), 0.0001)
	assert.InDelta(t, 3*math.Pi/2, vectors.New(1, 0).Angle(vectors.New(0, -1)), 0.0001)
	assert.InDelta(t, 2*math.Pi, vectors.New(1, 0).Angle(vectors.New(1000, -1)), 0.01)

	assert.InDelta(t, 0.0, vectors.New(0, -1).Angle(vectors.New(0, -1)), 0.0001)
}
