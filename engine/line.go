package engine

import "fmt"

type Line struct {
	A Vec
	B Vec
}

func (line Line) String() string {
	return fmt.Sprintf("Line{[%.4f, %.4f] -> [%.4f, %.4f]}", line.A.X, line.A.Y, line.B.X, line.B.Y)
}

func (line Line) Center() Vec {
	return line.A.Add(line.B.Sub(line.A).Scale(0.5))
}

func (line Line) Length() float64 {
	return line.B.Sub(line.A).Magnitude()
}
