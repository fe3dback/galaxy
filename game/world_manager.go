package game

import (
	"fmt"
	"runtime"

	"github.com/fe3dback/galaxy/engine/lib/event"
)

type (
	WorldProviderFn func() *World

	WorldManager struct {
		current       *World
		createWorldFn WorldProviderFn
		resetQueued   bool
	}
)

func NewWorldManager(createWorldFn WorldProviderFn, dispatcher *event.Dispatcher) *WorldManager {
	manager := &WorldManager{
		current:       createWorldFn(),
		createWorldFn: createWorldFn,
		resetQueued:   false,
	}
	dispatcher.OnKeyBoard(manager.handleKeyboard)

	return manager
}

func (w *WorldManager) handleKeyboard(keyboard event.EvKeyboard) error {
	if keyboard.PressType != event.KeyboardPressTypePressed {
		return nil
	}

	if keyboard.Key != event.KeyF4 {
		return nil
	}

	w.resetQueued = true
	return nil
}

func (w *WorldManager) OnFrameStart() {
	if !w.resetQueued {
		return
	}

	w.resetQueued = false
	w.Reset()
}

func (w *WorldManager) Reset() {
	fmt.Println("Resetting world..")

	w.current = nil
	runtime.GC()
	w.current = w.createWorldFn()
}

func (w *WorldManager) CurrentWorld() *World {
	return w.current
}
