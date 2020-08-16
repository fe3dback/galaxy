package movement

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/gm"
)

type Velocity struct {
	entity *entity.Entity

	acceleration engine.Vec
	velocity     engine.Vec
	maxVelocity  engine.Vec
}

func NewVelocity(entity *entity.Entity, acceleration, velocity, maxVelocity engine.Vec) *Velocity {
	return &Velocity{
		entity:       entity,
		acceleration: acceleration,
		velocity:     velocity,
		maxVelocity:  maxVelocity,
	}
}

func (v *Velocity) OnDraw(r engine.Renderer) error {
	if !r.Gizmos().Secondary() {
		return nil
	}

	r.DrawVector(engine.ColorYellow, 10, v.entity.Position(), v.velocity.Direction())

	return nil
}

func (v *Velocity) OnUpdate(s engine.State) error {
	// update velocity
	v.velocity = v.velocity.
		Add(v.acceleration.Scale(s.Moment().DeltaTime())).
		ClampAbs(v.maxVelocity)

	// update position
	v.entity.AddPosition(
		v.velocity.
			Scale(gm.PixelsPerMeter).
			Scale(s.Moment().DeltaTime()),
	)

	return nil
}
