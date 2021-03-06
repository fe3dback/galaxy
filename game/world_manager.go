package game

import (
	"log"
	"runtime"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/event"
)

type (
	WorldProviderFn func(creator engine.WorldCreator) *World

	WorldManager struct {
		current       *World
		worldCreator  engine.WorldCreator
		createWorldFn WorldProviderFn
		resetQueued   bool
	}
)

func NewWorldManager(createWorldFn WorldProviderFn, gameWorldCreator engine.WorldCreator, dispatcher *event.Dispatcher) *WorldManager {
	manager := &WorldManager{
		current:       createWorldFn(gameWorldCreator),
		worldCreator:  gameWorldCreator,
		createWorldFn: createWorldFn,
		resetQueued:   false,
	}
	dispatcher.OnKeyBoard(manager.handleKeyboard)

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

func (w *WorldManager) OnFrameStart() {
	if !w.resetQueued {
		return
	}

	w.resetQueued = false
	w.Reset()
}

func (w *WorldManager) Reset() {
	log.Println("Resetting world..")

	w.current = nil
	runtime.GC()
	w.current = w.createWorldFn(w.worldCreator)
}

func (w *WorldManager) CurrentWorld() *World {
	return w.current
}
