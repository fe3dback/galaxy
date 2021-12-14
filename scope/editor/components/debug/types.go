package debug

type (
	settingsPane interface {
		Extend(name string, priority int, behave func())
	}
)
