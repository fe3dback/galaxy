package weapon

import (
	"math/rand"

	"github.com/fe3dback/galaxy/generated"
)

func randomResource(list []generated.ResourcePath) *generated.ResourcePath {
	count := len(list)

	if count == 0 {
		return nil
	}

	return &list[rand.Intn(count)]
}
