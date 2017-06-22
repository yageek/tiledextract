package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	inputPath := flag.String("i", "", "The tileset file")
	outputPath := flag.String("o", "", "The location of generated tile")

	flag.Parse()

	if *inputPath == "" || *outputPath == "" {
		printUsage()
		os.Exit(-1)
	}

	ext := Extractor{}
	if err := ext.Process(*inputPath, *outputPath); err != nil {
		fmt.Printf("Errors during the conversion: %v \n", err)
		os.Exit(-1)
	}

}

func printUsage() {
	fmt.Printf("%s -i [input] -o [output] \n", os.Args[0])
}
