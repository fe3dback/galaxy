package game

import (
	"time"

	"github.com/fe3dback/galaxy/galx"
)

type LifeTime struct {
	entity    galx.GameObject
	destroyAt time.Time
	destroyed bool
}

func NewLifeTime(entity galx.GameObject, duration time.Duration) *LifeTime {
	return &LifeTime{
		entity:    entity,
		destroyAt: time.Now().Add(duration),
		destroyed: false,
	}
}

func (lt *LifeTime) OnDraw(_ galx.Renderer) error {
	return nil
}

func (lt *LifeTime) OnUpdate(_ galx.State) error {
	if lt.destroyed {
		return nil
	}

	if time.Now().After(lt.destroyAt) {
		lt.entity.Destroy()
		lt.destroyed = true
	}

	return nil
}
