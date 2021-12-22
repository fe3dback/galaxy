package shader

import (
	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/galx"
)

const vertTriangleCount = 3
const vertTriangleSizePos = galx.SizeOfVec2
const vertTriangleSizeColor = galx.SizeOfVec3
const vertTriangleSizeVertex = vertTriangleSizePos + vertTriangleSizeColor
const vertTriangleSizeTotal = vertTriangleSizeVertex * vertTriangleCount

type (
	VertInTriangle struct {
		Position [vertTriangleCount]galx.Vec2
		Color    [vertTriangleCount]galx.Vec3
	}
)

func (x *VertInTriangle) Size() uint64 {
	return vertTriangleSizeTotal
}

func (x *VertInTriangle) Data() []byte {
	r := make([]byte, 0, x.Size())
	for i := 0; i < vertTriangleCount; i++ {
		r = append(r, x.Position[i].Data()...)
		r = append(r, x.Color[i].Data()...)
	}

	return r
}

func (x *VertInTriangle) Bindings() []vulkan.VertexInputBindingDescription {
	return []vulkan.VertexInputBindingDescription{
		{
			Binding:   0,
			Stride:    vertTriangleSizeVertex,
			InputRate: vulkan.VertexInputRateVertex,
		},
	}
}

func (x *VertInTriangle) Attributes() []vulkan.VertexInputAttributeDescription {
	return []vulkan.VertexInputAttributeDescription{
		{
			Location: 0,
			Binding:  0,
			Format:   vulkan.FormatR32g32Sfloat,
			Offset:   0,
		},
		{
			Location: 1,
			Binding:  0,
			Format:   vulkan.FormatR32g32b32Sfloat,
			Offset:   vertTriangleSizePos,
		},
	}
}
