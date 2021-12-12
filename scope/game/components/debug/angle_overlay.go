package debug

import (
	"fmt"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
)

type AngleOverlay struct {
	entity galx.GameObject
}

func (td AngleOverlay) Id() string {
	return "e8cd7c22-0057-4362-9896-8bf0d4487763"
}

func (td AngleOverlay) Title() string {
	return "Debug.Angle overlay"
}

func (td AngleOverlay) Description() string {
	return "Draw angle overlay over entity, used for debug rotations"
}

func (td *AngleOverlay) OnCreated(entity galx.GameObject) {
	td.entity = entity
}

func (td *AngleOverlay) OnDraw(r galx.Renderer) error {
	// center
	r.DrawCrossLines(galx.ColorOrange, 10, td.entity.AbsPosition())
	r.DrawText(consts.AssetDefaultFont, galx.ColorOrange, td.entity.AbsPosition().String(), td.entity.AbsPosition().Add(galx.Vec{
		X: 0,
		Y: -20,
	}))

	// direction
	r.DrawVector(galx.ColorGreen, 300, td.entity.AbsPosition(), td.entity.Rotation())

	// real data
	direction := fmt.Sprintf("%.2f %.2f", td.entity.Rotation().Radians(), td.entity.Rotation().Degrees())
	r.DrawText(consts.AssetDefaultFont, galx.ColorGreen, direction, td.entity.AbsPosition())

	// test vectors
	// rotate around
	rotateAround := td.entity.AbsPosition().Add(galx.Vec{
		X: 100,
		Y: 0,
	}).RotateAround(td.entity.AbsPosition(), td.entity.Rotation())
	r.DrawCrossLines(galx.ColorGreen, 10, rotateAround)
	rotateAroundText := fmt.Sprintf("ra (%s)", rotateAround)
	r.DrawText(consts.AssetDefaultFont, galx.ColorOrange, rotateAroundText, td.entity.AbsPosition().Add(galx.Vec{
		X: 0,
		Y: 20,
	}))

	// polar offset
	polarOffset := td.entity.AbsPosition().PolarOffset(550, td.entity.Rotation())
	r.DrawCrossLines(galx.ColorGreen, 5, polarOffset)

	// rotate 90deg
	vec := galx.Vec{
		X: 1,
		Y: 0,
	}
	vec = vec.Rotate(td.entity.Rotation().Add(galx.NewAngle(90)))
	r.DrawVector(galx.ColorOrange, 50, td.entity.AbsPosition(), vec.Direction())
	rotateDir := fmt.Sprintf("S+90 (%.2f)", vec.Direction())
	r.DrawText(consts.AssetDefaultFont, galx.ColorOrange, rotateDir, td.entity.AbsPosition().Add(galx.Vec{
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
