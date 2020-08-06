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

	// calculated
	clcPreviousPosition engine.Vec

	// components
	motor  *motor
	wheels []*wheel
	spec   spec

	// results
	results struct {
		speed units.SpeedKmH
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

	for _, specWheel := range spec.wheels {
		wheels = append(wheels, &wheel{
			angle:  engine.NewAngle(0),
			torque: 0,
			radius: 30, // todo
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
		motor:  newMotor(float64(spec.mass.mass)),
		wheels: wheels,
	}
}

// return new position and rotation
func (mv *movements) update(s engine.State) (engine.Vec, engine.Angle) {
	mv.results.motor = mv.updateMotor(s)
	mv.updateWheels(s)
	mv.results.speed = mv.updateSpeed(s)

	// update velocities
	mv.velocity = mv.velocity.Add(mv.results.motor.acceleration.Scale(s.Moment().DeltaTime()))
	//mv.angularVelocity = mv.angularVelocity * engine.Angle(s.Moment().DeltaTime())

	fmt.Printf("            dt: %.2f\n", s.Moment().DeltaTime())
	fmt.Printf("      velocity: %s\n", mv.velocity)
	fmt.Printf(" ang. velocity: %.2f\n", mv.angularVelocity)

	mv.clcPreviousPosition = mv.position
	mv.position = mv.position.Add(mv.velocity.Scale(s.Moment().DeltaTime()))
	//mv.rotation = mv.rotation.Add(mv.angularVelocity)

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
	if s.Movement().Vector().Y < -0.1 {
		mv.motor.IncreaseForce(10000 * s.Moment().DeltaTime())
	}

	if s.Movement().Vector().Y > 0.1 {
		mv.motor.IncreaseForce(-10000 * s.Moment().DeltaTime())
	}

	if s.Movement().Space() {
		mv.motor.Brake()
	}

	return mv.motor.UpdateMotor(mv.velocity, mv.rotation)
}

func (mv *movements) updateWheels(s engine.State) {
	if s.Movement().Vector().X > 0.1 {
		mv.steeringAngle -= engine.Angle1
	}

	if s.Movement().Vector().X < -0.1 {
		mv.steeringAngle += engine.Angle1
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
	r.DrawVector(engine.ColorRed, 30, mv.position.Add(Vec{X: 20}), mv.steeringAngle)
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorOrange,
		fmt.Sprintf("Speed: %.2f km/h", mv.results.speed),
		mv.position.Add(Vec{Y: 75}),
	)

	r.DrawVector(engine.ColorCyan, 40, mv.position, mv.velocity.Direction())
}

func (mv *movements) drawWheels(r engine.Renderer) {
	for id, wheel := range mv.wheels {
		if !r.Gizmos().Secondary() {
			continue
		}

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
	motorPos := mv.position.Add(Vec{X: 100, Y: -30})

	if r.Gizmos().Secondary() {
		r.DrawText(
			generated.ResourcesFontsJetBrainsMonoRegular,
			engine.ColorPink,
			fmt.Sprintf("engine acceleration: %.2f", mv.results.motor.acceleration),
			motorPos.Add(Vec{Y: 20}),
		)
		r.DrawText(
			generated.ResourcesFontsJetBrainsMonoRegular,
			engine.ColorPink,
			fmt.Sprintf("force: %.2f (%.2f)", mv.results.motor.infoEngineForce, mv.results.motor.infoForceLongitudinal),
			motorPos.Add(Vec{Y: 40}),
		)
		//r.DrawText(
		//	generated.ResourcesFontsJetBrainsMonoRegular,
		//	engine.ColorPink,
		//	fmt.Sprintf(
		//		"acc: %.1f",
		//		mv.results.motor.infoAcceleration,
		//	),
		//	motorPos.Add(Vec{Y: 60}),
		//)
	}
}
