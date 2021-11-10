package level

const RoomWidth = 25
const RoomHeight = 10
const LayersCount = 4

type (
	SpriteID     = uint8
	BankID       = uint8
	EntityIndex  = uint8
	TileSetID    = uint8
	AtmosphereID = uint8
	MusicID      = uint8
	GradientID   = uint8

	World struct {
		Rooms []Room
	}

	Room struct {
		Position Position
		Sprites  [LayersCount]SpriteLayer
		Entities []Entity
		Options  Options
		Dummy    Dummy
	}

	SpriteLayer struct {
		Sprites [RoomWidth * RoomHeight]SpriteID
	}

	Entity struct {
		Layer    uint8
		Position Position
		BankID   BankID
		Index    EntityIndex
	}

	Options struct {
		TileSetIDA    TileSetID
		TileSetIDB    TileSetID
		AtmosphereIDA AtmosphereID
		AtmosphereIDB AtmosphereID
		MusicID       MusicID
		GradientID    GradientID
	}

	Dummy struct {
		UnknownInt uint32
	}

	Position struct {
		X int
		Y int
	}
)
