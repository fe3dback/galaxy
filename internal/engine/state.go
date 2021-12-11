package engine

import "github.com/fe3dback/galaxy/galx"

type GameState struct {
	moment       galx.Moment
	camera       galx.Camera
	mouse        galx.Mouse
	movement     galx.Movement
	engineState  galx.EngineState
	soundMixer   galx.SoundMixer
	sceneManager galx.SceneManager
}

func NewGameState(
	moment galx.Moment,
	camera galx.Camera,
	mouse galx.Mouse,
	movement galx.Movement,
	appState galx.EngineState,
	soundMixer galx.SoundMixer,
	sceneManager galx.SceneManager,
) *GameState {
	return &GameState{
		moment:       moment,
		camera:       camera,
		mouse:        mouse,
		movement:     movement,
		engineState:  appState,
		soundMixer:   soundMixer,
		sceneManager: sceneManager,
	}
}

func (g *GameState) Moment() galx.Moment {
	return g.moment
}

func (g *GameState) Camera() galx.Camera {
	return g.camera
}

func (g *GameState) Mouse() galx.Mouse {
	return g.mouse
}

func (g *GameState) Movement() galx.Movement {
	return g.movement
}

func (g *GameState) SoundMixer() galx.SoundMixer {
	return g.soundMixer
}

func (g *GameState) Scene() galx.Scene {
	return g.sceneManager.Current()
}

func (g *GameState) EngineState() galx.EngineState {
	return g.engineState
}
