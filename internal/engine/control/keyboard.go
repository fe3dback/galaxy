package control

import (
	"strings"

	"github.com/fe3dback/galaxy/internal/engine/event"
)

type (
	Keyboard struct {
		mapping keysMapping
	}

	keysMapping = map[string]state
)

func NewKeyboard(dispatcher *event.Dispatcher) *Keyboard {
	kb := &Keyboard{
		mapping: make(keysMapping),
	}
	kb.subscribeToFrameStart(dispatcher)
	kb.subscribeToKeyboard(dispatcher)

	return kb
}

func (kb *Keyboard) IsPressed(key rune) bool {
	if state, ok := kb.mapping[kb.unRune(key)]; ok {
		return state == statePressed
	}

	return false
}

func (kb *Keyboard) IsReleased(key rune) bool {
	if state, ok := kb.mapping[kb.unRune(key)]; ok {
		return state == stateReleased
	}

	return false
}

func (kb *Keyboard) IsDown(key rune) bool {
	if state, ok := kb.mapping[kb.unRune(key)]; ok {
		return state == stateDown
	}

	return false
}

func (kb *Keyboard) unRune(key rune) string {
	return strings.ToLower(string(key))
}

func (kb *Keyboard) subscribeToFrameStart(dispatcher *event.Dispatcher) {
	dispatcher.OnFrameStart(func(frameStartEvent event.FrameStartEvent) error {
		kb.updateFrameState()
		return nil
	})
}

func (kb *Keyboard) subscribeToKeyboard(dispatcher *event.Dispatcher) {
	dispatcher.OnKeyBoard(func(keyBoardEvent event.KeyBoardEvent) error {
		state := kb.mapState(keyBoardEvent.PressType)
		if state == stateUnknown {
			return nil
		}

		if state == statePressed && kb.IsDown(rune(keyBoardEvent.Key)) {
			// already down (exclude signal bags)
			return nil
		}

		kb.mapping[string(keyBoardEvent.Key)] = state
		return nil
	})
}

func (kb *Keyboard) mapState(pt event.KeyboardPressType) state {
	if pt == event.KeyboardPressTypePressed {
		return statePressed
	}

	if pt == event.KeyboardPressTypeReleased {
		return stateReleased
	}

	return stateUnknown
}

func (kb *Keyboard) updateFrameState() {
	for key, state := range kb.mapping {
		if state == statePressed {
			kb.mapping[key] = stateDown
			continue
		}

		if state == stateReleased {
			kb.mapping[key] = stateUp
			continue
		}
	}
}
