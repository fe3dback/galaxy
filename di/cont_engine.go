package di

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	engineeditor "github.com/fe3dback/galaxy/engine/editor"
	"github.com/fe3dback/galaxy/engine/lib"
	"github.com/fe3dback/galaxy/engine/lib/render"
	"github.com/fe3dback/galaxy/engine/lib/sound"
	"github.com/fe3dback/galaxy/engine/loader"
	"github.com/fe3dback/galaxy/engine/scene"
	"github.com/fe3dback/galaxy/generated"
)

func (c *Container) provideSDL() *lib.SDLLib {
	if c.memstate.renderer.sdl != nil {
		return c.memstate.renderer.sdl
	}

	sdlLib, err := lib.NewSDLLib(
		c.closer(),
		c.Flags().Width,
		c.Flags().Height,
		c.Flags().FullScreen,
	)

	if err != nil {
		panic(fmt.Sprintf("can`t create sdl: %v", err))
	}

	c.memstate.renderer.sdl = sdlLib
	return c.memstate.renderer.sdl
}

func (c *Container) ProvideEngineAppState() *engine.AppState {
	if c.memstate.engine.appState != nil {
		return c.memstate.engine.appState
	}

	c.memstate.engine.appState = engine.NewAppState(
		c.ProvideEventDispatcher(),
	)
	return c.memstate.engine.appState
}

func (c *Container) provideRenderFontsManager() *render.FontsManager {
	if c.memstate.renderer.fontsManager != nil {
		return c.memstate.renderer.fontsManager
	}

	fonts := render.NewFontsManager(c.closer())
	fonts.Load(generated.ResourcesFontsJetBrainsMonoRegular)

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
		c.ProvideGameOptions().Debug.System,
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
		c.ProvideEngineAppState(),
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
		c.provideEngineControlMovement(),
		c.ProvideEngineAppState(),
		c.provideSoundMixer(),
		c.ProvideEngineScenesManager(),
	)

	c.memstate.engine.gameState = gameState
	return c.memstate.engine.gameState
}

func (c *Container) provideEngineAssetsLoader() *loader.AssetsLoader {
	if c.memstate.engine.assetsLoader != nil {
		return c.memstate.engine.assetsLoader
	}

	assetsLoader := loader.NewAssetsLoader(
		c.provideSoundMixer(),
	)

	c.memstate.engine.assetsLoader = assetsLoader
	return c.memstate.engine.assetsLoader
}
