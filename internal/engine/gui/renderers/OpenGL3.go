package renderers

import (
	_ "embed" // using embed for the shader sources
	"fmt"

	"github.com/inkyblackness/imgui-go/v4"

	gl2 "github.com/fe3dback/galaxy/internal/engine/gui/renderers/gl/v3.2-core/gl"
)

//go:embed gl-shader/main.vert
var unversionedVertexShader string

//go:embed gl-shader/main.frag
var unversionedFragmentShader string

// OpenGL3 implements a renderer based on github.com/go-gl/gl (v3.2-core).
type OpenGL3 struct {
	imguiIO imgui.IO

	glslVersion            string
	fontTexture            uint32
	shaderHandle           uint32
	vertHandle             uint32
	fragHandle             uint32
	attribLocationTex      int32
	attribLocationProjMtx  int32
	attribLocationPosition int32
	attribLocationUV       int32
	attribLocationColor    int32
	vboHandle              uint32
	elementsHandle         uint32
}

// NewOpenGL3 attempts to initialize a renderer.
// An OpenGL context has to be established before calling this function.
func NewOpenGL3(io imgui.IO) (*OpenGL3, error) {
	err := gl2.Init()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize OpenGL: %w", err)
	}

	renderer := &OpenGL3{
		imguiIO:     io,
		glslVersion: "#version 150",
	}
	renderer.createDeviceObjects()

	io.SetBackendFlags(io.GetBackendFlags() | imgui.BackendFlagsRendererHasVtxOffset)

	return renderer, nil
}

// Dispose cleans up the resources.
func (renderer *OpenGL3) Dispose() {
	renderer.invalidateDeviceObjects()
}

// PreRender clears the framebuffer.
func (renderer *OpenGL3) PreRender(clearColor [3]float32) {
	gl2.ClearColor(clearColor[0], clearColor[1], clearColor[2], 1.0)
	gl2.Clear(gl2.COLOR_BUFFER_BIT)
}

