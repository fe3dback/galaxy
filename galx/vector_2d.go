package galx

import (
	"fmt"
	"math"
)

// SizeOfVec2 its size for low precision memory data dump (float32)
// dump have size of 8 bytes (x=4 + y=4)
const SizeOfVec2 = 8

// Vec2d is common vector data structure
type Vec2d struct {
	X, Y float64
}

func (v *Vec2d) String() string {
	return fmt.Sprintf("Vec2d{%.4f, %.4f}", v.X, v.Y)
}

// Data dump for low precision memory representation (GPU, shaders, etc..)
func (v *Vec2d) Data() []byte {
	var buf [SizeOfVec2]byte
	nX := math.Float32bits(float32(v.X))
	nY := math.Float32bits(float32(v.Y))

	buf[0] = byte(nX)
	buf[1] = byte(nX >> 8)
	buf[2] = byte(nX >> 16)
	buf[3] = byte(nX >> 24)

	buf[4] = byte(nY)
	buf[5] = byte(nY >> 8)
	buf[6] = byte(nY >> 16)
	buf[7] = byte(nY >> 24)

	return buf[:]
}

// =============================================
// Constructors
// =============================================

func Vector2dForward(y float64) Vec2d {
	return Vec2d{
		X: 0,
		Y: -y,
	}
}

func Vector2dRight(x float64) Vec2d {
	return Vec2d{
		X: x,
		Y: 0,
	}
}

func Vector2dTowards(a Angle) Vec2d {
	return Vec2d{
		X: math.Cos(a.Radians()),
		Y: -math.Sin(a.Radians()),
	}
}

// =============================================
// Simple Math
// =============================================

func (v Vec2d) Add(n Vec2d) Vec2d {
	return Vec2d{
		X: v.X + n.X,
		Y: v.Y + n.Y,
	}
}

func (v Vec2d) Sub(n Vec2d) Vec2d {
	return Vec2d{
		X: v.X - n.X,
		Y: v.Y - n.Y,
	}
}

func (v Vec2d) Mul(t Vec2d) Vec2d {
	return Vec2d{
		X: v.X * t.X,
		Y: v.Y * t.Y,
	}
}

func (v Vec2d) Div(t Vec2d) Vec2d {
	return Vec2d{
		X: v.X / t.X,
		Y: v.Y / t.Y,
	}
}

// =============================================
// Advanced Math
// =============================================

func (v Vec2d) Plus(n float64) Vec2d {
	return Vec2d{
		X: v.X + n,
		Y: v.Y + n,
	}
}

func (v Vec2d) Minus(n float64) Vec2d {
	return Vec2d{
		X: v.X - n,
		Y: v.Y - n,
	}
}

func (v Vec2d) Scale(n float64) Vec2d {
	return Vec2d{
		X: v.X * n,
		Y: v.Y * n,
	}
}

func (v Vec2d) Decrease(n float64) Vec2d {
	return Vec2d{
		X: v.X / n,
		Y: v.Y / n,
	}
}

func (v Vec2d) Cross(to Vec2d) float64 {
	return v.X*to.Y - v.Y*to.X
}

func (v Vec2d) Dot(to Vec2d) float64 {
	return v.X*to.X + v.Y*to.Y
}

// =============================================
// Simple properties
// =============================================

func (v Vec2d) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2d) RoundX() int {
	return int(math.Floor(v.X))
}

func (v Vec2d) RoundY() int {
	return int(math.Floor(v.Y))
}

func (v Vec2d) Zero() bool {
	return v.X == 0.0 && v.Y == 0.0
}

// =============================================
// Transforms
// =============================================

func (v Vec2d) Normalize() Vec2d {
	m := v.Magnitude()

	if m == 0 {
		return Vec2d{}
	}

	return Vec2d{
		X: v.X / m,
		Y: v.Y / m,
	}
}

func (v Vec2d) Round() Vec2d {
	return Vec2d{
		X: math.Round(v.X),
		Y: math.Round(v.Y),
	}
}

func (v Vec2d) Clamp(min, max float64) Vec2d {
	return Vec2d{
		X: Clamp(v.X, min, max),
		Y: Clamp(v.Y, min, max),
	}
}

func (v Vec2d) ClampAbs(to Vec2d) Vec2d {
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

func (v Vec2d) DistanceTo(to Vec2d) float64 {
	return math.Sqrt(
		(v.X-to.X)*(v.X-to.X) +
			(v.Y-to.Y)*(v.Y-to.Y),
	)
}

func (v Vec2d) Direction() Angle {
	return Angle(math.Atan2(-v.Y, v.X))
}

func (v Vec2d) AngleBetween(to Vec2d) Angle {
	return Angle(math.Atan2(v.Cross(to), v.Dot(to)))
}

func (v Vec2d) AngleTo(to Vec2d) Angle {
	return Angle(math.Atan2(to.Y-v.Y, v.X-to.X) + math.Pi)
}

func (v Vec2d) Rotate(angle Angle) Vec2d {
	sin := math.Sin(angle.Radians())
	cos := math.Cos(angle.Radians())

	return Vec2d{
		X: v.X*cos - v.Y*sin,
		Y: -(v.X*sin + v.Y*cos),
	}
}

func (v Vec2d) RotateAround(orig Vec2d, angle Angle) Vec2d {
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

func (v Vec2d) PolarOffset(distance float64, angle Angle) Vec2d {
	return Vec2d{
		X: v.X + distance*math.Cos(angle.Radians()),
		Y: v.Y - distance*math.Sin(angle.Radians()),
	}
}

// =============================================
// Utils
// =============================================

func Vector2dSum(list ...Vec2d) Vec2d {
	res := Vec2d{}

	for _, v := range list {
		res.X += v.X
		res.Y += v.Y
	}

	return res
}
