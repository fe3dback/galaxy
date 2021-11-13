package game

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/game/scenes"
)

const (
	sceneGame = "game"
)

type (
	ScenesLoader struct {
		sceneManager engine.SceneManager
	}
)

func NewScenesLoader(sceneManager engine.SceneManager) *ScenesLoader {
	return &ScenesLoader{
		sceneManager: sceneManager,
	}
}

func (l *ScenesLoader) LoadGameScenes() {
	for sceneName, blueprint := range l.blueprints() {
		l.sceneManager.AddBlueprint(sceneName, blueprint)
	}
}

func (l *ScenesLoader) EnterToFirstScene() {
	l.sceneManager.Switch(sceneGame)
}

func (l *ScenesLoader) blueprints() map[string]engine.SceneBlueprint {
	return map[string]engine.SceneBlueprint{
		sceneGame: scenes.SceneGame{},
	}
}
