package di

import (
	"fmt"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/internal/engine"
	engineeditor "github.com/fe3dback/galaxy/internal/engine/editor"
	"github.com/fe3dback/galaxy/internal/engine/lib"
	"github.com/fe3dback/galaxy/internal/engine/lib/render"
	"github.com/fe3dback/galaxy/internal/engine/lib/sound"
	"github.com/fe3dback/galaxy/internal/engine/loader"
	"github.com/fe3dback/galaxy/internal/engine/scene"
)

func (c *Container) provideSDL() *lib.SDLLib {
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

func (c *Container) provideRenderFontsManager() *render.FontsManager {
	if c.memstate.renderer.fontsManager != nil {
		return c.memstate.renderer.fontsManager
	}

	fonts := render.NewFontsManager(c.closer())
	fonts.Load(consts.AssetDefaultFont)

	c.memstate.renderer.fontsManager = fonts
	return c.memstate.renderer.fontsManager
}

func (c *Container) provideRenderTextureManager() *render.TextureManager {
	if c.memstate.renderer.textureManager != nil {
		return c.memstate.renderer.textureManager
	}

	c.memstate.renderer.textureManager = render.NewTextureManager(
		c.provideSDL().Renderer(),
		c.closer(),
	)
	return c.memstate.renderer.textureManager
}

func (c *Container) provideRenderCamera() *render.Camera {
	if c.memstate.renderer.camera != nil {
		return c.memstate.renderer.camera
	}

	c.memstate.renderer.camera = render.NewCamera(
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

func (c *Container) ProvideEngineRenderer() *render.Renderer {
	if c.memstate.renderer.renderer != nil {
		return c.memstate.renderer.renderer
	}

	renderer := render.NewRenderer(
		c.provideSDL().Window(),
		c.provideSDL().Renderer(),
		c.provideSDL().GUIRenderer(),
		c.provideSDL().GUI(),
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

func (c *Container) provideEngineAssetsLoader() *loader.AssetsLoader {
	return loader.NewAssetsLoader(
		c.provideSoundMixer(),
	)
}
