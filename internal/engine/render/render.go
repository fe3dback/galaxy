package render

import (
	"math"
	"time"

	"github.com/fe3dback/govgl"

	"github.com/fe3dback/galaxy/galx"
)

type Render struct {
	renderer *govgl.Render

	renderMode galx.RenderMode
	camera     *Camera
}

func NewRender(renderer *govgl.Render, camera *Camera) *Render {
	return &Render{
		renderer: renderer,
		camera:   camera,
	}
}

func (r *Render) Gizmos() galx.Gizmos {
	// TODO implement me
	panic("implement me")
}

func (r *Render) SetRenderTarget(id galx.RenderTarget) {
	// TODO implement me
	panic("implement me")
}

func (r *Render) StartEngineFrame() {
	r.renderMode = galx.RenderModeWorld
	r.renderer.FrameStart()
}

func (r *Render) EndEngineFrame() {
	r.renderer.FrameEnd()
}

// DrawTemporary todo: remove tmp
func (r *Render) DrawTemporary() {
	// r.renderer.DrawTmpTriangle()

	const rSizeX = 32
	const rSizeY = 48
	const rOffset = 8
	const rStartX = rSizeX
	const rEndX = 1280 - rSizeX - rOffset
	const rStartY = rSizeY
	const rEndY = 720 - rSizeY - rOffset

	dOffsetX := math.Sin(float64(time.Now().UnixMilli())*0.01) * rOffset
	dOffsetY := math.Cos(float64(time.Now().UnixMilli())*0.01) * rOffset

	gridX := 0
	gridY := 0

	for x := rStartX; x < rEndX; x += rSizeX + rOffset {
		gridX++
		for y := rStartY; y < rEndY; y += rSizeY + rOffset {
			gridY++
			r.DrawSquare(galx.ColorCyan, galx.Rect{
				TL: galx.Vec2d{
					X: float64(x) + ((dOffsetX * float64(gridX)) * 0.005),
					Y: float64(y) + ((dOffsetY * float64(gridY)) * 0.005),
				},
				BR: galx.Vec2d{
					X: float64(x+rSizeX) + ((dOffsetX * float64(gridX)) * 0.005),
					Y: float64(y+rSizeY) + ((dOffsetY * float64(gridY)) * 0.005),
				},
			})
		}
	}

	// r.renderer.DrawTmpTriangle()
}

func (r *Render) StartGUIFrame(color galx.Color) {
	r.renderMode = galx.RenderModeUI
	// TODO implement me
	panic("implement me")
}

func (r *Render) EndGUIFrame() {
	// do nothing here
}

func (r *Render) WaitGPUOperationsBeforeQuit() {
	r.renderer.WaitGPU()
}
