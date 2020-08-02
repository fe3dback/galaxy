package sprite2d

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/generated"
)

type Sprite2D struct {
	entity   *entity.Entity
	resource generated.ResourcePath
}

func NewSprite2D(entity *entity.Entity, resource generated.ResourcePath) *Sprite2D {
	return &Sprite2D{
		entity:   entity,
		resource: resource,
	}
}

func (s2d *Sprite2D) OnDraw(r engine.Renderer) error {
	fmt.Printf("rot = %f\n", s2d.entity.Rotation())
	// draw sprite
	r.DrawSpriteAngle(
		s2d.resource,
		s2d.entity.Position().ToPoint(),
		s2d.entity.Rotation(),
	)

	// draw center
	r.DrawCrossLines(engine.ColorCyan, 5, s2d.entity.Position().ToPoint())

	// draw vector
	r.DrawVector(engine.ColorCyan, 50, s2d.entity.Position(), s2d.entity.Rotation())

	return nil
}

func (s2d *Sprite2D) OnUpdate(_ engine.State) error {
	return nil
}
