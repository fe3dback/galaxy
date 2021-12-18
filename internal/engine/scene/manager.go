package scene

import (
	"fmt"
	"log"
	"runtime"

	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine/assets"
	"github.com/fe3dback/galaxy/internal/engine/event"
	"github.com/fe3dback/galaxy/internal/engine/node"
)

const (
	modeEdit mode = "edit"
	modeGame mode = "game"
)

type (
	ID        = string
	snapshots = map[ID]*SnapshotScene
	mode      string
)

type Manager struct {
	assetsManager     *assets.Manager
	componentRegistry *node.ComponentsRegistry
	editorIncluded    bool

	mode         mode
	snapshotID   ID
	snapshots    snapshots
	sceneMutable *Scene
	queueReset   bool
	queueSave    bool
}

func NewManager(
	dispatcher *event.Dispatcher,
	assetsManager *assets.Manager,
	componentRegistry *node.ComponentsRegistry,
	editorIncluded bool,
) *Manager {
	manager := &Manager{
		assetsManager:     assetsManager,
		componentRegistry: componentRegistry,
		snapshots:         make(snapshots),
		editorIncluded:    editorIncluded,
	}

	if editorIncluded {
		manager.mode = modeEdit
		manager.snapshotID = ""
	} else {
		manager.mode = modeGame
	}

	// load initial state from assets
	defaultSceneID, initialSnapshots := manager.loadSnapshots()
	manager.snapshots = initialSnapshots
	manager.Switch(defaultSceneID)

	// subscribe to events
	if editorIncluded {
		dispatcher.OnKeyBoard(manager.handleKeyboard)
	}

	dispatcher.OnFrameStart(manager.handleFrameStart)

	// return
	return manager
}

func (m *Manager) StateToGameMode() {
	if !m.editorIncluded {
		// ignore in exported game
		return
	}

	if m.mode == modeGame {
		// already in game mode
		return
	}

	if m.sceneMutable == nil {
		panic("failed snapshot editor scene, is not initialized yet")
	}

	// save current scene to snapshot
	m.snapshotID = m.sceneMutable.ID()
	m.snapshots[m.snapshotID] = m.encodeScene(*m.sceneMutable)

	// decode it back to game mode
	// also test encoded changes here
	m.sceneMutable = m.decodeScene(*m.snapshots[m.snapshotID])
	m.mode = modeGame
}

func (m *Manager) StateToEditorMode() {
	if !m.editorIncluded {
		// ignore in exported game
		return
	}

	if m.mode == modeEdit {
		// already in edit mode
		return
	}

	if m.snapshotID == "" {
		panic("failed restore scene from editor snapshot, because snapshot not exist")
	}

	// delete game scene
	m.sceneMutable.destroy()
	m.sceneMutable = nil
	runtime.GC()

	// restore editor scene back from snapshot
	m.sceneMutable = m.decodeScene(*m.snapshots[m.snapshotID])
	m.snapshotID = ""
	m.mode = modeEdit
}

func (m *Manager) CurrentSceneID() ID {
	return m.sceneMutable.ID()
}

func (m *Manager) Current() galx.Scene {
	return m.sceneMutable
}

func (m *Manager) Switch(nextID ID) {
	currentID := "nil"
	if m.sceneMutable != nil {
		currentID = m.sceneMutable.ID()
	}

	if _, ok := m.snapshots[nextID]; !ok {
		panic(fmt.Errorf("failed switch scene from '%s' to '%s'. Next scene not exist", currentID, nextID))
	}

	// destroy current
	if m.sceneMutable != nil {
		m.sceneMutable.destroy()
		runtime.GC()
	}

	// create from snapshot
	m.sceneMutable = m.decodeScene(*m.snapshots[nextID])
	log.Println(fmt.Sprintf("scene switched from '%s' to '%s'", currentID, nextID))
}

func (m *Manager) handleKeyboard(keyboard event.KeyBoardEvent) error {
	if keyboard.PressType != event.KeyboardPressTypePressed {
		return nil
	}

	if keyboard.Key == event.KeyF4 {
		m.queueReset = true
	}

	if keyboard.Key == event.KeyF6 {
		m.queueSave = true
	}

	return nil
}

func (m *Manager) handleFrameStart(_ event.FrameStartEvent) error {
	if m.queueReset {
		m.queueReset = false
		m.reset()
	}

	if m.queueSave {
		m.queueSave = false
		m.save()
	}

	return nil
}

func (m *Manager) guardEditor(cb func()) {
	if !m.editorIncluded {
		// ignore in exported game
		return
	}

	if m.mode != modeEdit {
		// ignore reset commands on game mode
		return
	}

	cb()
}

func (m *Manager) reset() {
	m.guardEditor(func() {
		log.Println(fmt.Sprintf("Resetting current scene '%s'..", m.sceneMutable.ID()))

		// switch to same scene
		// it will recreate it
		m.Switch(m.sceneMutable.ID())
	})
}

func (m *Manager) save() {
	m.guardEditor(func() {
		log.Println(fmt.Sprintf("Saving current scene '%s'..", m.sceneMutable.ID()))

		m.saveSnapshot(
			m.encodeScene(*m.sceneMutable),
		)
	})
}
