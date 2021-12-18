package scene

import (
	"fmt"
	"path/filepath"

	"github.com/fe3dback/galaxy/consts"
)

type (
	XMLScenes struct {
		Default ID       `xml:"default"`
		Refs    []string `xml:"refs>scene"`
	}
)

func (m *Manager) assetSceneFilePath(sceneID string) string {
	return filepath.Join(consts.AssetScenesRoot, sceneID, consts.AssetScenesObjectsFileName)
}

func (m *Manager) loadSnapshots() (defaultScene ID, result snapshots) {
	scenesMeta := XMLScenes{}
	m.assetsManager.LoadXML(consts.AssetScenesDefinitionXML, &scenesMeta)
	if err := scenesMeta.validate(); err != nil {
		panic(fmt.Sprintf("scenes '%s' not valid: %v", consts.AssetScenesDefinitionXML, err))
	}

	result = make(snapshots)
	for _, sceneID := range scenesMeta.Refs {
		result[sceneID] = m.loadSnapshot(sceneID)
	}

	return scenesMeta.Default, result
}

func (m *Manager) loadSnapshot(sceneID ID) (snapshot *SnapshotScene) {
	scenePath := m.assetSceneFilePath(sceneID)
	sceneXML := m.assetsManager.LoadFile(scenePath)
	snapshotScene, err := m.unmarshalScene(sceneXML)
	if err != nil {
		panic(fmt.Sprintf("failed load scene '%s' from '%s': %v", sceneID, scenePath, err))
	}

	snapshotScene.ID = sceneID
	return snapshotScene
}

func (m *Manager) saveSnapshot(scene *SnapshotScene) {
	scenePath := m.assetSceneFilePath(scene.ID)
	bytes, err := m.marshalScene(*scene)
	if err != nil {
		panic(fmt.Sprintf("failed marshal scene '%s': %v", scene.ID, err))
	}

	m.assetsManager.SaveFile(scenePath, bytes)
}

func (rs *XMLScenes) validate() error {
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
