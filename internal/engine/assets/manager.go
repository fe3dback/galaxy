package assets

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/internal/engine/lib/sound"
)

type Manager struct {
	soundManager *sound.Manager
}

func NewAssetsManager(soundManager *sound.Manager) *Manager {
	return &Manager{
		soundManager: soundManager,
	}
}

func (l *Manager) LoadYaml(res consts.AssetsPath, data interface{}) {
	buffer, err := ioutil.ReadFile(res)
	if err != nil {
		panic(fmt.Sprintf("can`t open `%s`: %v", res, err))
	}

	err = yaml.Unmarshal(buffer, data)
	if err != nil {
		panic(fmt.Sprintf("can`t parse `%s`: %v", res, err))
	}

	if err := validator.Validate(data); err != nil {
		panic(fmt.Sprintf("invalid spec '%s': %v", res, err))
	}
}

func (l *Manager) LoadXML(res consts.AssetsPath, data interface{}) {
	buffer, err := ioutil.ReadFile(res)
	if err != nil {
		panic(fmt.Sprintf("can`t open `%s`: %v", res, err))
	}

	err = xml.Unmarshal(buffer, data)
	if err != nil {
		panic(fmt.Sprintf("can`t parse `%s`: %v", res, err))
	}
}

func (l *Manager) LoadFile(res consts.AssetsPath) []byte {
	buffer, err := ioutil.ReadFile(res)
	if err != nil {
		panic(fmt.Sprintf("can`t open `%s`: %v", res, err))
	}

	return buffer
}

func (l *Manager) SaveFile(res consts.AssetsPath, data []byte) {
	err := ioutil.WriteFile(res, data, 0755)
	if err != nil {
		panic(fmt.Sprintf("can`t save `%s`: %v", res, err))
	}
}

func (l *Manager) CopyFile(src, dest consts.AssetsPath) {
	l.SaveFile(dest, l.LoadFile(src))
}

func (l *Manager) LoadSound(res consts.AssetsPath) {
	l.soundManager.LoadSound(res)
}
