package collision

import (
	"github.com/fe3dback/galaxy/galx"
)

func Rect2Rect(a, b galx.Rect) bool {
	if a.BR.X < b.TL.X {
		return false
	}

	if a.TL.X > b.BR.X {
		return false
	}

	if a.BR.Y < b.TL.Y {
		return false
	}

	if a.TL.Y > b.BR.Y {
		return false
	}

	return true
}

func Rect2Point(a galx.Rect, b galx.Vec2d) bool {
	if a.TL.X > b.X {
		return false
	}

	if a.BR.X < b.X {
		return false
	}

	if a.TL.Y > b.Y {
		return false
	}

	if a.BR.Y < b.Y {
		return false
	}

	return true
}

func Rect2Circle(r galx.Rect, c galx.Circle) bool {
	test := c.Pos

	if c.Pos.X < r.TL.X {
		test.X = r.TL.X
	} else if c.Pos.X > r.BR.X {
		test.X = r.BR.X
	}

	if c.Pos.Y < r.TL.Y {
		test.Y = r.TL.Y
	} else if c.Pos.Y > r.BR.Y {
		test.Y = r.BR.Y
	}

	distance := c.Pos.Sub(test).Magnitude()
	return distance <= c.Radius
}
