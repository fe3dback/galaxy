package editor

import "github.com/fe3dback/galaxy/galx"

type (
	Component interface {
		galx.Updater
		galx.Drawer
	}
)
