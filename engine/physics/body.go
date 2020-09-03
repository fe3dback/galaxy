package physics

import (
	"github.com/fe3dback/box2d"
	"github.com/fe3dback/galaxy/engine"
)

type ourBody struct {
	boxBody *box2d.B2Body
	shapes  []*ourShape // only for debug draw
}

func newOurBody(boxBody *box2d.B2Body, shapes ...*ourShape) *ourBody {
	return &ourBody{
		boxBody: boxBody,
		shapes:  shapes,
	}
}

func (b *ourBody) Position() engine.Vec {
	return vec2eng(b.boxBody.GetPosition())
}

func (b *ourBody) SetPosition(pos engine.Vec) {
	b.boxBody.SetTransform(
		vec2box(pos),
		b.boxBody.GetAngle(),
	)
}

func (b *ourBody) Rotation() engine.Angle {
	return engine.Angle(b.boxBody.GetAngle())
}

func (b *ourBody) SetRotation(rot engine.Angle) {
	b.boxBody.SetTransform(
		b.boxBody.GetPosition(),
		rot.Radians(),
	)
}

func (b *ourBody) ApplyForce(force engine.Vec, position engine.Vec) {
	b.boxBody.ApplyForce(vec2box(force), vec2box(position), true)
}
