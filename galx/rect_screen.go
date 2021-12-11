package galx

import "math"

func (r Rect) Screen() Rect {
	if r.Max.X < 0 {
		width := math.Abs(r.Max.X)

		// reset width
		r.Min.X -= width
		r.Max.X = width
	}

	if r.Max.Y < 0 {
		height := math.Abs(r.Max.Y)

		// reset width
		r.Min.Y -= height
		r.Max.Y = height
	}

	return Rect{
		Min: r.Min,
		Max: r.Max,
	}
}

func RectScreen(x, y, w, h int) Rect {
	return Rect{
		Min: Vec{
			X: float64(x),
			Y: float64(y),
		},
		Max: Vec{
			X: float64(w),
			Y: float64(h),
		},
	}.Screen()
}
