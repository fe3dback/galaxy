package physics

import (
	"github.com/fe3dback/box2d"
	"github.com/fe3dback/galaxy/engine"
)

const (
	colorBodyNotActive = 0x777744FF
	colorBodyStatic    = 0x77AA77FF
	colorBodyKinematic = 0x7777AAFF
	colorBodySleep     = 0x555555FF
	colorBodyAwake     = 0x885555FF
	colorJoints        = 0x22AA44FF
	colorContacts      = 0xAA33CCFF
	colorPairs         = 0x33AAAAFF
	colorBoundingBox   = 0x882288FF
	colorCenterOfMass  = 0xCCAA44FF
)

func debugDrawWorld(w *box2d.B2World, r engine.Renderer) {
	drawShapes := r.Gizmos().Debug()
	drawCenterOfMass := r.Gizmos().Debug()
	drawContacts := r.Gizmos().Debug()
	drawJoints := r.Gizmos().Spam()
	drawPairs := r.Gizmos().Spam()
	drawAABB := r.Gizmos().Spam()

	if drawShapes {
		for body := w.GetBodyList(); body != nil; body = body.GetNext() {
			transform := body.GetTransform()
			for fixture := body.GetFixtureList(); fixture != nil; fixture = fixture.GetNext() {
				if !body.IsActive() {
					drawShape(r, fixture, transform, colorBodyNotActive)
				} else if body.GetType() == bodyTypeStatic {
					drawShape(r, fixture, transform, colorBodyStatic)
				} else if body.GetType() == bodyTypeKinematic {
					drawShape(r, fixture, transform, colorBodyKinematic)
				} else if !body.IsAwake() {
					drawShape(r, fixture, transform, colorBodySleep)
				} else {
					drawShape(r, fixture, transform, colorBodyAwake)
				}
			}
		}
	}

	if drawJoints {
		for joint := w.GetJointList(); joint != nil; joint = joint.GetNext() {
			drawJoint(r, joint, colorJoints)
		}
	}

	if drawPairs || drawContacts {
		for contact := w.GetContactList(); contact != nil; contact = contact.GetNext() {
			if drawPairs {
				centerA := contact.GetFixtureA().GetAABB(0).GetCenter()
				centerB := contact.GetFixtureB().GetAABB(0).GetCenter()

				drawSolidSegment(r, centerA, centerB, colorPairs)
			}

			if drawContacts {
				contactCount := contact.GetManifold().PointCount
				manifold := box2d.MakeB2WorldManifold()
				contact.GetWorldManifold(&manifold)

				for i := 0; i < contactCount; i++ {
					drawSolidPoint(r, manifold.Points[i], colorContacts)
				}
			}
		}
	}

	if drawAABB {
		broad := &w.M_contactManager.M_broadPhase
		for body := w.GetBodyList(); body != nil; body = body.GetNext() {
			if !body.IsActive() {
				continue
			}

			for fixture := body.GetFixtureList(); fixture != nil; fixture = fixture.GetNext() {
				for i := 0; i < fixture.M_proxyCount; i++ {
					proxy := fixture.M_proxies[i]
					aabb := broad.GetFatAABB(proxy.ProxyId)
					vs := [4]box2d.B2Vec2{}
					vs[0] = box2d.B2Vec2{X: aabb.LowerBound.X, Y: aabb.LowerBound.Y}
					vs[1] = box2d.B2Vec2{X: aabb.UpperBound.X, Y: aabb.LowerBound.Y}
					vs[2] = box2d.B2Vec2{X: aabb.UpperBound.X, Y: aabb.UpperBound.Y}
					vs[3] = box2d.B2Vec2{X: aabb.LowerBound.X, Y: aabb.UpperBound.Y}
					drawSolidPolygon(r, vs[:], colorBoundingBox)
				}
			}
		}
	}

	if drawCenterOfMass {
		for body := w.GetBodyList(); body != nil; body = body.GetNext() {
			transform := body.GetTransform()
			transform.P = body.GetWorldCenter()

			drawSolidPoint(r, transform.P, colorCenterOfMass)
		}
	}
}

