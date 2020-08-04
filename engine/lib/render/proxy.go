package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/generated"
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

// Proxy user for engine input normalization, (for example convert rect to screen space)
// also only this renderer should check clip rect's

func (r *Renderer) transRect(rect engine.Rect) sdl.Rect {
	return Rect{
		X: int32(r.screenX(rect.Min.X)),
		Y: int32(r.screenY(rect.Min.Y)),
		W: int32(rect.Max.X),
		H: int32(rect.Max.Y),
	}
}

func (r *Renderer) transRectPtr(rect engine.Rect) *sdl.Rect {
	rRect := r.transRect(rect)
	return &rRect
}

func (r *Renderer) transPoint(point engine.Vec) sdl.Point {
	return Point{
		X: int32(r.screenX(point.X)),
		Y: int32(r.screenY(point.Y)),
	}
}

func (r *Renderer) transPointPtr(point engine.Vec) *sdl.Point {
	rPoint := r.transPoint(point)
	return &rPoint
}

func (r *Renderer) transLine(line engine.Line) []sdl.Point {
	return []sdl.Point{
		r.transPoint(line.A),
		r.transPoint(line.B),

		// close lines back will fix render glitches
		r.transPoint(line.B),
		r.transPoint(line.A),
	}
}

func transformColor(color engine.Color) sdl.Color {
	r, g, b, a := color.Split()
	return sdl.Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

// ==================================================
// Primitive
// ==================================================

func (r *Renderer) DrawSquare(color engine.Color, rect engine.Rect) {
	rect = rect.Screen()
	if !r.isRectInsideCamera(rect) {
		return
	}

	r.internalDrawSquare(color, r.transRect(rect))
}

func (r *Renderer) DrawSquareEx(color engine.Color, angle engine.Angle, rect engine.Rect) {
	rect = rect.Screen()
	if !r.isRectInsideCamera(rect) {
		return
	}

	center := engine.Vec{
		X: rect.Min.X + rect.Max.X/2,
		Y: rect.Min.Y + rect.Max.Y/2,
	}

	vertices := [4]engine.Vec{
		engine.Vec{
			X: rect.Min.X,
			Y: rect.Min.Y,
		}.RotateAround(center, angle),

		engine.Vec{
			X: rect.Min.X + rect.Max.X,
			Y: rect.Min.Y,
		}.RotateAround(center, angle),

		engine.Vec{
			X: rect.Min.X + rect.Max.X,
			Y: rect.Min.Y + rect.Max.Y,
		}.RotateAround(center, angle),

		engine.Vec{
			X: rect.Min.X,
			Y: rect.Min.Y + rect.Max.Y,
		}.RotateAround(center, angle),
	}

	r.internalDrawLines(color, []sdl.Point{
		r.transPoint(vertices[0]),
		r.transPoint(vertices[1]),
		r.transPoint(vertices[2]),
		r.transPoint(vertices[3]),
		r.transPoint(vertices[0]),
	})
}

func (r *Renderer) DrawLine(color engine.Color, line engine.Line) {
	if !r.isLineInsideCamera(line) {
		return
	}

	r.internalDrawLines(color, r.transLine(line))
}

func (r *Renderer) DrawPoint(color engine.Color, vec engine.Vec) {
	if !r.isPointInsideCamera(vec) {
		return
	}

	r.internalDrawPoint(color, r.transPoint(vec))
}

func (r *Renderer) DrawVector(color engine.Color, dist float64, vec engine.Vec, angle engine.Angle) {
	target := vec.PolarOffset(dist, angle)
	line := engine.Line{
		A: vec,
		B: target,
	}

	if !r.isLineInsideCamera(line) {
		return
	}

	// draw
	counterDeg := angle.Add(engine.NewAngle(180))
	arrowLeft := engine.Line{
		A: target,
		B: target.PolarOffset(6, counterDeg.Add(engine.NewAngle(-30))),
	}
	arrowRight := engine.Line{
		A: target,
		B: target.PolarOffset(6, counterDeg.Add(engine.NewAngle(+30))),
	}

	r.internalDrawLines(color, r.transLine(line))
	r.internalDrawLines(color, r.transLine(arrowLeft))
	r.internalDrawLines(color, r.transLine(arrowRight))
}

func (r *Renderer) DrawCrossLines(color engine.Color, size int, vec engine.Vec) {
	if !r.isPointInsideCamera(vec) {
		return
	}

	r.internalDrawCrossLines(color, int32(size), r.transPoint(vec))
}

// ==================================================
// Sprite
// ==================================================

func (r *Renderer) DrawSprite(res generated.ResourcePath, vec engine.Vec) {
	r.draw(res, engine.Rect{}, engine.Rect{Min: vec}, engine.NewAngle(0))
}

func (r *Renderer) DrawSpriteAngle(res generated.ResourcePath, vec engine.Vec, angle engine.Angle) {
	r.draw(res, engine.Rect{}, engine.Rect{Min: vec}, angle)
}

func (r *Renderer) DrawSpriteEx(res generated.ResourcePath, src, dest engine.Rect, angle engine.Angle) {
	r.draw(res, src, dest, angle)
}

func (r *Renderer) draw(res generated.ResourcePath, src, dest engine.Rect, angle engine.Angle) {
	defer utils.CheckPanic(fmt.Sprintf("draw sprite `%s`", res))

	texture := r.getTexture(res)

	if dest.Max.X == 0 {
		dest.Max.X = float64(texture.Width)
	}
	if dest.Max.Y == 0 {
		dest.Max.Y = float64(texture.Height)
	}

	// check is visible
	if !r.isRectInsideCamera(dest) {
		return
	}

	if src.Max.X == 0 {
		src.Max.X = dest.Max.X
	}
	if src.Max.Y == 0 {
		src.Max.Y = dest.Max.Y
	}

	r.internalDrawTexture(
		texture.Tex,
		sdl.Rect{
			X: int32(src.Min.X),
			Y: int32(src.Min.Y),
			W: int32(src.Max.X),
			H: int32(src.Max.Y),
		},
		r.transRect(dest),
		angle,
	)
}

// ==================================================
// Text
// ==================================================

func (r *Renderer) DrawText(fontId generated.ResourcePath, color engine.Color, text string, vec engine.Vec) {
	if !r.isRectInsideCamera(engine.Rect{
		Min: vec,
		Max: engine.Vec{
			X: avgTextWidthOptRender,  // todo magic numbers
			Y: avgTextHeightOptRender, // todo magic numbers
		},
	}) {
		return
	}

	r.internalDrawText(fontId, color, text, r.transPoint(vec))
}

// ==================================================
// Other
// ==================================================
