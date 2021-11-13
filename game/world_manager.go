package game

import (
	"log"
	"runtime"

	"github.com/fe3dback/galaxy/engine/event"
)

type (
	WorldManager struct {
		current     *World
		resetQueued bool
	}
)

func NewWorldManager(dispatcher *event.Dispatcher) *WorldManager {
	manager := &WorldManager{
		current:     NewWorld(),
		resetQueued: false,
	}
	dispatcher.OnKeyBoard(manager.handleKeyboard)
	dispatcher.OnFrameStart(manager.handleFrameStart)

	return manager
}

func (w *WorldManager) handleKeyboard(keyboard event.KeyBoardEvent) error {
	if keyboard.PressType != event.KeyboardPressTypePressed {
		return nil
	}

	if keyboard.Key != event.KeyF4 {
		return nil
	}

	w.resetQueued = true
	return nil
}

func (w *WorldManager) handleFrameStart(_ event.FrameStartEvent) error {
	if !w.resetQueued {
		return nil
	}

	w.resetQueued = false
	w.reset()
	return nil
}

func (w *WorldManager) reset() {
	log.Println("Resetting world..")

	w.current = nil
	runtime.GC()
	w.current = NewWorld()
}

func (w *WorldManager) CurrentWorld() *World {
	return w.current
}
