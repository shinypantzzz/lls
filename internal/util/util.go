package util

import (
	"fmt"
)

func HumanSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	fbytes := float64(bytes)
	index := -1
	for fbytes >= unit {
		fbytes /= unit
		index++
	}
	return fmt.Sprintf("%.1f %cB", fbytes, "KMGTPE"[index])
}
