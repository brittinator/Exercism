package bob // package name must match the package name in bob_test.go
import "strings"

const testVersion = 2 // same as targetTestVersion

//Hey is a method that talks to Bob, a teenager with
//four reliable responses.
func Hey(tellBob string) string {
	tellBob = strings.TrimSpace(tellBob)
	// Don't say anything to him
	if hasLetters(strings.ToUpper(tellBob)) == false && strings.TrimSpace(tellBob) == "" {
		return "Fine. Be that way!"
	}
	// Yell at Bob (Uppercase), it trumps a question response
	if tellBob == strings.ToUpper(tellBob) && hasLetters(strings.ToUpper(tellBob)) {
		hasLetters := false
		for _, char := range strings.ToUpper(tellBob) {
			if char > 'A' && char < 'Z' {
				hasLetters = true
			}
		}
		if hasLetters {
			return "Whoa, chill out!"
		}
	}
	// Ask Bob a question
	if tellBob[(len(tellBob)-1)] == '?' {
		return "Sure."
	}

	// Catchall
	return "Whatever."
}

func hasLetters(tellBob string) bool {
	for _, char := range tellBob {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}
