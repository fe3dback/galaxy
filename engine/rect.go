package engine

import (
	"fmt"
	"math"
)

type Rect struct {
	Min Vec
	Max Vec
}

type RectScreenX Rect

func RectScreen(x, y, w, h int) Rect {
	return Rect{
		Min: Vec{
			X: float64(x),
			Y: float64(y),
		},
		Max: Vec{
			X: float64(w),
			Y: float64(h),
		},
	}
}

func (r Rect) String() string {
	return fmt.Sprintf("Rect{%.4f, %.4f, %.4f, %.4f}", r.Min.X, r.Min.Y, r.Max.X, r.Max.Y)
}

func (r Rect) Normalize() Rect {
	return Rect{
		Min: Vec{
			X: math.Min(r.Min.X, r.Max.X),
			Y: math.Min(r.Min.Y, r.Max.Y),
		},
		Max: Vec{
			X: math.Max(r.Min.X, r.Max.X),
			Y: math.Max(r.Min.Y, r.Max.Y),
		},
	}
}

func (r Rect) Screen() Rect {
	// todo: RectScreenX
	if r.Max.X < 0 {
		width := math.Abs(r.Max.X)

		// reset width
		r.Min.X -= width
		r.Max.X = width
	}

	if r.Max.Y < 0 {
		height := math.Abs(r.Max.Y)

		// reset width
		r.Min.Y -= height
		r.Max.Y = height
	}

	return Rect{
		Min: r.Min,
		Max: r.Max,
	}
}

func (r Rect) Width() float64 {
	return r.Max.X - r.Min.X
}

func (r Rect) Height() float64 {
	return r.Max.Y - r.Min.Y
}

func (r Rect) Edges() [4]Line {
	corners := r.Vertices()

	return [4]Line{
		{A: corners[0], B: corners[1]},
		{A: corners[1], B: corners[2]},
		{A: corners[2], B: corners[3]},
		{A: corners[3], B: corners[0]},
	}
}

func (r Rect) Vertices() [4]Vec {
	return [4]Vec{
		r.Min,
		{
			X: r.Min.X,
			Y: r.Max.Y,
		},
		r.Max,
		{
			X: r.Max.X,
			Y: r.Min.Y,
		},
	}
}
