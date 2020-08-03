package engine

import "math"

const angleMod = 180.0

type Vector2D struct {
	X, Y float64
}

func VectorSum(list ...Vector2D) Vector2D {
	res := Vector2D{}

	for _, v := range list {
		res.X += v.X
		res.Y += v.Y
	}

	return res
}

func VectorForward(y float64) Vector2D {
	return Vector2D{
		X: 0,
		Y: -y,
	}
}

func VectorLeft(x float64) Vector2D {
	return Vector2D{
		X: -x,
		Y: 0,
	}
}

func VectorTowards(a Angle) Vector2D {
	rad := Radian(a)

	return Vector2D{
		X: floatPrecision(math.Cos(rad)),
		Y: -floatPrecision(math.Sin(rad)),
	}
}

func Radian(angle Angle) float64 {
	return angle.ToFloat() * math.Pi / angleMod
}

func Degrees(rad float64) float64 {
	return rad * angleMod / math.Pi
}

func (v Vector2D) Add(n Vector2D) Vector2D {
	return Vector2D{
		X: v.X + n.X,
		Y: v.Y + n.Y,
	}
}

func (v Vector2D) MulTo(t Vector2D) Vector2D {
	return Vector2D{
		X: v.X * t.X,
		Y: v.Y * t.Y,
	}
}

func (v Vector2D) Mul(n float64) Vector2D {
	return Vector2D{
		X: v.X * n,
		Y: v.Y * n,
	}
}

func (v Vector2D) Divide(n float64) Vector2D {
	return Vector2D{
		X: v.X / n,
		Y: v.Y / n,
	}
}

func (v Vector2D) Normalize() Vector2D {
	m := v.Length()

	if m == 0 {
		return Vector2D{}
	}

	return Vector2D{
		X: v.X / m,
		Y: v.Y / m,
	}
}

func (v Vector2D) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2D) DistanceTo(to Vector2D) float64 {
	return math.Sqrt(
		(v.X-to.X)*(v.X-to.X) +
			(v.Y-to.Y)*(v.Y-to.Y),
	)
}

func (v Vector2D) Cross(to Vector2D) float64 {
	return v.X*to.Y - v.Y*to.X
}

func (v Vector2D) Dot(to Vector2D) float64 {
	return v.X*to.X + v.Y*to.Y
}

func (v Vector2D) Direction() Angle {
	return Anglef(Degrees(math.Atan2(-v.Y, v.X)))
}

func (v Vector2D) AngleBetween(to Vector2D) Angle {
	return Anglef(Degrees(math.Atan2(v.Cross(to), v.Dot(to))))
}

func (v Vector2D) AngleTo(to Vector2D) Angle {
	return v.Add(to).Direction()
}

func (v Vector2D) Rotate(angle Angle) Vector2D {
	rad := Radian(angle)
	sin := math.Sin(rad)
	cos := math.Cos(rad)

	return Vector2D{
		X: floatPrecision(v.X*cos - v.Y*sin),
		Y: -floatPrecision(v.X*sin + v.Y*cos),
	}
}

func (v Vector2D) RotateAround(orig Vector2D, angle Angle) Vector2D {
	rad := Radian(angle)
	sin := math.Sin(rad)
	cos := math.Cos(rad)

	v.X -= orig.X
	v.Y -= orig.Y

	xx := v.X*cos + v.Y*sin
	yy := v.X*sin - v.Y*cos

	v.X = floatPrecision(xx + orig.X)
	v.Y = -floatPrecision(yy + orig.Y)

	return v
}

func (v Vector2D) PolarOffset(distance float64, angle Angle) Vector2D {
	rad := Radian(angle)

	return Vector2D{
		X: v.X + floatPrecision(distance*math.Cos(rad)),
		Y: v.Y - floatPrecision(distance*math.Sin(rad)),
	}
}

func (v Vector2D) ToPoint() Point {
	return Point{
		X: v.RoundX(),
		Y: v.RoundY(),
	}
}

func (v Vector2D) RoundX() int {
	return int(math.Floor(v.X))
}

func (v Vector2D) RoundY() int {
	return int(math.Floor(v.Y))
}
