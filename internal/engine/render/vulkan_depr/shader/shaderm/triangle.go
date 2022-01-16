package shaderm

import (
	_ "embed"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/galx"
)

const pTriangleVertexCount = 3
const pTriangleSizePos = galx.SizeOfVec2
const pTriangleSizeColor = galx.SizeOfVec3
const pTriangleSizeVertex = pTriangleSizePos + pTriangleSizeColor
const pTriangleSizeTotal = pTriangleSizeVertex * pTriangleVertexCount

type (
	Triangle struct {
		Position [pTriangleVertexCount]galx.Vec2
		Color    [pTriangleVertexCount]galx.Vec3
	}
)

var (
	//go:embed compiled/triangle.frag.spv
	triangleFrag []byte

	//go:embed compiled/triangle.vert.spv
	triangleVert []byte
)

func (x *Triangle) ID() string {
	return "triangle"
}

func (x *Triangle) ProgramFrag() []byte {
	return triangleFrag
}

func (x *Triangle) ProgramVert() []byte {
	return triangleVert
}

func (x *Triangle) Size() uint64 {
	return pTriangleSizeTotal
}

func (x *Triangle) VertexCount() uint32 {
	return pTriangleVertexCount
}

func (x *Triangle) Data() []byte {
	r := make([]byte, 0, x.Size())
	for i := 0; i < pTriangleVertexCount; i++ {
		r = append(r, x.Position[i].Data()...)
		r = append(r, x.Color[i].Data()...)
	}

	return r
}

func (x *Triangle) Bindings() []vulkan.VertexInputBindingDescription {
	return []vulkan.VertexInputBindingDescription{
		{
			Binding:   0,
			Stride:    pTriangleSizeVertex,
			InputRate: vulkan.VertexInputRateVertex,
		},
	}
}

func (x *Triangle) Attributes() []vulkan.VertexInputAttributeDescription {
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
			Offset:   pTriangleSizePos,
		},
	}
}
