package galx

import (
	"fmt"
	"math"
)

type Line struct {
	A Vec2d
	B Vec2d
}

func (line Line) String() string {
	return fmt.Sprintf("Line{[%.4f, %.4f] -> [%.4f, %.4f]}", line.A.X, line.A.Y, line.B.X, line.B.Y)
}

func (line Line) Center() Vec2d {
	return line.A.Add(line.B.Sub(line.A).Scale(0.5))
}

func (line Line) Length() float64 {
	return line.B.Sub(line.A).Magnitude()
}

func (line Line) Normalize() Line {
	return Line{
		A: Vec2d{
			X: math.Min(line.A.X, line.B.X),
			Y: math.Min(line.A.Y, line.B.Y),
		},
		B: Vec2d{
			X: math.Max(line.A.X, line.B.X),
			Y: math.Max(line.A.Y, line.B.Y),
		},
	}
}
