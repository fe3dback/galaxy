package engine

type GameWorldCreator struct {
	assetsLoader Loader
}

func NewGameWorldCreator(assetsLoader Loader) *GameWorldCreator {
	return &GameWorldCreator{
		assetsLoader: assetsLoader,
	}
}

func (g *GameWorldCreator) Loader() Loader {
	return g.assetsLoader
}
