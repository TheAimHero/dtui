package size

import (
	"fmt"
)

func formatFileSize(bytes int64) string {
	const (
		_ = iota
		// Constants for different file size units
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
	)

	size := float64(bytes)
	unit := ""

	switch {
	case size < KB:
		unit = "B"
	case size < MB:
		size /= KB
		unit = "KB"
	case size < GB:
		size /= MB
		unit = "MB"
	case size < TB:
		size /= GB
		unit = "GB"
	case size < PB:
		size /= TB
		unit = "TB"
	case size < EB:
		size /= PB
		unit = "PB"
	default:
		size /= EB
		unit = "EB"
	}

	return fmt.Sprintf("%.2f%s", size, unit)
}

func GetSize(size int64) string {
	return formatFileSize(size)
}
