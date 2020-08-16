package weapon

import "github.com/fe3dback/galaxy/generated"

type equip struct {
	weapons map[Id]*Weapon

	currentWeapon Id
	equipped      bool
}

func newEquip(loader *Loader) *equip {
	equip := &equip{
		weapons: map[Id]*Weapon{},
	}

	var lastId *Id

	for id, spec := range loader.Specs() {
		lastId = &id
		equip.weapons[id] = NewWeapon(spec, loader.creator.SoundMixer())

		// load all weapons sound to memory
		sounds := make([]generated.ResourcePath, 0)
		sounds = append(sounds, spec.Audio.ShotSounds...)
		sounds = append(sounds, spec.Audio.ReloadSounds...)

		for _, sound := range sounds {
			loader.creator.Loader().LoadSound(sound)
		}
	}

	if lastId != nil {
		// equip weapon by default
		equip.Equip(*lastId)
	}

	return equip
}

func (e *equip) Equip(id Id) {
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
