package car

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fe3dback/galaxy/game/gm"

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
		wheels  specWheels
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
	}

	specWheels struct {
		radius gm.Meter
		wheels []specWheel
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

	specPos struct {
		x int
		y int
	}

	specSize struct {
		width  int
		height int
	}
)

func (phys *Physics) assembleSpec(yaml YamlSpec) spec {
	return spec{
		mass:    phys.assembleSpecMass(yaml),
		model:   phys.assembleSpecModel(yaml),
		engine:  phys.assembleSpecEngine(yaml),
		wheels:  phys.assembleSpecWheels(yaml),
		weights: phys.assembleSpecWeights(yaml),
	}
}

func (phys *Physics) assembleSpecMass(yaml YamlSpec) specMass {
	return specMass{
		mass: yaml.Params.Mass,
	}
}

func (phys *Physics) assembleSpecModel(yaml YamlSpec) specModel {
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

func (phys *Physics) assembleSpecEngine(_ YamlSpec) specEngine {
	return specEngine{}
}

func (phys *Physics) assembleSpecWheels(yaml YamlSpec) specWheels {
	wheels := make([]specWheel, 0)

	offsetY := yaml.Wheels.Offset

	for id, offsetX := range yaml.Wheels.Axis {
		wheels = append(wheels, specWheel{
			size: specSize{
				width:  yaml.Wheels.Size.Width,
				height: yaml.Wheels.Size.Height,
			},
			axis: axis(id),
			posRelative: specPos{
				x: int(offsetX),
				y: -offsetY,
			},
			posAbsolute: specPos{
				x: yaml.Center.X + int(offsetX),
				y: yaml.Center.Y - offsetY,
			},
		})
		wheels = append(wheels, specWheel{
			size: specSize{
				width:  yaml.Wheels.Size.Width,
				height: yaml.Wheels.Size.Height,
			},
			axis: axis(id),
			posRelative: specPos{
				x: int(offsetX),
				y: offsetY,
			},
			posAbsolute: specPos{
				x: yaml.Center.X + int(offsetX),
				y: yaml.Center.Y + offsetY,
			},
		})
	}

	return specWheels{
		radius: yaml.Wheels.Radius,
		wheels: wheels,
	}
}

func (phys *Physics) assembleSpecWeights(yaml YamlSpec) []specWeight {
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

func (phys *Physics) assembleSpecWeightsPoints(yaml YamlSpec, points []string, x int) ([]specWeight, error) {
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

func (phys *Physics) assembleSpecWeight(yaml YamlSpec, raw string) (int, float32, error) {
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
