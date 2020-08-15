package weapon

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/generated"
	"gopkg.in/validator.v2"
)

const (
	Remi16 = Id(generated.ResourcesWeaponsRemi16)
)

type (
	Specs  = map[Id]YamlSpec
	Id     = string
	Loader struct {
		creator engine.WorldCreator
		specs   Specs
	}
)

func NewLoader(creator engine.WorldCreator) *Loader {
	loader := &Loader{
		creator: creator,
		specs:   make(Specs),
	}
	loader.loadSpecs()

	return loader
}

func (l *Loader) Specs() Specs {
	return l.specs
}

func (l *Loader) loadSpecs() {
	l.specs[Remi16] = l.load(Remi16)
}

func (l *Loader) load(id Id) YamlSpec {
	spec := YamlSpec{}

	l.creator.Loader().LoadYaml(generated.ResourcePath(id), &spec)

	if err := validator.Validate(spec); err != nil {
		panic(fmt.Sprintf("invalid weapon spec '%s': %v", id, err))
	}

	return spec
}

func (l *Loader) prepareSpec(spec YamlSpec) YamlSpec {
	// convert deg to rad
	spec.Fire.SpreadMin = engine.NewAngle(spec.Fire.SpreadMin).Radians()
	spec.Fire.SpreadMax = engine.NewAngle(spec.Fire.SpreadMax).Radians()

	return spec
}
