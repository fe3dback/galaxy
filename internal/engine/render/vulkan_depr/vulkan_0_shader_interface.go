package vulkan_depr

import (
	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/engine/render/vulkan_depr/shader/shaderm"
)

type (
	shaderProgram interface {
		// shader bytecode

		ID() string
		ProgramFrag() []byte
		ProgramVert() []byte

		// attributes

		Size() uint64
		Data() []byte
		Bindings() []vulkan.VertexInputBindingDescription
		Attributes() []vulkan.VertexInputAttributeDescription
	}
)

var shaderModules = []shaderProgram{
	&shaderm.Triangle{},
}
