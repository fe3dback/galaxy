package ui

import (
	"github.com/fe3dback/galaxy/engine"
)

type (
	Layer interface {
		engine.Drawer
		engine.Updater
	}
)
