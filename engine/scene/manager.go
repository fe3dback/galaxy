package scene

import (
	"fmt"
	"log"
	"runtime"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/event"
)

const blueprintIDDefaultScene = "_empty"

type sceneID = string

type Manager struct {
	blueprints   map[sceneID]engine.SceneBlueprint
	currentID    sceneID
	currentScene *Scene

	resetQueued bool
}

func NewManager(dispatcher *event.Dispatcher) *Manager {
	defBlueprint := emptySceneBlueprint{}

	blueprints := make(map[string]engine.SceneBlueprint)
	blueprints[blueprintIDDefaultScene] = defBlueprint

	manager := &Manager{
		blueprints:   blueprints,
		currentID:    blueprintIDDefaultScene,
		currentScene: createSceneFromBlueprint(defBlueprint),
	}
	dispatcher.OnKeyBoard(manager.handleKeyboard)
	dispatcher.OnFrameStart(manager.handleFrameStart)

	return manager
}

func (m *Manager) CurrentSceneID() string {
	return m.currentID
}

func (m *Manager) AddBlueprint(ID string, blueprint engine.SceneBlueprint) {
	if _, ok := m.blueprints[ID]; ok {
		panic(fmt.Errorf("scene blueprint '%s' already exist", ID))
	}

	m.blueprints[ID] = blueprint
}

func (m *Manager) Current() engine.Scene {
	return m.currentScene
}

func (m *Manager) Switch(nextID string) {
	if _, ok := m.blueprints[nextID]; !ok {
		panic(fmt.Errorf("failed switch scene from '%s' to '%s'. Next scene not exist", m.currentID, nextID))
	}

	previousID := m.currentID

	// destroy current
	m.currentScene.destroy()
	runtime.GC()

	// create from blueprint
	m.currentID = nextID
	m.currentScene = createSceneFromBlueprint(
		m.blueprints[nextID],
	)

	log.Println(fmt.Sprintf("scene switched from '%s' to '%s'", previousID, nextID))
}

func (m *Manager) handleKeyboard(keyboard event.KeyBoardEvent) error {
	if keyboard.PressType != event.KeyboardPressTypePressed {
		return nil
	}

	if keyboard.Key != event.KeyF4 {
		return nil
	}

	m.resetQueued = true
	return nil
}

func (m *Manager) handleFrameStart(_ event.FrameStartEvent) error {
	if !m.resetQueued {
		return nil
	}

	m.resetQueued = false
	m.reset()
	return nil
}

func (m *Manager) reset() {
	log.Println(fmt.Sprintf("Resetting current scene '%s'..", m.currentID))

	// switch to same scene
	// it will recreate it
	m.Switch(m.currentID)
}

func createSceneFromBlueprint(p engine.SceneBlueprint) *Scene {
	return NewScene(
		p.CreateEntities(),
	)
}
