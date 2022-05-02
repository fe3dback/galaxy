package collision

import (
	"math"

	"github.com/fe3dback/galaxy/galx"
)

func Line2Line(l1, l2 galx.Line) bool {
	x1 := l1.A.X
	x2 := l1.B.X
	x3 := l2.A.X
	x4 := l2.B.X

	y1 := l1.A.Y
	y2 := l1.B.Y
	y3 := l2.A.Y
	y4 := l2.B.Y

	uA := ((x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)) / ((y4-y3)*(x2-x1) - (x4-x3)*(y2-y1))
	uB := ((x2-x1)*(y1-y3) - (y2-y1)*(x1-x3)) / ((y4-y3)*(x2-x1) - (x4-x3)*(y2-y1))

	if uA >= 0 && uA <= 1 && uB >= 0 && uB <= 1 {
		return true
	}

	if math.IsNaN(uA) || math.IsNaN(uB) {
		return true
	}

	return false
}

func Line2Point(line galx.Line, p galx.Vec2d) bool {
	d1 := line.A.DistanceTo(p)
	d2 := line.B.DistanceTo(p)
	lineLen := line.A.Sub(line.B).Magnitude()

	if d1+d2 >= lineLen-0.000001 && d1+d2 <= lineLen+0.000001 {
		return true
	}

	return false
}

func Line2Circle(line galx.Line, circle galx.Circle) bool {
	if Circle2Point(circle, line.A) || Circle2Point(circle, line.B) {
		return true
	}

	dist := line.A.Sub(line.B)
	dot := (((circle.Pos.X - line.A.X) * (line.B.X - line.A.X)) +
		((circle.Pos.Y - line.A.Y) * (line.B.Y - line.A.Y))) / math.Pow(dist.Magnitude(), 2)

	closest := galx.Vec2d{
		X: line.A.X + (dot * (line.B.X - line.A.X)),
		Y: line.A.Y + (dot * (line.B.Y - line.A.Y)),
	}

	onSegment := Line2Point(line, closest)
	if !onSegment {
		return false
	}

	distX := closest.X - circle.Pos.X
	distY := closest.Y - circle.Pos.Y

	return math.Sqrt((distX*distX)+(distY*distY)) <= circle.Radius
}
