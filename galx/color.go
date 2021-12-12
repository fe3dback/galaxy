package galx

import (
	"encoding/binary"
)

type Color uint32

//goland:noinspection GoUnusedConst
const (
	ColorBackground Color = 0x282A36FF
	ColorGray1      Color = 0x333333FF
	ColorGray2      Color = 0x555555FF
	ColorForeground Color = 0xF8F8F2FF
	ColorSelection  Color = 0x44475AFF
	ColorRed        Color = 0xFF5555FF
	ColorGreen      Color = 0x50FA7BFF
	ColorCyan       Color = 0x8BE9FDFF
	ColorOrange     Color = 0xFFB86CFF
	ColorPink       Color = 0xFF79C6FF
	ColorPurple     Color = 0xBD93F9FF
	ColorYellow     Color = 0xF1FA8CFF
)

func NewColor(r, g, b, a uint8) Color {
	return Color(binary.LittleEndian.Uint32([]byte{r, g, b, a}))
}

func (c Color) Split() (r uint8, g uint8, b uint8, a uint8) {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(c))
	return bytes[3], bytes[2], bytes[1], bytes[0]
}
