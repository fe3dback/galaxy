package weapon

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/game/units"
)

type (
	YamlSpec struct {
		Fire     YamlFire     `yaml:"fire" validate:"nonnil"`
		Magazine YamlMagazine `yaml:"magazine" validate:"nonnil"`
		Trail    YamlTrail    `yaml:"trail" validate:"nonnil"`
		Muzzle   YamlMuzzle   `yaml:"muzzle" validate:"nonnil"`
		Bullet   YamlBullet   `yaml:"bullet" validate:"nonnil"`
	}

	YamlFire struct {
		Bullets   units.Count `yaml:"bullets" validate:"nonnil,min=1,max=32"`       // Bullets per shot
		Rate      units.Rate  `yaml:"rate" validate:"nonnil,min=0.001,max=1024"`    // shots per second
		SpreadMin float64     `yaml:"spreadMinDeg" validate:"nonnil,min=0,max=180"` // min spread deg
		SpreadMax float64     `yaml:"spreadMaxDeg" validate:"nonnil,min=0,max=180"` // max spread deg
	}

	YamlMagazine struct {
		StartAmmo  units.Count  `yaml:"startAmmoCount" validate:"nonnil,min=0"`            // default ammo on first equip
		MaxAmmo    units.Count  `yaml:"maxAmmoCount" validate:"nonnil,min=1"`              // max ammo
		Stack      units.Count  `yaml:"stackCount" validate:"nonnil,min=1"`                // count ammo in single magazine
		ReloadTime units.Second `yaml:"reloadTimeSec" validate:"nonnil,min=0.01,max=3600"` // reload magazine time in sec
	}

	YamlTrail struct {
		Has   bool         `yaml:"has" validate:"nonnil"`      // Has trail path (traced ammo)
		Color engine.Color `yaml:"colorHex" validate:"nonnil"` // trace Color
	}

	YamlMuzzle struct {
		Offset units.Pixel     `yaml:"offsetPx" validate:"nonnil,min=0,max=256"` // muzzle right Offset in px
		Flash  YamlMuzzleFlash `yaml:"flash" validate:"nonnil"`
	}

	YamlMuzzleFlash struct {
		Color  engine.Color `yaml:"colorHex" validate:"nonnil"`                     // muzzle flash Color
		Radius units.Meter  `yaml:"radiusMeters" validate:"nonnil,min=0.01,max=10"` // muzzle flash Radius
	}

	YamlBullet struct {
		Damage      units.Count   `yaml:"damage" validate:"nonnil"`                         // Damage in units
		MaxLifeTime units.Second  `yaml:"maxLifeTimeSec" validate:"nonnil,min=0.01,max=60"` // max bullet life in sec
		Air         YamlBulletAir `yaml:"air" validate:"nonnil"`
	}

	YamlBulletAir struct {
		StartAcceleration units.SpeedMpS `yaml:"startAccelerationMps" validate:"nonnil,min=-32,max=32"`
		StartVelocity     units.SpeedMpS `yaml:"startVelocityMps" validate:"nonnil,min=-32,max=32"`
		MaxVelocity       units.SpeedMpS `yaml:"maxVelocityMps" validate:"nonnil,min=-32,max=32"`
	}
)
