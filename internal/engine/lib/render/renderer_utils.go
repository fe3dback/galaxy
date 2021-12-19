package render

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/utils"
)

func (r *Renderer) FillRect(rect galx.Rect) {
	utils.Check("fill", r.ref.FillRect(r.transRectPtr(rect)))
}

func (r *Renderer) StartEngineFrame(color galx.Color) {
	var err error

	// clear all custom surfaces
	for i := galx.RenderTarget(0); i < surfacesCount; i++ {
		r.SetRenderTarget(i)
		r.SetDrawColor(color)

		err = r.ref.SetDrawBlendMode(sdl.BLENDMODE_NONE)
		utils.Check("set surface clear blendMode", err)

		err = r.ref.Clear()
		utils.Check("clear primary surface", err)
	}

	// set render target to scene
	r.SetRenderTarget(galx.RenderTargetMain)
}

func (r *Renderer) EndEngineFrame() {
	// set render target to scene
	r.SetRenderTarget(galx.RenderTargetMain)

	// copy other surfaces to scene
	for i := galx.RenderTarget(1); i < surfacesCount; i++ {
		err := r.ref.Copy(r.renderTarget.textureLayers[i], nil, nil)
		utils.Check("copy surface to main layer", err)
	}
}

func (r *Renderer) StartGUIFrame(color galx.Color) {
	r.renderTo(r.renderTarget.engineGUI)

	err := r.ref.SetDrawBlendMode(sdl.BLENDMODE_NONE)
	utils.Check("set gui surface clear blendMode", err)

	r.SetDrawColor(color)

	err = r.ref.Clear()
	utils.Check("clear gui surface", err)
}

func (r *Renderer) EndGUIFrame() {
	// do nothing here
}

func (r *Renderer) UpdateGPU() {
	// set target to screen texture
	err := r.ref.SetRenderTarget(r.renderTarget.screenTexture)
	utils.Check("set render target to screen texture", err)

	err = r.ref.CopyEx(r.renderTarget.engineGUI, nil, nil, 0, nil, sdl.FLIP_VERTICAL)
	utils.Check("copy GUI to main layer", err)

	// render
	r.ref.Present()
}
