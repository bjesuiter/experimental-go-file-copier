package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"codemonument.com/overnight-copier/utils"
)

/**
 * Note: This method reads the file in 32 KiB Blocks, since it uses io.CopyBuffer under the hood
 */

func main() {
	bytesCopied, err := copy("example/source/image.jpg", "example/target/image.jpg")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Bytes Copied: %s", utils.ReadableByteCount(bytesCopied))
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
