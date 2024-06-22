package asciiart

import (
	"unicode"
)

func IsPrintable(str string) bool {
	// Iterates through each character in the input string
	/* Checks if the current character is a backslash \ and if there's another character after it.
	   If so, checks if the next character is one of the escape characters: \t, \a, \b, \v, \f, or \r */
	for _, value := range str {
	
		if !unicode.IsPrint(value) {
			return false
		}
	}
	return true
}
	