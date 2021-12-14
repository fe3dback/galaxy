package control

type (
	settingsPane interface {
		Extend(name string, priority int, behave func())
	}
)
