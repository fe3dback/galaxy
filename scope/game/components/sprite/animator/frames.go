package animator

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
)

type sequenceWalkFunc func(frame *frame)

type (
	frame struct {
		index      int
		x, y, w, h int
		column     int
		row        int
		rect       *galx.Rect
	}
)

func (f *frame) TextureRect() galx.Rect {
	if f.rect == nil {
		r := galx.Rect{
			TL: galx.Vec{
				X: float64(f.x),
				Y: float64(f.y),
			},
			BR: galx.Vec{
				X: float64(f.x + f.w),
				Y: float64(f.y + f.h),
			},
		}
		f.rect = &r
	}

	return *f.rect
}

func sliceFrames(slice SequenceSlice, w, h int) []*frame {
	assertSliceValid(slice, w, h)

	frames := make([]*frame, 0)
	walk(slice, w, h, func(frame *frame) {
		frames = append(frames, frame)
	})

	return frames
}

func assertSliceValid(slice SequenceSlice, w, h int) {
	if slice.FramesCount > 128 {
		// this is soft limit, can by changed to any number
		panic("Not supported animations with 128 or more frames")
	}

	if slice.FramesCount <= 0 {
		panic("Not supported animations without frames")
	}

	assetPointInside(slice.FirstX, slice.FirstY, w, h)
}

func walk(slice SequenceSlice, texW, texH int, fn sequenceWalkFunc) {
	ind := 0
	left := slice.FramesCount

	startX := slice.FirstX
	startY := slice.FirstY
	row := 0
	column := 0

	for {
		if left <= 0 {
			break
		}

		// ----------------------
		// Get current points
		// ----------------------

		curLX := startX + (column * slice.FrameWidth)
		curTY := startY + (row * slice.FrameHeight)

		curRX := curLX + slice.FrameWidth
		curBY := curTY + slice.FrameHeight

		// ----------------------
		// Assert is valid
		// ----------------------

		assetPointInside(curLX, curTY, texW, texH)
		assetPointInside(curRX, curBY, texW, texH)

		// ----------------------
		// Send
		// ----------------------

		fn(&frame{
			index:  ind,
			x:      curLX,
			y:      curTY,
			w:      slice.FrameWidth,
			h:      slice.FrameHeight,
			column: column,
			row:    row,
		})

		// ----------------------
		// Check next point possible
		// ----------------------

		switch slice.SliceType {
		case SequenceSliceDirectionToRightToBottom:
			nextLX := curLX + slice.FrameWidth
			if nextLX >= texW {
				// next image in first col, next row
				startX = 0
				column = 0
				row++
			} else {
				column++
			}
		case SequenceSliceDirectionToBottomToRight:
			nextTY := curTY + slice.FrameHeight
			if nextTY >= texH {
				// next image in first col, next row
				startY = 0
				row = 0
				column++
			} else {
				row++
			}
		default:
			panic(fmt.Sprintf("unknown slice type %d", slice.SliceType))
		}

		ind++
		left--
	}
}

func assetPointInside(x, y, w, h int) {
	if x < 0 || y < 0 {
		panic(fmt.Sprintf("invalid seq X (%d) or Y (%d) (<= 0)", x, y))
	}

	if x > w || y > h {
		panic(fmt.Sprintf("invalid seq X (%d) or Y (%d) (>= width or height)", x, y))
	}
}
