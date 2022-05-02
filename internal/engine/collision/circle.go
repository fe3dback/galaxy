package collision

import (
	"github.com/fe3dback/galaxy/galx"
)

func Circle2Circle(a, b galx.Circle) bool {
	return a.Radius+b.Radius >= a.DistanceTo(b)
}

func Circle2Point(c galx.Circle, p galx.Vec2d) bool {
	return c.Radius >= c.Pos.DistanceTo(p)
}
