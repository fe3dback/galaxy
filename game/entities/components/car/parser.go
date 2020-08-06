package car

import (
	"fmt"
	"io/ioutil"

	"github.com/fe3dback/galaxy/generated"
	"gopkg.in/yaml.v2"
)

type (
	yamlSpec struct {
		Center  yamlCenter       `yaml:"center"`
		Size    yamlSize         `yaml:"size"`
		Wheels  yamlWheels       `yaml:"wheels"`
		Weights []yamlWeightLine `yaml:"weights"`
		Params  yamlParams       `yaml:"params"`
	}

	yamlParams struct {
		Mass int `yaml:"mass"`
	}

	yamlCenter struct {
		X int `yaml:"x"`
		Y int `yaml:"y"`
	}

	yamlSize struct {
		Width  int `yaml:"width"`
		Height int `yaml:"height"`
	}

	yamlAxisId      string
	yamlAxisOffsetY int
	yamlWheels      struct {
		Size   yamlSize                       `yaml:"size"`
		Offset int                            `yaml:"offset"`
		Axis   map[yamlAxisId]yamlAxisOffsetY `yaml:"axis"`
	}

	yamlWeightLine struct {
		X      int      `yaml:"x"`
		Points []string `yaml:"points"`
		Mirror bool     `yaml:"mirror"`
	}
)

func (phys *Physics) parse(res generated.ResourcePath) yamlSpec {
	spec := yamlSpec{}

	buffer, err := ioutil.ReadFile(string(res))
	if err != nil {
		panic(fmt.Sprintf("can`t open car physics spec `%s`: %v", res, err))
	}

	err = yaml.Unmarshal(buffer, &spec)
	if err != nil {
		panic(fmt.Sprintf("can`t parse car physics spec `%s`: %v", res, err))
	}

	return spec
}
