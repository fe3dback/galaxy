package ui

import (
	"sort"

	"github.com/inkyblackness/imgui-go/v4"

	"github.com/fe3dback/galaxy/galx"
)

type (
	LayerSettings struct {
		open   bool
		panels map[string]Panel
	}

	Panel struct {
		priority int
		behave   func()
		active   bool
	}
)

func NewLayerSettings() *LayerSettings {
	return &LayerSettings{
		open:   true,
		panels: map[string]Panel{},
	}
}

func (l *LayerSettings) Extend(name string, priority int, behave func()) {
	if panel, exist := l.panels[name]; exist {
		panel.active = true
		panel.priority = priority
		return
	}

	l.panels[name] = Panel{
		priority: priority,
		behave:   behave,
		active:   true,
	}
}

func (l *LayerSettings) OnUpdate(_ galx.State) error {
	sortedPanels := make([]string, 0, len(l.panels))
	for id := range l.panels {
		sortedPanels = append(sortedPanels, id)
	}

	sort.Slice(sortedPanels, func(i, j int) bool {
		pI := l.panels[sortedPanels[i]]
		pJ := l.panels[sortedPanels[j]]
		return pI.priority >= pJ.priority
	})

	imgui.BeginV("Settings", &l.open, 0)

	for _, id := range sortedPanels {
		panel := l.panels[id]
		if !panel.active {
			continue
		}

		if imgui.CollapsingHeader(id) {
			panel.behave()
		}
	}

	imgui.End()

	for _, panel := range l.panels {
		panel.active = false
	}

	return nil
}

func (l *LayerSettings) OnDraw(_ galx.Renderer) (err error) {
	return nil
}
