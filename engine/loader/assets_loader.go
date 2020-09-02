package loader

import (
	"fmt"
	"io/ioutil"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/generated"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

type AssetsLoader struct {
	loaderSound engine.LoaderSound
}

func NewAssetsLoader(loaderSound engine.LoaderSound) *AssetsLoader {
	return &AssetsLoader{
		loaderSound: loaderSound,
	}
}

func (l *AssetsLoader) LoadYaml(res generated.ResourcePath, data interface{}) {
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

func (l *AssetsLoader) LoadSound(res generated.ResourcePath) {
	l.loaderSound.LoadSound(res)
}
