package scene

import (
	"encoding/xml"
	"fmt"
)

func (m *Manager) marshalScene(snapshotScene SnapshotScene) ([]byte, error) {
	m.marshalCleanSnapshot(&snapshotScene)
	bytes, err := xml.MarshalIndent(snapshotScene, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("failed marshal snapshot to xml: %w", err)
	}

	return bytes, nil
}

func (m *Manager) marshalCleanSnapshot(snapshotScene *SnapshotScene) {
	snapshotScene.ID = "" // clean ID in xml, because it loaded from asset file name
}
