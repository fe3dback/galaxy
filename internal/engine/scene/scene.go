package scene

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
)

type Scene struct {
	id         string
	destroyed  bool
	entities   []galx.GameObject
	spawnQueue []galx.GameObject
}

func NewScene(id string, entities []galx.GameObject) *Scene {
	return &Scene{
		id:         id,
		destroyed:  false,
		entities:   entities,
		spawnQueue: []galx.GameObject{},
	}
}

func (s *Scene) ID() string {
	return s.id
}

func (s *Scene) Entities() []galx.GameObject {
	return s.entities
}

func (s *Scene) OnUpdate(state galx.State) error {
	needGc := false

	if s.destroyed {
		return nil
	}

	// spawn new entities
	if len(s.spawnQueue) > 0 {
		s.entities = append(s.entities, s.spawnQueue...)
		s.spawnQueue = s.spawnQueue[:0]
	}

	// update game
	for _, e := range s.entities {
		if !e.IsRoot() {
			// scene will update only root entities
			// all child entities will be updated from parent
			continue
		}

		if e.IsDestroyed() {
			needGc = true
			continue
		}

		err := e.OnUpdate(state)
		if err != nil {
			return fmt.Errorf("can`t update world entity `%T`: %w", e, err)
		}
	}

	if needGc {
		s.garbageCollect()
	}

	return nil
}

func (s *Scene) OnDraw(r galx.Renderer) error {
	if s.destroyed {
		return nil
	}

	// draw world
	for _, e := range s.entities {
		if !e.IsRoot() {
			// scene will draw only root entities
			// all child entities will be drawn from parent
			continue
		}

		if e.IsDestroyed() {
			continue
		}

		err := e.OnDraw(r)
		if err != nil {
			return fmt.Errorf("can`t draw world entity `%T`: %w", e, err)
		}
	}

	return nil
}

func (s *Scene) garbageCollect() {
	list := make([]galx.GameObject, 0, len(s.entities))

	for _, e := range s.entities {
		if e.IsDestroyed() {
			continue
		}

		list = append(list, e)
	}

	s.entities = list
}

func (s *Scene) destroy() {
	for _, e := range s.entities {
		e.Destroy()
	}

	s.garbageCollect()
}
