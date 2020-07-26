package game

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
)

type entityList []*engine.Entity

type World struct {
	entities entityList
}

func NewWorld() *World {
	return &World{
		entities: make(entityList, 0),
	}
}

func (w *World) AddEntity(e *engine.Entity) {
	w.entities = append(w.entities, e)
}

func (w *World) OnUpdate(deltaTime float64) error {
	needGc := false

	for _, entity := range w.entities {
		if entity.IsDestroyed() {
			needGc = true
			continue
		}

		err := entity.OnUpdate(deltaTime)
		if err != nil {
			return fmt.Errorf("can`t update world entity `%T`: %v", entity, err)
		}
	}

	if needGc {
		w.garbageCollect()
	}

	return nil
}

func (w *World) OnDraw() error {
	for _, entity := range w.entities {
		if entity.IsDestroyed() {
			continue
		}

		err := entity.OnDraw()
		if err != nil {
			return fmt.Errorf("can`t draw world entity `%T`: %v", entity, err)
		}
	}

	return nil
}

func (w *World) garbageCollect() {
	list := make(entityList, 0, len(w.entities))

	for _, entity := range w.entities {
		if entity.IsDestroyed() {
			continue
		}

		list = append(list, entity)
	}

	w.entities = list
}
