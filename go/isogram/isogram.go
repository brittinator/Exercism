package isogram

import "unicode"

// IsIsogram determines whether or not any alpha
// character repeats. If so returns false.
func IsIsogram(input string) bool {
	letters := make(map[rune]int, len(input))
	for _, char := range input {
		if char == '-' || char == ' ' {
			continue
		}
		char = unicode.ToLower(char)
		if _, found := letters[char]; found {
			return false
		}
		letters[char] = 1
	}
	return true
}
