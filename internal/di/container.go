package di

import (
	"github.com/fe3dback/govgl"
	"go.uber.org/zap"

	"github.com/fe3dback/galaxy/cfg"
	"github.com/fe3dback/galaxy/internal/engine"
	"github.com/fe3dback/galaxy/internal/engine/assets"
	"github.com/fe3dback/galaxy/internal/engine/control"
	engineEditor "github.com/fe3dback/galaxy/internal/engine/editor"
	"github.com/fe3dback/galaxy/internal/engine/event"
	"github.com/fe3dback/galaxy/internal/engine/gui"
	"github.com/fe3dback/galaxy/internal/engine/lib"
	oldRender "github.com/fe3dback/galaxy/internal/engine/lib/render"
	"github.com/fe3dback/galaxy/internal/engine/lib/sound"
	"github.com/fe3dback/galaxy/internal/engine/node"
	"github.com/fe3dback/galaxy/internal/engine/render"
	"github.com/fe3dback/galaxy/internal/engine/scene"
	"github.com/fe3dback/galaxy/internal/engine/windows"
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
			fontsManager   *oldRender.FontsManager
			textureManager *oldRender.TextureManager
			oldCamera      *oldRender.Camera
			renderer       *oldRender.Renderer
			gui            *gui.Gui
		}
		render struct {
			windowManager *windows.Manager
			inst          *render.Render
			vglRender     *govgl.Render
			camera        *render.Camera
		}
		engine struct {
			editorGizmos  *engineEditor.DrawGizmos
			appState      *engine.State
			gameState     *engine.GameState
			soundMixer    *sound.Manager
			assetsLoader  *assets.Manager
			scenesManager *scene.Manager
		}
		control struct {
			mouse    *control.Mouse
			keyboard *control.Keyboard
			movement *control.Movement
		}
		node struct {
			componentsRegistry *node.ComponentsRegistry
			query              *node.ObjectQuery
		}
		editor struct {
			manager *editor.Manager
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
