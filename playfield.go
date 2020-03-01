package main

import "fmt"

type spawner func(l int) int

// Playfield is the grid into which tetrominoes fall
type Playfield struct {
	store   [][]Mino
	current *CurrentTetromino
	Shapes  []string
}

// CurrentTetromino location and pointer
type CurrentTetromino struct {
	x   int
	y   int
	obj Tetromino
}

type Coordinates struct {
	x int
	y int
}

// NewPlayfield generator
func NewPlayfield(rows int, cols int) *Playfield {
	m := make([][]Mino, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]Mino, cols)
	}

	pf := &Playfield{
		store:  m,
		Shapes: []string{"I", "O", "T", "S", "Z", "J", "L"},
	}
	return pf
}

// Cap of Playfield
func (pf *Playfield) Cap() (rows int, cols int) {
	rows = cap(pf.store)
	cols = cap(pf.store[0])

	return rows, cols
}

// Spawn a Tetromino in the playfield;
// spawner function can return a static or dynamic (random) value
func (pf *Playfield) Spawn(spwn spawner) error {
	var obj Tetromino
	i := spwn(len(pf.Shapes))

	if i < 0 || i >= len(pf.Shapes) {
		return fmt.Errorf("value out of bounds")
	}

	switch pf.Shapes[i] {
	case "I":
		obj = &ITetromino{}
	case "O":
		obj = &OTetromino{}
	case "T":
		obj = &TTetromino{}
	case "S":
		obj = &STetromino{}
	case "Z":
		obj = &ZTetromino{}
	case "J":
		obj = &JTetromino{}
	case "L":
		obj = &LTetromino{}
	}

	pf.current = &CurrentTetromino{
		x:   0,
		y:   0,
		obj: obj,
	}

	return nil
}

func (pf *Playfield) Move(x, y int) error {
	if pf.outOfBounds(pf.current.x+x, pf.current.y+y) {
		return fmt.Errorf("value out of bounds")
	}
	pf.current.x = x
	pf.current.y = y
	return nil
}

func (pf *Playfield) outOfBounds(x, y int) bool {
	r, c := pf.Cap()
	if (x > r || y > c) || (x < 0 || y < 0) {
		return true
	}

	return false
}

func (pf *Playfield) Rotate() error {
	return nil
}