// Render translates the ImGui draw data to OpenGL3 commands.
func (renderer *OpenGL3) Render(displaySize [2]float32, framebufferSize [2]float32, drawData imgui.DrawData) {
	// Avoid rendering when minimized, scale coordinates for retina displays (screen coordinates != framebuffer coordinates)
	displayWidth, displayHeight := displaySize[0], displaySize[1]
	fbWidth, fbHeight := framebufferSize[0], framebufferSize[1]
	if (fbWidth <= 0) || (fbHeight <= 0) {
		return
	}
	drawData.ScaleClipRects(imgui.Vec2{
		X: fbWidth / displayWidth,
		Y: fbHeight / displayHeight,
	})

	// Backup GL state
	var lastActiveTexture int32
	gl2.GetIntegerv(gl2.ACTIVE_TEXTURE, &lastActiveTexture)
	gl2.ActiveTexture(gl2.TEXTURE0)
	var lastProgram int32
	gl2.GetIntegerv(gl2.CURRENT_PROGRAM, &lastProgram)
	var lastTexture int32
	gl2.GetIntegerv(gl2.TEXTURE_BINDING_2D, &lastTexture)
	var lastSampler int32
	gl2.GetIntegerv(gl2.SAMPLER_BINDING, &lastSampler)
	var lastArrayBuffer int32
	gl2.GetIntegerv(gl2.ARRAY_BUFFER_BINDING, &lastArrayBuffer)
	var lastElementArrayBuffer int32
	gl2.GetIntegerv(gl2.ELEMENT_ARRAY_BUFFER_BINDING, &lastElementArrayBuffer)
	var lastVertexArray int32
	gl2.GetIntegerv(gl2.VERTEX_ARRAY_BINDING, &lastVertexArray)
	var lastPolygonMode [2]int32
	gl2.GetIntegerv(gl2.POLYGON_MODE, &lastPolygonMode[0])
	var lastViewport [4]int32
	gl2.GetIntegerv(gl2.VIEWPORT, &lastViewport[0])
	var lastScissorBox [4]int32
	gl2.GetIntegerv(gl2.SCISSOR_BOX, &lastScissorBox[0])
	var lastBlendSrcRgb int32
	gl2.GetIntegerv(gl2.BLEND_SRC_RGB, &lastBlendSrcRgb)
	var lastBlendDstRgb int32
	gl2.GetIntegerv(gl2.BLEND_DST_RGB, &lastBlendDstRgb)
	var lastBlendSrcAlpha int32
	gl2.GetIntegerv(gl2.BLEND_SRC_ALPHA, &lastBlendSrcAlpha)
	var lastBlendDstAlpha int32
	gl2.GetIntegerv(gl2.BLEND_DST_ALPHA, &lastBlendDstAlpha)
	var lastBlendEquationRgb int32
	gl2.GetIntegerv(gl2.BLEND_EQUATION_RGB, &lastBlendEquationRgb)
	var lastBlendEquationAlpha int32
	gl2.GetIntegerv(gl2.BLEND_EQUATION_ALPHA, &lastBlendEquationAlpha)
	lastEnableBlend := gl2.IsEnabled(gl2.BLEND)
	lastEnableCullFace := gl2.IsEnabled(gl2.CULL_FACE)
	lastEnableDepthTest := gl2.IsEnabled(gl2.DEPTH_TEST)
	lastEnableScissorTest := gl2.IsEnabled(gl2.SCISSOR_TEST)

	// Setup render state: alpha-blending enabled, no face culling, no depth testing, scissor enabled, polygon fill
	gl2.Enable(gl2.BLEND)
	gl2.BlendEquation(gl2.FUNC_ADD)
	gl2.BlendFunc(gl2.SRC_ALPHA, gl2.ONE_MINUS_SRC_ALPHA)
	gl2.Disable(gl2.CULL_FACE)
	gl2.Disable(gl2.DEPTH_TEST)
	gl2.Enable(gl2.SCISSOR_TEST)
	gl2.PolygonMode(gl2.FRONT_AND_BACK, gl2.FILL)

	// Setup viewport, orthographic projection matrix
	// Our visible imgui space lies from draw_data->DisplayPos (top left) to draw_data->DisplayPos+data_data->DisplaySize (bottom right).
	// DisplayMin is typically (0,0) for single viewport apps.
	gl2.Viewport(0, 0, int32(fbWidth), int32(fbHeight))
	orthoProjection := [4][4]float32{
		{2.0 / displayWidth, 0.0, 0.0, 0.0},
		{0.0, 2.0 / -displayHeight, 0.0, 0.0},
		{0.0, 0.0, -1.0, 0.0},
		{-1.0, 1.0, 0.0, 1.0},
	}
	gl2.UseProgram(renderer.shaderHandle)
	gl2.Uniform1i(renderer.attribLocationTex, 0)
	gl2.UniformMatrix4fv(renderer.attribLocationProjMtx, 1, false, &orthoProjection[0][0])
	gl2.BindSampler(0, 0) // Rely on combined texture/sampler state.

	// Recreate the VAO every time
	// (This is to easily allow multiple GL contexts. VAO are not shared among GL contexts, and
	// we don't track creation/deletion of windows so we don't have an obvious key to use to cache them.)
	var vaoHandle uint32
	gl2.GenVertexArrays(1, &vaoHandle)
	gl2.BindVertexArray(vaoHandle)
	gl2.BindBuffer(gl2.ARRAY_BUFFER, renderer.vboHandle)
	gl2.EnableVertexAttribArray(uint32(renderer.attribLocationPosition))
	gl2.EnableVertexAttribArray(uint32(renderer.attribLocationUV))
	gl2.EnableVertexAttribArray(uint32(renderer.attribLocationColor))
	vertexSize, vertexOffsetPos, vertexOffsetUv, vertexOffsetCol := imgui.VertexBufferLayout()
	gl2.VertexAttribPointerWithOffset(uint32(renderer.attribLocationPosition), 2, gl2.FLOAT, false, int32(vertexSize), uintptr(vertexOffsetPos))
	gl2.VertexAttribPointerWithOffset(uint32(renderer.attribLocationUV), 2, gl2.FLOAT, false, int32(vertexSize), uintptr(vertexOffsetUv))
	gl2.VertexAttribPointerWithOffset(uint32(renderer.attribLocationColor), 4, gl2.UNSIGNED_BYTE, true, int32(vertexSize), uintptr(vertexOffsetCol))
	indexSize := imgui.IndexBufferLayout()
	drawType := gl2.UNSIGNED_SHORT
	const bytesPerUint32 = 4
	if indexSize == bytesPerUint32 {
		drawType = gl2.UNSIGNED_INT
	}

	// Draw
	for _, list := range drawData.CommandLists() {
		vertexBuffer, vertexBufferSize := list.VertexBuffer()
		gl2.BindBuffer(gl2.ARRAY_BUFFER, renderer.vboHandle)
		gl2.BufferData(gl2.ARRAY_BUFFER, vertexBufferSize, vertexBuffer, gl2.STREAM_DRAW)

		indexBuffer, indexBufferSize := list.IndexBuffer()
		gl2.BindBuffer(gl2.ELEMENT_ARRAY_BUFFER, renderer.elementsHandle)
		gl2.BufferData(gl2.ELEMENT_ARRAY_BUFFER, indexBufferSize, indexBuffer, gl2.STREAM_DRAW)

		for _, cmd := range list.Commands() {
			if cmd.HasUserCallback() {
				cmd.CallUserCallback(list)
			} else {
				gl2.BindTexture(gl2.TEXTURE_2D, uint32(cmd.TextureID()))
				clipRect := cmd.ClipRect()
				gl2.Scissor(int32(clipRect.X), int32(fbHeight)-int32(clipRect.W), int32(clipRect.Z-clipRect.X), int32(clipRect.W-clipRect.Y))
				gl2.DrawElementsBaseVertexWithOffset(gl2.TRIANGLES, int32(cmd.ElementCount()), uint32(drawType),
					uintptr(cmd.IndexOffset()*indexSize), int32(cmd.VertexOffset()))
			}
		}
	}
	gl2.DeleteVertexArrays(1, &vaoHandle)

	// Restore modified GL state
	gl2.UseProgram(uint32(lastProgram))
	gl2.BindTexture(gl2.TEXTURE_2D, uint32(lastTexture))
	gl2.BindSampler(0, uint32(lastSampler))
	gl2.ActiveTexture(uint32(lastActiveTexture))
	gl2.BindVertexArray(uint32(lastVertexArray))
	gl2.BindBuffer(gl2.ARRAY_BUFFER, uint32(lastArrayBuffer))
	gl2.BindBuffer(gl2.ELEMENT_ARRAY_BUFFER, uint32(lastElementArrayBuffer))
	gl2.BlendEquationSeparate(uint32(lastBlendEquationRgb), uint32(lastBlendEquationAlpha))
	gl2.BlendFuncSeparate(uint32(lastBlendSrcRgb), uint32(lastBlendDstRgb), uint32(lastBlendSrcAlpha), uint32(lastBlendDstAlpha))
	if lastEnableBlend {
		gl2.Enable(gl2.BLEND)
	} else {
		gl2.Disable(gl2.BLEND)
	}
	if lastEnableCullFace {
		gl2.Enable(gl2.CULL_FACE)
	} else {
		gl2.Disable(gl2.CULL_FACE)
	}
	if lastEnableDepthTest {
		gl2.Enable(gl2.DEPTH_TEST)
	} else {
		gl2.Disable(gl2.DEPTH_TEST)
	}
	if lastEnableScissorTest {
		gl2.Enable(gl2.SCISSOR_TEST)
	} else {
		gl2.Disable(gl2.SCISSOR_TEST)
	}
	gl2.PolygonMode(gl2.FRONT_AND_BACK, uint32(lastPolygonMode[0]))
	gl2.Viewport(lastViewport[0], lastViewport[1], lastViewport[2], lastViewport[3])
	gl2.Scissor(lastScissorBox[0], lastScissorBox[1], lastScissorBox[2], lastScissorBox[3])
}

