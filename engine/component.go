package engine

type ComponentId string

type Component interface {
	Drawer
	Updater

	Id() ComponentId
}
