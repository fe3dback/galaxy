package control

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine/event"
)

const (
	statePressed state = iota
	stateReleased
	stateDown
	stateUp
)

type (
	state = uint8

	guiIO interface {
		WantCaptureMouse() bool
	}

	Mouse struct {
		guiIO guiIO

		scrollPosition   float64
		scrollLastOffset float64

		leftState               state
		rightState              state
		propagationLockPriority int
		propagationPriority     int
	}
)

func NewMouse(dispatcher *event.Dispatcher, guiIO guiIO) *Mouse {
	m := &Mouse{
		guiIO:                   guiIO,
		leftState:               stateUp,
		rightState:              stateUp,
		propagationLockPriority: 0,
		propagationPriority:     0,
	}
	m.subscribeToMouse(dispatcher)
	m.subscribeToMouseButton(dispatcher)
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

func (m *Mouse) subscribeToMouseButton(dispatcher *event.Dispatcher) {
	dispatcher.OnMouseButton(func(mouseButtonEvent event.MouseButtonEvent) error {
		if m.guiIO.WantCaptureMouse() {
			// click consumed by GUI layer, we need stop propagate click next to engine
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
		m.propagationLockPriority = 0
		m.scrollLastOffset = 0

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

func (m *Mouse) SetPriority(priority int) {
	m.propagationPriority = priority
}

func (m *Mouse) ResetPriority() {
	m.propagationPriority = 0
}

func (m *Mouse) StopPropagation() {
	if m.propagationLockPriority > m.propagationPriority {
		return
	}

	m.propagationLockPriority = m.propagationPriority
}

func (m *Mouse) ScrollPosition() float64 {
	return m.scrollPosition
}

func (m *Mouse) ScrollLastOffset() float64 {
	return m.scrollLastOffset
}

func (m *Mouse) LeftPressed() bool {
	if m.propagationPriority < m.propagationLockPriority {
		return false
	}

	return m.leftState == statePressed
}

func (m *Mouse) LeftReleased() bool {
	if m.propagationPriority < m.propagationLockPriority {
		return false
	}

	return m.leftState == stateReleased
}

func (m *Mouse) LeftDown() bool {
	if m.propagationPriority < m.propagationLockPriority {
		return false
	}

	return m.leftState == stateDown
}

func (m *Mouse) RightPressed() bool {
	if m.propagationPriority < m.propagationLockPriority {
		return false
	}

	return m.rightState == statePressed
}

func (m *Mouse) RightReleased() bool {
	if m.propagationPriority < m.propagationLockPriority {
		return false
	}

	return m.rightState == stateReleased
}

func (m *Mouse) RightDown() bool {
	if m.propagationPriority < m.propagationLockPriority {
		return false
	}

	return m.rightState == stateDown
}
