package engine

type GameState struct {
	moment     Moment
	camera     Camera
	mouse      Mouse
	movement   Movement
	appState   *AppState
	soundMixer SoundMixer
}

func NewGameState(
	moment Moment,
	camera Camera,
	mouse Mouse,
	movement Movement,
	appState *AppState,
	soundMixer SoundMixer,
) *GameState {
	return &GameState{
		moment:     moment,
		camera:     camera,
		mouse:      mouse,
		movement:   movement,
		appState:   appState,
		soundMixer: soundMixer,
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

func (g *GameState) SoundMixer() SoundMixer {
	return g.soundMixer
}

func (g *GameState) InEditorMode() bool {
	return g.appState.InEditorState()
}
