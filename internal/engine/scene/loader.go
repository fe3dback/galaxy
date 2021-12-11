package scene

import (
	"fmt"
	"path/filepath"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine/entity"
)

type (
	ResScene struct {
		Default ID       `xml:"default"`
		Refs    []string `xml:"refs>scene"`
	}

	ResObjects struct {
		Objects []ResObject `xml:"objects>object"`
	}

	ResObject struct {
		ID         string         `xml:"id,attr"`
		Transform  ResTransform   `xml:"transform"`
		Components []ResComponent `xml:"components>component"`
	}

	ResTransform struct {
		Position ResPosition `xml:"position"`
		Rotation float64     `xml:"rotation"`
		Scale    float64     `xml:"scale"`
	}

	ResPosition struct {
		X float64 `xml:"x"`
		Y float64 `xml:"y"`
	}

	ResComponent struct {
		ID    string             `xml:"id,attr"`
		Props []ResComponentProp `xml:"props>prop"`
	}

	ResComponentProp struct {
		Name  string `xml:"name,attr"`
		Value string `xml:"value,attr"`
	}
)

func (rs *ResScene) validate() error {
	if rs.Default == "" {
		return fmt.Errorf("should contain 'scenes.default'")
	}

	if len(rs.Refs) == 0 {
		return fmt.Errorf("should contain at least one item in 'scenes.refs'")
	}

	hasDefault := false
	for _, sceneID := range rs.Refs {
		if sceneID == rs.Default {
			hasDefault = true
		}
	}

	if !hasDefault {
		return fmt.Errorf("default scene '%s' not defined in 'scenes.refs'", rs.Default)
	}

	return nil
}

func (m *Manager) LoadScenes() {
	scenes := ResScene{}
	m.assetsLoader.LoadXML(consts.AssetScenesDefinitionXML, &scenes)
	if err := scenes.validate(); err != nil {
		panic(fmt.Sprintf("scenes '%s' not valid: %v", consts.AssetScenesDefinitionXML, err))
	}

	for _, sceneID := range scenes.Refs {
		bp, err := m.loadScene(sceneID)
		if err != nil {
			panic(fmt.Sprintf("scenes '%s', failed load '%s': %v", consts.AssetScenesDefinitionXML, sceneID, err))
		}

		m.blueprints[sceneID] = bp
	}

	m.Switch(scenes.Default)
}

func (m *Manager) loadScene(sceneID ID) (bp blueprint, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("%v", v)
			return
		}
	}()

	objectsPath := filepath.Join(consts.AssetScenesRoot, sceneID, consts.AssetScenesObjectsFileName)
	sceneRoot := ResObjects{}

	m.assetsLoader.LoadXML(objectsPath, &sceneRoot)
	return m.createBlueprint(sceneRoot), nil
}

func (m *Manager) createBlueprint(scene ResObjects) blueprint {
	return func() []galx.GameObject {
		objects := make([]galx.GameObject, 0, len(scene.Objects))

		for _, object := range scene.Objects {
			objects = append(objects, m.createGameObject(object))
		}

		return objects
	}
}

func (m *Manager) createGameObject(res ResObject) galx.GameObject {
	object := entity.NewEntity(
		res.ID,
		galx.Vec{
			X: res.Transform.Position.X,
			Y: res.Transform.Position.Y,
		},
		galx.Angle(res.Transform.Rotation),
	)

	for _, resComponent := range res.Components {
		cpm := m.createGameComponent(resComponent)

		if creatingComponent, ok := cpm.(entity.ComponentCycleCreated); ok {
			creatingComponent.OnCreated(object)
		}

		object.AddComponent(cpm)
	}

	return object
}

func (m *Manager) createGameComponent(res ResComponent) entity.Component {
	props := make(map[string]string, len(res.Props))
	for _, resProp := range res.Props {
		props[resProp.Name] = resProp.Value
	}

	return m.componentRegistry.CreateComponentWithProps(res.ID, props)
}
