package utils

import (
	"fmt"
	"io"
	"os"
)

func OpenFileSafely(src string) (io.Reader, error) {
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
	defer source.Close()

	return source, nil
}
