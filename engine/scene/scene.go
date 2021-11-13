package scene

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
)

type Scene struct {
	destroyed  bool
	entities   []engine.GameObject
	spawnQueue []engine.GameObject
}

func NewScene(entities []engine.GameObject) *Scene {
	return &Scene{
		destroyed:  false,
		entities:   entities,
		spawnQueue: []engine.GameObject{},
	}
}

func (w *Scene) Entities() []engine.GameObject {
	return w.entities
}

func (w *Scene) OnUpdate(s engine.State) error {
	needGc := false

	if w.destroyed {
		return nil
	}

	// spawn new entities
	if len(w.spawnQueue) > 0 {
		for _, e := range w.spawnQueue {
			w.entities = append(w.entities, e)
		}
		w.spawnQueue = w.spawnQueue[:0]
	}

	// update game
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

	if needGc {
		w.garbageCollect()
	}

	return nil
}

func (w *Scene) OnDraw(r engine.Renderer) error {
	if w.destroyed {
		return nil
	}

	// draw world
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

func (w *Scene) garbageCollect() {
	list := make([]engine.GameObject, 0, len(w.entities))

	for _, e := range w.entities {
		if e.IsDestroyed() {
			continue
		}

		list = append(list, e)
	}

	w.entities = list
}

func (w *Scene) destroy() {
	for _, e := range w.entities {
		e.Destroy()
	}

	w.garbageCollect()
}
