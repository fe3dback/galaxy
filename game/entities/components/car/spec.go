package car

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fe3dback/galaxy/utils"
)

type (
	kg     = int
	KMpS   = int
	axis   = string
	weight = float32

	spec struct {
		mass    specMass
		model   specModel
		engine  specEngine
		wheels  []specWheel
		weights []specWeight
	}

	specMass struct {
		mass kg
	}

	specModel struct {
		center specPos
		size   specSize
	}

	specEngine struct {
		gears []specGear
	}

	specWheel struct {
		size        specSize
		axis        axis
		posRelative specPos
		posAbsolute specPos
	}

	specWeight struct {
		weight      weight
		posRelative specPos
		posAbsolute specPos
	}

	specGear struct {
		maxSpeed KMpS
	}

	specPos struct {
		x int
		y int
	}

	specSize struct {
		width  int
		height int
	}
)

func (phys *Physics) assembleSpec() {
	yaml := phys.parse(phys.resource)
	phys.spec = spec{
		mass:    phys.assembleSpecMass(yaml),
		model:   phys.assembleSpecModel(yaml),
		engine:  phys.assembleSpecEngine(yaml),
		wheels:  phys.assembleSpecWheels(yaml),
		weights: phys.assembleSpecWeights(yaml),
	}
}

func (phys *Physics) assembleSpecMass(yaml yamlSpec) specMass {
	return specMass{
		mass: yaml.Params.Mass,
	}
}

func (phys *Physics) assembleSpecModel(yaml yamlSpec) specModel {
	return specModel{
		center: specPos{
			x: yaml.Center.X,
			y: yaml.Center.Y,
		},
		size: specSize{
			width:  yaml.Size.Width,
			height: yaml.Size.Height,
		},
	}
}

func (phys *Physics) assembleSpecEngine(yaml yamlSpec) specEngine {
	return specEngine{
		gears: phys.assembleSpecEngineGears(yaml),
	}
}

func (phys *Physics) assembleSpecEngineGears(yaml yamlSpec) []specGear {
	// todo: gear speed calculation based on params, rpm, horse power, mass..
	return []specGear{
		{
			maxSpeed: KMpS(float32(yaml.Params.MaxSpeed) * 0.5),
		},
		{
			maxSpeed: KMpS(float32(yaml.Params.MaxSpeed) * 0.75),
		},
		{
			maxSpeed: KMpS(float32(yaml.Params.MaxSpeed) * 1),
		},
	}

}

func (phys *Physics) assembleSpecWheels(yaml yamlSpec) []specWheel {
	wheels := make([]specWheel, 0)

	carWidth := yaml.Size.Width
	offsetX := carWidth / 2

	for id, offsetY := range yaml.Wheels.Axis {
		wheels = append(wheels, specWheel{
			size: specSize{
				width:  yaml.Wheels.Size.Width,
				height: yaml.Wheels.Size.Height,
			},
			axis: axis(id),
			posRelative: specPos{
				x: -offsetX + (yaml.Wheels.Size.Width / 2),
				y: int(offsetY),
			},
			posAbsolute: specPos{
				x: yaml.Center.X - offsetX + (yaml.Wheels.Size.Width / 2),
				y: yaml.Center.Y + int(offsetY),
			},
		})
		wheels = append(wheels, specWheel{
			size: specSize{
				width:  yaml.Wheels.Size.Width,
				height: yaml.Wheels.Size.Height,
			},
			axis: axis(id),
			posRelative: specPos{
				x: offsetX - (yaml.Wheels.Size.Width / 2),
				y: int(offsetY),
			},
			posAbsolute: specPos{
				x: yaml.Center.X + offsetX - (yaml.Wheels.Size.Width / 2),
				y: yaml.Center.Y + int(offsetY),
			},
		})
	}

	return wheels
}

func (phys *Physics) assembleSpecWeights(yaml yamlSpec) []specWeight {
	weights := make([]specWeight, 0)

	for _, line := range yaml.Weights {
		assembled, err := phys.assembleSpecWeightsPoints(yaml, line.Points, line.X)
		utils.Check("assemble left weights", err)
		weights = append(weights, assembled...)

		if line.Mirror {
			assembled, err := phys.assembleSpecWeightsPoints(yaml, line.Points, -line.X)
			utils.Check("assemble right weights", err)
			weights = append(weights, assembled...)
		}
	}

	return weights
}

func (phys *Physics) assembleSpecWeightsPoints(yaml yamlSpec, points []string, x int) ([]specWeight, error) {
	xLeft := 0 - (yaml.Size.Width / 2)
	xRight := 0 + (yaml.Size.Width / 2)

	if !(xLeft <= x && x <= xRight) {
		return nil, fmt.Errorf("point x `%d` outside car X (%d .. %d)", x, xLeft, xRight)
	}

	weights := make([]specWeight, 0)

	for _, rawPoint := range points {
		y, weight, err := phys.assembleSpecWeight(yaml, rawPoint)

		if err != nil {
			return nil, fmt.Errorf("assemble weight line X `%d` failed: %v", x, err)
		}

		weights = append(weights, specWeight{
			weight: weight,
			posRelative: specPos{
				x: x,
				y: y,
			},
			posAbsolute: specPos{
				x: yaml.Center.X + x,
				y: yaml.Center.Y + y,
			},
		})
	}

	return weights, nil
}

func (phys *Physics) assembleSpecWeight(yaml yamlSpec, raw string) (int, float32, error) {
	parts := strings.Split(raw, ",")

	rawY := strings.Trim(parts[0], " ")
	rawWeight := strings.Trim(parts[1], " ")

	y, err := strconv.Atoi(rawY)
	if err != nil {
		return 0, 0, fmt.Errorf("parse y: %v", err)
	}

	weight, err := strconv.ParseFloat(rawWeight, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("parse weight: %v", err)
	}

	xTop := 0 - (yaml.Size.Height / 2)
	xBottom := 0 + (yaml.Size.Height / 2)

	if !(xTop <= y && y <= xBottom) {
		return 0, 0, fmt.Errorf("point y `%d` outside car Y (%d .. %d)", y, xTop, xBottom)
	}

	if weight > 1.0 {
		return 0, 0, fmt.Errorf("weight `%f` can`t be more than 1", weight)
	}

	if weight < 0 {
		return 0, 0, fmt.Errorf("weight `%f` can`t be less than 0", weight)
	}

	return y, float32(weight), nil
}
