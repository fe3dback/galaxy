package control

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/event"
)

type Movement struct {
	vector engine.Vec

	pressedTop    bool
	pressedBottom bool
	pressedLeft   bool
	pressedRight  bool
	pressedShift  bool
	pressedSpace  bool
}

func NewMovement(dispatcher *event.Dispatcher) *Movement {
	m := &Movement{
		vector:        engine.Vec{},
		pressedTop:    false,
		pressedBottom: false,
		pressedLeft:   false,
		pressedRight:  false,
		pressedShift:  false,
		pressedSpace:  false,
	}
	m.subscribeToKeyboard(dispatcher)

	return m
}

func (m *Movement) Vector() engine.Vec {
	return m.vector
}

func (m *Movement) Shift() bool {
	return m.pressedShift
}

func (m *Movement) Space() bool {
	return m.pressedSpace
}

func (m *Movement) subscribeToKeyboard(dispatcher *event.Dispatcher) {
	dispatcher.OnKeyBoard(func(keyboard event.KeyBoardEvent) error {
		var pressed bool

		if keyboard.PressType == event.KeyboardPressTypePressed {
			pressed = true
		}
		if keyboard.PressType == event.KeyboardPressTypeReleased {
			pressed = false
		}

		if keyboard.Key == event.KeyA {
			m.pressedLeft = pressed
		}
		if keyboard.Key == event.KeyD {
			m.pressedRight = pressed
		}
		if keyboard.Key == event.KeyW {
			m.pressedTop = pressed
		}
		if keyboard.Key == event.KeyS {
			m.pressedBottom = pressed
		}
		if keyboard.Key == event.KeyLshift {
			m.pressedShift = pressed
		}
		if keyboard.Key == event.KeySpace || keyboard.Key == event.KeyKpSpace {
			m.pressedSpace = pressed
		}

		m.update()
		return nil
	})
}

func (m *Movement) update() {
	m.vector.X = 0
	m.vector.Y = 0

	if m.pressedLeft {
		m.vector.X--
	}
	if m.pressedRight {
		m.vector.X++
	}
	if m.pressedTop {
		m.vector.Y--
	}
	if m.pressedBottom {
		m.vector.Y++
	}

	m.vector = m.vector.Normalize()
}
