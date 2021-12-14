package loader

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/internal/engine/lib/sound"
)

type AssetsLoader struct {
	soundManager *sound.Manager
}

func NewAssetsLoader(soundManager *sound.Manager) *AssetsLoader {
	return &AssetsLoader{
		soundManager: soundManager,
	}
}

func (l *AssetsLoader) LoadYaml(res consts.AssetsPath, data interface{}) {
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

func (l *AssetsLoader) LoadXML(res consts.AssetsPath, data interface{}) {
	buffer, err := ioutil.ReadFile(res)
	if err != nil {
		panic(fmt.Sprintf("can`t open `%s`: %v", res, err))
	}

	err = xml.Unmarshal(buffer, data)
	if err != nil {
		panic(fmt.Sprintf("can`t parse `%s`: %v", res, err))
	}
}

func (l *AssetsLoader) LoadFile(res consts.AssetsPath) []byte {
	buffer, err := ioutil.ReadFile(res)
	if err != nil {
		panic(fmt.Sprintf("can`t open `%s`: %v", res, err))
	}

	return buffer
}

func (l *AssetsLoader) LoadSound(res consts.AssetsPath) {
	l.soundManager.LoadSound(res)
}
