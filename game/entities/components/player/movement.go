package player

import (
	"fmt"
	"math"

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

	previousPos   engine.Vec
	previousAngle engine.Angle
	walkSpeed     gm.SpeedMpS
	runSpeed      gm.SpeedMpS

	velocity engine.Vec
}

func NewMovement(entity *entity.Entity, walkSpeed gm.SpeedMpS, runSpeed gm.SpeedMpS) *Movement {
	return &Movement{
		entity:    entity,
		walkSpeed: walkSpeed * gm.PixelsPerMeter,
		runSpeed:  runSpeed * gm.PixelsPerMeter,
	}
}

func (r *Movement) OnCollide(target engine.Entity, targetLayer uint8) {
	fmt.Printf("Collision %d -> %d on layer %d\n", r.entity.Id(), target.Id(), targetLayer)

	r.entity.SetPosition(r.previousPos)
	r.entity.SetRotation(r.previousAngle)
}

func (r *Movement) OnDraw(_ engine.Renderer) error {
	return nil
}

func (r *Movement) OnUpdate(state engine.State) error {
	r.previousPos = r.entity.Position()
	r.previousAngle = r.entity.Rotation()

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
	r.entity.AddPosition(r.velocity.Scale(dt))

	return nil
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
