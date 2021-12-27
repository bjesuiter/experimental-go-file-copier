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

	readableSize := utils.ByteCountBase2(stats.Size())

	//Prints stats of the file
	fmt.Printf("Permission: %s\n", stats.Mode())
	fmt.Printf("Name: %s\n", stats.Name())
	fmt.Printf("Size: %s\n", readableSize)
	fmt.Printf("Modification Time: %s\n", stats.ModTime())
}
