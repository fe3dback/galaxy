package engine

type RenderTech string

const (
	RenderTechOpenGL2 RenderTech = "OpenGL2"
	RenderTechOpenGL3 RenderTech = "OpenGL3"
	RenderTechVulkan  RenderTech = "Vulkan"
)
