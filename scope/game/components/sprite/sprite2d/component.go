package sprite2d

import (
	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
)

type Sprite2D struct {
	entity      galx.GameObject
	resource    consts.AssetsPath
	textureInfo *galx.TextureInfo
}

// todo: refactor to new components system

func NewSprite2D(entity galx.GameObject, resource consts.AssetsPath) *Sprite2D {
	return &Sprite2D{
		entity:   entity,
		resource: resource,
	}
}

func (s2d *Sprite2D) OnDraw(r galx.Renderer) error {
	if s2d.textureInfo == nil {
		info := r.TextureQuery(s2d.resource)
		s2d.textureInfo = &info
	}

	// draw sprite
	r.DrawSpriteAngle(
		s2d.resource,
		s2d.entity.AbsPosition().Sub(galx.Vec{
			X: float64(s2d.textureInfo.Width / 2),
			Y: float64(s2d.textureInfo.Height / 2),
		}),
		s2d.entity.Rotation(),
	)

	return nil
}

func (s2d *Sprite2D) OnUpdate(_ galx.State) error {
	return nil
}
