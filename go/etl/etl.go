package etl

import "strings"

const testVersion = 1

// Transform converts an old map of scores: letters to letter: score map
func Transform(legacy map[int][]string) map[string]int {
	// for each score, break out into a map entry of letter: score inside 1 map
	newFormat := map[string]int{}
	for score, letters := range legacy {
		for _, letter := range letters {
			newFormat[strings.ToLower(letter)] = score
		}
	}
	return newFormat
}
