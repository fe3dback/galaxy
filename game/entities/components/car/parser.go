package car

type (
	YamlSpec struct {
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
		Radius float64                        `yaml:"radius"`
		Axis   map[yamlAxisId]yamlAxisOffsetY `yaml:"axis"`
	}

	yamlWeightLine struct {
		X      int      `yaml:"x"`
		Points []string `yaml:"points"`
		Mirror bool     `yaml:"mirror"`
	}
)
