package game

import (
	"time"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
)

type LifeTime struct {
	entity    *entity.Entity
	destroyAt time.Time
	destroyed bool
}

func NewLifeTime(entity *entity.Entity, duration time.Duration) *LifeTime {
	return &LifeTime{
		entity:    entity,
		destroyAt: time.Now().Add(duration),
		destroyed: false,
	}
}

func (lt *LifeTime) OnDraw(_ engine.Renderer) error {
	return nil
}

func (lt *LifeTime) OnUpdate(s engine.State) error {
	if lt.destroyed {
		return nil
	}

	if time.Now().After(lt.destroyAt) {
		lt.entity.Destroy()
		lt.destroyed = true
	}

	return nil
}
