package handlers

import "unicode"

func ValidateAPICount(input string) bool {
	if len(input) != 32 {
		return false
	}

	// Use a map to store the count of each character
	charCount := make(map[rune]int)

	// Iterate over each character in the string
	for _, char := range input {
		// Check if the character is a letter or a number
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			// Increment the count of the character in the map
			charCount[char]++

			// Check if the character is repeated 32 times
			if charCount[char] == 32 {
				return true
			}
		}
	}

	// If no character is repeated 32 times, return false
	return false
}
