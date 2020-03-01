package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestITetrominoShape(t *testing.T) {
	out := `
0001
0001
0001
0001`
	s := NewITetromino()
	out2 := shapeToString(s.Shape())
	assert.Equal(t, strings.TrimSpace(out), out2)
	assert.Equal(t, &Size{w: 1, h: 4}, s.Size())
}

func TestOTetrominoShape(t *testing.T) {
	out := `
0011
0011
0000
0000
`
	s := NewOTetromino()
	out2 := shapeToString(s.Shape())
	assert.Equal(t, strings.TrimSpace(out), out2)
	assert.Equal(t, &Size{w: 2, h: 2}, s.Size())
}

func TestTTetrominoShape(t *testing.T) {
	out := `
0010
0111
0000
0000
`
	s := NewTTetromino()
	out2 := shapeToString(s.Shape())
	assert.Equal(t, strings.TrimSpace(out), out2)
	assert.Equal(t, &Size{w: 3, h: 2}, s.Size())
}

func TestSTetrominoShape(t *testing.T) {
	out := `
0011
0110
0000
0000
`
	s := NewSTetromino()
	out2 := shapeToString(s.Shape())
	assert.Equal(t, strings.TrimSpace(out), out2)
	assert.Equal(t, &Size{w: 3, h: 2}, s.Size())
}

func TestZTetrominoShape(t *testing.T) {
	out := `
0110
0011
0000
0000
`
	s := NewZTetromino()
	out2 := shapeToString(s.Shape())
	assert.Equal(t, strings.TrimSpace(out), out2)
	assert.Equal(t, &Size{w: 3, h: 2}, s.Size())
}

func TestJTetrominoShape(t *testing.T) {
	out := `
0001
0111
0000
0000
`
	s := NewJTetromino()
	out2 := shapeToString(s.Shape())
	assert.Equal(t, strings.TrimSpace(out), out2)
	assert.Equal(t, &Size{w: 3, h: 2}, s.Size())
}

func TestLTetrominoShape(t *testing.T) {
	out := `
0100
0111
0000
0000
`
	s := NewLTetromino()
	out2 := shapeToString(s.Shape())
	assert.Equal(t, strings.TrimSpace(out), out2)
	assert.Equal(t, &Size{w: 3, h: 2}, s.Size())
}

func shapeToString(s *Shape) string {
	out := make([]string, len(s.Definition))
	for i := 0; i < len(s.Definition); i++ {
		out[i] = fmt.Sprintf("%04b", s.Definition[i])
	}
	return strings.Join(out, "\n")
}
