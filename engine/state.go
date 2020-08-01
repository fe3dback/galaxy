package engine

type GameState struct {
	moment Moment
	camera Camera
	mouse  Mouse
}

func NewGameState(
	moment Moment,
	camera Camera,
	mouse Mouse,
) *GameState {
	return &GameState{
		moment: moment,
		camera: camera,
		mouse:  mouse,
	}
}

func (g *GameState) Moment() Moment {
	return g.moment
}

func (g *GameState) Camera() Camera {
	return g.camera
}

func (g *GameState) Mouse() Mouse {
	return g.mouse
}
