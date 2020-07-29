package ui

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"

	"github.com/fe3dback/galaxy/render"
)

type UI struct {
	layers []Layer
}

func NewUI(layers ...Layer) *UI {
	ui := &UI{
		layers: layers,
	}

	return ui
}

func (u *UI) OnDraw(r *render.Renderer) error {
	for _, layer := range u.layers {
		err := layer.OnDraw(r)
		if err != nil {
			return fmt.Errorf("draw ui: %v", err)
		}
	}

	return nil
}

func (u *UI) OnUpdate(moment engine.Moment) error {
	for _, layer := range u.layers {
		err := layer.OnUpdate(moment)
		if err != nil {
			return fmt.Errorf("draw ui: %v", err)
		}
	}

	return nil
}
