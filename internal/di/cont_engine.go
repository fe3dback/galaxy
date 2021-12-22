package di

import (
	"fmt"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/internal/engine"
	"github.com/fe3dback/galaxy/internal/engine/assets"
	engineeditor "github.com/fe3dback/galaxy/internal/engine/editor"
	"github.com/fe3dback/galaxy/internal/engine/gui"
	"github.com/fe3dback/galaxy/internal/engine/lib"
	oldRender "github.com/fe3dback/galaxy/internal/engine/lib/render"
	"github.com/fe3dback/galaxy/internal/engine/lib/sound"
	"github.com/fe3dback/galaxy/internal/engine/render"
	"github.com/fe3dback/galaxy/internal/engine/render/vulkan"
	"github.com/fe3dback/galaxy/internal/engine/scene"
	"github.com/fe3dback/galaxy/internal/engine/windows"
)

func (c *Container) provideWindowsManager() *windows.Manager {
	if c.memstate.render.windowManager != nil {
		return c.memstate.render.windowManager
	}

	manager := windows.NewManager(
		c.closer(),
		c.ProvideEventDispatcher(),
		engine.RenderTechVulkan,
		c.Flags().ScreenWidth(),
		c.Flags().ScreenHeight(),
		c.Flags().IsFullscreen(),
	)

	c.memstate.render.windowManager = manager
	return c.memstate.render.windowManager
}

func (c *Container) provideSDL() *lib.SDLLib {
	// todo: remove SDL render
	panic("sdl used")

	if c.memstate.renderer.sdl != nil {
		return c.memstate.renderer.sdl
	}

	sdlLib, err := lib.NewSDLLib(
		c.closer(),
		c.Flags().ScreenWidth(),
		c.Flags().ScreenHeight(),
		c.Flags().IsFullscreen(),
	)

	if err != nil {
		panic(fmt.Sprintf("can`t create sdl: %v", err))
	}

	c.memstate.renderer.sdl = sdlLib
	return c.memstate.renderer.sdl
}

func (c *Container) ProvideEngineState() *engine.State {
	if c.memstate.engine.appState != nil {
		return c.memstate.engine.appState
	}

	c.memstate.engine.appState = engine.NewEngineState(
		c.ProvideEventDispatcher(),
		c.ProvideEngineScenesManager(),
		c.flags.IsDefaultGameMode(),
		c.flags.IsIncludeEditor(),
	)
	return c.memstate.engine.appState
}

func (c *Container) provideRenderFontsManager() *oldRender.FontsManager {
	if c.memstate.renderer.fontsManager != nil {
		return c.memstate.renderer.fontsManager
	}

	fonts := oldRender.NewFontsManager(c.closer())
	fonts.Load(consts.AssetDefaultFont)

	c.memstate.renderer.fontsManager = fonts
	return c.memstate.renderer.fontsManager
}

func (c *Container) provideRenderTextureManager() *oldRender.TextureManager {
	if c.memstate.renderer.textureManager != nil {
		return c.memstate.renderer.textureManager
	}

	c.memstate.renderer.textureManager = oldRender.NewTextureManager(
		c.provideSDL().Renderer(),
		c.closer(),
	)
	return c.memstate.renderer.textureManager
}

func (c *Container) provideRenderCamera() *oldRender.Camera {
	if c.memstate.renderer.camera != nil {
		return c.memstate.renderer.camera
	}

	c.memstate.renderer.camera = oldRender.NewCamera(
		c.ProvideEventDispatcher(),
	)
	return c.memstate.renderer.camera
}

func (c *Container) provideEditorGizmos() *engineeditor.DrawGizmos {
	if c.memstate.engine.editorGizmos != nil {
		return c.memstate.engine.editorGizmos
	}

	gizmos := engineeditor.NewDrawGizmos(
		c.ProvideEventDispatcher(),
		c.flags.DebugOpts().System,
	)

	c.memstate.engine.editorGizmos = gizmos
	return c.memstate.engine.editorGizmos
}

func (c *Container) provideSoundMixer() *sound.Manager {
	if c.memstate.engine.soundMixer != nil {
		return c.memstate.engine.soundMixer
	}

	soundMixer := sound.NewManager(c.closer())

	c.memstate.engine.soundMixer = soundMixer
	return c.memstate.engine.soundMixer
}

func (c *Container) ProvideEngineScenesManager() *scene.Manager {
	if c.memstate.engine.scenesManager != nil {
		return c.memstate.engine.scenesManager
	}

	scenesManager := scene.NewManager(
		c.ProvideEventDispatcher(),
		c.provideEngineAssetsLoader(),
		c.provideEngineNodeComponentsRegistry(),
		c.flags.IsIncludeEditor(),
	)

	c.memstate.engine.scenesManager = scenesManager
	return c.memstate.engine.scenesManager
}

func (c *Container) ProvideEngineRenderer() *render.Render {
	if c.memstate.render.inst != nil {
		return c.memstate.render.inst
	}

	renderer := render.NewRender(
		c.ProvideEngineRendererVulkan(),
	)

	c.memstate.render.inst = renderer
	return c.memstate.render.inst
}

func (c *Container) ProvideEngineRendererVulkan() *vulkan.Vk {
	if c.memstate.render.libVulkan != nil {
		return c.memstate.render.libVulkan
	}

	opts := c.flags.VulkanOpts()
	cfg := vulkan.NewConfig(
		vulkan.WithDebug(opts.Debug),
		vulkan.WithVSync(opts.VSync),
	)

	vk := vulkan.NewVulkanApi(
		c.provideWindowsManager().Window(),
		c.ProvideEventDispatcher(),
		cfg,
		c.closer(),
	)
	c.memstate.render.libVulkan = vk
	return c.memstate.render.libVulkan
}

func (c *Container) ProvideEngineRendererOLD() *oldRender.Renderer {
	if c.memstate.renderer.renderer != nil {
		return c.memstate.renderer.renderer
	}

	renderer := oldRender.NewRenderer(
		c.provideSDL().Window(),
		c.provideSDL().Renderer(),
		c.provideRenderFontsManager(),
		c.provideRenderTextureManager(),
		c.provideRenderCamera(),
		c.ProvideEventDispatcher(),
		c.provideEditorGizmos(),
		c.ProvideEngineState(),
	)

	c.memstate.renderer.renderer = renderer
	return c.memstate.renderer.renderer
}

func (c *Container) ProvideEngineGUI() *gui.Gui {
	if c.memstate.renderer.gui != nil {
		return c.memstate.renderer.gui
	}

	engineGUI := gui.NewGUI(
		c.closer(),
		c.ProvideEngineRenderer(),
		c.ProvideEventDispatcher(),
	)

	c.memstate.renderer.gui = engineGUI
	return c.memstate.renderer.gui
}

func (c *Container) ProvideEngineGameState() *engine.GameState {
	if c.memstate.engine.gameState != nil {
		return c.memstate.engine.gameState
	}

	gameState := engine.NewGameState(
		c.ProvideFrames(),
		c.provideRenderCamera(),
		c.provideEngineControlMouse(),
		c.provideEngineControlKeyboard(),
		c.provideEngineControlMovement(),
		c.ProvideEngineState(),
		c.provideSoundMixer(),
		c.ProvideEngineScenesManager(),
		c.provideEngineNodeQuery(),
	)

	c.memstate.engine.gameState = gameState
	return c.memstate.engine.gameState
}

func (c *Container) provideEngineAssetsLoader() *assets.Manager {
	return assets.NewAssetsManager(
		c.provideSoundMixer(),
	)
}
