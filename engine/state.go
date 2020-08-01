package engine

type GameState struct {
	moment   Moment
	camera   Camera
	mouse    Mouse
	movement Movement
}

func NewGameState(
	moment Moment,
	camera Camera,
	mouse Mouse,
	movement Movement,
) *GameState {
	return &GameState{
		moment:   moment,
		camera:   camera,
		mouse:    mouse,
		movement: movement,
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

func (g *GameState) Movement() Movement {
	return g.movement
}
