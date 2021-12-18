package control

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine/event"
)

type (
	Mouse struct {
		engineGUI engineGUI

		scrollPosition   float64
		scrollLastOffset float64

		leftState               state
		rightState              state
		propagationLockPriority int

		lastX int
		lastY int
	}
)

func NewMouse(dispatcher *event.Dispatcher, engineGUI engineGUI) *Mouse {
	m := &Mouse{
		engineGUI:               engineGUI,
		leftState:               stateUp,
		rightState:              stateUp,
		propagationLockPriority: 0,
	}
	m.subscribeToMouse(dispatcher)
	m.subscribeToMouseButton(dispatcher)
	m.subscribeToFrameStart(dispatcher)

	return m
}

func (m *Mouse) subscribeToMouse(dispatcher *event.Dispatcher) {
	dispatcher.OnMouseWheel(func(mouseWheel event.MouseWheelEvent) error {
		if m.engineGUI.CursorOnWindow() {
			// ignore mouse wheel event
			// because it already captured in engine window
			return nil
		}

		m.scrollPosition += mouseWheel.ScrollOffset
		m.scrollLastOffset = mouseWheel.ScrollOffset

		return nil
	})
}

func (m *Mouse) subscribeToMouseButton(dispatcher *event.Dispatcher) {
	dispatcher.OnMouseButton(func(mouseButtonEvent event.MouseButtonEvent) error {
		if m.engineGUI.CaptureMouse() {
			// click consumed by engine GUI layer, we need stop propagate click next to engine
			return nil
		}

		if mouseButtonEvent.IsLeft {
			switch {
			case mouseButtonEvent.IsPressed:
				m.leftState = statePressed
			case mouseButtonEvent.IsReleased:
				m.leftState = stateReleased
			}
			return nil
		}

		if mouseButtonEvent.IsRight {
			switch {
			case mouseButtonEvent.IsPressed:
				m.rightState = statePressed
			case mouseButtonEvent.IsReleased:
				m.rightState = stateReleased
			}
			return nil
		}

		return nil
	})
}

func (m *Mouse) subscribeToFrameStart(dispatcher *event.Dispatcher) {
	dispatcher.OnFrameStart(func(_ event.FrameStartEvent) error {
		// mouse state
		m.propagationLockPriority = 0

		// mouse scroll
		m.scrollLastOffset = 0

		// mouse buttons
		if m.leftState == statePressed {
			m.leftState = stateDown
		}
		if m.leftState == stateReleased {
			m.leftState = stateUp
		}

		if m.rightState == statePressed {
			m.rightState = stateDown
		}
		if m.rightState == stateReleased {
			m.rightState = stateUp
		}

		// mouse position
		x, y, _ := sdl.GetMouseState()
		if int32(m.lastX) != x || int32(m.lastY) != y {
			m.lastX = int(x)
			m.lastY = int(y)
			dispatcher.PublishEventMouseMove(event.MouseMoveEvent{
				X: m.lastX,
				Y: m.lastY,
			})
		}

		return nil
	})
}

func (m *Mouse) MouseCoords() galx.Vec {
	return galx.Vec{
		X: float64(m.lastX),
		Y: float64(m.lastY),
	}
}

func (m *Mouse) IsButtonsAvailable(priority int) bool {
	return priority >= m.propagationLockPriority
}

func (m *Mouse) StopPropagation(priority int) {
	if priority <= m.propagationLockPriority {
		return
	}

	m.propagationLockPriority = priority
}

func (m *Mouse) ScrollPosition() float64 {
	return m.scrollPosition
}

func (m *Mouse) ScrollLastOffset() float64 {
	return m.scrollLastOffset
}

func (m *Mouse) LeftPressed() bool {
	return m.leftState == statePressed
}

func (m *Mouse) LeftReleased() bool {
	return m.leftState == stateReleased
}

func (m *Mouse) LeftDown() bool {
	return m.leftState == stateDown
}

func (m *Mouse) RightPressed() bool {
	return m.rightState == statePressed
}

func (m *Mouse) RightReleased() bool {
	return m.rightState == stateReleased
}

func (m *Mouse) RightDown() bool {
	return m.rightState == stateDown
}
