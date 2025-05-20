package util

import (
	"fmt"
	"strconv"
	"strings"
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

func ParseRGB(text string) (int, int, int, error) {
	splitRes := strings.Split(text, ",")
	err := fmt.Errorf("text should be of form 'r,g,b'")

	if len(splitRes) != 3 {
		return -1, -1, -1, err
	}

	r, errR := strconv.Atoi(splitRes[0])
	g, errG := strconv.Atoi(splitRes[1])
	b, errB := strconv.Atoi(splitRes[2])

	if errR != nil || errG != nil || errB != nil {
		return -1, -1, -1, err
	}

	return r, g, b, nil
}
