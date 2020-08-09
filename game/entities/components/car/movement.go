package car

import (
	"fmt"

	"github.com/fe3dback/galaxy/game/units"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/generated"
)

const wheelAxisTop = "top"

type movements struct {
	position        engine.Vec
	rotation        engine.Angle
	steeringAngle   engine.Angle
	velocity        engine.Vec
	angularVelocity engine.Angle
	speed           units.SpeedKmH

	// calculated
	clcPreviousPosition engine.Vec

	// components
	motor  *motor
	wheels []*wheel
	spec   spec

	// results
	results struct {
		motor motorResult
	}
}

type wheel struct {
	angle  engine.Angle
	torque float64
	radius float64
	spec   specWheel
}

func newMovements(position engine.Vec, angle engine.Angle, spec spec) *movements {
	wheels := make([]*wheel, 0)
	wheelsRadius := 0.34 // todo to spec, in meters

	for _, specWheel := range spec.wheels {
		wheels = append(wheels, &wheel{
			angle:  engine.NewAngle(0),
			torque: 0, // todo?
			radius: wheelsRadius,
			spec:   specWheel,
		})
	}

	return &movements{
		position:        position,
		rotation:        angle,
		velocity:        Vec{},
		angularVelocity: engine.Angle0,

		spec:          spec,
		steeringAngle: engine.Angle0,

		// car
		motor: newMotor(
			float64(spec.mass.mass),
			0.7, // todo to spec,
			wheelsRadius,
		),
		wheels: wheels,
	}
}

// return new position and rotation
func (mv *movements) update(s engine.State) (engine.Vec, engine.Angle) {
	mv.results.motor = mv.updateMotor(s)
	mv.updateWheels(s)
	mv.speed = mv.updateSpeed(s)

	// update velocities
	mv.velocity = mv.velocity.Add(
		mv.results.motor.acceleration.Rotate(mv.rotation).Scale(s.Moment().DeltaTime()),
	)
	mv.angularVelocity = mv.angularVelocity.Add(
		mv.results.motor.angularAcceleration * Angle(s.Moment().DeltaTime()),
	)

	fmt.Printf("            dt: %.2f\n", s.Moment().DeltaTime())
	fmt.Printf("      velocity: %s\n", mv.velocity)
	fmt.Printf(" ang. velocity: %.2f\n", mv.angularVelocity)

	mv.clcPreviousPosition = mv.position
	mv.position = mv.position.Add(
		mv.velocity.
			Scale(units.PixelsPerMeter).
			Scale(s.Moment().DeltaTime()),
	)
	mv.rotation = mv.rotation.Add(
		mv.angularVelocity * units.PixelsPerMeter * Angle(s.Moment().DeltaTime()),
	)

	return mv.position, mv.rotation
}

func (mv *movements) updateSpeed(s engine.State) units.SpeedKmH {
	pixelsPerFrame := mv.position.Sub(mv.clcPreviousPosition).Magnitude()
	pixelsPerSecond := pixelsPerFrame * float64(s.Moment().TargetFPS())
	metersPerSecond := pixelsPerSecond / units.PixelsPerMeter
	metersPerHour := metersPerSecond * 3600

	return metersPerHour / 1000
}

// =======================================================================

func (mv *movements) updateMotor(s engine.State) motorResult {
	mv.motor.GasPedalPushPercent(-s.Movement().Vector().Y)

	if s.Movement().Space() {
		mv.motor.Brake()
	}

	return mv.motor.UpdateMotor(mv.speed, mv.velocity, mv.rotation, mv.steeringAngle, mv.angularVelocity)
}

func (mv *movements) updateWheels(s engine.State) {
	if s.Movement().Vector().X > 0.1 {
		mv.steeringAngle -= engine.Angle1
		//mv.rotation -= engine.Angle1 // todo: rem
	}

	if s.Movement().Vector().X < -0.1 {
		mv.steeringAngle += engine.Angle1
		//mv.rotation += engine.Angle1 // todo: rem
	}

	mv.steeringAngle = engine.Angle(
		engine.Clamp(mv.steeringAngle.Radians(), -engine.Angle25, engine.Angle25),
	)

	for _, wheel := range mv.wheels {
		if wheel.spec.axis != wheelAxisTop {
			continue
		}

		wheel.angle = mv.steeringAngle
	}
}

func (mv *movements) draw(r engine.Renderer) {
	mv.drawBoundingBox(r)
	mv.drawMotor(r)
	mv.drawWheels(r)
}

func (mv *movements) drawBoundingBox(r engine.Renderer) {
	if !r.Gizmos().Secondary() {
		return
	}

	r.DrawVector(engine.ColorRed, 30, mv.position, mv.rotation+mv.steeringAngle)
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorOrange,
		fmt.Sprintf("Speed: %.2f km/h", mv.speed),
		mv.position.Add(Vec{Y: 75}),
	)

	r.DrawVector(engine.ColorCyan, 40, mv.position, mv.velocity.Direction())
	r.DrawVector(
		engine.ColorOrange,
		30*mv.results.motor.infoGasPedal,
		mv.position.PolarOffset(50, mv.rotation),
		mv.rotation,
	)
}

func (mv *movements) drawWheels(r engine.Renderer) {
	if !r.Gizmos().Secondary() {
		return
	}

	for id, wheel := range mv.wheels {
		pos := mv.position.Add(Vec{
			X: float64(wheel.spec.posRelative.x),
			Y: float64(wheel.spec.posRelative.y),
		}).RotateAround(mv.position, mv.rotation)

		// draw bounding box
		r.DrawPoint(engine.ColorPink, pos)
		r.DrawSquareEx(engine.ColorOrange, mv.rotation+wheel.angle, engine.RectScreen(
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
}

func (mv *movements) drawMotor(r engine.Renderer) {
	if !r.Gizmos().Secondary() {
		return
	}

	motorPos := mv.position.Add(Vec{X: 150, Y: -100})

	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorPink,
		fmt.Sprintf("engine acceleration: %.2f", mv.results.motor.acceleration),
		motorPos.Add(Vec{Y: 20}),
	)
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorPink,
		fmt.Sprintf("force: %s", mv.results.motor.infoForce.longitudinalForce),
		motorPos.Add(Vec{Y: 40}),
	)
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorPink,
		fmt.Sprintf("trac: %s, drag: %s, roll.rt: %s",
			mv.results.motor.infoForce.infoTraction,
			mv.results.motor.infoForce.infoDrag,
			mv.results.motor.infoForce.infoRollingResistance,
		),
		motorPos.Add(Vec{Y: 60}),
	)
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorPink,
		fmt.Sprintf(
			"rpm: %.1f (tq.max: %.1f, tq: %.1f)",
			mv.results.motor.infoDrive.infoRPM,
			mv.results.motor.infoDrive.infoMaxTorque,
			mv.results.motor.infoDrive.infoEngineTorque,
		),
		motorPos.Add(Vec{Y: 80}),
	)
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorPink,
		fmt.Sprintf("drv.force: %s", mv.results.motor.infoDrive.driveForce),
		motorPos.Add(Vec{Y: 100}),
	)
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorPink,
		fmt.Sprintf("gear: %d (%.1f)",
			mv.results.motor.infoDrive.infoGear,
			mv.results.motor.infoDrive.infoGearRatio,
		),
		motorPos.Add(Vec{Y: 120}),
	)
}
