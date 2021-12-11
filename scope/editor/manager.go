package editor

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
)

type Manager struct {
	components []Component
}

func NewManager(components []Component) *Manager {
	return &Manager{
		components: components,
	}
}

func (m *Manager) OnUpdate(s galx.State) error {
	for _, component := range m.components {
		err := component.OnUpdate(s)
		if err != nil {
			return fmt.Errorf("can`t update editor component '%T': %w", component, err)
		}
	}

	return nil
}

func (m *Manager) OnDraw(r galx.Renderer) error {
	for _, component := range m.components {
		err := component.OnDraw(r)
		if err != nil {
			return fmt.Errorf("can`t draw editor component '%T': %w", component, err)
		}
	}

	return nil
}
