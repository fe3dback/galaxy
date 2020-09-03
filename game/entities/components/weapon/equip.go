package weapon

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/game/loader/weaponloader"
	"github.com/fe3dback/galaxy/generated"
)

type (
	id = generated.ResourcePath

	equip struct {
		weapons map[id]*Weapon

		currentWeapon id
		equipped      bool
	}
)

func newEquip(weaponsLoader *weaponloader.Loader, mixer engine.SoundMixer) *equip {
	equip := &equip{
		weapons: map[id]*Weapon{},
	}

	var lastId *id

	for id, spec := range weaponsLoader.LoadedSpecs() {
		lastId = &id
		equip.weapons[id] = NewWeapon(spec, mixer)
	}

	if lastId != nil {
		// equip weapon by default
		equip.Equip(*lastId)
	}

	return equip
}

func (e *equip) Equip(id id) {
	e.currentWeapon = id
	e.equipped = true
}

func (e *equip) UnEquip() {
	e.equipped = false
}

func (e *equip) CurrentWeapon() (*Weapon, bool) {
	if !e.equipped {
		return nil, false
	}

	return e.weapons[e.currentWeapon], true
}
