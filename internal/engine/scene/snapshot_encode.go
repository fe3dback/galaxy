package scene

import (
	"github.com/fe3dback/galaxy/galx"
)

func (m *Manager) encodeScene(scene Scene) *SnapshotScene {
	snapshotGameObjects := make([]SnapshotGameObject, 0, len(scene.entities))

	for _, gameObject := range scene.entities {
		snapshotGameObjects = append(snapshotGameObjects, m.encodeGameObject(gameObject))
	}

	return &SnapshotScene{
		ID:          scene.id,
		GameObjects: snapshotGameObjects,
	}
}

func (m *Manager) encodeGameObject(gameObject galx.GameObject) SnapshotGameObject {
	components := gameObject.Components()
	child := gameObject.Child()

	snapshotComponents := make([]SnapshotComponent, 0, len(components))
	for _, component := range components {
		snapshotComponents = append(snapshotComponents, m.encodeComponent(component))
	}

	snapshotGameObjectChild := make([]SnapshotGameObject, 0, len(child))
	for _, gameObjectChild := range child {
		snapshotGameObjectChild = append(snapshotGameObjectChild, m.encodeGameObject(gameObjectChild))
	}

	return SnapshotGameObject{
		ID:   gameObject.Id(),
		Name: gameObject.Name(),
		Transform: SnapshotTransform{
			Position: SnapshotPosition{
				X: gameObject.RelativePosition().X,
				Y: gameObject.RelativePosition().Y,
			},
			Rotation: gameObject.Rotation().Radians(),
			Scale:    gameObject.Scale(),
		},
		Components: snapshotComponents,
		Child:      snapshotGameObjectChild,
	}
}

func (m *Manager) encodeComponent(component galx.Component) SnapshotComponent {
	componentProps := m.componentRegistry.ExtractComponentProps(component)
	snapshotComponentProperties := make([]SnapshotComponentProperty, 0, len(componentProps))

	for name, value := range componentProps {
		snapshotComponentProperties = append(snapshotComponentProperties, SnapshotComponentProperty{
			Name:  name,
			Value: value,
		})
	}

	return SnapshotComponent{
		ID:    component.Id(),
		Props: snapshotComponentProperties,
	}
}
