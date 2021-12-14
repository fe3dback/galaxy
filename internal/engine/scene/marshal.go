package scene

import (
	"github.com/fe3dback/galaxy/galx"
)

func (m *Manager) marshalScene(scene *Scene) (*SerializedScene, error) {
	sGameObjects := make([]SerializedGameObject, 0, len(scene.entities))

	for _, entity := range scene.entities {
		sGameObjects = append(sGameObjects, m.marshalGameObject(entity))
	}

	return &SerializedScene{
		GameObjects: sGameObjects,
	}, nil
}

func (m *Manager) marshalGameObject(gameObject galx.GameObject) SerializedGameObject {
	components := gameObject.Components()
	child := gameObject.Child()

	sComponents := make([]SerializedComponent, 0, len(components))
	for _, component := range components {
		sComponents = append(sComponents, m.marshalComponent(component))
	}

	sChild := make([]SerializedGameObject, 0, len(child))
	for _, childObject := range child {
		sChild = append(sChild, m.marshalGameObject(childObject))
	}

	return SerializedGameObject{
		ID:   gameObject.Id(),
		Name: gameObject.Name(),
		Transform: SerializedTransform{
			Position: SerializedPosition{
				X: gameObject.RelativePosition().X,
				Y: gameObject.RelativePosition().Y,
			},
			Rotation: gameObject.Rotation().Radians(),
			Scale:    gameObject.Scale(),
		},
		Components: sComponents,
		Child:      sChild,
	}
}

func (m *Manager) marshalComponent(component galx.Component) SerializedComponent {
	props := m.componentRegistry.ExtractComponentProps(component)
	sProps := make([]SerializedComponentProperty, 0, len(props))

	for name, value := range props {
		sProps = append(sProps, SerializedComponentProperty{
			Name:  name,
			Value: value,
		})
	}

	return SerializedComponent{
		ID:    component.Id(),
		Props: sProps,
	}
}
