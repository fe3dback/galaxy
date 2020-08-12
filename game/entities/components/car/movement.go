package car

import (
	"fmt"
	"math"

	"github.com/fe3dback/galaxy/game/units"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/generated"
)

const wheelAxisTop = "top"
const gearShiftMinWaitTimeSec = 3
const defaultRoadFriction = 0.02
const wheelAxisDistance = 15.0

type movements struct {
	// spec
	spec spec

	// primary
	position        engine.Vec
	rotation        engine.Angle
	velocity        engine.Vec
	angularVelocity engine.Angle
	steeringAngle   engine.Angle

	// motor
	motorGearInd gearInd

	// calculated
	clcEngineTorque     float64
	clcPreviousPosition engine.Vec
	clcSpeed            units.SpeedKmH
	clcGasPedal         float64 // input gas 0 .. 1

	// components
	wheels []*wheel
}

type wheel struct {
	angle  engine.Angle
	torque float64
	spec   specWheel
}

func newMovements(position engine.Vec, angle engine.Angle, spec spec) *movements {
	wheels := make([]*wheel, 0)

	for _, specWheel := range spec.wheels.wheels {
		wheels = append(wheels, &wheel{
			angle:  engine.NewAngle(0),
			torque: 0,
			spec:   specWheel,
		})
	}

	return &movements{
		// spec
		spec: spec,

		// primary
		position:        position,
		rotation:        angle,
		velocity:        Vec{},
		angularVelocity: engine.Angle0,
		steeringAngle:   engine.Angle0,

		// motor
		motorGearInd: 1,

		// calculated
		clcEngineTorque:     0,
		clcPreviousPosition: position,
		clcSpeed:            0,
		clcGasPedal:         0,

		// components
		wheels: wheels,
	}
}

// return new position and rotation
func (mv *movements) update(s engine.State) (engine.Vec, engine.Angle) {
	mv.updateMotor(s)
	mv.updateWheels(s)
	mv.clcSpeed = mv.updateSpeed(s)

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

func (mv *movements) updateMotor(s engine.State) {
	rpm := engineRpm(mv.motorGearInd, mv.spec.wheels.radius, mv.clcSpeed)

	if s.Moment().FrameId()%(s.Moment().TargetFPS()*gearShiftMinWaitTimeSec) == 0 {
		// allow change gear only one time per 3 seconds
		mv.motorGearInd = automaticTransmission(mv.motorGearInd, rpm)
	}

	mv.clcEngineTorque = engineTorque(rpm)

	// calculate wheels torque
	wheelsCount := float64(len(mv.wheels))
	carUnit := mv.rotation.Unit()

	// car forward force
	acceleration := Vec{}
	friction := 0.00001
	angularAcceleration := engine.Angle0

	// calculate acceleration (wheels force)
	for _, wheel := range mv.wheels {
		// calculate wheel torque
		wheelDirection := mv.rotation.Add(wheel.angle)
		wheelUnit := wheelDirection.Unit()
		forwardFactor := math.Abs(wheelUnit.Dot(carUnit))

		maxTorque := (mv.clcEngineTorque / wheelsCount) * forwardFactor
		wheel.torque = maxTorque * mv.clcGasPedal

		// sum all wheels acceleration
		acceleration = acceleration.Add(wheelUnit.Scale(wheel.torque * s.Moment().DeltaTime()))

		// subtract friction
		wheelFriction := defaultRoadFriction*2 - (defaultRoadFriction * forwardFactor)
		friction += wheelFriction

		// calculate angular velocity
		angularAcceleration += engine.Lerpf(-engine.Angle360, engine.Angle360, -1, 1, wheel.angle.Radians())
	}

	// apply wheelbase friction
	forwardFactor := math.Abs(mv.rotation.Unit().Dot(mv.velocity.Direction().Unit()))
	friction += 1 - engine.Clamp(forwardFactor, 0, 1)

	// subtract friction for each wheel
	friction = engine.Clamp(friction, 0, 1)

	if s.Movement().Space() {
		// break
		friction = 1
	}

	// free acceleration
	fmt.Println(friction)

	// update velocity
	mv.velocity = mv.velocity.
		Add(acceleration.Scale(s.Moment().DeltaTime())).
		Scale(1 - (friction * s.Moment().DeltaTime()))

	speed := math.Abs(mv.velocity.Magnitude())
	mv.angularVelocity = Angle(speed / (wheelAxisDistance / math.Sin(mv.steeringAngle.Radians())))

	// update position
	mv.clcPreviousPosition = mv.position
	mv.position = mv.position.Add(mv.velocity)
	mv.rotation = mv.rotation.Add(mv.angularVelocity * Angle(s.Moment().DeltaTime()))
}

func (mv *movements) updateWheels(s engine.State) {
	if s.Movement().Vector().X > 0.1 {
		mv.steeringAngle -= Angle(engine.Angle45 * s.Moment().DeltaTime())
	}

	if s.Movement().Vector().X < -0.1 {
		mv.steeringAngle += Angle(engine.Angle45 * s.Moment().DeltaTime())
	}

	if s.Movement().Vector().X == 0 {
		// reset steering back
		mv.steeringAngle *= Angle(1 - (2 * s.Moment().DeltaTime()))
	}

	mv.steeringAngle = engine.Angle(
		engine.Clamp(mv.steeringAngle.Radians(), -engine.Angle45, engine.Angle45),
	)

	for _, wheel := range mv.wheels {
		if wheel.spec.axis != wheelAxisTop {
			continue
		}

		wheel.angle = mv.steeringAngle
	}

	mv.clcGasPedal = engine.Lerpf(-1, 1, 1, -1, s.Movement().Vector().Y)
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
		fmt.Sprintf("Speed: %.2f km/h", mv.clcSpeed),
		mv.position.Add(Vec{Y: 75}),
	)

	r.DrawVector(engine.ColorCyan, 40, mv.position, mv.velocity.Direction())
	r.DrawVector(
		engine.ColorOrange,
		30*mv.clcGasPedal,
		mv.position.PolarOffset(50, mv.rotation),
		mv.rotation,
	)
}

