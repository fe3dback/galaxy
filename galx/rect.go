package galx

import (
	"fmt"
	"math"
)

type Rect struct {
	TL Vec
	BR Vec
}

func (r Rect) String() string {
	return fmt.Sprintf("Rect{%.4f, %.4f, %.4f, %.4f}", r.TL.X, r.TL.Y, r.BR.X, r.BR.Y)
}

func (r Rect) Valid() bool {
	return r.BR.X >= r.TL.X && r.BR.Y >= r.TL.Y
}

func (r Rect) Normalize() Rect {
	if r.Valid() {
		return r
	}

	if r.BR.X < r.TL.X {
		r.TL.X, r.BR.X = r.BR.X, r.TL.X
	}

	if r.BR.Y < r.TL.Y {
		r.TL.Y, r.BR.Y = r.BR.Y, r.TL.Y
	}

	return r
}

func (r Rect) Width() float64 {
	return math.Abs(r.BR.X - r.TL.X)
}

func (r Rect) Height() float64 {
	return math.Abs(r.BR.Y - r.TL.Y)
}

func (r Rect) Center() Vec {
	return Vec{
		X: r.TL.X + ((r.BR.X - r.TL.X) / 2),
		Y: r.TL.Y + ((r.BR.Y - r.TL.Y) / 2),
	}
}

func (r Rect) Scale(s float64) Rect {
	r = r.Normalize()

	center := r.Center()
	wh := (r.Width() * s) / 2
	hh := (r.Height() * s) / 2

	return Rect{
		TL: Vec{
			X: center.X - wh,
			Y: center.Y - hh,
		},
		BR: Vec{
			X: center.X + wh,
			Y: center.Y + hh,
		},
	}
}

func (r Rect) Increase(size float64) Rect {
	r = r.Normalize()

	return Rect{
		TL: Vec{
			X: r.TL.X - size,
			Y: r.TL.Y - size,
		},
		BR: Vec{
			X: r.BR.X + size,
			Y: r.BR.Y + size,
		},
	}
}

func (r Rect) Contains(v Vec) bool {
	if v.X < r.TL.X {
		return false
	}
	if v.Y < r.TL.Y {
		return false
	}
	if v.X > r.BR.X {
		return false
	}
	if v.Y > r.BR.Y {
		return false
	}

	return true
}

func (r Rect) Edges() [4]Line {
	corners := r.VerticesClockWise()

	return [4]Line{
		{A: corners[0], B: corners[1]},
		{A: corners[1], B: corners[2]},
		{A: corners[2], B: corners[3]},
		{A: corners[3], B: corners[0]},
	}
}

func (r Rect) VerticesClockWise() [4]Vec {
	return [4]Vec{
		r.TL,
		{
			X: r.BR.X,
			Y: r.TL.Y,
		},
		r.BR,
		{
			X: r.TL.X,
			Y: r.BR.Y,
		},
	}
}

func SurroundRect(boxes ...Rect) Rect {
	minX := float64(math.MaxInt32)
	minY := float64(math.MaxInt32)
	maxX := -float64(math.MaxInt32)
	maxY := -float64(math.MaxInt32)

	for _, box := range boxes {
		box = box.Normalize()

		if box.TL.X < minX {
			minX = box.TL.X
		}
		if box.TL.Y < minY {
			minY = box.TL.Y
		}
		if box.BR.X > maxX {
			maxX = box.BR.X
		}
		if box.BR.Y > maxY {
			maxY = box.BR.Y
		}
	}

	return Rect{
		TL: Vec{
			X: minX,
			Y: minY,
		},
		BR: Vec{
			X: maxX,
			Y: maxY,
		},
	}
}
