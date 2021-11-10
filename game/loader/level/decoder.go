package level

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/nielsAD/gowarcraft3/protocol"
)

func WorldDecode(bin []byte) (World, error) {
	world := World{}

	var b protocol.Buffer
	if _, err := io.Copy(&b, bytes.NewReader(bin)); err != nil {
		return world, fmt.Errorf("failed copy binary to buffer: %w", err)
	}

	roomInd := -1
	for {
		if b.Size() == 0 {
			break
		}

		roomInd++
		room := Room{}

		// read position
		posStr := readTo(&b, 0x00) // read all chars until pos delimiter
		pos, err := decodePositionStr(string(posStr))
		if err != nil {
			return world, fmt.Errorf("failed decode room '%d' position: %w", roomInd, err)
		}

		room.Position = pos

		// dummy
		room.Dummy.UnknownInt = b.ReadUInt32()

		// sprite layers
		for layerInd := 0; layerInd < LayersCount; layerInd++ {
			for tileInd := 0; tileInd < RoomWidth*RoomHeight; tileInd++ {
				room.Sprites[layerInd].Sprites[tileInd] = b.ReadUInt8()
			}
		}

		// bank layers
		type index = int
		for layerInd := 0; layerInd < LayersCount; layerInd++ {
			objectsMap := make(map[index]EntityIndex)
			banksMap := make(map[index]BankID)

			// read objects ids
			for j := 0; j < RoomWidth*RoomHeight; j++ {
				objectsMap[j] = b.ReadUInt8()
			}

			// read bank ids
			for j := 0; j < RoomWidth*RoomHeight; j++ {
				banksMap[j] = b.ReadUInt8()
			}

			// assemble not empty bank data to entities
			for ind, bankId := range banksMap {
				if bankId == 0 {
					// skip empty bank object, for compression
					// 0 - nothing
					continue
				}

				objectID := objectsMap[ind]

				// ind = 249 for example
				// ceil(249 / 25) = 10
				entRowY := int(math.Ceil(float64(ind) / float64(RoomWidth)))

				// 25 - ((10 * 25 = 250) - 249 = 1) = 24
				entRowX := RoomWidth - ((entRowY * RoomWidth) - ind)

				room.Entities = append(room.Entities, Entity{
					Layer: uint8(layerInd),
					Position: Position{
						X: entRowX - 1,
						Y: entRowY - 1,
					},
					BankID: bankId,
					Index:  objectID,
				})
			}
		}

		// read options
		room.Options.TileSetIDA = b.ReadUInt8()
		room.Options.TileSetIDB = b.ReadUInt8()
		room.Options.AtmosphereIDA = b.ReadUInt8()
		room.Options.AtmosphereIDB = b.ReadUInt8()
		room.Options.MusicID = b.ReadUInt8()
		room.Options.GradientID = b.ReadUInt8()

		// apply room
		world.Rooms = append(world.Rooms, room)
	}

	return world, nil
}

func readTo(b *protocol.Buffer, stopChar byte) []byte {
	limit := 65536
	result := make([]byte, 0)

	for {
		limit--
		if limit <= 0 {
			break
		}

		current := b.ReadUInt8()
		if current == stopChar {
			return result
		}

		result = append(result, current)
	}

	// not found end char, return all consumed buffer
	return result
}

// decode x520y650 into {520,650}
func decodePositionStr(pos string) (Position, error) {
	parts := strings.Split(pos, "y")
	if len(parts) != 2 {
		return Position{}, fmt.Errorf("failed decode pos '%s', expected exactly two parts", pos)
	}

	x, err := strconv.Atoi(strings.TrimLeft(parts[0], "x"))
	if err != nil {
		return Position{}, fmt.Errorf("failed extract X pos: %w", err)
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return Position{}, fmt.Errorf("failed extract Y pos: %w", err)
	}

	return Position{
		X: x,
		Y: y,
	}, nil
}
