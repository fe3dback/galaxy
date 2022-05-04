package galx

import (
	"fmt"
)

type Mat4 struct {
	A Vec4d
	B Vec4d
	C Vec4d
	D Vec4d
}

func (v *Mat4) String() string {
	return fmt.Sprintf(`Mat4[
     %0.3f, %0.3f, %0.3f, %0.3f,
     %0.3f, %0.3f, %0.3f, %0.3f,
     %0.3f, %0.3f, %0.3f, %0.3f,
     %0.3f, %0.3f, %0.3f, %0.3f,
]`,

		v.A.X, v.A.Y, v.A.Z, v.A.A,
		v.B.X, v.B.Y, v.B.Z, v.B.A,
		v.C.X, v.C.Y, v.C.Z, v.C.A,
		v.D.X, v.D.Y, v.D.Z, v.D.A,
	)
}

// Data dump for low precision memory representation (GPU, shaders, etc..)
func (v *Mat4) Data() []byte {
	buf := make([]byte, 0, SizeOfVec4*4)

	buf = append(buf, v.A.Data()...)
	buf = append(buf, v.B.Data()...)
	buf = append(buf, v.C.Data()...)
	buf = append(buf, v.D.Data()...)

	return buf
}

func Mat4Identity() Mat4 {
	return Mat4{
		A: Vec4d{1, 0, 0, 0},
		B: Vec4d{0, 1, 0, 0},
		C: Vec4d{0, 0, 1, 0},
		D: Vec4d{0, 0, 0, 1},
	}
}
