package editor

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/event"
)

type Manager struct {
	dispatcher *event.Dispatcher
	components []Component
}

func NewManager(components []Component) *Manager {
	return &Manager{
		components: components,
	}
}

func (m *Manager) OnUpdate(s engine.State) error {
	for _, component := range m.components {
		err := component.OnUpdate(s)
		if err != nil {
			return fmt.Errorf("can`t update editor component '%T': %v", component, err)
		}
	}

	return nil
}

func (m *Manager) OnDraw(r engine.Renderer) error {
	for _, component := range m.components {
		err := component.OnDraw(r)
		if err != nil {
			return fmt.Errorf("can`t draw editor component '%T': %v", component, err)
		}
	}

	return nil
}
