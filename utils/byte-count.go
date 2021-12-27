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
	return fmt.Sprintf("%.2f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func ByteCountBase2(byteCount int64) string {
	const unit = 1024

	if byteCount < unit {
		return fmt.Sprintf("%d B", byteCount)
	}
	div, exp := unit, 0
	for n := byteCount / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %ciB",
		float64(byteCount)/float64(div), "KMGTPE"[exp])
}

func ByteCountBase2And10(byteCount int64) string {
	sizeBase2 := ByteCountBase2(byteCount)
	sizeBase10 := ByteCountBase10(byteCount)

	return fmt.Sprintf("%s / %s\n", sizeBase2, sizeBase10)
}

/**
 * @param template a Sprintf style string which must have exactly one %s placeholder
 */
func FormatByteCount(template string, byteCount int64, format ByteCountFormat) string {
	switch format {
	case Binary:
		return fmt.Sprintf(template, ByteCountBase2(byteCount))
	case Decimal:
		return fmt.Sprintf(template, ByteCountBase10(byteCount))
	case Both:
		return fmt.Sprintf(template, ByteCountBase2And10(byteCount))
	}

	return "unknown ByteCountFormat"
}
