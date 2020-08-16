package engine

type GameWorldCreator struct {
	assetsLoader Loader
	soundMixer   SoundMixer
}

func NewGameWorldCreator(assetsLoader Loader, soundMixer SoundMixer) *GameWorldCreator {
	return &GameWorldCreator{
		assetsLoader: assetsLoader,
		soundMixer:   soundMixer,
	}
}

func (g *GameWorldCreator) Loader() Loader {
	return g.assetsLoader
}

func (g *GameWorldCreator) SoundMixer() SoundMixer {
	return g.soundMixer
}
