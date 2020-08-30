package weapon

import (
	"time"

	"github.com/fe3dback/galaxy/generated"

	"github.com/fe3dback/galaxy/engine"
)

type (
	YamlSpec struct {
		Fire     YamlFire     `yaml:"fire" validate:"nonnil"`
		Magazine YamlMagazine `yaml:"magazine" validate:"nonnil"`
		Trail    YamlTrail    `yaml:"trail" validate:"nonnil"`
		Muzzle   YamlMuzzle   `yaml:"muzzle" validate:"nonnil"`
		Bullet   YamlBullet   `yaml:"bullet" validate:"nonnil"`
		Audio    YamlAudio    `yaml:"audio" validate:"nonnil"`
	}

	YamlFire struct {
		Bullets   engine.Count `yaml:"bullets" validate:"nonnil,min=1,max=32"`       // Bullets per shot
		Rate      engine.Rate  `yaml:"rate" validate:"nonnil,min=0.001,max=1024"`    // shots per second
		SpreadMin float64      `yaml:"spreadMinDeg" validate:"nonnil,min=0,max=180"` // min spread deg
		SpreadMax float64      `yaml:"spreadMaxDeg" validate:"nonnil,min=0,max=180"` // max spread deg
	}

	YamlMagazine struct {
		StartAmmo  engine.Count  `yaml:"startAmmoCount" validate:"nonnil,min=0"` // default ammo on first equip
		MaxAmmo    engine.Count  `yaml:"maxAmmoCount" validate:"nonnil,min=1"`   // max ammo
		Stack      engine.Count  `yaml:"stackCount" validate:"nonnil,min=1"`     // count ammo in single magazine
		ReloadTime time.Duration `yaml:"reloadTime" validate:"nonnil"`           // reload magazine time in sec
	}

	YamlTrail struct {
		Has   bool         `yaml:"has" validate:"nonnil"`      // Has trail path (traced ammo)
		Color engine.Color `yaml:"colorHex" validate:"nonnil"` // trace Color
	}

	YamlMuzzle struct {
		Offset engine.Pixel    `yaml:"offsetPx" validate:"nonnil,min=0,max=256"` // muzzle right Offset in px
		Flash  YamlMuzzleFlash `yaml:"flash" validate:"nonnil"`
	}

	YamlMuzzleFlash struct {
		Color  engine.Color `yaml:"colorHex" validate:"nonnil"`                     // muzzle flash Color
		Radius engine.Meter `yaml:"radiusMeters" validate:"nonnil,min=0.01,max=10"` // muzzle flash Radius
	}

	YamlBullet struct {
		Damage   engine.Count  `yaml:"damage" validate:"nonnil"`   // Damage in units
		LifeTime time.Duration `yaml:"lifeTime" validate:"nonnil"` // max bullet life in sec
		Air      YamlBulletAir `yaml:"air" validate:"nonnil"`
	}

	YamlBulletAir struct {
		StartAcceleration engine.SpeedMpS `yaml:"startAccelerationMps" validate:"nonnil,min=-1024,max=1024"`
		StartVelocity     engine.SpeedMpS `yaml:"startVelocityMps" validate:"nonnil,min=-1024,max=1024"`
		MaxVelocity       engine.SpeedMpS `yaml:"maxVelocityMps" validate:"nonnil,min=-1024,max=1024"`
	}

	YamlAudio struct {
		ShotSounds   []generated.ResourcePath `yaml:"shot" validate:"nonnil"`
		ReloadSounds []generated.ResourcePath `yaml:"reload" validate:"nonnil"`
	}
)
