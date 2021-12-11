package utils

import (
	"fmt"
)

func FormatBytes(b uint64) string {
	const unit = uint64(1000)
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := unit, 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1d %cB", b/div, "kMGTPE"[exp])
}
