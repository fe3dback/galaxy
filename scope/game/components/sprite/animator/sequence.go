package animator

import (
	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
)

type SequenceDirection uint8
type SequenceSliceDirection uint8
type Sequences map[string]*Sequence

const (
	SequenceDirectionForward  SequenceDirection = 0
	SequenceDirectionBackward SequenceDirection = 1
)

const (
	SequenceSliceDirectionToRightToBottom SequenceSliceDirection = 0
	SequenceSliceDirectionToBottomToRight SequenceSliceDirection = 1
)

const emptySequence string = "<empty>"

type (
	Sequence struct {
		// static
		resource  consts.AssetsPath
		slice     SequenceSlice
		frames    []*frame
		fps       int
		repeat    bool
		bounce    bool
		direction SequenceDirection
		offsetX   int
		offsetY   int

		// mutable data
		progress     float64
		currentFrame int
		firstFrame   int
		lastFrame    int
		progressMod  float64
		finished     bool
	}

	SequenceSetupFunc func(*Sequence)

	SequenceSlice struct {
		FrameWidth  int
		FrameHeight int
		FirstX      int
		FirstY      int
		FramesCount int
		SliceType   SequenceSliceDirection
	}
)

func NewSequence(texId consts.AssetsPath, slice SequenceSlice, initializers ...SequenceSetupFunc) *Sequence {
	seq := &Sequence{
		resource:     texId,
		slice:        slice,
		frames:       make([]*frame, 0),
		fps:          0,
		repeat:       true,
		bounce:       true,
		direction:    SequenceDirectionForward,
		offsetX:      0,
		offsetY:      0,
		currentFrame: 0,
		firstFrame:   0,
		lastFrame:    0,
		progressMod:  0,
		finished:     false,
	}

	for _, init := range initializers {
		init(seq)
	}

	return seq
}

func (seq *Sequence) initialize(renderer galx.Renderer) {
	texInfo := renderer.TextureQuery(seq.resource)
	frames := sliceFrames(seq.slice, texInfo.Width, texInfo.Height)

	seq.setFrames(frames)
}

func (seq *Sequence) clearFrames() {
	seq.setFrames(seq.frames)
}

func (seq *Sequence) setFrames(frames []*frame) {
	seq.finished = false
	seq.currentFrame = 0
	seq.firstFrame = 0
	seq.lastFrame = len(frames) - 1
	seq.frames = frames
	seq.progressMod = float64(seq.fps) / float64(seq.lastFrame)

	if seq.fps == 0 {
		seq.fps = len(frames)
	}
}

//goland:noinspection GoUnusedExportedFunction
func WithFps(fps int) SequenceSetupFunc {
	return func(seq *Sequence) {
		seq.fps = fps
	}
}

//goland:noinspection GoUnusedExportedFunction
func WithCustomPlayback(repeat bool, bounce bool, direction SequenceDirection) SequenceSetupFunc {
	return func(seq *Sequence) {
		seq.repeat = repeat
		seq.bounce = bounce
		seq.direction = direction
	}
}

//goland:noinspection GoUnusedExportedFunction
func WithOffset(x, y int) SequenceSetupFunc {
	return func(seq *Sequence) {
		seq.offsetX = x
		seq.offsetY = y
	}
}
