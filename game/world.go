package game

import (
	"fmt"

	"github.com/fe3dback/galaxy/render"

	"github.com/fe3dback/galaxy/engine"
)

type EntityList []*engine.Entity

type World struct {
	entities EntityList
}

func NewWorld() *World {
	return &World{
		entities: make(EntityList, 0),
	}
}

func (w *World) AddEntity(e *engine.Entity) {
	w.entities = append(w.entities, e)
}

func (w *World) Entities() EntityList {
	return w.entities
}

func (w *World) OnUpdate(moment engine.Moment) error {
	needGc := false

	for _, entity := range w.entities {
		if entity.IsDestroyed() {
			needGc = true
			continue
		}

		err := entity.OnUpdate(moment)
		if err != nil {
			return fmt.Errorf("can`t update world entity `%T`: %v", entity, err)
		}
	}

	if needGc {
		w.garbageCollect()
	}

	return nil
}

func (w *World) OnDraw(r *render.Renderer) error {
	for _, entity := range w.entities {
		if entity.IsDestroyed() {
			continue
		}

		err := entity.OnDraw(r)
		if err != nil {
			return fmt.Errorf("can`t draw world entity `%T`: %v", entity, err)
		}
	}

	return nil
}

func (w *World) garbageCollect() {
	list := make(EntityList, 0, len(w.entities))

	for _, entity := range w.entities {
		if entity.IsDestroyed() {
			continue
		}

		list = append(list, entity)
	}

	w.entities = list
}
