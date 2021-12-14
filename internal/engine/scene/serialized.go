package scene

import "fmt"

type (
	SerializedScenes struct {
		Default ID       `xml:"default"`
		Refs    []string `xml:"refs>scene"`
	}

	SerializedScene struct {
		GameObjects []SerializedGameObject `xml:"objects>object"`
	}

	SerializedGameObject struct {
		ID         string                 `xml:"id,attr"`
		Name       string                 `xml:"name,attr"`
		Transform  SerializedTransform    `xml:"transform"`
		Components []SerializedComponent  `xml:"components>component"`
		Child      []SerializedGameObject `xml:"child>object"`
	}

	SerializedTransform struct {
		Position SerializedPosition `xml:"position"`
		Rotation float64            `xml:"rotation"`
		Scale    float64            `xml:"scale"`
	}

	SerializedPosition struct {
		X float64 `xml:"x"`
		Y float64 `xml:"y"`
	}

	SerializedComponent struct {
		ID    string                        `xml:"id,attr"`
		Props []SerializedComponentProperty `xml:"props>prop"`
	}

	SerializedComponentProperty struct {
		Name  string `xml:"name,attr"`
		Value string `xml:"value,attr"`
	}
)

func (rs *SerializedScenes) validate() error {
	if rs.Default == "" {
		return fmt.Errorf("should contain 'scenes.default'")
	}

	if len(rs.Refs) == 0 {
		return fmt.Errorf("should contain at least one item in 'scenes.refs'")
	}

	hasDefault := false
	for _, sceneID := range rs.Refs {
		if sceneID == rs.Default {
			hasDefault = true
		}
	}

	if !hasDefault {
		return fmt.Errorf("default scene '%s' not defined in 'scenes.refs'", rs.Default)
	}

	return nil
}
