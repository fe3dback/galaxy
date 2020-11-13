package game

import (
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities/factory"
	"github.com/fe3dback/galaxy/game/entities/factory/schemefactory"
)

type factoryMethod func(entity.Scheme) entity.FactoryFn

var schemeFactoryMap = map[entity.SchemeID]factoryMethod{
	// todo: code generation
	// todo: generate schemeFactory const
	schemefactory.SchemeGameBullet: func(scheme entity.Scheme) entity.FactoryFn {
		if typed, ok := scheme.(schemefactory.Bullet); ok {
			return factory.BulletFactoryFn(typed)
		}

		panic("invalid scheme")
	},
	schemefactory.SchemeGameCar: func(scheme entity.Scheme) entity.FactoryFn {
		if typed, ok := scheme.(schemefactory.Car); ok {
			return factory.CarFactoryFn(typed)
		}

		panic("invalid scheme")
	},
	schemefactory.SchemeGameUnit: func(scheme entity.Scheme) entity.FactoryFn {
		if typed, ok := scheme.(schemefactory.Unit); ok {
			return factory.UnitFactoryFn(typed)
		}

		panic("invalid scheme")
	},
	schemefactory.SchemeGameWall: func(scheme entity.Scheme) entity.FactoryFn {
		if typed, ok := scheme.(schemefactory.Wall); ok {
			return factory.WallFactoryFn(typed)
		}

		panic("invalid scheme")
	},
}
