package render

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/utils"
)

func (r *Renderer) FillRect(rect galx.Rect) {
	utils.Check("fill", r.ref.FillRect(r.transRectPtr(rect)))
}

func (r *Renderer) Clear(color galx.Color) {
	var err error

	// draw all surfaces
	for i := 1; i <= 1; i++ { // todo: surfacesCount
		r.SetRenderTarget(uint8(i))
		err = r.ref.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		utils.Check("set surface clear blendMode", err)

		r.SetDrawColor(0x00000050)
		err = r.ref.FillRect(&sdl.Rect{
			X: 0,
			Y: 0,
			W: r.renderTarget.width,
			H: r.renderTarget.height,
		})
		utils.Check("clear surface", err)
	}

	// draw primary surface
	r.SetRenderTarget(0)
	r.SetDrawColor(color)

	err = r.ref.SetDrawBlendMode(sdl.BLENDMODE_NONE)
	utils.Check("reset clear blendMode", err)

	err = r.ref.Clear()
	utils.Check("clear primary surface", err)
}

func (r *Renderer) EndEngineFrame() {
	var err error
	r.SetRenderTarget(0)

	// draw all surfaces
	for i := 0; i < 1; i++ { // todo: surfacesCount
		// copy to main texture
		err = r.ref.Copy(r.renderTarget.secondary[i], nil, nil)
		utils.Check("copy surface to main layer", err)
	}

	// draw primary
	err = r.ref.SetDrawBlendMode(sdl.BLENDMODE_NONE)
	utils.Check("set blendMode for present", err)
}

func (r *Renderer) UpdateGPU() {
	r.ref.Present()
}