func drawShape(
	r engine.Renderer,
	fixture *box2d.B2Fixture,
	transform box2d.B2Transform,
	color engine.Color,
) {
	switch fixture.GetType() {
	case shapeTypeCircle:
		shapeCircle := fixture.GetShape().(*box2d.B2CircleShape)

		center := box2d.B2TransformVec2Mul(transform, shapeCircle.M_p)
		radius := shapeCircle.M_radius
		axis := box2d.B2RotVec2Mul(transform.Q, box2d.B2Vec2{X: 1.0, Y: 0.0})

		drawSolidCircle(r, center, radius, axis, color)

	case shapeTypeEdge:
		shapeEdge := fixture.GetShape().(*box2d.B2EdgeShape)
		v1 := box2d.B2TransformVec2Mul(transform, shapeEdge.M_vertex1)
		v2 := box2d.B2TransformVec2Mul(transform, shapeEdge.M_vertex2)

		drawSolidSegment(r, v1, v2, color)

	case shapeTypeTypeCount:
		// not draw

	case shapeTypeChain:
		// todo
		// 		{
		// 			b2ChainShape* chain = (b2ChainShape*)fixture.GetShape();
		// 			int count = chain.m_count;
		// 			const b2Vec2* vertices = chain.m_vertices;

		// 			b2Color ghostColor(0.75f * color.r, 0.75f * color.g, 0.75f * color.b, color.a);

		// 			b2Vec2 v1 = b2Mul(xf, vertices[0]);
		// 			g_debugDraw.DrawPoint(v1, 4.0, color);

		// 			if (chain.m_hasPrevVertex)
		// 			{
		// 				b2Vec2 vp = b2Mul(xf, chain.m_prevVertex);
		// 				g_debugDraw.DrawSegment(vp, v1, ghostColor);
		// 				g_debugDraw.DrawCircle(vp, 0.1f, ghostColor);
		// 			}

		// 			for (int i = 1; i < count; ++i)
		// 			{
		// 				b2Vec2 v2 = b2Mul(xf, vertices[i]);
		// 				g_debugDraw.DrawSegment(v1, v2, color);
		// 				g_debugDraw.DrawPoint(v2, 4.0, color);
		// 				v1 = v2;
		// 			}

		// 			if (chain.m_hasNextVertex)
		// 			{
		// 				b2Vec2 vn = b2Mul(xf, chain.m_nextVertex);
		// 				g_debugDraw.DrawSegment(v1, vn, ghostColor);
		// 				g_debugDraw.DrawCircle(vn, 0.1f, ghostColor);
		// 			}
		// 		}
	case shapeTypePolygon:
		shapePolygon := fixture.GetShape().(*box2d.B2PolygonShape)
		vertexCount := shapePolygon.M_count
		vertices := make([]box2d.B2Vec2, 0)

		for i := 0; i < vertexCount; i++ {
			vertices = append(vertices, box2d.B2TransformVec2Mul(
				transform,
				shapePolygon.M_vertices[i],
			))
		}

		drawSolidPolygon(r, vertices, color)

	default:
		return
	}
}

func drawJoint(r engine.Renderer, joint box2d.B2JointInterface, color engine.Color) {
	transform1 := joint.GetBodyA().GetTransform()
	transform2 := joint.GetBodyB().GetTransform()
	a1, a2 := resolveJointAnchors(joint)

	switch joint.GetType() {
	case jointTypeDistance:
		drawSolidSegment(r, a1, a2, color)

	case jointTypePulley:
		pulleyJoint := joint.(*box2d.B2PulleyJoint)
		s1 := pulleyJoint.GetGroundAnchorA()
		s2 := pulleyJoint.GetGroundAnchorB()

		drawSolidSegment(r, s1, a1, color)
		drawSolidSegment(r, s2, a2, color)
		drawSolidSegment(r, s1, a2, color)

	case jointTypeMouse:
	case jointTypeUnknown:
		// don't draw this
		break

	default:
		x1 := transform1.P
		x2 := transform2.P

		drawSolidSegment(r, x1, a1, color)
		drawSolidSegment(r, a1, a2, color)
		drawSolidSegment(r, x2, a2, color)
	}
}

