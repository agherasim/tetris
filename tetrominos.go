package main

type Size struct {
	w int
	h int
}

type Shape struct {
	// definition of a 4x4 points and 3 colors / point
	Definition []byte
	Color      []int
}

type Tetromino interface {
	Shape() *Shape
	Size() *Size
}

type BaseTetromino struct {
	shape *Shape
	size  *Size
}

type ITetromino struct {
	*BaseTetromino
}

type OTetromino struct {
	*BaseTetromino
}

type TTetromino struct {
	*BaseTetromino
}

type STetromino struct {
	*BaseTetromino
}

type ZTetromino struct {
	*BaseTetromino
}

type JTetromino struct {
	*BaseTetromino
}

type LTetromino struct {
	*BaseTetromino
}

func (bt *BaseTetromino) Size() *Size {
	var w, h int = 0, 0

	if bt.size == nil {
		for i := 0; i < len(bt.shape.Definition); i++ {
			oc := bitsLen(bt.shape.Definition[i])
			w = intMax(w, oc)

			if oc > 0 {
				h++
			}
		}

		bt.size = &Size{
			w: w,
			h: h,
		}
	}

	return bt.size
}

func (bt *BaseTetromino) Shape() *Shape {
	return bt.shape
}

func NewITetromino() Tetromino {
	s := &Shape{}
	s.Definition = []byte{0x1, 0x1, 0x1, 0x1}
	s.Color = []int{0, 240, 240}

	bt := &BaseTetromino{
		shape: s,
	}

	t := &ITetromino{
		bt,
	}

	return t
}

func NewOTetromino() Tetromino {
	s := &Shape{}
	s.Definition = []byte{0x3, 0x3, 0x0, 0x0}
	s.Color = []int{240, 240, 0}

	bt := &BaseTetromino{
		shape: s,
	}

	t := &OTetromino{
		bt,
	}
	return t
}

func NewTTetromino() Tetromino {
	s := &Shape{}
	s.Definition = []byte{0x2, 0x7, 0x0, 0x0}
	s.Color = []int{160, 28, 240}

	bt := &BaseTetromino{
		shape: s,
	}

	t := &TTetromino{
		bt,
	}
	return t
}

func NewSTetromino() Tetromino {
	s := &Shape{}
	s.Definition = []byte{0x3, 0x6, 0x0, 0x0}
	s.Color = []int{0, 240, 0}

	bt := &BaseTetromino{
		shape: s,
	}

	t := &STetromino{
		bt,
	}
	return t
}

func NewZTetromino() Tetromino {
	s := &Shape{}
	s.Definition = []byte{0x6, 0x3, 0x0, 0x0}
	s.Color = []int{240, 0, 0}

	bt := &BaseTetromino{
		shape: s,
	}

	t := &ZTetromino{
		bt,
	}
	return t
}

func NewJTetromino() Tetromino {
	s := &Shape{}
	s.Definition = []byte{0x1, 0x7, 0x0, 0x0}
	s.Color = []int{240, 160, 0}

	bt := &BaseTetromino{
		shape: s,
	}

	t := &JTetromino{
		bt,
	}
	return t
}

func NewLTetromino() Tetromino {
	s := &Shape{}
	s.Definition = []byte{0x4, 0x7, 0x0, 0x0}
	s.Color = []int{0, 28, 240}

	bt := &BaseTetromino{
		shape: s,
	}

	t := &LTetromino{
		bt,
	}
	return t
}
