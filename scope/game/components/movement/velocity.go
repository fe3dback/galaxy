package movement

import (
	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
)

type Velocity struct {
	entity galx.GameObject

	Acceleration galx.Vec
	Velocity     galx.Vec
	MaxVelocity  galx.Vec
}

func (v Velocity) Id() string {
	return "a47f4771-2de7-4f2b-85e0-d35425bb6994"
}

func (v Velocity) Title() string {
	return "Movement.Velocity"
}

func (v Velocity) Description() string {
	return "Set stable entity velocity, will move and speedup entity each frame toward velocity direction"
}

func (v *Velocity) OnCreated(entity galx.GameObject) {
	v.entity = entity
}

func (v *Velocity) OnDraw(r galx.Renderer) error {
	if !r.Gizmos().Secondary() {
		return nil
	}

	r.DrawVector(galx.ColorYellow, 10, v.entity.AbsPosition(), v.Velocity.Direction())

	return nil
}

func (v *Velocity) OnUpdate(s galx.State) error {
	// update velocity
	v.Velocity = v.Velocity.
		Add(v.Acceleration.Scale(s.Moment().DeltaTime())).
		ClampAbs(v.MaxVelocity)

	// update position
	v.entity.AddPosition(
		v.Velocity.
			Scale(consts.PixelsPerMeter).
			Scale(s.Moment().DeltaTime()),
	)

	return nil
}
