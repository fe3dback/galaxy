package galx

import "math"

type Circle struct {
	Pos    Vec
	Radius float64
}

func (c Circle) BoundingBox() Rect {
	c.Radius = math.Abs(c.Radius)

	return Rect{
		Min: Vec{
			X: c.Pos.X - c.Radius,
			Y: c.Pos.Y - c.Radius,
		},
		Max: Vec{
			X: c.Pos.X + c.Radius,
			Y: c.Pos.Y + c.Radius,
		},
	}
}

func (c Circle) DistanceTo(to Circle) float64 {
	return c.Pos.DistanceTo(to.Pos)
}
