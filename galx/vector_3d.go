package galx

import (
	"fmt"
	"math"
)

// SizeOfVec3 its size for low precision memory data dump (float32)
// dump have size of 8 bytes (x=4 + y=4 + z=4)
const SizeOfVec3 = 12

// Vec3d is common vector data structure
type Vec3d struct {
	X, Y, Z float64
}

func (v *Vec3d) String() string {
	return fmt.Sprintf("Vec3d{%.4f, %.4f, %.4f}", v.X, v.Y, v.Z)
}

// Data dump for low precision memory representation (GPU, shaders, etc..)
func (v *Vec3d) Data() []byte {
	var buf [SizeOfVec3]byte
	nX := math.Float32bits(float32(v.X))
	nY := math.Float32bits(float32(v.Y))
	nZ := math.Float32bits(float32(v.Z))

	buf[0] = byte(nX)
	buf[1] = byte(nX >> 8)
	buf[2] = byte(nX >> 16)
	buf[3] = byte(nX >> 24)

	buf[4] = byte(nY)
	buf[5] = byte(nY >> 8)
	buf[6] = byte(nY >> 16)
	buf[7] = byte(nY >> 24)

	buf[8] = byte(nZ)
	buf[9] = byte(nZ >> 8)
	buf[10] = byte(nZ >> 16)
	buf[11] = byte(nZ >> 24)

	return buf[:]
}
