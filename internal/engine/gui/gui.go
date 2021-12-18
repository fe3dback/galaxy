package gui

import (
	"github.com/inkyblackness/imgui-go/v4"

	"github.com/fe3dback/galaxy/internal/engine"
	"github.com/fe3dback/galaxy/internal/engine/event"
	"github.com/fe3dback/galaxy/internal/engine/gui/renderers"
	"github.com/fe3dback/galaxy/internal/utils"
)

// Renderer covers rendering imgui draw data.
type Renderer interface {
	// PreRender causes the display buffer to be prepared for new output.
	PreRender(clearColor [3]float32)
	// Render draws the provided imgui draw data.
	Render(displaySize [2]float32, framebufferSize [2]float32, drawData imgui.DrawData)
	// Dispose close renderer
	Dispose()
}

type Gui struct {
	io       imgui.IO
	renderer Renderer

	// engine state from updates
	windowWidth       float32
	windowHeight      float32
	deltaTime         float32
	mouseX            float32
	mouseY            float32
	mouseWheel        float32
	mousePressedLeft  bool
	mousePressedRight bool
}

func NewGUI(closer *utils.Closer, renderTech engine.RenderTech, dispatcher *event.Dispatcher) *Gui {
	gui := createGUI(closer, renderTech)
	gui.subscribe(dispatcher)

	return gui
}

func createGUI(closer *utils.Closer, renderTech engine.RenderTech) *Gui {
	// fallback settings for first frame
	gui := &Gui{
		windowWidth:  320,
		windowHeight: 240,
		deltaTime:    1 / 60.0,
	}

	context := imgui.CreateContext(nil)
	closer.EnqueueFree(context.Destroy)
	gui.io = imgui.CurrentIO()

	// create gui renderer
	var renderer Renderer
	var err error

	switch renderTech {
	case engine.RenderTechOpenGL2:
		renderer, err = renderers.NewOpenGL2(gui.io)
	case engine.RenderTechOpenGL3:
		renderer, err = renderers.NewOpenGL3(gui.io)
	}

	utils.Check("create GUI renderer", err)
	closer.EnqueueFree(renderer.Dispose)
	gui.renderer = renderer

	return gui
}

func (g *Gui) subscribe(dispatcher *event.Dispatcher) {
	dispatcher.OnCameraUpdate(func(cameraUpdateEvent event.CameraUpdateEvent) error {
		g.windowWidth = float32(cameraUpdateEvent.Width)
		g.windowHeight = float32(cameraUpdateEvent.Height)
		return nil
	})

	dispatcher.OnMouseButton(func(mouseButtonEvent event.MouseButtonEvent) error {
		if mouseButtonEvent.IsLeft {
			g.mousePressedLeft = mouseButtonEvent.IsPressed
			return nil
		}

		if mouseButtonEvent.IsRight {
			g.mousePressedRight = mouseButtonEvent.IsPressed
			return nil
		}

		return nil
	})

	dispatcher.OnMouseMove(func(mouseMoveEvent event.MouseMoveEvent) error {
		g.mouseX = float32(mouseMoveEvent.X)
		g.mouseY = float32(mouseMoveEvent.Y)
		return nil
	})

	dispatcher.OnMouseWheel(func(mouseWheelEvent event.MouseWheelEvent) error {
		g.mouseWheel = float32(mouseWheelEvent.ScrollOffset)
		return nil
	})

	dispatcher.OnFrameEnd(func(frameEndEvent event.FrameEndEvent) error {
		if frameEndEvent.DeltaTime > 0 {
			g.deltaTime = float32(frameEndEvent.DeltaTime)
			return nil
		}

		// fallback
		g.deltaTime = 1 / 60.0
		return nil
	})
}

func (g *Gui) CaptureMouse() bool {
	return g.io.WantCaptureMouse()
}

func (g *Gui) CaptureKeyboard() bool {
	return g.io.WantCaptureKeyboard()
}

func (g *Gui) CursorOnWindow() bool {
	return imgui.IsWindowHoveredV(imgui.HoveredFlagsAnyWindow)
}

func (g *Gui) StartGUIFrame() {
	g.io.SetDeltaTime(g.deltaTime)
	g.io.SetMousePosition(imgui.Vec2{X: g.mouseX, Y: g.mouseY})
	g.io.SetMouseButtonDown(0, g.mousePressedLeft)
	g.io.SetMouseButtonDown(1, g.mousePressedRight)
	g.io.SetDisplaySize(imgui.Vec2{X: g.windowWidth, Y: g.windowHeight})
	g.io.AddMouseWheelDelta(0, g.mouseWheel)

	imgui.NewFrame()
}

func (g *Gui) EndGUIFrame() {
	imgui.Render()
	g.renderer.Render(
		[2]float32{g.windowWidth, g.windowHeight},
		[2]float32{g.windowWidth, g.windowHeight},
		imgui.RenderedDrawData(),
	)
}
