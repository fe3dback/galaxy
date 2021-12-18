package scene

type (
	SnapshotScene struct {
		ID          string               `xml:",omitempty"`
		GameObjects []SnapshotGameObject `xml:"objects>object"`
	}

	SnapshotGameObject struct {
		ID         string               `xml:"id,attr"`
		Name       string               `xml:"name,attr"`
		Transform  SnapshotTransform    `xml:"transform"`
		Components []SnapshotComponent  `xml:"components>component,omitempty"`
		Child      []SnapshotGameObject `xml:"child>object,omitempty"`
	}

	SnapshotTransform struct {
		Position SnapshotPosition `xml:"position"`
		Rotation float64          `xml:"rotation"`
		Scale    float64          `xml:"scale"`
	}

	SnapshotPosition struct {
		X float64 `xml:"x"`
		Y float64 `xml:"y"`
	}

	SnapshotComponent struct {
		ID    string                      `xml:"id,attr"`
		Props []SnapshotComponentProperty `xml:"props>prop,omitempty"`
	}

	SnapshotComponentProperty struct {
		Name  string `xml:"name,attr"`
		Value string `xml:"value,attr"`
	}
)
