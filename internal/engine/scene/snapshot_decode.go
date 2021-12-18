package scene

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine/node"
)

type (
	trackedIDs = map[string]struct{}
)

func (m *Manager) decodeScene(snapshotScene SnapshotScene) *Scene {
	trackedIDs := make(trackedIDs)
	gameObjects := make([]galx.GameObject, 0, len(snapshotScene.GameObjects))

	for _, snapshotGameObject := range snapshotScene.GameObjects {
		gameObjects = append(gameObjects, m.decodeGameObject(trackedIDs, snapshotGameObject))
	}

	return NewScene(snapshotScene.ID, gameObjects)
}

func (m *Manager) decodeGameObject(trackedIDs trackedIDs, snapshotGameObject SnapshotGameObject) galx.GameObject {
	// track ids
	uniqueID := snapshotGameObject.ID
	if _, exist := trackedIDs[uniqueID]; exist {
		panic(fmt.Sprintf("failed create gameObject, entity with id '%s' already loaded. (duplicate ID in scene file)", uniqueID))
	}
	trackedIDs[uniqueID] = struct{}{}

	// create
	gameObject := node.NewNode(uniqueID)
	gameObject.SetPosition(galx.Vec{
		X: snapshotGameObject.Transform.Position.X,
		Y: snapshotGameObject.Transform.Position.Y,
	})
	gameObject.SetRotation(galx.Angle(snapshotGameObject.Transform.Rotation)) // in rad
	gameObject.SetName(snapshotGameObject.Name)

	// apply components
	for _, snapshotComponent := range snapshotGameObject.Components {
		gameComponent := m.decodeGameComponent(snapshotComponent)

		if creatingComponent, ok := gameComponent.(galx.ComponentCycleCreated); ok {
			creatingComponent.OnCreated(gameObject)
		}

		gameObject.AddComponent(gameComponent)
	}

	// apply child
	for _, snapshotGameObjectChild := range snapshotGameObject.Child {
		gameObjectChild := m.decodeGameObject(trackedIDs, snapshotGameObjectChild)
		gameObjectChild.SetParent(gameObject)
		gameObject.AddChild(gameObjectChild)
	}

	return gameObject
}

func (m *Manager) decodeGameComponent(snapshotComponent SnapshotComponent) galx.Component {
	props := make(map[string]string, len(snapshotComponent.Props))
	for _, resProp := range snapshotComponent.Props {
		props[resProp.Name] = resProp.Value
	}

	return m.componentRegistry.CreateComponentWithProps(snapshotComponent.ID, props)
}
