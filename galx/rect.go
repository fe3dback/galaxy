package galx

import (
	"fmt"
	"math"
)

type Rect struct {
	Min Vec
	Max Vec
}

func (r Rect) String() string {
	return fmt.Sprintf("Rect{%.4f, %.4f, %.4f, %.4f}", r.Min.X, r.Min.Y, r.Max.X, r.Max.Y)
}

func (r Rect) Valid() bool {
	return r.Max.X >= r.Min.X && r.Max.Y >= r.Min.Y
}

func (r Rect) Normalize() Rect {
	if r.Valid() {
		return r
	}

	if r.Max.X < r.Min.X {
		r.Min.X, r.Max.X = r.Max.X, r.Min.X
	}

	if r.Max.Y < r.Min.Y {
		r.Min.Y, r.Max.Y = r.Max.Y, r.Min.Y
	}

	return r
}

func (r Rect) Width() float64 {
	return math.Abs(r.Max.X - r.Min.X)
}

func (r Rect) Height() float64 {
	return math.Abs(r.Max.Y - r.Min.Y)
}

func (r Rect) Center() Vec {
	return Vec{
		X: r.Min.X + ((r.Max.X - r.Min.X) / 2),
		Y: r.Min.Y + ((r.Max.Y - r.Min.Y) / 2),
	}
}

func (r Rect) Scale(s float64) Rect {
	r = r.Normalize()

	center := r.Center()
	wh := (r.Width() * s) / 2
	hh := (r.Height() * s) / 2

	return Rect{
		Min: Vec{
			X: center.X - wh,
			Y: center.Y - hh,
		},
		Max: Vec{
			X: center.X + wh,
			Y: center.Y + hh,
		},
	}
}

func (r Rect) Increase(size float64) Rect {
	r = r.Normalize()

	return Rect{
		Min: Vec{
			X: r.Min.X - size,
			Y: r.Min.Y - size,
		},
		Max: Vec{
			X: r.Max.X + size,
			Y: r.Max.Y + size,
		},
	}
}

func (r Rect) Contains(v Vec) bool {
	if v.X < r.Min.X {
		return false
	}
	if v.Y < r.Min.Y {
		return false
	}
	if v.X > r.Max.X {
		return false
	}
	if v.Y > r.Max.Y {
		return false
	}

	return true
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

func SurroundRect(boxes ...Rect) Rect {
	minX := float64(math.MaxInt32)
	minY := float64(math.MaxInt32)
	maxX := -float64(math.MaxInt32)
	maxY := -float64(math.MaxInt32)

	for _, box := range boxes {
		box = box.Normalize()
		
		if box.Min.X < minX {
			minX = box.Min.X
		}
		if box.Min.Y < minY {
			minY = box.Min.Y
		}
		if box.Max.X > maxX {
			maxX = box.Max.X
		}
		if box.Max.Y > maxY {
			maxY = box.Max.Y
		}
	}

	return Rect{
		Min: Vec{
			X: minX,
			Y: minY,
		},
		Max: Vec{
			X: maxX,
			Y: maxY,
		},
	}
}
