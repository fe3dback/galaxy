package collision

import "github.com/fe3dback/galaxy/engine"

func Circle2Circle(a, b engine.Circle) bool {
	return a.Radius+b.Radius >= a.DistanceTo(b)
}

func Circle2Point(c engine.Circle, p engine.Vec) bool {
	return c.Radius >= c.Pos.DistanceTo(p)
}
