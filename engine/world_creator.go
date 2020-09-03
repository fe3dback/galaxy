package engine

type GameWorldCreator struct {
	assetsLoader Loader
	soundMixer   SoundMixer
	physics      Physics
}

func NewGameWorldCreator(
	assetsLoader Loader,
	soundMixer SoundMixer,
	physics Physics,
) *GameWorldCreator {
	return &GameWorldCreator{
		assetsLoader: assetsLoader,
		soundMixer:   soundMixer,
		physics:      physics,
	}
}

func (g *GameWorldCreator) Loader() Loader {
	return g.assetsLoader
}

func (g *GameWorldCreator) SoundMixer() SoundMixer {
	return g.soundMixer
}

func (g *GameWorldCreator) Physics() Physics {
	return g.physics
}
