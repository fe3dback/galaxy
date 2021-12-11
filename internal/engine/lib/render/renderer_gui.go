package render

import (
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/veandco/go-sdl2/sdl"
)

func (r *Renderer) StartGUIFrame() {
	sizeW, sizeH := r.window.GetSize()

	// Setup display size (every frame to accommodate for window resizing)
	r.gui.SetDisplaySize(imgui.Vec2{X: float32(sizeW), Y: float32(sizeH)})

	// Setup time step (we don't use SDL_GetTicks() because it is using millisecond resolution)
	frequency := sdl.GetPerformanceFrequency()
	currentTime := sdl.GetPerformanceCounter()
	if r.guiTime > 0 {
		r.gui.SetDeltaTime(float32(currentTime-r.guiTime) / float32(frequency))
	} else {
		const fallbackDelta = 1.0 / 60.0
		r.gui.SetDeltaTime(fallbackDelta)
	}
	r.guiTime = currentTime

	// If a mouse press event came, always pass it as "mouse held this frame", so we don't miss click-release events that are shorter than 1 frame.
	x, y, _ := sdl.GetMouseState()
	r.gui.SetMousePosition(imgui.Vec2{X: float32(x), Y: float32(y)})
	r.gui.SetMouseButtonDown(0, r.guiMousePressedLeft)
	r.gui.SetMouseButtonDown(1, r.guiMousePressedRight)

	imgui.NewFrame()
}

func (r *Renderer) EndGUIFrame() {
	// render GUI to buffer
	imgui.Render()

	// get window vars
	sizeW, sizeH := r.window.GetSize()
	fbW, fbH := r.window.GLGetDrawableSize()

	// copy GUI buffer to openGL buffer
	r.guiRenderer.Render(
		[2]float32{float32(sizeW), float32(sizeH)},
		[2]float32{float32(fbW), float32(fbH)},
		imgui.RenderedDrawData(),
	)
}
