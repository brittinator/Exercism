package acronym

import "strings"

const testVersion = 2

//Abbreviate comment ha ha
func Abbreviate(input string) string {
	var acronym string
	wordsNoSpaces := splitter(input, " ")
	for _, word := range wordsNoSpaces {
		wordsNoDashes := splitter(word, "-")
		for _, word := range wordsNoDashes {
			// edge case: if colon return everything before colon
			// last one
			end := len(word) - 1
			if word[end] == ':' {
				// return all but the colon
				return word[0 : len(word)-1]
			}

			// add first letter
			acronym = acronym + strings.ToUpper(word[:1])
			for i := 1; i < len(word)-1; i++ {
				if word[i:i+1] == strings.ToUpper(word[i:i+1]) {
					acronym = acronym + word[i:i+1]
				}
			}
		}
	}

	return acronym
}

func splitter(input, splitOn string) []string {
	return strings.Split(input, splitOn)
}
