package engine

import "github.com/fe3dback/galaxy/galx"

const (
	RenderModeWorld galx.RenderMode = iota
	RenderModeUI
)

type RenderTech string

const (
	RenderTechOpenGL2 RenderTech = "OpenGL2"
	RenderTechOpenGL3 RenderTech = "OpenGL3"
	RenderTechVulkan  RenderTech = "Vulkan"
)
