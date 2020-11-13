package weapon

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities/factory/schemefactory"
	"github.com/fe3dback/galaxy/game/loader/weaponloader"
)

type CharacterInventory struct {
	entitySpawner entity.Spawner
	entity        *entity.Entity
	equip         *equip
}

func NewCharacterInventory(
	entity *entity.Entity,
	weaponsLoader *weaponloader.Loader,
	mixer engine.SoundMixer,
	entitySpawner entity.Spawner,
) *CharacterInventory {
	return &CharacterInventory{
		entitySpawner: entitySpawner,
		entity:        entity,
		equip:         newEquip(weaponsLoader, mixer),
	}
}

func (r *CharacterInventory) OnDraw(_ engine.Renderer) error {
	return nil
}

func (r *CharacterInventory) OnUpdate(s engine.State) error {
	if !s.Movement().Space() {
		return nil
	}

	// test shot
	currentWeapon, ok := r.equip.CurrentWeapon()
	if !ok {
		return nil
	}

	currentWeapon.Shot(func(params bulletSpawnParams) {
		r.SpawnBullet(params)
	})

	return nil
}

func (r *CharacterInventory) SpawnBullet(params bulletSpawnParams) {
	spawnPos := r.entity.Position().PolarOffset(
		float64(params.muzzle.Offset),
		r.entity.Rotation(),
	)

	// create bullet
	direction := r.entity.Rotation().Add(params.spread)

	// add velocity component
	startAccelerationVec := engine.VectorRight(params.bullet.Air.StartAcceleration).Rotate(direction)
	startVelocityVec := engine.
		VectorRight(params.bullet.Air.StartVelocity).
		Rotate(direction)
	maxVelocityVec := engine.Vec{
		X: params.bullet.Air.MaxVelocity,
		Y: params.bullet.Air.MaxVelocity,
	}

	r.entitySpawner.SpawnEntity(spawnPos, direction, schemefactory.Bullet{
		StartAccelerationVec: startAccelerationVec,
		StartVelocityVec:     startVelocityVec,
		MaxVelocityVec:       maxVelocityVec,
		LifeTime:             params.bullet.LifeTime,
		HasTrail:             params.trail.Has,
		TrailColor:           params.trail.Color,
	})
}
