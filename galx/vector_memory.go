package galx

import (
	"unsafe"
)

// utility vectors for working with memory, GPU, shaders, etc..

const (
	SizeOfVec2 = 8
	SizeOfVec3 = 12
	SizeOfVec4 = 16
)

type (
	Vec2 struct {
		X, Y float32
	}

	Vec3 struct {
		R, G, B float32
	}

	Vec4 struct {
		R, G, B, A float32
	}
)

func (v *Vec2) Size() uint64 {
	return SizeOfVec2
}

func (v *Vec2) Data() []byte {
	return (*(*[SizeOfVec2]byte)(unsafe.Pointer(v)))[:]
}

func (v *Vec3) Size() uint64 {
	return SizeOfVec3
}

func (v *Vec3) Data() []byte {
	return (*(*[SizeOfVec3]byte)(unsafe.Pointer(v)))[:]
}

func (v *Vec4) Size() uint64 {
	return SizeOfVec4
}

func (v *Vec4) Data() []byte {
	return (*(*[SizeOfVec4]byte)(unsafe.Pointer(&v)))[:]
}
