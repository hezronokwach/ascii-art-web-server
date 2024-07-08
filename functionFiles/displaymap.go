package asciiart

import (
	"fmt"
	"strings"
)

func DisplayAsciiArt(characterMap map[rune][]string, input string) string {
	// Check if the character map is empty
	// Split the input string by "\n" to handle newline characters.
	inputRep := strings.Replace(input, "\r\n", "\\n", -1)
	inputSlice := strings.Split(inputRep, "\\n")
	var result strings.Builder
	for _, value := range inputSlice {
		if value == "" {
			result.WriteString("\n") // Ensure empty lines are maintained.
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range value {
				line, ok := characterMap[char]
				if ok {
					result.WriteString(fmt.Sprintf("%s ", line[i]))
				} else {
					result.WriteString("Characters not found")
					return ""
				}
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
