package physics

import (
	"github.com/fe3dback/box2d"
	"github.com/fe3dback/galaxy/engine"
)

func vec2box(v engine.Vec) box2d.B2Vec2 {
	return box2d.B2Vec2{
		X: v.X / engine.PixelsPerMeter,
		Y: v.Y / engine.PixelsPerMeter,
	}
}

func vec2eng(v box2d.B2Vec2) engine.Vec {
	return engine.Vec{
		X: v.X * engine.PixelsPerMeter,
		Y: v.Y * engine.PixelsPerMeter,
	}
}
