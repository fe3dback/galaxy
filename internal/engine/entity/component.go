package entity

import (
	"github.com/fe3dback/galaxy/galx"
)

type Component interface {
	galx.Drawer
	galx.Updater
}
