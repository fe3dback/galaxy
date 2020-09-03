package engine

import (
	"fmt"
	"math"
)

type Vec struct {
	X, Y float64
}

func (v Vec) String() string {
	return fmt.Sprintf("Vec{%.4f, %.4f}", v.X, v.Y)
}

// =============================================
// Constructors
// =============================================

func VectorForward(y float64) Vec {
	return Vec{
		X: 0,
		Y: -y,
	}
}

func VectorRight(x float64) Vec {
	return Vec{
		X: x,
		Y: 0,
	}
}

func VectorTowards(a Angle) Vec {
	return Vec{
		X: math.Cos(a.Radians()),
		Y: -math.Sin(a.Radians()),
	}
}

// =============================================
// Simple Math
// =============================================

func (v Vec) Add(n Vec) Vec {
	return Vec{
		X: v.X + n.X,
		Y: v.Y + n.Y,
	}
}

func (v Vec) Sub(n Vec) Vec {
	return Vec{
		X: v.X - n.X,
		Y: v.Y - n.Y,
	}
}

func (v Vec) Mul(t Vec) Vec {
	return Vec{
		X: v.X * t.X,
		Y: v.Y * t.Y,
	}
}

func (v Vec) Div(t Vec) Vec {
	return Vec{
		X: v.X / t.X,
		Y: v.Y / t.Y,
	}
}

// =============================================
// Advanced Math
// =============================================

func (v Vec) Plus(n float64) Vec {
	return Vec{
		X: v.X + n,
		Y: v.Y + n,
	}
}

func (v Vec) Minus(n float64) Vec {
	return Vec{
		X: v.X - n,
		Y: v.Y - n,
	}
}

func (v Vec) Scale(n float64) Vec {
	return Vec{
		X: v.X * n,
		Y: v.Y * n,
	}
}

func (v Vec) Decrease(n float64) Vec {
	return Vec{
		X: v.X / n,
		Y: v.Y / n,
	}
}

func (v Vec) Cross(to Vec) float64 {
	return v.X*to.Y - v.Y*to.X
}

func (v Vec) Dot(to Vec) float64 {
	return v.X*to.X + v.Y*to.Y
}

// =============================================
// Simple properties
// =============================================

func (v Vec) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec) RoundX() int {
	return int(math.Floor(v.X))
}

func (v Vec) RoundY() int {
	return int(math.Floor(v.Y))
}

func (v Vec) Zero() bool {
	return v.X == 0.0 && v.Y == 0.0
}

// =============================================
// Transforms
// =============================================

func (v Vec) Normalize() Vec {
	m := v.Magnitude()

	if m == 0 {
		return Vec{}
	}

	return Vec{
		X: v.X / m,
		Y: v.Y / m,
	}
}

func (v Vec) Round() Vec {
	return Vec{
		X: math.Round(v.X),
		Y: math.Round(v.Y),
	}
}

func (v Vec) Clamp(min, max float64) Vec {
	return Vec{
		X: Clamp(v.X, min, max),
		Y: Clamp(v.Y, min, max),
	}
}

func (v Vec) ClampAbs(to Vec) Vec {
	absX := math.Abs(to.X)
	absY := math.Abs(to.Y)

	if v.X > absX {
		v.X = absX
	} else if v.X < -absX {
		v.X = -absX
	}

	if v.Y > absY {
		v.Y = absY
	} else if v.Y < -absY {
		v.Y = -absY
	}

	return v
}

// =============================================
// Trigonometry
// =============================================

func (v Vec) DistanceTo(to Vec) float64 {
	return math.Sqrt(
		(v.X-to.X)*(v.X-to.X) +
			(v.Y-to.Y)*(v.Y-to.Y),
	)
}

func (v Vec) Direction() Angle {
	return Angle(math.Atan2(-v.Y, v.X))
}

func (v Vec) AngleBetween(to Vec) Angle {
	return Angle(math.Atan2(v.Cross(to), v.Dot(to)))
}

func (v Vec) AngleTo(to Vec) Angle {
	return Angle(math.Atan2(to.Y-v.Y, v.X-to.X) + math.Pi)
}

func (v Vec) Rotate(angle Angle) Vec {
	sin := math.Sin(angle.Radians())
	cos := math.Cos(angle.Radians())

	return Vec{
		X: v.X*cos - v.Y*sin,
		Y: -(v.X*sin + v.Y*cos),
	}
}

func (v Vec) RotateAround(orig Vec, angle Angle) Vec {
	sin := math.Sin(angle.Radians())
	cos := math.Cos(angle.Radians())

	v.X -= orig.X
	v.Y -= orig.Y

	xx := v.X*cos + v.Y*sin
	yy := -(v.X*sin - v.Y*cos)

	v.X = xx + orig.X
	v.Y = yy + orig.Y

	return v
}

func (v Vec) PolarOffset(distance float64, angle Angle) Vec {
	return Vec{
		X: v.X + distance*math.Cos(angle.Radians()),
		Y: v.Y - distance*math.Sin(angle.Radians()),
	}
}

// =============================================
// Utils
// =============================================

func VectorSum(list ...Vec) Vec {
	res := Vec{}

	for _, v := range list {
		res.X += v.X
		res.Y += v.Y
	}

	return res
}
