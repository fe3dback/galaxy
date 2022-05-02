package galx

import (
	"fmt"
	"math"
)

// SizeOfVec4 its size for low precision memory data dump (float32)
// dump have size of 8 bytes (x=4 + y=4 + z=4 + a=4)
const SizeOfVec4 = 16

// Vec4d is common vector data structure
type Vec4d struct {
	X, Y, Z, A float64
}

func (v *Vec4d) String() string {
	return fmt.Sprintf("Vec4d{%.4f, %.4f, %.4f, %.4f}", v.X, v.Y, v.Z, v.A)
}

// Data dump for low precision memory representation (GPU, shaders, etc..)
func (v *Vec4d) Data() []byte {
	var buf [SizeOfVec4]byte
	nX := math.Float32bits(float32(v.X))
	nY := math.Float32bits(float32(v.Y))
	nZ := math.Float32bits(float32(v.Z))
	nA := math.Float32bits(float32(v.A))

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

	buf[12] = byte(nA)
	buf[13] = byte(nA >> 8)
	buf[14] = byte(nA >> 16)
	buf[15] = byte(nA >> 24)

	return buf[:]
}
