package di

import (
	"go.uber.org/zap"

	"github.com/fe3dback/galaxy/cfg"
	"github.com/fe3dback/galaxy/internal/engine"
	"github.com/fe3dback/galaxy/internal/engine/control"
	engineEditor "github.com/fe3dback/galaxy/internal/engine/editor"
	"github.com/fe3dback/galaxy/internal/engine/entity"
	"github.com/fe3dback/galaxy/internal/engine/event"
	"github.com/fe3dback/galaxy/internal/engine/lib"
	"github.com/fe3dback/galaxy/internal/engine/lib/render"
	"github.com/fe3dback/galaxy/internal/engine/lib/sound"
	"github.com/fe3dback/galaxy/internal/engine/loader"
	"github.com/fe3dback/galaxy/internal/engine/scene"
	"github.com/fe3dback/galaxy/internal/frames"
	"github.com/fe3dback/galaxy/internal/utils"
	"github.com/fe3dback/galaxy/scope/editor"
	"github.com/fe3dback/galaxy/scope/shared/ui"
)

type Container struct {
	flags    *cfg.InitFlags
	memstate struct {
		closer   *utils.Closer
		logger   *zap.SugaredLogger
		renderer struct {
			sdl            *lib.SDLLib
			fontsManager   *render.FontsManager
			textureManager *render.TextureManager
			camera         *render.Camera
			renderer       *render.Renderer
		}
		engine struct {
			editorGizmos  *engineEditor.DrawGizmos
			appState      *engine.EngineState
			gameState     *engine.GameState
			soundMixer    *sound.Manager
			assetsLoader  *loader.AssetsLoader
			scenesManager *scene.Manager
		}
		control struct {
			mouse    *control.Mouse
			movement *control.Movement
		}
		entity struct {
			componentsRegistry *entity.ComponentsRegistry
		}
		editor struct {
			manager *editor.Manager
			ui      *ui.UI
		}
		game struct {
			ui *ui.UI
		}
		eventDispatcher *event.Dispatcher
		frames          *frames.Frames
	}
}

func NewContainer(flags *cfg.InitFlags) *Container {
	c := &Container{
		flags: flags,
	}

	// this log used only for logger initialization at this place
	c.logger().Info("game container created")
	return c
}

func (c *Container) Flags() *cfg.InitFlags {
	return c.flags
}