func (renderer *OpenGL3) createDeviceObjects() {
	// Backup GL state
	var lastTexture int32
	var lastArrayBuffer int32
	var lastVertexArray int32
	gl2.GetIntegerv(gl2.TEXTURE_BINDING_2D, &lastTexture)
	gl2.GetIntegerv(gl2.ARRAY_BUFFER_BINDING, &lastArrayBuffer)
	gl2.GetIntegerv(gl2.VERTEX_ARRAY_BINDING, &lastVertexArray)

	vertexShader := renderer.glslVersion + "\n" + unversionedVertexShader
	fragmentShader := renderer.glslVersion + "\n" + unversionedFragmentShader

	renderer.shaderHandle = gl2.CreateProgram()
	renderer.vertHandle = gl2.CreateShader(gl2.VERTEX_SHADER)
	renderer.fragHandle = gl2.CreateShader(gl2.FRAGMENT_SHADER)

	glShaderSource := func(handle uint32, source string) {
		csource, free := gl2.Strs(source + "\x00")
		defer free()

		gl2.ShaderSource(handle, 1, csource, nil)
	}

	glShaderSource(renderer.vertHandle, vertexShader)
	glShaderSource(renderer.fragHandle, fragmentShader)
	gl2.CompileShader(renderer.vertHandle)
	gl2.CompileShader(renderer.fragHandle)
	gl2.AttachShader(renderer.shaderHandle, renderer.vertHandle)
	gl2.AttachShader(renderer.shaderHandle, renderer.fragHandle)
	gl2.LinkProgram(renderer.shaderHandle)

	renderer.attribLocationTex = gl2.GetUniformLocation(renderer.shaderHandle, gl2.Str("Texture"+"\x00"))
	renderer.attribLocationProjMtx = gl2.GetUniformLocation(renderer.shaderHandle, gl2.Str("ProjMtx"+"\x00"))
	renderer.attribLocationPosition = gl2.GetAttribLocation(renderer.shaderHandle, gl2.Str("Position"+"\x00"))
	renderer.attribLocationUV = gl2.GetAttribLocation(renderer.shaderHandle, gl2.Str("UV"+"\x00"))
	renderer.attribLocationColor = gl2.GetAttribLocation(renderer.shaderHandle, gl2.Str("Color"+"\x00"))

	gl2.GenBuffers(1, &renderer.vboHandle)
	gl2.GenBuffers(1, &renderer.elementsHandle)

	renderer.createFontsTexture()

	// Restore modified GL state
	gl2.BindTexture(gl2.TEXTURE_2D, uint32(lastTexture))
	gl2.BindBuffer(gl2.ARRAY_BUFFER, uint32(lastArrayBuffer))
	gl2.BindVertexArray(uint32(lastVertexArray))
}

