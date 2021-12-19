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

func vkStringsToStringLabels(src []string) []string {
	labels := make([]string, 0, len(src))
	for _, s := range src {
		labels = append(labels, vkLabelToString(vkStringToLabel(s)))
	}
	return labels
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

func vkClampUint(n, min, max uint32) uint32 {
	if n <= min {
		return min
	}

	if n >= max {
		return max
	}

	return n
}
