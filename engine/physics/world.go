package physics

import (
	"fmt"

	"github.com/fe3dback/galaxy/utils"

	"github.com/fe3dback/box2d"
	"github.com/fe3dback/galaxy/engine"
)

// The suggested iteration count for Box2D is 8 for velocity and 3 for position.
// You can tune this number to your liking, just keep in mind that this
// has a trade-off between performance and accuracy.
// Using fewer iterations increases performance but accuracy suffers.
// Likewise, using more iterations decreases performance but improves the
// quality of your simulation. For this simple example, we don't need much
// iteration. Here are our chosen iteration counts.
const (
	velocityIterations = 8
	positionIterations = 3
)

const (
	bodyTypeStatic    = 0
	bodyTypeKinematic = 1
	bodyTypeDynamic   = 2
)

type World struct {
	world *box2d.B2World
}

func NewWorld(closer *utils.Closer) *World {
	world := box2d.MakeB2World(box2d.B2Vec2{
		X: 0,
		Y: 0,
	})

	closer.EnqueueFree(func() {
		world.Destroy()
	})

	return &World{
		world: &world,
	}
}

func (w *World) Update(deltaTime float64) {
	w.world.Step(deltaTime, velocityIterations, positionIterations)
}

func (w *World) CreateShapeBox(width, height engine.Pixel) engine.PhysicsShape {
	shape := box2d.NewB2PolygonShape()
	shape.SetAsBox(
		(float64(width)/engine.PixelsPerMeter)/2,
		(float64(height)/engine.PixelsPerMeter)/2,
	)

	return newOurShape(shape)
}

func (w *World) AddBodyStatic(
	pos engine.Vec,
	rot engine.Angle,
	shape engine.PhysicsShape,
	categoryBits uint16,
	maskBits uint16,
) engine.PhysicsBody {
	sh, ok := shape.(*ourShape)
	if !ok {
		panic(fmt.Sprintf("can`t add static body, shape %T should by instance of %T", shape, ourShape{}))
	}

	ref := w.newBoxBodyStatic(pos, rot)

	fixtureDef := box2d.MakeB2FixtureDef()
	fixtureDef.Shape = sh.boxShape
	fixtureDef.Density = 0
	fixtureDef.Filter.CategoryBits = categoryBits
	fixtureDef.Filter.MaskBits = maskBits

	ref.CreateFixtureFromDef(&fixtureDef)

	return newOurBody(ref, sh)
}

func (w *World) AddBodyDynamic(
	pos engine.Vec,
	rot engine.Angle,
	mass engine.Kilogram,
	shape engine.PhysicsShape,
	categoryBits uint16,
	maskBits uint16,
) engine.PhysicsBody {
	sh, ok := shape.(*ourShape)
	if !ok {
		panic(fmt.Sprintf("can`t add static body, shape %T should by instance of %T", shape, ourShape{}))
	}

	ref := w.newBoxBodyDynamic(pos, rot)

	fixtureDef := box2d.MakeB2FixtureDef()
	fixtureDef.Shape = sh.boxShape
	fixtureDef.Density = mass // todo: calculate density
	fixtureDef.Friction = 0.1
	fixtureDef.Filter.CategoryBits = categoryBits
	fixtureDef.Filter.MaskBits = maskBits

	ref.CreateFixtureFromDef(&fixtureDef)

	return newOurBody(ref, sh)
}

func (w *World) newBoxBodyStatic(pos engine.Vec, rot engine.Angle) *box2d.B2Body {
	def := box2d.NewB2BodyDef()
	def.Position = vec2box(pos)
	def.Angle = rot.Radians()
	def.Type = bodyTypeStatic
	def.Active = true

	return w.world.CreateBody(def)
}

func (w *World) newBoxBodyDynamic(pos engine.Vec, rot engine.Angle) *box2d.B2Body {
	def := box2d.NewB2BodyDef()
	def.Position = vec2box(pos)
	def.Angle = rot.Radians()
	def.Type = bodyTypeDynamic
	def.Active = true
	def.AllowSleep = false

	return w.world.CreateBody(def)
}
