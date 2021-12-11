package sound

import (
	"fmt"
	"path"
	"time"

	"github.com/veandco/go-sdl2/mix"

	"github.com/fe3dback/galaxy/consts"
	utils2 "github.com/fe3dback/galaxy/internal/utils"
)

const maxChannels = 64 // how much sounds can play at same time

type (
	sounds    map[consts.AssetsPath]*mix.Chunk
	durations map[consts.AssetsPath]time.Duration

	Manager struct {
		closer    *utils2.Closer
		loaded    sounds
		durations durations
		channels  [maxChannels]bool
	}
)

func NewManager(closer *utils2.Closer) *Manager {
	err := mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE)
	utils2.Check("mix open audio", err)
	closer.EnqueueClose(func() error {
		mix.CloseAudio()
		return nil
	})

	return &Manager{
		closer:    closer,
		loaded:    make(sounds),
		durations: make(durations),
		channels:  [maxChannels]bool{},
	}
}

func (m *Manager) Play(res consts.AssetsPath) {
	channel, ok := m.acquireChannel()
	if !ok {
		return
	}

	m.playEx(res, channel)

	// unlock channel
	time.AfterFunc(m.durations[res], func() {
		m.freeChannel(channel)
	})
}

func (m *Manager) LoadSound(res consts.AssetsPath) {
	if _, ok := m.loaded[res]; ok {
		return
	}

	var sound *mix.Chunk

	ext := path.Ext(string(res))
	switch ext {
	case ".wav":
		sound = m.loadWav(res)
	default:
		panic(fmt.Sprintf("Failed to load sound '%s', unknown format '%s'", res, ext))
	}

	// store to memory
	m.loaded[res] = sound
	m.durations[res] = time.Millisecond * time.Duration(sound.LengthInMs())
}

func (m *Manager) acquireChannel() (int, bool) {
	channelId := m.findFreeChannelId()
	if channelId == nil {
		// max sounds is already played
		return 0, false
	}

	m.lockChannel(*channelId)
	return *channelId, true
}

func (m *Manager) findFreeChannelId() *int {
	for i, locked := range m.channels {
		if locked {
			continue
		}

		return &i
	}

	return nil
}

func (m *Manager) lockChannel(id int) {
	m.channels[id] = true
}

func (m *Manager) freeChannel(id int) {
	m.channels[id] = false
}

func (m *Manager) playEx(res consts.AssetsPath, channel int) {
	chunk, ok := m.loaded[res]
	if !ok {
		panic(fmt.Sprintf("Failed to play '%s', sound not loaded yet", res))
	}

	_, err := chunk.Play(channel, 0)
	utils2.Check(fmt.Sprintf("play '%s'", res), err)
}

func (m *Manager) loadWav(res consts.AssetsPath) *mix.Chunk {
	chunk, err := mix.LoadWAV(string(res))
	utils2.Check(fmt.Sprintf("load wav file '%s'", res), err)

	return chunk
}
