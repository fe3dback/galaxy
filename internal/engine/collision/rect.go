package collision

import (
	"github.com/fe3dback/galaxy/galx"
)

func Rect2Rect(a, b galx.Rect) bool {
	if a.Max.X < b.Min.X {
		return false
	}

	if a.Min.X > b.Max.X {
		return false
	}

	if a.Max.Y < b.Min.Y {
		return false
	}

	if a.Min.Y > b.Max.Y {
		return false
	}

	return true
}

func Rect2Point(a galx.Rect, b galx.Vec) bool {
	if a.Min.X > b.X {
		return false
	}

	if a.Max.X < b.X {
		return false
	}

	if a.Min.Y > b.Y {
		return false
	}

	if a.Max.Y < b.Y {
		return false
	}

	return true
}

func Rect2Circle(r galx.Rect, c galx.Circle) bool {
	test := c.Pos

	if c.Pos.X < r.Min.X {
		test.X = r.Min.X
	} else if c.Pos.X > r.Max.X {
		test.X = r.Max.X
	}

	if c.Pos.Y < r.Min.Y {
		test.Y = r.Min.Y
	} else if c.Pos.Y > r.Max.Y {
		test.Y = r.Max.Y
	}

	distance := c.Pos.Sub(test).Magnitude()
	return distance <= c.Radius
}
