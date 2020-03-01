package main

import (
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntMax(t *testing.T) {
	assert.Equal(t, 5, intMax(3, 5))
	assert.Equal(t, -5, intMax(-6, -5))
	assert.Equal(t, 0, intMax(0, 0))
}

func TestCountBits(t *testing.T) {
	assert.Equal(t, 0, bits.Len(0x0))
	assert.Equal(t, 3, bits.Len(0x7))
}
