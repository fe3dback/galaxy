package debug

import (
	"fmt"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
)

type AngleOverlay struct {
	entity galx.GameObject
}

func NewAngleOverlay(entity galx.GameObject) *AngleOverlay {
	comp := &AngleOverlay{
		entity: entity,
	}

	return comp
}

func (td *AngleOverlay) OnDraw(r galx.Renderer) error {
	// center
	r.DrawCrossLines(galx.ColorOrange, 10, td.entity.Position())
	r.DrawText(consts.DefaultFont, galx.ColorOrange, td.entity.Position().String(), td.entity.Position().Add(galx.Vec{
		X: 0,
		Y: -20,
	}))

	// direction
	r.DrawVector(galx.ColorGreen, 300, td.entity.Position(), td.entity.Rotation())

	// real data
	direction := fmt.Sprintf("%.2f %.2f", td.entity.Rotation().Radians(), td.entity.Rotation().Degrees())
	r.DrawText(consts.DefaultFont, galx.ColorGreen, direction, td.entity.Position())

	// test vectors
	// rotate around
	rotateAround := td.entity.Position().Add(galx.Vec{
		X: 100,
		Y: 0,
	}).RotateAround(td.entity.Position(), td.entity.Rotation())
	r.DrawCrossLines(galx.ColorGreen, 10, rotateAround)
	rotateAroundText := fmt.Sprintf("ra (%s)", rotateAround)
	r.DrawText(consts.DefaultFont, galx.ColorOrange, rotateAroundText, td.entity.Position().Add(galx.Vec{
		X: 0,
		Y: 20,
	}))

	// polar offset
	polarOffset := td.entity.Position().PolarOffset(550, td.entity.Rotation())
	r.DrawCrossLines(galx.ColorGreen, 5, polarOffset)

	// rotate 90deg
	vec := galx.Vec{
		X: 1,
		Y: 0,
	}
	vec = vec.Rotate(td.entity.Rotation().Add(galx.NewAngle(90)))
	r.DrawVector(galx.ColorOrange, 50, td.entity.Position(), vec.Direction())
	rotateDir := fmt.Sprintf("S+90 (%.2f)", vec.Direction())
	r.DrawText(consts.DefaultFont, galx.ColorOrange, rotateDir, td.entity.Position().Add(galx.Vec{
		X: -20,
		Y: 40,
	}))

	return nil
}

func (td *AngleOverlay) OnUpdate(s galx.State) error {
	if s.Movement().Vector().Y < -0.1 {
		td.entity.AddRotation(galx.Angle5)
	}

	if s.Movement().Vector().Y > 0.1 {
		td.entity.AddRotation(-galx.Angle5)
	}

	return nil
}
