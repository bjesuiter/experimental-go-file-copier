package main

import (
	"fmt"
	"log"
	"os"

	"codemonument.com/overnight-copier/utils"
)

func main() {
	stats, err := os.Stat("example/source/image.jpg")

	if err != nil {
		log.Fatal(err)
	}

	//Prints stats of the file
	fmt.Printf("isDir: %t\n", stats.IsDir())
	fmt.Printf("Permission: %s\n", stats.Mode())
	fmt.Printf("Name: %s\n", stats.Name())
	fmt.Printf("Size: %s\n", utils.ReadableByteCount(stats.Size()))
	fmt.Printf("Modification Time: %s\n", stats.ModTime())
}
