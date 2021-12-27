package utils

import (
	"fmt"
	"os"
)

func OpenFileSafely(src string) (*os.File, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return nil, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return nil, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	return source, nil
}
