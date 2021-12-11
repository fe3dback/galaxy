package loader

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
)

type AssetsLoader struct {
	loaderSound galx.LoaderSound
}

func NewAssetsLoader(loaderSound galx.LoaderSound) *AssetsLoader {
	return &AssetsLoader{
		loaderSound: loaderSound,
	}
}

func (l *AssetsLoader) LoadYaml(res consts.AssetsPath, data interface{}) {
	buffer, err := ioutil.ReadFile(string(res))
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

func (l *AssetsLoader) LoadSound(res consts.AssetsPath) {
	l.loaderSound.LoadSound(res)
}
