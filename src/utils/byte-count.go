package utils

import "fmt"

/**
 * These utility functions convert a size in bytes to a human-readable string in either
 * SI (decimal, aka KB, GB, etc. - converted with 1000) or
 * IEC (binary, aka KiB, GiB, etc. - converted with 1024)
 * Source: https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
 */

func ByteCountBase10(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func ByteCountBase2(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
