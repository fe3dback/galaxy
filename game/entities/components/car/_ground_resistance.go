package car

const resistanceMod = 0.1

type groundResistance = float64

const (
	groundResistanceAsphaltGood groundResistance = 0.1 * resistanceMod
	groundResistanceAsphalt     groundResistance = 0.15 * resistanceMod
	groundResistanceGravel      groundResistance = 0.25 * resistanceMod
	groundResistanceRock        groundResistance = 0.30 * resistanceMod
	groundResistanceIce         groundResistance = 0.01 * resistanceMod
)

func groundResist(groundResistance groundResistance, speed float64) float64 {
	return speed * (1 - groundResistance)
}
