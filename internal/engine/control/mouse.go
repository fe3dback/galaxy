package control

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/galx"
	event2 "github.com/fe3dback/galaxy/internal/engine/event"
)

type Mouse struct {
	scrollPosition   float64
	scrollLastOffset float64
}

func NewMouse(dispatcher *event2.Dispatcher) *Mouse {
	m := &Mouse{}
	m.subscribeToMouse(dispatcher)
	m.subscribeToFrameStart(dispatcher)

	return m
}

func (m *Mouse) subscribeToMouse(dispatcher *event2.Dispatcher) {
	dispatcher.OnMouseWheel(func(mouseWheel event2.MouseWheelEvent) error {
		m.scrollPosition += mouseWheel.ScrollOffset
		m.scrollLastOffset = mouseWheel.ScrollOffset

		return nil
	})
}

func (m *Mouse) subscribeToFrameStart(dispatcher *event2.Dispatcher) {
	dispatcher.OnFrameStart(func(_ event2.FrameStartEvent) error {
		m.scrollLastOffset = 0

		return nil
	})
}

func (m *Mouse) MouseCoords() galx.Vec {
	x, y, _ := sdl.GetMouseState()

	return galx.Vec{
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
