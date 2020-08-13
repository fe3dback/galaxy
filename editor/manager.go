package editor

import (
	"github.com/fe3dback/galaxy/engine"
)

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func (m Manager) OnUpdate(_ engine.State) error {
	return nil
}

func (m Manager) OnDraw(_ engine.Renderer) error {
	return nil
}
