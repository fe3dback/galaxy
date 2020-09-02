package physics

import (
	"github.com/fe3dback/box2d"
	"github.com/fe3dback/galaxy/engine"
)

const colorShape = engine.ColorOrange

func debugDrawPolygon(body *ourBody, sh *box2d.B2PolygonShape, r engine.Renderer) {
	if sh.M_count <= 0 {
		return
	}

	var next *box2d.B2Vec2

	for i := 0; i <= sh.M_count-1; i++ {
		current := sh.GetVertex(i)

		if i == sh.M_count-1 {
			next = sh.GetVertex(0)
		} else {
			next = sh.GetVertex(i + 1)
		}

		r.DrawLine(colorShape, engine.Line{
			A: vec2eng(*current).Add(body.Position()),
			B: vec2eng(*next).Add(body.Position()),
		})
	}
}
