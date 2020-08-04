package car

import (
	"fmt"

	"github.com/fe3dback/galaxy/game/units"

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

	// todo: remove/refactor
	engineTurnoverRate    float64
	engineTorque          float64
	wheelsTorque          float64
	wheelDirection        engine.Angle
	distanceFromLastFrame float64
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
		calc:               Calculator{},
		position:           position,
		velocity:           engine.VectorTowards(angle),
		direction:          angle,
		wheels:             wheels,
		spec:               spec,
		engineTurnoverRate: 4000,
		engineTorque:       0,
		wheelsTorque:       0,
		wheelDirection:     engine.NewAngle(0), // todo: remove
	}
}

// return new position and direction
func (mv *movements) update(s engine.State) (engine.Vec, engine.Angle) {
	mv.updateWheels(s)
	mv.updateDirection()

	// calculate all
	return mv.updateVelocity(s)
}

func (mv *movements) getSpeed() units.SpeedKmH {
	pixelsPerSeconds := mv.distanceFromLastFrame * 30 // 30 as fps target // todo
	metersPerSeconds := pixelsPerSeconds / units.PixelsPerMeter

	return units.TransformSpeed(metersPerSeconds)
}

func (mv *movements) draw(r engine.Renderer) {
	wheelDebPos := mv.position.Sub(Vec{X: 250})

	r.DrawVector(engine.ColorRed, 30, wheelDebPos, mv.wheelDirection)
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorRed,
		fmt.Sprintf("%.2f (%.2f)", mv.engineTurnoverRate, mv.engineTorque),
		wheelDebPos.Add(Vec{Y: 10}),
	)
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorRed,
		fmt.Sprintf("%.2f (%.2f)", mv.wheelsTorque, mv.wheelDirection),
		wheelDebPos.Add(Vec{Y: 30}),
	)
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorOrange,
		fmt.Sprintf("Speed: %.2f km/h", mv.getSpeed()),
		wheelDebPos.Add(Vec{Y: 50}),
	)

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

	mv.engineTurnoverRate += s.Movement().Vector().X * 100
	if mv.engineTurnoverRate < 0 {
		mv.engineTurnoverRate = 0
	}

	if mv.engineTurnoverRate > 6500 {
		mv.engineTurnoverRate = 6500 // todo: car max engine limit
	}
}

func (mv *movements) updateDirection() {
	mv.direction = mv.velocity.Direction()
}

func (mv *movements) updateVelocity(s engine.State) (engine.Vec, engine.Angle) {
	c := mv.calc

	// todo: move to wheels
	mv.engineTorque = mv.calc.engineAngularVelocity(mv.engineTurnoverRate)
	mv.wheelsTorque = mv.calc.wheelsTorque(mv.engineTorque, 1, 1)

	mv.wheelDirection += engine.Angle(mv.wheelsTorque * s.Moment().DeltaTime())
	// todo ^ end

	wheelTorque := mv.wheelsTorque

	isBraking := false // todo

	if s.Movement().Vector().Y > 0.1 {
		// is not braking, just release gas pedal
		isBraking = true
	}

	brakingFactor := 0.5 // todo 0 .. 1 // scale direction down
	wheelRadius := 0.18  // todo 0 .. 1 // speed = (wheelTorque / wheelRadius)
	frontalDrag := 0.1   // todo 0 .. 1 // air resistance (will slow car)
	roadSurface := 0.5   // todo 0 .. 1 // ground resistance (will slow car)
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

	nextPos := c.nextPosition(
		mv.position,
		mv.velocity,
		s.Moment().DeltaTime(),
	)
	mv.distanceFromLastFrame = nextPos.Sub(mv.position).Magnitude()
	mv.position = nextPos

	return mv.position, mv.direction
}
