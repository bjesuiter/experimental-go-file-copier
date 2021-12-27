package utils

type ByteCountFormat int64

const (
	// since iota starts with 0, the first value
	// defined here will be the default
	Decimal ByteCountFormat = iota // KB, MB, GB (base 10 => 1000)
	Binary                         //KiB, MiB, etc. (base 2 => 1024)
	Both                           // outputs xx.xxx MiB / xx.xxx MB
)
