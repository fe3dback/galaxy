package weapon

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
)

type CharacterInventory struct {
	entity *entity.Entity
	equip  *equip
}

func NewCharacterInventory(entity *entity.Entity, loader *Loader) *CharacterInventory {
	return &CharacterInventory{
		entity: entity,
		equip:  newEquip(loader),
	}
}

func (r *CharacterInventory) OnDraw(_ engine.Renderer) error {
	return nil
}

func (r *CharacterInventory) OnUpdate(_ engine.State) error {
	return nil
}
