package control

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/event"
	"github.com/veandco/go-sdl2/sdl"
)

type Mouse struct {
	scrollPosition   float64
	scrollLastOffset float64
}

func NewMouse(dispatcher *event.Dispatcher) *Mouse {
	m := &Mouse{}
	m.subscribeToMouse(dispatcher)
	m.subscribeToFrameStart(dispatcher)

	return m
}

func (m *Mouse) subscribeToMouse(dispatcher *event.Dispatcher) {
	dispatcher.OnMouseWheel(func(mouseWheel event.MouseWheelEvent) error {
		m.scrollPosition += mouseWheel.ScrollOffset
		m.scrollLastOffset = mouseWheel.ScrollOffset

		return nil
	})
}

func (m *Mouse) subscribeToFrameStart(dispatcher *event.Dispatcher) {
	dispatcher.OnFrameStart(func(_ event.FrameStartEvent) error {
		m.scrollLastOffset = 0

		return nil
	})
}

func (m *Mouse) MouseCoords() engine.Vec {
	x, y, _ := sdl.GetMouseState()

	return engine.Vec{
		X: float64(x),
		Y: float64(y),
	}
}

func (m *Mouse) ScrollPosition() float64 {
	return m.scrollPosition
}

func (m *Mouse) ScrollLastOffset() float64 {
	return m.scrollLastOffset
}