func resolveJointAnchors(joint box2d.B2JointInterface) (box2d.B2Vec2, box2d.B2Vec2) {
	switch joint.GetType() {
	case jointTypeRevolute:
		casted := joint.(*box2d.B2RevoluteJoint)
		return casted.GetAnchorA(), casted.GetAnchorB()
	case jointTypePrismatic:
		casted := joint.(*box2d.B2PrismaticJoint)
		return casted.GetAnchorA(), casted.GetAnchorB()
	case jointTypeDistance:
		casted := joint.(*box2d.B2DistanceJoint)
		return casted.GetAnchorA(), casted.GetAnchorB()
	case jointTypePulley:
		casted := joint.(*box2d.B2PulleyJoint)
		return casted.GetAnchorA(), casted.GetAnchorB()
	case jointTypeMouse:
		casted := joint.(*box2d.B2MouseJoint)
		return casted.GetAnchorA(), casted.GetAnchorB()
	case jointTypeGear:
		casted := joint.(*box2d.B2GearJoint)
		return casted.GetAnchorA(), casted.GetAnchorB()
	case jointTypeWheel:
		casted := joint.(*box2d.B2WheelJoint)
		return casted.GetAnchorA(), casted.GetAnchorB()
	case jointTypeWeld:
		casted := joint.(*box2d.B2WeldJoint)
		return casted.GetAnchorA(), casted.GetAnchorB()
	case jointTypeFriction:
		casted := joint.(*box2d.B2FrictionJoint)
		return casted.GetAnchorA(), casted.GetAnchorB()
	case jointTypeRope:
		casted := joint.(*box2d.B2RopeJoint)
		return casted.GetAnchorA(), casted.GetAnchorB()
	case jointTypeMotor:
		casted := joint.(*box2d.B2MotorJoint)
		return casted.GetAnchorA(), casted.GetAnchorB()
	default:
		return box2d.B2Vec2{}, box2d.B2Vec2{}
	}
}

func drawSolidPolygon(r engine.Renderer, vec2s []box2d.B2Vec2, color engine.Color) {
	if len(vec2s) == 0 {
		return
	}

	var next box2d.B2Vec2

	for i, current := range vec2s {
		if i == len(vec2s)-1 {
			next = vec2s[0]
		} else {
			next = vec2s[i+1]
		}

		r.DrawLine(color, engine.Line{
			A: vec2eng(current),
			B: vec2eng(next),
		})
	}
}

func drawSolidSegment(r engine.Renderer, v1 box2d.B2Vec2, v2 box2d.B2Vec2, color engine.Color) {
	r.DrawLine(color, engine.Line{
		A: vec2eng(v1),
		B: vec2eng(v2),
	})
}

func drawSolidCircle(r engine.Renderer, center box2d.B2Vec2, radius float64, axis box2d.B2Vec2, color engine.Color) {
	r.DrawCircle(color, engine.Circle{
		Pos:    vec2eng(center),
		Radius: radius,
	})

	r.DrawLine(color, engine.Line{
		A: vec2eng(center),
		B: vec2eng(center).PolarOffset(radius, vec2eng(axis).Direction()),
	})
}

func drawSolidPoint(r engine.Renderer, pos box2d.B2Vec2, color engine.Color) {
	worldPos := vec2eng(pos)

	r.DrawSquare(color, engine.RectScreen(int(worldPos.X-2), int(worldPos.Y-2), 4, 4))
}
