package scene

import (
	"fmt"
	"path/filepath"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
)

func (m *Manager) LoadScenes() {
	scenes := SerializedScenes{}
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
	sceneXML := m.assetsLoader.LoadFile(objectsPath)
	return m.createBlueprint(sceneXML), nil
}

func (m *Manager) createBlueprint(sceneXML []byte) blueprint {
	return func() []galx.GameObject {
		scene, err := m.unmarshalSceneXML(sceneXML)
		if err != nil {
			panic(fmt.Errorf("failed unmarshal scene: %w", err))
		}

		return scene.entities
	}
}
