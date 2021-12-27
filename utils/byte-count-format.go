package utils

import "log"

type ByteCountFormat int64

const (
	// since iota starts with 0, the first value
	// defined here will be the default
	Decimal ByteCountFormat = iota // KB, MB, GB (base 10 => 1000)
	Binary                         //KiB, MiB, etc. (base 2 => 1024)
	Both                           // outputs xx.xxx MiB / xx.xxx MB
)

func (f ByteCountFormat) String() string {
	switch f {
	case Decimal:
		return "Decimal (base 10, conversion number 1000)"
	case Binary:
		return "Binary (base 2, conversion number 1024)"
	case Both:
		return "Binary Size / Decimal Size"
	}

	log.Fatalf("Illegal ByteCountFormat! %s", f)
	return "unknown"
}
