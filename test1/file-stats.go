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

	sizeBase2 := utils.ByteCountBase2(stats.Size())
	sizeBase10 := utils.ByteCountBase10(stats.Size())

	//Prints stats of the file
	fmt.Printf("Permission: %s\n", stats.Mode())
	fmt.Printf("Name: %s\n", stats.Name())
	fmt.Printf("Size: %s / %s\n", sizeBase2, sizeBase10)
	fmt.Printf("Modification Time: %s\n", stats.ModTime())
}
