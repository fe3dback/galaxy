package game

import (
	"github.com/fe3dback/galaxy/galx"
)

type LifeTime struct {
	TimeToLiveSeconds int `editable:"time-to-live-seconds"`

	entity      galx.GameObject
	ticksLeft   int
	initialized bool
}

func (lt LifeTime) Id() string {
	return "89270693-9d58-4d8f-ad4f-1b70779f5939"
}

func (lt LifeTime) Title() string {
	return "Game.Life time"
}

func (lt LifeTime) Description() string {
	return "Will automatic destroy entity after life time elapsed"
}

func (lt *LifeTime) OnCreated(entity galx.GameObject) {
	lt.entity = entity

}

func (lt *LifeTime) OnDraw(_ galx.Renderer) error {
	return nil
}

func (lt *LifeTime) OnUpdate(state galx.State) error {
	if lt.ticksLeft == -1 {
		return nil
	}

	if !lt.initialized {
		lt.initialized = true
		lt.ticksLeft = lt.TimeToLiveSeconds * state.Moment().TargetFPS()
	}

	lt.ticksLeft--
	if lt.ticksLeft <= 0 {
		lt.entity.Destroy()
		lt.ticksLeft = -1
	}

	return nil
}
