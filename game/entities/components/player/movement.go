package player

import (
	"fmt"
	"math"
	"time"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/gm"
)

const accelerationMul = 4
const deAccelerationMul = 0.95
const deAccelerationStopMul = 0.65
const deAccelerationChangeDirectionMul = 0.45

type Movement struct {
	entity *entity.Entity

	walkSpeed gm.SpeedMpS
	runSpeed  gm.SpeedMpS

	velocity     engine.Vec
	isCollide    bool
	collideAngle engine.Angle
	backAngle    engine.Angle
}

func NewMovement(entity *entity.Entity, walkSpeed gm.SpeedMpS, runSpeed gm.SpeedMpS) *Movement {
	return &Movement{
		entity:    entity,
		walkSpeed: walkSpeed * gm.PixelsPerMeter,
		runSpeed:  runSpeed * gm.PixelsPerMeter,
	}
}

func (r *Movement) OnCollide(target engine.Entity, targetLayer uint8) {
	fmt.Printf("[%s] Collision %d (%s) -> %d on layer %d\n",
		time.Now().String(),
		r.entity.Id(),
		r.entity.Position(),
		target.Id(),
		targetLayer,
	)

	r.collideAngle = r.entity.Position().AngleTo(target.Position())
	r.backAngle = r.collideAngle.Flip()
	r.velocity = engine.VectorRight(25).Rotate(r.backAngle)

	r.isCollide = true
}

func (r *Movement) OnDraw(d engine.Renderer) error {
	d.DrawVector(engine.ColorGreen, 30, r.entity.Position(), r.collideAngle)
	d.DrawVector(engine.ColorOrange, 25, r.entity.Position(), r.backAngle)
	return nil
}

func (r *Movement) OnUpdate(state engine.State) error {
	if !r.isCollide {
		r.updateWalkVelocity(state)
	}
	r.isCollide = false

	r.entity.AddPosition(
		r.velocity.Scale(state.Moment().DeltaTime()),
	)

	return nil
}
func (r *Movement) updateWalkVelocity(state engine.State) {
	speedPerSecond := r.walkSpeed
	if state.Movement().Shift() {
		speedPerSecond = r.runSpeed
	}

	dt := state.Moment().DeltaTime()
	movVec := state.Movement().Vector()

	if movVec.Zero() {
		r.velocity = r.velocity.Scale(deAccelerationStopMul)
	} else {
		isSameDir := r.velocity.Normalize().Dot(movVec) > 0
		r.velocity = r.velocity.
			Scale(deAccelerationMul).
			Add(movVec.Scale(speedPerSecond * accelerationMul * dt))

		if !isSameDir {
			r.velocity = r.velocity.Scale(deAccelerationChangeDirectionMul)
		}
	}

	r.velocity = r.softClamp(r.velocity, speedPerSecond)
}

func (r *Movement) softClamp(vec engine.Vec, unit float64) engine.Vec {
	diffX := math.Abs(vec.X) - unit
	diffY := math.Abs(vec.Y) - unit

	if vec.X > unit {
		vec.X -= diffX * 0.25
	} else if vec.X < -unit {
		vec.X += diffX * 0.25
	}

	if vec.Y > unit {
		vec.Y -= diffY * 0.25
	} else if vec.Y < -unit {
		vec.Y += diffY * 0.25
	}

	return vec
}
