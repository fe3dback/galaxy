package collision

import "github.com/fe3dback/galaxy/engine"

func Rect2Rect(a, b engine.Rect) bool {
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

func Rect2Point(a engine.Rect, b engine.Vec) bool {
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

func Rect2Circle(r engine.Rect, c engine.Circle) bool {
	// fast rough check
	if !Rect2Rect(r, c.BoundingBox()) {
		return false
	}

	// slow check circle inside rect..
	if Rect2Point(r, c.Pos) {
		return true
	}

	// or rect edges touch or overlap circle
	for _, edge := range r.Edges() {
		if Ray2Circle(edge, c) {
			return true
		}
	}

	return false
}
