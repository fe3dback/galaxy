package schemefactory

import (
	"github.com/fe3dback/galaxy/engine/entity"
)

type Wall struct {
	BoxWidth int32
}

func (b Wall) SchemeID() entity.SchemeID {
	return SchemeGameWall
}
