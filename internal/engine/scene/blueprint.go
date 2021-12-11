package scene

import (
	"github.com/fe3dback/galaxy/galx"
)

type emptySceneBlueprint struct {
}

func (l emptySceneBlueprint) CreateEntities() []galx.GameObject {
	return []galx.GameObject{}
}
