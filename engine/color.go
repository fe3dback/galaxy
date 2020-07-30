package engine

import (
	"encoding/binary"
)

type Color uint32

const (
	ColorBackground = 0x282A36FF
	ColorForeground = 0xF8F8F2FF
	ColorSelection  = 0x44475AFF
	ColorRed        = 0xFF5555FF
	ColorGreen      = 0x50FA7BFF
	ColorCyan       = 0x8BE9FDFF
	ColorOrange     = 0xFFB86CFF
	ColorPink       = 0xFF79C6FF
	ColorPurple     = 0xBD93F9FF
	ColorYellow     = 0xF1FA8CFF
)

func NewColor(r, g, b, a uint8) Color {
	return Color(binary.LittleEndian.Uint32([]byte{r, g, b, a}))
}

func (c Color) Split() (r uint8, g uint8, b uint8, a uint8) {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(c))
	return bytes[3], bytes[2], bytes[1], bytes[0]
}
