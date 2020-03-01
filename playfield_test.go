package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayfieldCap(t *testing.T) {
	pf := NewPlayfield(16, 10)
	r, c := pf.Cap()

	assert.Equal(t, r, 16)
	assert.Equal(t, c, 10)
}

func TestPlayfieldSpawn(t *testing.T) {
	pf := NewPlayfield(16, 10)
	err := pf.Spawn(func(l int) int { return 0 })
	assert.NoError(t, err)
	assert.IsType(t, &ITetromino{}, pf.current.obj)
}

func TestPlayfieldSpawnUpperBound(t *testing.T) {
	pf := NewPlayfield(16, 10)
	var ll spawner = func(l int) int {
		return l - 1
	}
	err := pf.Spawn(ll)
	assert.NoError(t, err)
	assert.IsType(t, &LTetromino{}, pf.current.obj)
}

func TestPlayfieldSpawnOutOfBounds(t *testing.T) {
	pf := NewPlayfield(16, 10)
	err := pf.Spawn(func(l int) int { return 100 })
	assert.EqualError(t, err, "value out of bounds")
	assert.Nil(t, pf.current)
}

func TestPlayfieldSpawnOutOfBoundsNegative(t *testing.T) {
	pf := NewPlayfield(16, 10)
	err := pf.Spawn(func(l int) int { return -1 })
	assert.EqualError(t, err, "value out of bounds")
	assert.Nil(t, pf.current)
}

func TestPlayfieldMoveDown(t *testing.T) {
	pf := NewPlayfield(16, 10)
	err := pf.Spawn(func(l int) int { return 0 })

	assert.NoError(t, err)

	pf.Move(0, 1)
	assert.Equal(t, 1, pf.current.y)
}

func TestPlayfieldMoveDownOutOfBounds(t *testing.T) {
	var err error

	pf := NewPlayfield(16, 10)
	err = pf.Spawn(func(l int) int { return 0 })
	assert.NoError(t, err)
	prevY := pf.current.y

	err = pf.Move(0, 32)
	assert.EqualError(t, err, "value out of bounds")
	assert.Equal(t, prevY, pf.current.y)
}
