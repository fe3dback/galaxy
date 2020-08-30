package physics

import (
	"github.com/fe3dback/box2d"
	"github.com/fe3dback/galaxy/engine"
)

type ourShape struct {
	boxShape box2d.B2ShapeInterface
}

func newOurShape(boxShape box2d.B2ShapeInterface) *ourShape {
	return &ourShape{
		boxShape: boxShape,
	}
}

func (s *ourShape) debugDraw(body *ourBody, r engine.Renderer) {
	switch shape := s.boxShape.(type) {
	case *box2d.B2PolygonShape:
		debugDrawPolygon(body, shape, r)
	}
}
