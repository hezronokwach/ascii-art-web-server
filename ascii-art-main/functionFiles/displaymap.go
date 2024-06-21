package asciiart

import (
	"fmt"
	"strings"
)

func DisplayAsciiArt(characterMap map[rune][]string, input string) {
	// Check if the character map is empty
	if len(characterMap) == 0 {
		fmt.Println("Error: Character map is empty. Please provide a valid character map.")
		return
	}
	// Split the input string by "\n" to handle newline characters.
	inputSlice := strings.Split(input, "\\n")
	count := 0
	// Iterate through the split input segments.
	// If the segment is empty (indicating a newline), it prints a newline character, ensuring proper formatting.
	/* If the segment contains characters:
       Iterates through each character in the segment.
       Retrieves the corresponding ASCII art lines for the character from the character map.
       Prints each line of ASCII art for the character.
       If the character is not found in the map, it prints an error message and returns.*/
	for _, value := range inputSlice {
		if value == "" {
			count++
			if count < len(inputSlice) {
				fmt.Println()
			}
		} else {
			for i := 0; i < 8; i++ {
				for _, char := range value {
					line, ok := characterMap[char]
					if ok {
						fmt.Printf("%s ", line[i])
					} else {
						fmt.Println("Characters not found")
						return
					}
				}
				fmt.Println()
			}
		}
	}
}
