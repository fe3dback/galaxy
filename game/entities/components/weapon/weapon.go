package weapon

import (
	"math/rand"
	"time"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/game/loader/weaponloader"
)

type (
	Weapon struct {
		spec       weaponloader.YamlSpec
		soundMixer engine.SoundMixer

		// state
		ammo   ammo
		reload reload
		fire   fire
	}

	ammo struct {
		magazine int
		total    int
	}
	reload struct {
		reloading bool
	}
	fire struct {
		coolingTo time.Time
	}

	bulletSpawnParams struct {
		spread engine.Angle
		bullet weaponloader.YamlBullet
		muzzle weaponloader.YamlMuzzle
		trail  weaponloader.YamlTrail
	}
	bulletSpawnFn func(bulletSpawnParams)
)

func NewWeapon(spec weaponloader.YamlSpec, soundMixer engine.SoundMixer) *Weapon {
	weapon := &Weapon{
		spec:       spec,
		soundMixer: soundMixer,
		ammo: ammo{
			magazine: 0,
			total:    engine.ClampInt(int(spec.Magazine.StartAmmo), 0, int(spec.Magazine.MaxAmmo)),
		},
		reload: reload{
			reloading: false,
		},
		fire: fire{
			coolingTo: time.Now(),
		},
	}

	weapon.reloadMagazine()

	return weapon
}

func (w *Weapon) Shot(spawnBullet bulletSpawnFn) bool {
	if w.ammo.magazine == 0 {
		return false
	}

	if time.Now().Before(w.fire.coolingTo) {
		return false
	}

	loaded := engine.ClampInt(int(w.spec.Fire.Bullets), 0, w.ammo.magazine)
	for loaded > 0 {
		loaded--

		// define spread
		spreadSign := 0.0
		if rand.Float64() > 0.5 {
			spreadSign = 1
		} else {
			spreadSign = -1
		}

		spread := engine.NewAngle(
			spreadSign * engine.RandomRange(w.spec.Fire.SpreadMin, w.spec.Fire.SpreadMax),
		)

		// spawn bullet
		spawnBullet(bulletSpawnParams{
			spread: spread,
			bullet: w.spec.Bullet,
			muzzle: w.spec.Muzzle,
			trail:  w.spec.Trail,
		})

		// play sound
		sound := randomResource(w.spec.Audio.ShotSounds)
		if sound != nil {
			w.soundMixer.Play(*sound)
		}

		// pop ammo
		w.ammo.magazine--
	}

	// set new cooling time based on rate
	w.fire.coolingTo = time.Now().Add(
		time.Second / time.Duration(w.spec.Fire.Rate),
	)

	// auto reload
	if w.ammo.magazine == 0 {
		w.Reload()
	}

	return true
}

func (w *Weapon) Reload() bool {
	if w.reload.reloading {
		// already reloading
		return false
	}

	if w.ammo.total <= 0 {
		// no ammo left
		return false
	}

	needAmmo := w.needAmmoInMagazine()
	if needAmmo <= 0 {
		// magazine is full
		return false
	}

	// update state to off
	w.reload.reloading = true

	// play sound
	sound := randomResource(w.spec.Audio.ReloadSounds)
	if sound != nil {
		w.soundMixer.Play(*sound)
	}

	// enqueue state on
	time.AfterFunc(w.spec.Magazine.ReloadTime, func() {
		w.reloadMagazine()
		w.reload.reloading = false
	})

	return true
}

func (w *Weapon) needAmmoInMagazine() int {
	return int(w.spec.Magazine.Stack) - w.ammo.magazine
}

func (w *Weapon) reloadMagazine() {
	needAmmo := w.needAmmoInMagazine()
	if needAmmo == 0 {
		return
	}

	load := engine.ClampInt(needAmmo, 0, w.ammo.total)
	w.ammo.magazine += load
	w.ammo.total -= load
}
