package ui

import (
	"github.com/fe3dback/galaxy/galx"
)

type (
	Layer interface {
		galx.Drawer
		galx.Updater
	}
)
