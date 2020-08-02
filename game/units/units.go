package units

// in meters
type Meters = float64

// in m/s
type SpeedMpS = float64

// map pixels to meters
const PixelsPerMeter = 128.0

// calculate basic distances
const (
	DistanceMeter Meters = PixelsPerMeter
	DistanceCm    Meters = DistanceMeter / 100
	DistanceKm    Meters = DistanceMeter * 1000
)

// world
const Gravity = 9.8
