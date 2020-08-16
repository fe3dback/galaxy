package weapon

import (
	"fmt"

	"github.com/fe3dback/galaxy/game/entities/components/sprite/trail"

	"github.com/fe3dback/galaxy/game/entities/components/game"

	"github.com/fe3dback/galaxy/game/entities/components/movement"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
)

type CharacterInventory struct {
	entity *entity.Entity
	equip  *equip
}

func NewCharacterInventory(entity *entity.Entity, loader *Loader) *CharacterInventory {
	return &CharacterInventory{
		entity: entity,
		equip:  newEquip(loader),
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
	weapon, ok := r.equip.CurrentWeapon()
	if !ok {
		fmt.Println("No weapon")
		return nil
	}

	ok = weapon.Shot(func(params bulletSpawnParams) {
		r.SpawnBullet(s, params)
	})

	if ok {
		fmt.Printf("Ammo: %+v\n", weapon.ammo)
		fmt.Printf("Fire: %+v\n", weapon.fire)
		fmt.Printf("Reload: %+v\n", weapon.reload)
	}

	return nil
}

func (r *CharacterInventory) SpawnBullet(s engine.State, params bulletSpawnParams) {
	// todo replace to weapon bullets factory (DI + different bullets from different weapons)
	fmt.Println("SHOT")
	fmt.Printf("Bullet: %+v\n", params)

	spawnPos := r.entity.Position().PolarOffset(
		float64(params.muzzle.Offset),
		r.entity.Rotation(),
	)

	// create bullet
	direction := r.entity.Rotation().Add(params.spread)
	bullet := entity.NewEntity(
		spawnPos,
		direction,
	)

	// add velocity component
	startAccelerationVec := engine.VectorRight(params.bullet.Air.StartAcceleration).Rotate(direction)
	startVelocityVec := engine.
		VectorRight(params.bullet.Air.StartVelocity).
		Rotate(direction)
	maxVelocityVec := engine.Vec{
		X: params.bullet.Air.MaxVelocity,
		Y: params.bullet.Air.MaxVelocity,
	}

	bullet.AddComponent(movement.NewVelocity(
		bullet,
		startAccelerationVec,
		startVelocityVec,
		maxVelocityVec,
	))

	// add LifeTime component
	bullet.AddComponent(game.NewLifeTime(bullet, params.bullet.LifeTime))

	// add trail
	if params.trail.Has {
		bullet.AddComponent(trail.NewTrail(bullet, params.trail.Color))
	}

	// spawn
	s.EntitySpawner().SpawnEntity(bullet)
}
