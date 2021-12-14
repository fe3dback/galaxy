package scene

import (
	"fmt"
	"log"
	"runtime"

	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine/event"
	"github.com/fe3dback/galaxy/internal/engine/loader"
	"github.com/fe3dback/galaxy/internal/engine/node"
)

type (
	ID         = string
	blueprint  = func() []galx.GameObject
	blueprints = map[ID]blueprint
)

type Manager struct {
	assetsLoader      *loader.AssetsLoader
	componentRegistry *node.ComponentsRegistry

	blueprints   blueprints
	currentID    ID
	currentScene *Scene

	snapshotScene SerializedScene
	snapshotId    ID
	hasSnapshot   bool

	resetQueued           bool
	snapshotQueued        bool
	snapshotRestoreQueued bool
}

func NewManager(
	dispatcher *event.Dispatcher,
	assetsLoader *loader.AssetsLoader,
	componentRegistry *node.ComponentsRegistry,
	includeEditor bool,
) *Manager {
	manager := &Manager{
		assetsLoader:      assetsLoader,
		componentRegistry: componentRegistry,
		blueprints:        make(blueprints),
	}

	if includeEditor {
		dispatcher.OnKeyBoard(manager.handleKeyboard)
	}

	dispatcher.OnFrameStart(manager.handleFrameStart)

	return manager
}

func (m *Manager) SaveSnapshot(force bool) {
	if force {
		m.snapshot()
		return
	}

	m.snapshotQueued = true
}

func (m *Manager) RestoreFromSnapshot(force bool) {
	if force {
		m.snapshotRestore()
		return
	}

	m.snapshotRestoreQueued = true
}

func (m *Manager) CurrentSceneID() ID {
	return m.currentID
}

func (m *Manager) Current() galx.Scene {
	return m.currentScene
}

func (m *Manager) Switch(nextID ID) {
	if _, ok := m.blueprints[nextID]; !ok {
		panic(fmt.Errorf("failed switch scene from '%s' to '%s'. Next scene not exist", m.currentID, nextID))
	}

	previousID := m.currentID

	// destroy current
	if m.currentScene != nil {
		m.currentScene.destroy()
		runtime.GC()
	}

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
	if m.resetQueued {
		m.resetQueued = false
		m.reset()
	}

	if m.snapshotQueued {
		m.snapshotQueued = false
		m.snapshot()
	}

	if m.snapshotRestoreQueued {
		m.snapshotRestoreQueued = false
		m.snapshotRestore()
	}

	return nil
}

func (m *Manager) reset() {
	log.Println(fmt.Sprintf("Resetting current scene '%s'..", m.currentID))

	// switch to same scene
	// it will recreate it
	m.Switch(m.currentID)
}

func (m *Manager) snapshot() {
	log.Println(fmt.Sprintf("Snaphoting scene '%s'..", m.currentID))

	serialized, err := m.marshalScene(m.currentScene)
	if err != nil {
		panic(fmt.Errorf("failed marshal scene for snapshot: %w", err))
	}

	m.snapshotId = m.currentID
	m.snapshotScene = *serialized
	m.hasSnapshot = true
}

func (m *Manager) snapshotRestore() {
	log.Println(fmt.Sprintf("Restore scene '%s' from snapshot..", m.snapshotId))

	if !m.hasSnapshot {
		// if we not have snapshot, just rollback to last saved state on disk
		m.reset()
		return
	}

	if m.currentID != m.snapshotId {
		// if scene changed during game, just rollback to last saved state on disk
		m.reset()
		return
	}

	scene, err := m.unmarshalScene(&m.snapshotScene)
	if err != nil {
		panic(fmt.Errorf("failed restore from snaphot: unmarshal: %w", err))
	}

	m.currentScene = scene
	m.snapshotId = ""
	m.snapshotScene = SerializedScene{}
	m.hasSnapshot = false
}

func createSceneFromBlueprint(bp blueprint) *Scene {
	return NewScene(bp())
}
