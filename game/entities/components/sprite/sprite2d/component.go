package sprite2d

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/generated"
)

type Sprite2D struct {
	entity      *entity.Entity
	resource    generated.ResourcePath
	textureInfo *engine.TextureInfo
}

func NewSprite2D(entity *entity.Entity, resource generated.ResourcePath) *Sprite2D {
	return &Sprite2D{
		entity:   entity,
		resource: resource,
	}
}

func (s2d *Sprite2D) OnDraw(r engine.Renderer) error {
	if s2d.textureInfo == nil {
		info := r.TextureQuery(s2d.resource)
		s2d.textureInfo = &info
	}

	// draw sprite
	r.DrawSpriteAngle(
		s2d.resource,
		s2d.entity.Position().Sub(engine.Vec{
			X: float64(s2d.textureInfo.Width / 2),
			Y: float64(s2d.textureInfo.Height / 2),
		}),
		s2d.entity.Rotation(),
	)

	return nil
}

func (s2d *Sprite2D) OnUpdate(_ engine.State) error {
	return nil
}
