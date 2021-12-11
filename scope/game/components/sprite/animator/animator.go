package animator

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
)

type (
	Animator struct {
		// static
		entity galx.GameObject

		// mutable
		sequences        Sequences
		activeSequence   *Sequence
		activeSequenceId string
		paused           bool
		initialized      bool
	}
)

// todo: refactor to new components system

func NewAnimator(entity galx.GameObject) *Animator {
	return &Animator{
		entity:           entity,
		sequences:        make(Sequences),
		activeSequence:   nil,
		activeSequenceId: emptySequence,
		paused:           false,
		initialized:      false,
	}
}

func (anim *Animator) AddSequence(name string, seq *Sequence) {
	if _, ok := anim.sequences[name]; ok {
		panic(fmt.Sprintf("animator sequence `%s` already added", name))
	}

	anim.sequences[name] = seq

	// play first sequence
	if anim.activeSequenceId == emptySequence {
		anim.PlaySequence(name)
	}
}

func (anim *Animator) PlaySequence(name string) {
	seq, exist := anim.sequences[name]
	if !exist {
		panic(fmt.Sprintf("animator sequence `%s` not exist", name))
	}

	// set active
	anim.activeSequenceId = name
	anim.activeSequence = seq

	// reset
	seq.clearFrames()
}

func (anim *Animator) Play() {
	anim.paused = false
}

func (anim *Animator) Pause() {
	anim.paused = true
}

func (anim *Animator) initialize(renderer galx.Renderer) {
	if anim.initialized {
		return
	}

	for _, sequence := range anim.sequences {
		sequence.initialize(renderer)
	}

	anim.initialized = true
}
