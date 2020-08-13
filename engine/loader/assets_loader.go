package loader

import (
	"fmt"
	"io/ioutil"

	"github.com/fe3dback/galaxy/generated"
	"gopkg.in/yaml.v2"
)

type AssetsLoader struct {
}

func NewAssetsLoader() *AssetsLoader {
	return &AssetsLoader{}
}

func (l AssetsLoader) LoadYaml(res generated.ResourcePath, data interface{}) {
	buffer, err := ioutil.ReadFile(string(res))
	if err != nil {
		panic(fmt.Sprintf("can`t open `%s`: %v", res, err))
	}

	err = yaml.Unmarshal(buffer, data)
	if err != nil {
		panic(fmt.Sprintf("can`t parse `%s`: %v", res, err))
	}
}
