package shaderm

import (
	_ "embed"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/galx"
)

const pRectTriangleCount = 2
const pRectVertexCount = 4
const pRectSizePos = galx.SizeOfVec2
const pRectSizeColor = galx.SizeOfVec3
const pRectSizeVertex = pRectSizePos + pRectSizeColor
const pRectSizeTotal = pRectSizeVertex * (pRectTriangleCount * 3)

type (
	Rect struct {
		Position [pRectVertexCount]galx.Vec2
		Color    [pRectVertexCount]galx.Vec3
	}
)

var (
	//go:embed compiled/rect.frag.spv
	rectFrag []byte

	//go:embed compiled/rect.vert.spv
	rectVert []byte
)

func (x *Rect) ID() string {
	return "rect"
}

func (x *Rect) ProgramFrag() []byte {
	return rectFrag
}

func (x *Rect) ProgramVert() []byte {
	return rectVert
}

func (x *Rect) Size() uint64 {
	return pRectSizeTotal
}

func (x *Rect) TriangleCount() uint32 {
	return pRectTriangleCount
}

func (x *Rect) VertexCount() uint32 {
	return pRectVertexCount
}

func (x *Rect) Data() []byte {
	r := make([]byte, 0, x.Size())
	list := []int{2, 1, 0, 2, 0, 3}

	for _, vInd := range list {
		r = append(r, x.Position[vInd].Data()...)
		r = append(r, x.Color[vInd].Data()...)
	}

	return r
}

func (x *Rect) Bindings() []vulkan.VertexInputBindingDescription {
	return []vulkan.VertexInputBindingDescription{
		{
			Binding:   0,
			Stride:    pRectSizeVertex,
			InputRate: vulkan.VertexInputRateVertex,
		},
	}
}

func (x *Rect) Attributes() []vulkan.VertexInputAttributeDescription {
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
			Offset:   pRectSizePos,
		},
	}
}
