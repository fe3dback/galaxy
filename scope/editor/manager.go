package editor

import (
	"fmt"
	"reflect"

	"github.com/fe3dback/galaxy/galx"
)

type Manager struct {
	components []component
}

func NewManager(components ...component) *Manager {
	m := &Manager{
		components: components,
	}

	m.onCreate()
	return m
}

func (m *Manager) onCreate() {
	for _, component := range m.components {
		lc, needInit := component.(galx.EditorComponentCycleCreated)
		if !needInit {
			continue
		}

		lc.OnCreated(m.resolveComponentDependency)
	}
}

// should check slot type
// and return component, by slot.Id
func (m *Manager) resolveComponentDependency(slot galx.EditorComponentIdentifiable) galx.EditorComponentIdentifiable {
	if _, ok := slot.(component); !ok {
		panic(fmt.Sprintf("editor component require slot with type '%T', but only another components allowed", slot))
	}

	requiredComponent := reflect.Zero(reflect.TypeOf(slot).Elem()).Interface().(galx.EditorComponentIdentifiable)
	for _, editorComponent := range m.components {
		if editorComponent.Id() == requiredComponent.Id() {
			return editorComponent
		}
	}

	panic(fmt.Sprintf("editor component require slot with type '%T', but this component not registered in editor", slot))
}

func (m *Manager) OnUpdate(s galx.State) error {
	for _, component := range m.components {
		if lc, ok := component.(componentLifeCycleUpdate); ok {
			err := lc.OnUpdate(s)
			if err != nil {
				return fmt.Errorf("can`t update editor component '%T': %w", component, err)
			}
		}
	}

	return nil
}

func (m *Manager) OnBeforeDraw(r galx.Renderer) error {
	for _, component := range m.components {
		if lc, ok := component.(componentLifeCycleBeforeDraw); ok {
			err := lc.OnBeforeDraw(r)
			if err != nil {
				return fmt.Errorf("can`t before draw editor component '%T': %w", component, err)
			}
		}
	}

	return nil
}

func (m *Manager) OnAfterDraw(r galx.Renderer) error {
	for _, component := range m.components {
		if lc, ok := component.(componentLifeCycleAfterDraw); ok {
			err := lc.OnAfterDraw(r)
			if err != nil {
				return fmt.Errorf("can`t after draw editor component '%T': %w", component, err)
			}
		}
	}

	return nil
}
