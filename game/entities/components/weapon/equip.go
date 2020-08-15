package weapon

type equip struct {
	weapons map[Id]*Weapon
}

func newEquip(loader *Loader) *equip {
	equip := &equip{
		weapons: map[Id]*Weapon{},
	}

	for id, spec := range loader.Specs() {
		equip.weapons[id] = NewWeapon(spec)
	}

	return equip
}

func (e *equip) Weapon(id Id) *Weapon {
	return e.weapons[id]
}
