package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	stats, err := os.Stat("example/source/image.jpg")

	if err != nil {
		log.Fatal(err)
	}

	//Prints stats of the file
	fmt.Printf("Permission: %s\n", stats.Mode())
	fmt.Printf("Name: %s\n", stats.Name())
	fmt.Printf("Size: %d\n", stats.Size())
	fmt.Printf("Modification Time: %s\n", stats.ModTime())
}
