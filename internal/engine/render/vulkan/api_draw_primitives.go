package vulkan

import (
	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine/render/vulkan/shader/shaderm"
)

func (vk *Vk) DrawTmpTriangle() {
	for i := float64(-1); i < 1; i += 0.1 {
		vk.appendToRenderQueue(&shaderm.Triangle{
			Position: [3]galx.Vec2d{
				{X: i, Y: -0.5},
				{X: 0.5, Y: 0.5},
				{X: -0.5, Y: 0.5},
			},
			Color: [3]galx.Vec3d{
				{X: (i + 1) / 2, Y: 0, Z: 0},
				{X: 0, Y: 1, Z: 0},
				{X: 0, Y: 0, Z: 1},
			},
		})
	}
}

func (vk *Vk) DrawRect(vertexPos [4]galx.Vec2d, vertexColor [4]galx.Vec3d) {
	vk.appendToRenderQueue(&shaderm.Rect{
		Position: vertexPos,
		Color:    vertexColor,
	})
}
