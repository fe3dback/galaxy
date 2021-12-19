package gui

import (
	"fmt"
	"sort"

	"github.com/inkyblackness/imgui-go/v4"

	"github.com/fe3dback/galaxy/galx"
)

type (
	Settings struct {
		enabled bool
		panels  map[string]Panel
	}

	Panel struct {
		priority int
		behave   func()
		active   bool
	}
)

func (g Settings) Id() string {
	return "4e866e2a-5451-4069-b524-0d022c82bc8c"
}

func (g *Settings) OnCreated(_ galx.EditorComponentResolver) {
	g.panels = map[string]Panel{}
}

func (g *Settings) Extend(name string, priority int, behave func()) {
	if panel, exist := g.panels[name]; exist {
		panel.active = true
		panel.priority = priority
		return
	}

	g.panels[name] = Panel{
		priority: priority,
		behave:   behave,
		active:   true,
	}
}

func (g *Settings) OnUpdate(s galx.State) error {
	sortedPanels := make([]string, 0, len(g.panels))
	for id := range g.panels {
		sortedPanels = append(sortedPanels, id)
	}

	sort.Slice(sortedPanels, func(i, j int) bool {
		pI := g.panels[sortedPanels[i]]
		pJ := g.panels[sortedPanels[j]]
		return pI.priority >= pJ.priority
	})

	windowTitle := fmt.Sprintf("Settings [%d / %dms]",
		s.Moment().FPS(),
		s.Moment().FrameDuration().Milliseconds(),
	)
	imgui.BeginV(windowTitle+"###Settings", &g.enabled, 0)

	for _, id := range sortedPanels {
		panel := g.panels[id]
		if !panel.active {
			continue
		}

		if imgui.CollapsingHeader(id) {
			panel.behave()
		}
	}

	imgui.End()

	for _, panel := range g.panels {
		panel.active = false
	}

	return nil
}
