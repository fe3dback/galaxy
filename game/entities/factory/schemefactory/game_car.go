package schemefactory

import (
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/generated"
)

type Car struct {
	TextureRes generated.ResourcePath
	PhysicsRes generated.ResourcePath
}

func (b Car) SchemeID() entity.SchemeID {
	return SchemeGameCar
}
