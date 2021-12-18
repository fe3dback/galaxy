package scene

import (
	"encoding/xml"
	"fmt"
)

func (m *Manager) unmarshalScene(sceneXML []byte) (*SnapshotScene, error) {
	snapshotScene := &SnapshotScene{}
	err := xml.Unmarshal(sceneXML, snapshotScene)
	if err != nil {
		return nil, fmt.Errorf("can`t unmarshal scene xml to snapshot: %w", err)
	}

	return snapshotScene, nil
}
