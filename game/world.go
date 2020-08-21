package game

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
)

type EntityList []*entity.Entity

type World struct {
	entities   EntityList
	spawnQueue EntityList
}

func NewWorld() *World {
	return &World{
		entities:   make(EntityList, 0),
		spawnQueue: make(EntityList, 0),
	}
}

func (w *World) AddEntity(e *entity.Entity) {
	w.entities = append(w.entities, e)
}

func (w *World) SpawnInGameEntity(e *entity.Entity) {
	w.spawnQueue = append(w.spawnQueue, e)
}

func (w *World) Entities() EntityList {
	return w.entities
}

func (w *World) OnUpdate(s engine.State) error {
	needGc := false

	if len(w.spawnQueue) > 0 {
		for _, e := range w.spawnQueue {
			w.AddEntity(e)
		}
		w.spawnQueue = w.spawnQueue[:0]
	}

	for _, e := range w.entities {
		if e.IsDestroyed() {
			needGc = true
			continue
		}

		err := e.OnUpdate(s)
		if err != nil {
			return fmt.Errorf("can`t update world entity `%T`: %v", e, err)
		}
	}

	for _, eA := range w.entities {
		if eA.IsDestroyed() {
			continue
		}

		for _, eB := range w.entities {
			if eB.IsDestroyed() {
				continue
			}

			if eA.IsCollideWith(eB) {
				// todo: collision masks
				eA.OnCollide(eB, 0)
			}
		}
	}

	if needGc {
		w.garbageCollect()
	}

	return nil
}

func (w *World) OnDraw(r engine.Renderer) error {
	for _, e := range w.entities {
		if e.IsDestroyed() {
			continue
		}

		err := e.OnDraw(r)
		if err != nil {
			return fmt.Errorf("can`t draw world entity `%T`: %v", e, err)
		}
	}

	return nil
}

func (w *World) garbageCollect() {
	list := make(EntityList, 0, len(w.entities))

	for _, e := range w.entities {
		if e.IsDestroyed() {
			continue
		}

		list = append(list, e)
	}

	w.entities = list
}
