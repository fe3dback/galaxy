package units

// count of Something
type Pixel = int32

// count of Something
type Count = int32

// rate in X per second
type Rate = float64

// in meters
type Meter = float64

// in seconds
type Second = float64

// in m/s
type SpeedMpS = float64

// in km/h
type SpeedKmH = float64

// map pixels to meters
const PixelsPerMeter = 128.0

// calculate basic distances
const (
	DistanceMeter Meter = PixelsPerMeter
	DistanceCm    Meter = DistanceMeter / 100
	DistanceKm    Meter = DistanceMeter * 1000
)

// world
const Gravity = 9.8

func TransformSpeed(mps SpeedMpS) SpeedKmH {
	metersPerHour := mps * 3600

	return metersPerHour / 1000
}
