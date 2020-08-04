package car

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/generated"
)

const wheelAxisTop = "top"

type movements struct {
	calc      Calculator
	position  engine.Vec
	velocity  engine.Vec
	direction engine.Angle
	wheels    []*wheel
	spec      spec
}

type wheel struct {
	angle  engine.Angle
	torque float64
	radius float64
	spec   specWheel
}

func newMovements(position engine.Vec, angle engine.Angle, spec spec) *movements {
	wheels := make([]*wheel, 0)

	for _, specWheel := range spec.wheels {
		wheels = append(wheels, &wheel{
			angle:  engine.NewAngle(0),
			torque: 0,
			radius: 30, // todo
			spec:   specWheel,
		})
	}

	return &movements{
		spec:      spec,
		calc:      Calculator{},
		velocity:  engine.VectorTowards(angle),
		position:  position,
		direction: angle,
		wheels:    wheels,
	}
}

// return new position and direction
func (mv *movements) update(s engine.State) (engine.Vec, engine.Angle) {
	mv.updateWheels(s)
	mv.updateDirection()

	// calculate all
	return mv.updateVelocity(s)
}

func (mv *movements) draw(r engine.Renderer) {
	for id, wheel := range mv.wheels {
		if !r.Gizmos().Secondary() {
			continue
		}

		pos := mv.position.Add(Vec{
			X: float64(wheel.spec.posRelative.x),
			Y: float64(wheel.spec.posRelative.y),
		}).RotateAround(mv.position, mv.direction)

		// draw bounding box
		r.DrawPoint(engine.ColorPink, pos)
		r.DrawSquareEx(engine.ColorOrange, mv.direction+wheel.angle, engine.RectScreen(
			pos.RoundX()-wheel.spec.size.width/2,
			pos.RoundY()-wheel.spec.size.height/2,
			wheel.spec.size.width,
			wheel.spec.size.height,
		))

		if r.Gizmos().Debug() {
			r.DrawText(
				generated.ResourcesFontsJetBrainsMonoRegular,
				engine.ColorPink,
				fmt.Sprintf("#%d %s", id, wheel.spec.axis),
				pos,
			)
		}
	}

	// draw velocity
	r.DrawVector(engine.ColorPink, 35, mv.position, mv.velocity.Direction())
}

func (mv *movements) updateWheels(s engine.State) {
	for _, wheel := range mv.wheels {
		if wheel.spec.axis != wheelAxisTop {
			continue
		}

		wheel.angle = engine.NewAngle(float64(-25) * s.Movement().Vector().X)
	}
}

func (mv *movements) updateDirection() {
	mv.direction = mv.velocity.Direction()
}

func (mv *movements) updateVelocity(s engine.State) (engine.Vec, engine.Angle) {
	c := mv.calc

	isBraking := true  // todo
	wheelTorque := 0.0 // todo

	if s.Movement().Vector().Y > 0.1 {
		isBraking = true
		wheelTorque = 0
	} else if s.Movement().Vector().Y < -0.1 {
		isBraking = false  // todo
		wheelTorque = 5000 // todo
	}

	brakingFactor := 0.5 // todo
	wheelRadius := 30.0  // todo
	frontalDrag := 50.0  // todo
	roadSurface := 0.5   // todo
	slopeAngle := 0.0    // todo

	// calculate current velocity
	currentVelocity := mv.velocity

	// calculate traction
	traction := c.traction(wheelTorque, wheelRadius)

	// calculate drag
	drag := c.drag(
		currentVelocity,
		frontalDrag,
	)

	// calculate rollingResistance
	rollingResistance := c.rollingResistance(
		roadSurface,
		currentVelocity,
	)

	// calculate gravity
	gravity := c.gravityForce(
		float64(mv.spec.mass.mass),
		slopeAngle,
	)

	// calculate direction
	direction := engine.VectorTowards(mv.direction)

	// calculate net force
	netForce := c.netForce(
		traction,
		drag,
		rollingResistance,
		gravity,
		direction,
		isBraking,
		brakingFactor,
	)

	// calculate acceleration
	acceleration := c.acceleration(
		netForce,
		float64(mv.spec.mass.mass),
	)

	// calculate next velocity
	mv.velocity = c.velocity(
		currentVelocity,
		acceleration,
		s.Moment().DeltaTime(),
	)

	// ---------------------------------------
	// calculate next position
	// ---------------------------------------

	fmt.Printf("----------------\n")
	fmt.Printf("    traction: %v\n", traction)
	fmt.Printf("        drag: %v\n", drag)
	fmt.Printf("roll. resist: %v\n", rollingResistance)
	fmt.Printf("     gravity: %v\n", gravity)
	fmt.Printf("   direction: %v\n", direction)
	fmt.Printf("           --\n")
	fmt.Printf("   net force: %v\n", netForce)
	fmt.Printf("acceleration: %v\n", acceleration)

	mv.position = c.nextPosition(
		mv.position,
		mv.velocity,
		s.Moment().DeltaTime(),
	)

	return mv.position, mv.direction
}
