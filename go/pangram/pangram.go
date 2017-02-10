package pangram

import "strings"

// This solution created in 20 minutes
const testVersion = 1

// IsPangram returns true if the incoming sentence uses all 26 letters
// of the English alphabet, and false if otherwise
func IsPangram(input string) bool {
	if len(input) < 26 {
		return false
	}

	input = strings.ToUpper(input)
	// create the map that will house character count
	alphaCount := make(map[rune]int)
	for char := 'A'; char <= 'Z'; char++ {
		alphaCount[char] = 0
	}

	// iterate through input and tally up alphabetical runes
	for _, char := range input {
		val, present := alphaCount[char]
		if present {
			alphaCount[char] = val + 1
		}
	}

	// check to see if any alpha values are still zero
	for _, value := range alphaCount {
		if value == 0 {
			return false
		}
	}

	return true
}
