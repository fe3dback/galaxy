package game

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
)

type (
	EntityList []*entity.Entity

	World struct {
		entities     EntityList
		spawnQueue   EntityList
		physics      engine.Physics
		worldCreator engine.WorldCreator
	}
)

func NewWorld(creator engine.WorldCreator) *World {
	return &World{
		entities:     make(EntityList, 0),
		spawnQueue:   make(EntityList, 0),
		physics:      creator.Physics(),
		worldCreator: creator,
	}
}

func (w *World) AddEntity(e *entity.Entity) {
	w.entities = append(w.entities, e)
}

func (w *World) SpawnEntity(pos engine.Vec, angle engine.Angle, scheme entity.Scheme) {
	defer func() {
		if data := recover(); data != nil {
			zap.S().Errorf("panic: failed to spawn entity by scheme `%s`: %s", scheme.SchemeID(), data)
			return
		}
	}()

	e := entity.NewEntity(pos, angle)
	schemeFactoryBuilder, ok := schemeFactoryMap[scheme.SchemeID()]
	if !ok {
		zap.S().Errorf("Can`t spawn entity: not found scheme factory for scheme: %s", scheme.SchemeID())
		return
	}

	factoryMethod := schemeFactoryBuilder(scheme)
	if !ok {
		zap.S().Errorf("Can`t spawn entity: scheme factory '%s' not implement factory method", scheme.SchemeID())
		return
	}

	e = factoryMethod(e, w.worldCreator)

	w.spawnQueue = append(w.spawnQueue, e)
}

func (w *World) Entities() EntityList {
	return w.entities
}

func (w *World) OnUpdate(s engine.State) error {
	needGc := false

	// spawn new entities
	if len(w.spawnQueue) > 0 {
		for _, e := range w.spawnQueue {
			w.AddEntity(e)
		}
		w.spawnQueue = w.spawnQueue[:0]
	}

	// update physics
	w.physics.Update(s.Moment().DeltaTime())

	// update game
	for _, e := range w.entities {
		if e.IsDestroyed() {
			w.physics.DestroyBody(e.PhysicsBody())
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

func (w *World) OnDraw(r engine.Renderer) error {
	// draw physics gizmos
	w.physics.Draw(r)

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
