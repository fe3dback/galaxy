package galx

import "math"

type Circle struct {
	Pos    Vec2d
	Radius float64
}

func (c Circle) BoundingBox() Rect {
	c.Radius = math.Abs(c.Radius)

	return Rect{
		TL: Vec2d{
			X: c.Pos.X - c.Radius,
			Y: c.Pos.Y - c.Radius,
		},
		BR: Vec2d{
			X: c.Pos.X + c.Radius,
			Y: c.Pos.Y + c.Radius,
		},
	}
}

func (c Circle) Contains(p Vec2d) bool {
	return c.Radius >= c.Pos.DistanceTo(p)
}

func (c Circle) IncreaseRadius(r float64) Circle {
	return Circle{
		Pos:    c.Pos,
		Radius: c.Radius + r,
	}
}

func (c Circle) DistanceTo(to Circle) float64 {
	return c.Pos.DistanceTo(to.Pos)
}
