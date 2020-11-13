package schemefactory

import (
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/loader/weaponloader"
	"github.com/fe3dback/galaxy/generated"
)

type Unit struct {
	EntitySpawner entity.Spawner
	TextureRes    generated.ResourcePath
	WeaponsLoader *weaponloader.Loader
}

func (b Unit) SchemeID() entity.SchemeID {
	return SchemeGameUnit
}
