package debug

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/generated"
)

type AngleOverlay struct {
	entity *entity.Entity
}

func NewAngleOverlay(entity *entity.Entity) *AngleOverlay {
	comp := &AngleOverlay{
		entity: entity,
	}

	return comp
}

func (td *AngleOverlay) OnDraw(r engine.Renderer) error {

	// center
	r.DrawCrossLines(engine.ColorOrange, 10, td.entity.Position().ToPoint())

	// helper
	r.DrawSprite(generated.ResourcesSystemAngles, td.entity.Position().ToPoint())

	// direction
	r.DrawVector(engine.ColorGreen, 300, td.entity.Position(), td.entity.Rotation())

	// real data
	direction := fmt.Sprintf("%.2f %.2f", td.entity.Rotation(), engine.Radian(td.entity.Rotation()))
	r.DrawText(generated.ResourcesFontsJetBrainsMonoRegular, engine.ColorGreen, direction, td.entity.Position().ToPoint())

	// test vectors
	// rotate around
	rotateAround := td.entity.Position().Add(engine.Vector2D{
		X: 500,
		Y: 0,
	}).RotateAround(td.entity.Position(), td.entity.Rotation())
	r.DrawCrossLines(engine.ColorGreen, 10, rotateAround.ToPoint())

	// polar offset
	polarOffset := td.entity.Position().PolarOffset(550, td.entity.Rotation())
	r.DrawCrossLines(engine.ColorGreen, 5, polarOffset.ToPoint())

	// rotate 90deg
	vec := engine.Vector2D{
		X: 1,
		Y: 0,
	}
	vec = vec.Rotate(td.entity.Rotation().Add(90))
	r.DrawVector(engine.ColorOrange, 50, td.entity.Position(), vec.Direction())
	rotateDir := fmt.Sprintf("S+90 (%.2f)", vec.Direction())
	r.DrawText(generated.ResourcesFontsJetBrainsMonoRegular, engine.ColorOrange, rotateDir, td.entity.Position().Add(engine.Vector2D{
		X: -20,
		Y: 20,
	}).ToPoint())

	return nil
}

func (td *AngleOverlay) OnUpdate(s engine.State) error {
	if s.Movement().Vector().Y < -0.1 {
		td.entity.AddRotation(5)
	}

	if s.Movement().Vector().Y > 0.1 {
		td.entity.AddRotation(-5)
	}

	return nil
}
