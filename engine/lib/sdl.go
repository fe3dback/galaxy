package lib

import (
	"fmt"

	"github.com/inkyblackness/imgui-go/v4"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

	"github.com/fe3dback/galaxy/engine/lib/renderers"
	"github.com/fe3dback/galaxy/utils"
)

// SDLClientAPI identifies the render system that shall be initialized.
type SDLClientAPI string

// This is a list of SDLClientAPI constants.
const (
	SDLClientAPIOpenGL2 SDLClientAPI = "OpenGL2"
	SDLClientAPIOpenGL3 SDLClientAPI = "OpenGL3"
)

// GUIRenderer covers rendering imgui draw data.
type GUIRenderer interface {
	// PreRender causes the display buffer to be prepared for new output.
	PreRender(clearColor [3]float32)
	// Render draws the provided imgui draw data.
	Render(displaySize [2]float32, framebufferSize [2]float32, drawData imgui.DrawData)
	// Dispose close renderer
	Dispose()
}

type SDLLib struct {
	window      *sdl.Window
	renderer    *sdl.Renderer
	guiRenderer GUIRenderer
	io          imgui.IO
}

func (s *SDLLib) Window() *sdl.Window {
	return s.window
}

func (s *SDLLib) Renderer() *sdl.Renderer {
	return s.renderer
}

func (s *SDLLib) GUIRenderer() GUIRenderer {
	return s.guiRenderer
}

func (s *SDLLib) GUI() imgui.IO {
	return s.io
}

func (s *SDLLib) quit() {
	if s.renderer != nil {
		_ = s.renderer.Destroy()
	}

	if s.window != nil {
		_ = s.window.Destroy()
	}

	sdl.Quit()
}

func NewSDLLib(closer *utils.Closer, defaultWidth, defaultHeight int, fullscreen bool) (*SDLLib, error) {
	platform := &SDLLib{}
	closer.EnqueueFree(platform.quit)

	defer utils.CheckPanicWith("sdl lib", func() {
		platform.quit()
	})

	// restrict to opengl2 for now
	const clientAPI = SDLClientAPIOpenGL2

	// create imgui
	context := imgui.CreateContext(nil)
	closer.EnqueueFree(context.Destroy)
	platform.io = imgui.CurrentIO()

	// init sdl
	err := sdl.Init(sdl.INIT_VIDEO & sdl.INIT_EVENTS)
	utils.Check("sdl init", err)

	err = ttf.Init()
	utils.Check("ttf init", err)

	// create main window
	var winFlags uint32
	winFlags |= sdl.WINDOW_OPENGL
	winFlags |= sdl.WINDOW_SHOWN
	winFlags |= sdl.WINDOW_ALLOW_HIGHDPI
	winFlags |= sdl.WINDOW_BORDERLESS

	winWidth := defaultWidth
	winHeight := defaultHeight

	if fullscreen {
		mode, displayModeErr := sdl.GetCurrentDisplayMode(0)
		utils.Check("get display mode", displayModeErr)

		winFlags &= sdl.WINDOW_FULLSCREEN
		winWidth = int(mode.W)
		winHeight = int(mode.H)
	}

	window, err := sdl.CreateWindow(
		"Galaxy",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		int32(winWidth), int32(winHeight),
		winFlags,
	)
	utils.Check("create window", err)
	closer.EnqueueClose(window.Destroy)
	platform.window = window

	// set mapping sdl -> gui
	setGUIKeyMapping(platform.io)

	// set openGL attributes
	// currently we support only opengl renderer
	switch clientAPI {
	case SDLClientAPIOpenGL2:
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 2)
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	case SDLClientAPIOpenGL3:
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 2)
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_FLAGS, sdl.GL_CONTEXT_FORWARD_COMPATIBLE_FLAG)
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	default:
		panic(fmt.Errorf("unknown render client API: %s", clientAPI))
	}

	// set additional openGL attributes
	_ = sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)
	_ = sdl.GLSetAttribute(sdl.GL_DEPTH_SIZE, 24)
	_ = sdl.GLSetAttribute(sdl.GL_STENCIL_SIZE, 8)

	// create openGL context
	glContext, err := window.GLCreateContext()
	utils.Check("failed to create OpenGL context", err)

	err = window.GLMakeCurrent(glContext)
	utils.Check("failed to set current OpenGL context", err)

	_ = sdl.GLSetSwapInterval(1)

	// create engine renderer
	surface, err := window.GetSurface()
	utils.Check("window get surface", err)

	err = surface.FillRect(nil, 0)
	utils.Check("clear window surface", err)

	err = window.UpdateSurface()
	utils.Check("update window surface", err)

	renderer, err := window.GetRenderer()
	if renderer == nil {
		renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	}

	utils.Check("create engine renderer", err)
	closer.EnqueueClose(renderer.Destroy)
	platform.renderer = renderer

	// create gui renderer
	var guiRenderer GUIRenderer

	switch clientAPI {
	case SDLClientAPIOpenGL2:
		guiRenderer, err = renderers.NewOpenGL2(platform.io)
	case SDLClientAPIOpenGL3:
		guiRenderer, err = renderers.NewOpenGL3(platform.io)
	}

	utils.Check("create GUI renderer", err)
	closer.EnqueueFree(guiRenderer.Dispose)
	platform.guiRenderer = guiRenderer

	return platform, nil
}

func setGUIKeyMapping(io imgui.IO) {
	keys := map[int]int{
		imgui.KeyTab:        sdl.SCANCODE_TAB,
		imgui.KeyLeftArrow:  sdl.SCANCODE_LEFT,
		imgui.KeyRightArrow: sdl.SCANCODE_RIGHT,
		imgui.KeyUpArrow:    sdl.SCANCODE_UP,
		imgui.KeyDownArrow:  sdl.SCANCODE_DOWN,
		imgui.KeyPageUp:     sdl.SCANCODE_PAGEUP,
		imgui.KeyPageDown:   sdl.SCANCODE_PAGEDOWN,
		imgui.KeyHome:       sdl.SCANCODE_HOME,
		imgui.KeyEnd:        sdl.SCANCODE_END,
		imgui.KeyInsert:     sdl.SCANCODE_INSERT,
		imgui.KeyDelete:     sdl.SCANCODE_DELETE,
		imgui.KeyBackspace:  sdl.SCANCODE_BACKSPACE,
		imgui.KeySpace:      sdl.SCANCODE_BACKSPACE,
		imgui.KeyEnter:      sdl.SCANCODE_RETURN,
		imgui.KeyEscape:     sdl.SCANCODE_ESCAPE,
		imgui.KeyA:          sdl.SCANCODE_A,
		imgui.KeyC:          sdl.SCANCODE_C,
		imgui.KeyV:          sdl.SCANCODE_V,
		imgui.KeyX:          sdl.SCANCODE_X,
		imgui.KeyY:          sdl.SCANCODE_Y,
		imgui.KeyZ:          sdl.SCANCODE_Z,
	}

	// Keyboard mapping. ImGui will use those indices to peek into the io.KeysDown[] array.
	for imguiKey, nativeKey := range keys {
		io.KeyMap(imguiKey, nativeKey)
	}
}
