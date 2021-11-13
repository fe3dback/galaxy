package scene

import (
	"github.com/fe3dback/galaxy/engine"
)

type emptySceneBlueprint struct {
}

func (l emptySceneBlueprint) CreateEntities() []engine.GameObject {
	return []engine.GameObject{}
}
