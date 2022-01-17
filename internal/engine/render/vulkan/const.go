package vulkan

import "time"

// GPU timeout for render. After this, app will be crashed
const swapChainTimeout = time.Second * 10

// How many frames can be failed continuously before crash
const maxPresetFails = 100

// Capacity of each vertex buffer page
const vertexBufferSize = 1 * 1024 // 1 MB

const (
	shaderEntryPoint = "main"
	shaderTypeVert   = ".vert"
	shaderTypeFrag   = ".frag"
)
