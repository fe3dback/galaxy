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
	// todo tests
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
