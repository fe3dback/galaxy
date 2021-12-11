package control

import (
	"github.com/fe3dback/galaxy/galx"
	event2 "github.com/fe3dback/galaxy/internal/engine/event"
)

type Movement struct {
	vector galx.Vec

	pressedTop    bool
	pressedBottom bool
	pressedLeft   bool
	pressedRight  bool
	pressedShift  bool
	pressedSpace  bool
}

func NewMovement(dispatcher *event2.Dispatcher) *Movement {
	m := &Movement{
		vector:        galx.Vec{},
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

func (m *Movement) Vector() galx.Vec {
	return m.vector
}

func (m *Movement) Shift() bool {
	return m.pressedShift
}

func (m *Movement) Space() bool {
	return m.pressedSpace
}

func (m *Movement) subscribeToKeyboard(dispatcher *event2.Dispatcher) {
	dispatcher.OnKeyBoard(func(keyboard event2.KeyBoardEvent) error {
		var pressed bool

		if keyboard.PressType == event2.KeyboardPressTypePressed {
			pressed = true
		}
		if keyboard.PressType == event2.KeyboardPressTypeReleased {
			pressed = false
		}

		if keyboard.Key == event2.KeyA {
			m.pressedLeft = pressed
		}
		if keyboard.Key == event2.KeyD {
			m.pressedRight = pressed
		}
		if keyboard.Key == event2.KeyW {
			m.pressedTop = pressed
		}
		if keyboard.Key == event2.KeyS {
			m.pressedBottom = pressed
		}
		if keyboard.Key == event2.KeyLshift {
			m.pressedShift = pressed
		}
		if keyboard.Key == event2.KeySpace || keyboard.Key == event2.KeyKpSpace {
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
