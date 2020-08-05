package car

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fe3dback/galaxy/game/units"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/generated"
)

const wheelAxisTop = "top"

type movements struct {
	calc          Calculator
	position      engine.Vec
	rotation      engine.Angle
	velocity      engine.Vec
	angularSpeed  float64
	steeringAngle engine.Angle

	// todo: remove
	wheels []*wheel
	spec   spec

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
		rotation:           angle,
		wheels:             wheels,
		spec:               spec,
		engineTurnoverRate: 4000,
		engineTorque:       0,
		wheelsTorque:       0,
		wheelDirection:     engine.NewAngle(0), // todo: remove
	}
}

// return new position and rotation
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

	// draw velocity
	r.DrawVector(engine.ColorPink, 35, mv.position, mv.velocity.Direction())
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

func (mv *movements) updateDirection() {
	mv.rotation = mv.velocity.Direction()
}

func (mv *movements) updateVelocity(s engine.State) (engine.Vec, engine.Angle) {
	c := mv.calc

	gearRatio := 1.0  // todo mul
	driveRatio := 1.0 // todo mul
	distanceToFrontAxle := 1.0
	distanceToRearAxle := 1.2
	corneringStiffness := 0.5

	//brakingFactor := 0.5 // todo 0 .. 1 // scale rotation down
	wheelsRadius := 0.18 // todo 0 .. 1 // speed = (wheelTorque / wheelRadius)
	dragResistance := Vec{
		X: 0.5, // todo 0 .. 1 // air resistance (will slow car)
		Y: 0.1, // todo 0 .. 1 // air resistance (will slow car)
	}
	roadSurface := 0.5 // todo 0 .. 1 // ground resistance (will slow car)
	//slopeAngle := 0.0    // todo

	isBraking := false // todo

	if s.Movement().Vector().Y > 0.1 {
		// is not braking, just release gas pedal
		isBraking = true
	}

	// calculate current velocity

	// -----------------------------------

	computed := c.compute(
		mv.velocity,
		mv.angularSpeed,

		// input:
		mv.steeringAngle,
		gearRatio,
		isBraking,

		// env:
		roadSurface,

		// const:
		float64(mv.spec.mass.mass),
		distanceToFrontAxle,
		distanceToRearAxle,
		corneringStiffness,
		driveRatio,
		wheelsRadius,
		dragResistance,
	)

	//22. Transform the acceleration from car reference frame to world
	//reference frame

	//23. Integrate the acceleration to get the velocity (in world reference
	//frame)
	//Vwc += dt * a

	mv.velocity = mv.velocity.Add(computed.acceleration.Scale(s.Moment().DeltaTime()))
	mv.angularSpeed = mv.angularSpeed + computed.angularAcceleration*s.Moment().DeltaTime()

	//24. Integrate the velocity to get the new position in world coordinate
	//Pwc += dt * Vwc

	// swap velocity vectors, because -Y in compute is car forward axle
	// in engine car forward default in right (0 deg), +X

	nextVelocityInv := Vec{
		X: mv.velocity.Y,
		Y: mv.velocity.X,
	}

	nextPos := mv.position.Add(nextVelocityInv.Scale(s.Moment().DeltaTime()))
	nextRot := mv.rotation.Add(engine.Angle(mv.angularSpeed * s.Moment().DeltaTime()))
	mv.distanceFromLastFrame = nextPos.Sub(mv.position).Magnitude()

	mv.position = nextPos
	mv.rotation = nextRot

	// ---------------------------------------
	// calculate next position
	// ---------------------------------------

	comm := exec.Command("clear")
	comm.Stdout = os.Stdout
	_ = comm.Run()

	fmt.Printf("steering deg: %v\n", mv.steeringAngle.Degrees())
	fmt.Printf("    velocity: %v\n", nextVelocityInv)
	fmt.Printf("----------------\n")
	fmt.Printf("acceleration: %v\n", computed.acceleration)
	fmt.Printf("     angular: %v\n", computed.angularAcceleration)
	fmt.Printf("----------------\n")
	fmt.Printf("    traction: %v\n", computed.infoTraction)
	fmt.Printf("d.resistance: %v\n", computed.infoDragResistance)
	fmt.Printf("r.resistance: %v\n", computed.infoRollingResistance)
	fmt.Printf("  resistance: %v\n", computed.infoResistance)
	fmt.Printf("     en. rpm: %v\n", computed.infoEngineRPM)
	fmt.Printf("en. turnover: %v\n", computed.infoEngineTurnoverRate)
	fmt.Printf("  en. torque: %v\n", computed.infoEngineTorque)
	fmt.Printf("wheel torque: %v\n", computed.infoWheelsTorque)
	fmt.Printf(" body torque: %v\n", computed.infoBodyTorque)
	fmt.Printf("       force: %v\n", computed.infoForce)
	fmt.Printf(" force front: %v\n", computed.infoForceLatFront)
	fmt.Printf("  force rear: %v\n", computed.infoForceLatRear)
	fmt.Printf("  slip front: %v\n", computed.infoWheelsSlipFront)
	fmt.Printf("   slip rear: %v\n", computed.infoWheelsSlipRear)

	// debug, todo: remove
	mv.engineTorque = computed.infoEngineTorque
	mv.engineTurnoverRate = computed.infoEngineTurnoverRate
	mv.wheelsTorque = computed.infoWheelsTorque

	return mv.position, mv.rotation
}
