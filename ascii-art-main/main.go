package main

import (
	"fmt"
	"os"

	asciiart "asciiart/functionFiles"
)

func main() {
	// Check if command-line arguments are appropriate, providing usage instructions if not.
	
	// Read the ASCII art map from the specified file.
	characterMap, err := asciiart.CreateMap(filePath)
	if err != nil {
		fmt.Println("Error reading map:", err)
		return
	}
	// Display ASCII art corresponding to the input string using the provided map
	if len(characterMap) == 0 {
		fmt.Println("File is empty")
		return
	}
	asciiart.DisplayAsciiArt(characterMap, input)
}
