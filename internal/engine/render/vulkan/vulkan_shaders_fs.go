package vulkan

import _ "embed"

// all common shaders include as byte code inside engine binary
// custom shaders will be loaded from assets, and compile on the fly

//go:embed shader/triangle.frag.spv
var commonShaderTriangleFrag []byte

//go:embed shader/triangle.vert.spv
var commonShaderTriangleVert []byte