func (mv *movements) drawWheels(r engine.Renderer) {
	if !r.Gizmos().Secondary() {
		return
	}

	torqueSum := 0.0
	for _, wheel := range mv.wheels {
		torqueSum += wheel.torque
	}

	for id, wheel := range mv.wheels {
		wheelPos := mv.position.Add(Vec{
			X: float64(wheel.spec.posRelative.x),
			Y: float64(wheel.spec.posRelative.y),
		}).RotateAround(mv.position, mv.rotation)

		wheelDirection := mv.rotation + wheel.angle
		wheelForwardPos := wheelPos.PolarOffset(float64(wheel.spec.size.width), wheelDirection)

		// draw bounding box
		r.DrawPoint(engine.ColorYellow, wheelPos)
		r.DrawSquareEx(engine.ColorOrange, wheelDirection, engine.RectScreen(
			wheelPos.RoundX()-wheel.spec.size.width/2,
			wheelPos.RoundY()-wheel.spec.size.height/2,
			wheel.spec.size.width,
			wheel.spec.size.height,
		))

		// draw torque and direction
		maxWheelTorque := (torqueSum / float64(len(mv.wheels))) + 0.0001
		vectorSize := engine.Lerpf(0, maxWheelTorque, 5, 25, wheel.torque)
		r.DrawVector(engine.ColorYellow, vectorSize, wheelForwardPos, wheelDirection)

		if r.Gizmos().Debug() {
			r.DrawText(
				generated.ResourcesFontsJetBrainsMonoRegular,
				engine.ColorOrange,
				fmt.Sprintf("#%d %s (%.2f)", id, wheel.spec.axis, wheel.torque),
				wheelPos,
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
		fmt.Sprintf("engine torque: %.2f", mv.clcEngineTorque),
		motorPos.Add(Vec{Y: 20}),
	)
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorPink,
		fmt.Sprintf("gear: %d",
			mv.motorGearInd,
		),
		motorPos.Add(Vec{Y: 40}),
	)
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorPink,
		fmt.Sprintf("velocity: %s",
			mv.velocity,
		),
		motorPos.Add(Vec{Y: 60}),
	)
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorPink,
		fmt.Sprintf("angular: %.2f",
			mv.angularVelocity,
		),
		motorPos.Add(Vec{Y: 80}),
	)
}
