package weaponloader

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/generated"
)

var knownIds = []id{
	generated.ResourcesWeaponsRemi16,
}

type (
	id     = generated.ResourcePath
	Specs  = map[id]YamlSpec
	Loader struct {
		specs Specs
	}
)

func NewLoader(loader engine.Loader) *Loader {
	specs := make(Specs)

	for _, knownId := range knownIds {
		specs[knownId] = load(loader, knownId)
	}

	return &Loader{
		specs: specs,
	}
}

func (l *Loader) LoadedSpecs() Specs {
	return l.specs
}

func load(loader engine.Loader, id id) YamlSpec {
	spec := YamlSpec{}

	// load spec
	loader.LoadYaml(id, &spec)

	spec = prepareSpec(spec)

	// load sounds to memory
	sounds := make([]generated.ResourcePath, 0)
	sounds = append(sounds, spec.Audio.ShotSounds...)
	sounds = append(sounds, spec.Audio.ReloadSounds...)

	for _, sound := range sounds {
		loader.LoadSound(sound)
	}

	return spec
}

func prepareSpec(spec YamlSpec) YamlSpec {
	// convert deg to rad
	spec.Fire.SpreadMin = engine.NewAngle(spec.Fire.SpreadMin).Radians()
	spec.Fire.SpreadMax = engine.NewAngle(spec.Fire.SpreadMax).Radians()

	return spec
}
