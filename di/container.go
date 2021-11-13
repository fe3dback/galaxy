package di

import (
	"go.uber.org/zap"

	"github.com/fe3dback/galaxy/editor"
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/control"
	engineEditor "github.com/fe3dback/galaxy/engine/editor"
	"github.com/fe3dback/galaxy/engine/event"
	"github.com/fe3dback/galaxy/engine/lib"
	"github.com/fe3dback/galaxy/engine/lib/render"
	"github.com/fe3dback/galaxy/engine/lib/sound"
	"github.com/fe3dback/galaxy/engine/loader"
	"github.com/fe3dback/galaxy/engine/scene"
	"github.com/fe3dback/galaxy/game"
	"github.com/fe3dback/galaxy/shared/ui"
	"github.com/fe3dback/galaxy/system"
	"github.com/fe3dback/galaxy/utils"
)

type Container struct {
	flags    *InitFlags
	memstate struct {
		closer      *utils.Closer
		logger      *zap.SugaredLogger
		gameOptions *system.GameOptions
		renderer    struct {
			sdl            *lib.SDLLib
			fontsManager   *render.FontsManager
			textureManager *render.TextureManager
			camera         *render.Camera
			renderer       *render.Renderer
		}
		engine struct {
			editorGizmos  *engineEditor.DrawGizmos
			appState      *engine.AppState
			gameState     *engine.GameState
			soundMixer    *sound.Manager
			assetsLoader  *loader.AssetsLoader
			scenesManager *scene.Manager
		}
		control struct {
			mouse    *control.Mouse
			movement *control.Movement
		}
		editor struct {
			manager *editor.Manager
			ui      *ui.UI
		}
		game struct {
			ScenesLoader *game.ScenesLoader
			ui           *ui.UI
		}
		eventDispatcher *event.Dispatcher
		frames          *system.Frames
	}
}

func NewContainer(flags *InitFlags) *Container {
	return &Container{
		flags: flags,
	}
}

func (c *Container) Flags() *InitFlags {
	return c.flags
}
