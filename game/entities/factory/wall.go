package factory

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities/factory/schemefactory"
	"github.com/fe3dback/galaxy/game/utils/physics"
)

func WallFactoryFn(params schemefactory.Wall) entity.FactoryFn {
	return func(e *entity.Entity, creator engine.WorldCreator) *entity.Entity {
		shapeBox := creator.Physics().CreateShapeBox(
			params.BoxWidth,
			params.BoxWidth,
		)

		staticBody := creator.Physics().AddBodyStatic(
			e.Position(),
			e.Rotation(),
			shapeBox,
			physics.LayerGround.Category(),
			physics.LayerGround.Mask(),
		)

		e.AttachPhysicsBody(staticBody)

		return e
	}
}
