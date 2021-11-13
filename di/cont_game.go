package di

import (
	"go.uber.org/zap"

	"github.com/fe3dback/galaxy/game"
	"github.com/fe3dback/galaxy/shared/ui"
	"github.com/fe3dback/galaxy/system"
)

func (c *Container) ProvideGameOptions() *system.GameOptions {
	if c.memstate.gameOptions != nil {
		return c.memstate.gameOptions
	}

	c.memstate.gameOptions = &system.GameOptions{
		Debug: system.DebugOpt{
			InProfiling: c.Flags().IsProfiling,
			System:      true,
			Memory:      false,
			Frames:      false,
			World:       false,
		},
		Frames: system.FramesOpt{
			TargetFps: c.Flags().TargetFPS,
		},
	}

	c.logger().Infow("system registered",
		zap.Int64("seed", c.flags.Seed),
		zap.Int("targetFPS", c.memstate.gameOptions.Frames.TargetFps),
	)

	return c.memstate.gameOptions
}

func (c *Container) ProvideGameUI() *ui.UI {
	if c.memstate.game.ui != nil {
		return c.memstate.game.ui
	}

	c.memstate.game.ui = ui.NewUI(
		c.createUILayerFPS(),
	)
	return c.memstate.game.ui
}

func (c *Container) ProvideGameScenesLoader() *game.ScenesLoader {
	if c.memstate.game.ScenesLoader != nil {
		return c.memstate.game.ScenesLoader
	}

	c.memstate.game.ScenesLoader = game.NewScenesLoader(
		c.ProvideEngineScenesManager(),
	)
	return c.memstate.game.ScenesLoader
}
