package render

type (
	renderer interface {
		// Draw calls

		Clear(color uint32)

		// System

		FrameStart()
		FrameEnd()
		Draw()
		DrawTriangle() // todo: tmp
		GPUWait()
	}
)
