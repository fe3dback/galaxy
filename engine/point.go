package engine

type Point struct {
	X, Y int
}

func (p Point) Add(t Point) Point {
	return Point{
		X: p.X + t.X,
		Y: p.Y + t.Y,
	}
}