func (renderer *OpenGL3) createFontsTexture() {
	// Build texture atlas
	io := imgui.CurrentIO()
	image := io.Fonts().TextureDataAlpha8()

	// Upload texture to graphics system
	var lastTexture int32
	gl2.GetIntegerv(gl2.TEXTURE_BINDING_2D, &lastTexture)
	gl2.GenTextures(1, &renderer.fontTexture)
	gl2.BindTexture(gl2.TEXTURE_2D, renderer.fontTexture)
	gl2.TexParameteri(gl2.TEXTURE_2D, gl2.TEXTURE_MIN_FILTER, gl2.LINEAR)
	gl2.TexParameteri(gl2.TEXTURE_2D, gl2.TEXTURE_MAG_FILTER, gl2.LINEAR)
	gl2.PixelStorei(gl2.UNPACK_ROW_LENGTH, 0)
	gl2.TexImage2D(gl2.TEXTURE_2D, 0, gl2.RED, int32(image.Width), int32(image.Height),
		0, gl2.RED, gl2.UNSIGNED_BYTE, image.Pixels)

	// Store our identifier
	io.Fonts().SetTextureID(imgui.TextureID(renderer.fontTexture))

	// Restore state
	gl2.BindTexture(gl2.TEXTURE_2D, uint32(lastTexture))
}

func (renderer *OpenGL3) invalidateDeviceObjects() {
	if renderer.vboHandle != 0 {
		gl2.DeleteBuffers(1, &renderer.vboHandle)
	}
	renderer.vboHandle = 0
	if renderer.elementsHandle != 0 {
		gl2.DeleteBuffers(1, &renderer.elementsHandle)
	}
	renderer.elementsHandle = 0

	if (renderer.shaderHandle != 0) && (renderer.vertHandle != 0) {
		gl2.DetachShader(renderer.shaderHandle, renderer.vertHandle)
	}
	if renderer.vertHandle != 0 {
		gl2.DeleteShader(renderer.vertHandle)
	}
	renderer.vertHandle = 0

	if (renderer.shaderHandle != 0) && (renderer.fragHandle != 0) {
		gl2.DetachShader(renderer.shaderHandle, renderer.fragHandle)
	}
	if renderer.fragHandle != 0 {
		gl2.DeleteShader(renderer.fragHandle)
	}
	renderer.fragHandle = 0

	if renderer.shaderHandle != 0 {
		gl2.DeleteProgram(renderer.shaderHandle)
	}
	renderer.shaderHandle = 0

	if renderer.fontTexture != 0 {
		gl2.DeleteTextures(1, &renderer.fontTexture)
		imgui.CurrentIO().Fonts().SetTextureID(0)
		renderer.fontTexture = 0
	}
}
