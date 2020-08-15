package editor

import "github.com/fe3dback/galaxy/engine"

type (
	Component interface {
		engine.Updater
		engine.Drawer
	}
)
