package movement

import (
	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
)

type Velocity struct {
	entity galx.GameObject

	acceleration galx.Vec
	velocity     galx.Vec
	maxVelocity  galx.Vec
}

func NewVelocity(entity galx.GameObject, acceleration, velocity, maxVelocity galx.Vec) *Velocity {
	return &Velocity{
		entity:       entity,
		acceleration: acceleration,
		velocity:     velocity,
		maxVelocity:  maxVelocity,
	}
}

func (v *Velocity) OnDraw(r galx.Renderer) error {
	if !r.Gizmos().Secondary() {
		return nil
	}

	r.DrawVector(galx.ColorYellow, 10, v.entity.Position(), v.velocity.Direction())

	return nil
}

func (v *Velocity) OnUpdate(s galx.State) error {
	// update velocity
	v.velocity = v.velocity.
		Add(v.acceleration.Scale(s.Moment().DeltaTime())).
		ClampAbs(v.maxVelocity)

	// update position
	v.entity.AddPosition(
		v.velocity.
			Scale(consts.PixelsPerMeter).
			Scale(s.Moment().DeltaTime()),
	)

	return nil
}
