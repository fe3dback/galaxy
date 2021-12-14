package scene

import (
	"encoding/xml"
	"fmt"

	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine/node"
)

type (
	trackedIDs = map[string]struct{}
)

func (m *Manager) unmarshalSceneXML(sceneXML []byte) (*Scene, error) {
	sScene := &SerializedScene{}
	err := xml.Unmarshal(sceneXML, sScene)
	if err != nil {
		return nil, fmt.Errorf("can`t decode scene xml: %w", err)
	}

	return m.unmarshalScene(sScene)
}

func (m *Manager) unmarshalScene(sScene *SerializedScene) (*Scene, error) {
	trackedIDs := make(trackedIDs)
	objects := make([]galx.GameObject, 0, len(sScene.GameObjects))

	for _, object := range sScene.GameObjects {
		objects = append(objects, m.unmarshalGameObject(trackedIDs, object))
	}

	return NewScene(objects), nil
}

func (m *Manager) unmarshalGameObject(trackedIDs trackedIDs, res SerializedGameObject) galx.GameObject {
	// track ids
	uniqueID := res.ID
	if _, exist := trackedIDs[uniqueID]; exist {
		panic(fmt.Sprintf("failed create gameObject, entity with id '%s' already loaded. (duplicate ID in scene file)", uniqueID))
	}
	trackedIDs[uniqueID] = struct{}{}

	// create
	object := node.NewNode(uniqueID)
	object.SetPosition(galx.Vec{
		X: res.Transform.Position.X,
		Y: res.Transform.Position.Y,
	})
	object.SetRotation(galx.Angle(res.Transform.Rotation)) // in rad
	object.SetName(res.Name)

	// apply components
	for _, resComponent := range res.Components {
		cpm := m.unmarshalGameComponent(resComponent)

		if creatingComponent, ok := cpm.(galx.ComponentCycleCreated); ok {
			creatingComponent.OnCreated(object)
		}

		object.AddComponent(cpm)
	}

	// apply child
	for _, resChild := range res.Child {
		child := m.unmarshalGameObject(trackedIDs, resChild)
		child.SetParent(object)
		object.AddChild(child)
	}

	return object
}

func (m *Manager) unmarshalGameComponent(sComponent SerializedComponent) galx.Component {
	props := make(map[string]string, len(sComponent.Props))
	for _, resProp := range sComponent.Props {
		props[resProp.Name] = resProp.Value
	}

	return m.componentRegistry.CreateComponentWithProps(sComponent.ID, props)
}
