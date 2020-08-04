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
	r.DrawCrossLines(engine.ColorOrange, 10, td.entity.Position())
	r.DrawText(generated.ResourcesFontsJetBrainsMonoRegular, engine.ColorOrange, td.entity.Position().String(), td.entity.Position().Add(engine.Vec{
		X: 0,
		Y: -20,
	}))

	// helper
	r.DrawSprite(generated.ResourcesSystemAngles, td.entity.Position())

	// direction
	r.DrawVector(engine.ColorGreen, 300, td.entity.Position(), td.entity.Rotation())

	// real data
	direction := fmt.Sprintf("%.2f %.2f", td.entity.Rotation().Radians(), td.entity.Rotation().Degrees())
	r.DrawText(generated.ResourcesFontsJetBrainsMonoRegular, engine.ColorGreen, direction, td.entity.Position())

	// test vectors
	// rotate around
	rotateAround := td.entity.Position().Add(engine.Vec{
		X: 100,
		Y: 0,
	}).RotateAround(td.entity.Position(), td.entity.Rotation())
	r.DrawCrossLines(engine.ColorGreen, 10, rotateAround)
	rotateAroundText := fmt.Sprintf("ra (%s)", rotateAround)
	r.DrawText(generated.ResourcesFontsJetBrainsMonoRegular, engine.ColorOrange, rotateAroundText, td.entity.Position().Add(engine.Vec{
		X: 0,
		Y: 20,
	}))

	// polar offset
	polarOffset := td.entity.Position().PolarOffset(550, td.entity.Rotation())
	r.DrawCrossLines(engine.ColorGreen, 5, polarOffset)

	// rotate 90deg
	vec := engine.Vec{
		X: 1,
		Y: 0,
	}
	vec = vec.Rotate(td.entity.Rotation().Add(engine.NewAngle(90)))
	r.DrawVector(engine.ColorOrange, 50, td.entity.Position(), vec.Direction())
	rotateDir := fmt.Sprintf("S+90 (%.2f)", vec.Direction())
	r.DrawText(generated.ResourcesFontsJetBrainsMonoRegular, engine.ColorOrange, rotateDir, td.entity.Position().Add(engine.Vec{
		X: -20,
		Y: 40,
	}))

	return nil
}

func (td *AngleOverlay) OnUpdate(s engine.State) error {
	if s.Movement().Vector().Y < -0.1 {
		td.entity.AddRotation(engine.Angle5)
	}

	if s.Movement().Vector().Y > 0.1 {
		td.entity.AddRotation(-engine.Angle5)
	}

	return nil
}
