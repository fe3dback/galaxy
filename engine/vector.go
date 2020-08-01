package engine

import "math"

type Vector2D struct {
	X, Y float64
}

func (v Vector2D) Add(n Vector2D) Vector2D {
	return Vector2D{
		X: v.X + n.X,
		Y: v.Y + n.Y,
	}
}

func (v Vector2D) Mul(n float64) Vector2D {
	return Vector2D{
		X: v.X * n,
		Y: v.Y * n,
	}
}

func (v Vector2D) RoundX() int {
	return int(math.Floor(v.X))
}

func (v Vector2D) RoundY() int {
	return int(math.Floor(v.Y))
}
