package physics

import "fmt"

const (
	maskAll  Mask = 0xFFFF
	maskNone Mask = 0xFFFF
)

const (
	LayerGround Layer = 1 << iota
	LayerPlayer
	LayerEntities
)

type Layer uint16
type Mask = uint16

func (l Layer) Category() uint16 {
	return uint16(l)
}

func (l Layer) Mask() uint16 {
	switch l {
	case LayerGround:
		return maskAll
	case LayerPlayer:
		return uint16(LayerGround) | uint16(LayerEntities)
	case LayerEntities:
		return uint16(LayerGround) | uint16(LayerPlayer)
	default:
		panic(fmt.Sprintf("Collision mask for layer '%d' is not set", l))
	}
}
