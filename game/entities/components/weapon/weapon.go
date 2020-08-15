package weapon

type Weapon struct {
	spec YamlSpec
}

func NewWeapon(spec YamlSpec) *Weapon {
	return &Weapon{
		spec: spec,
	}
}
