package render

import (
	"fmt"
	"strings"

	"github.com/vulkan-go/vulkan"
)

func vkAssert(result vulkan.Result, err error) {
	if result == vulkan.Success {
		return
	}

	panic(err)
}

func vkMapToList(src map[string]struct{}) []string {
	r := make([]string, 0, len(src))

	for s := range src {
		r = append(r, s)
	}

	return r
}

func vkLabelToString(src [256]byte) string {
	r := strings.Builder{}

	for _, b := range src {
		r.WriteByte(b)
	}

	return strings.TrimSpace(r.String())
}

func vkStringToLabel(src string) [256]byte {
	if len(src) > 256 {
		panic(fmt.Errorf("failed convert string '%s' to vulkan label: len is greater to 256", src))
	}

	label := [256]byte{}

	for i, b := range src {
		label[i] = byte(b)
	}

	return label
}
