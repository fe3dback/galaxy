package physics

import (
	"github.com/fe3dback/box2d"
)

const (
	shapeTypeCircle    = 0
	shapeTypeEdge      = 1
	shapeTypePolygon   = 2
	shapeTypeChain     = 3
	shapeTypeTypeCount = 4
)

type ourShape struct {
	boxShape box2d.B2ShapeInterface
}

func newOurShape(boxShape box2d.B2ShapeInterface) *ourShape {
	return &ourShape{
		boxShape: boxShape,
	}
}
