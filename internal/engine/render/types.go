package render

import "github.com/fe3dback/galaxy/galx"

type (
	renderer interface {
		// Draw calls

		Clear(color galx.Color)

		// System

		FrameStart()
		FrameEnd()
		Draw()
		DrawTmpTriangle() // todo: tmp
		DrawRect(vertexPos [4]galx.Vec2d, vertexColor [4]galx.Vec3d)
		GPUWait()
	}
)
