package render

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/utils"
)

// Proxy user for engine input normalization, (for example convert rect to screen space)
// also only this renderer should check clip rect's

func (r *Renderer) transRect(rect galx.Rect) sdl.Rect {
	rect = rect.MaxToSize()
	return Rect{
		X: int32(r.screenX(rect.Min.X)),
		Y: int32(r.screenY(rect.Min.Y)),
		W: int32(rect.Max.X),
		H: int32(rect.Max.Y),
	}
}

func (r *Renderer) transCircle(circle galx.Circle) (sdl.Point, float64) {
	return r.transPoint(circle.Pos), circle.Radius
}

func (r *Renderer) transRectPtr(rect galx.Rect) *sdl.Rect {
	rRect := r.transRect(rect)
	return &rRect
}

func (r *Renderer) transPoint(point galx.Vec) sdl.Point {
	return Point{
		X: int32(r.screenX(point.X)),
		Y: int32(r.screenY(point.Y)),
	}
}

func (r *Renderer) transLine(line galx.Line) []sdl.Point {
	return []sdl.Point{
		r.transPoint(line.A),
		r.transPoint(line.B),

		// close lines back will fix render glitches
		r.transPoint(line.B),
		r.transPoint(line.A),
	}
}

func transformColor(color galx.Color) sdl.Color {
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

func (r *Renderer) DrawSquare(color galx.Color, rect galx.Rect) {
	rect = rect.Screen()
	if !r.isRectInsideCamera(rect) {
		return
	}

	r.internalDrawSquare(color, r.transRect(rect))
}

func (r *Renderer) DrawSquareFilled(color galx.Color, rect galx.Rect) {
	rect = rect.Screen()
	if !r.isRectInsideCamera(rect) {
		return
	}

	r.internalDrawSquareFilled(color, r.transRect(rect))
}

func (r *Renderer) DrawCircle(color galx.Color, circle galx.Circle) {
	rect := circle.BoundingBox().Screen()
	if !r.isRectInsideCamera(rect) {
		return
	}

	pos, radius := r.transCircle(circle)
	r.internalDrawCircle(color, pos, radius)
}

func (r *Renderer) DrawSquareEx(color galx.Color, angle galx.Angle, rect galx.Rect) {
	rect = rect.Screen()
	if !r.isRectInsideCamera(rect) {
		return
	}

	center := galx.Vec{
		X: rect.Min.X + rect.Max.X/2,
		Y: rect.Min.Y + rect.Max.Y/2,
	}

	vertices := [4]galx.Vec{
		galx.Vec{
			X: rect.Min.X,
			Y: rect.Min.Y,
		}.RotateAround(center, angle),

		galx.Vec{
			X: rect.Min.X + rect.Max.X,
			Y: rect.Min.Y,
		}.RotateAround(center, angle),

		galx.Vec{
			X: rect.Min.X + rect.Max.X,
			Y: rect.Min.Y + rect.Max.Y,
		}.RotateAround(center, angle),

		galx.Vec{
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

func (r *Renderer) DrawLine(color galx.Color, line galx.Line) {
	if !r.isLineInsideCamera(line) {
		return
	}

	r.internalDrawLines(color, r.transLine(line))
}

func (r *Renderer) DrawPoint(color galx.Color, vec galx.Vec) {
	if !r.isPointInsideCamera(vec) {
		return
	}

	r.internalDrawPoint(color, r.transPoint(vec))
}

func (r *Renderer) DrawVector(color galx.Color, dist float64, vec galx.Vec, angle galx.Angle) {
	target := vec.PolarOffset(dist, angle)
	line := galx.Line{
		A: vec,
		B: target,
	}

	if !r.isLineInsideCamera(line) {
		return
	}

	// draw
	counterDeg := angle.Add(galx.NewAngle(180))
	arrowLeft := galx.Line{
		A: target,
		B: target.PolarOffset(6, counterDeg.Add(galx.NewAngle(-30))),
	}
	arrowRight := galx.Line{
		A: target,
		B: target.PolarOffset(6, counterDeg.Add(galx.NewAngle(+30))),
	}

	r.internalDrawLines(color, r.transLine(line))
	r.internalDrawLines(color, r.transLine(arrowLeft))
	r.internalDrawLines(color, r.transLine(arrowRight))
}

func (r *Renderer) DrawCrossLines(color galx.Color, size int, vec galx.Vec) {
	if !r.isPointInsideCamera(vec) {
		return
	}

	r.internalDrawCrossLines(color, int32(size), r.transPoint(vec))
}

// ==================================================
// Sprite
// ==================================================

func (r *Renderer) DrawSprite(res consts.AssetsPath, vec galx.Vec) {
	r.draw(res, galx.Rect{}, galx.Rect{Min: vec}, galx.NewAngle(0))
}

func (r *Renderer) DrawSpriteAngle(res consts.AssetsPath, vec galx.Vec, angle galx.Angle) {
	r.draw(res, galx.Rect{}, galx.Rect{Min: vec}, angle)
}

func (r *Renderer) DrawSpriteEx(res consts.AssetsPath, src, dest galx.Rect, angle galx.Angle) {
	r.draw(res, src, dest, angle)
}

func (r *Renderer) draw(res consts.AssetsPath, src, dest galx.Rect, angle galx.Angle) {
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
		src.Max.X = float64(texture.Width)
	}
	if src.Max.Y == 0 {
		src.Max.Y = float64(texture.Height)
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

func (r *Renderer) DrawText(fontId consts.AssetsPath, color galx.Color, text string, vec galx.Vec) {
	if !r.isRectInsideCamera(galx.Rect{
		Min: vec,
		Max: galx.Vec{
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
