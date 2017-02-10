package raindrops

import "strconv"

const testVersion = 2

// Convert returns a string of combos of Pling, Plang and Plong depending on the number's primes
func Convert(incomingNumber int) string {
	raindrops := ""

	if incomingNumber%3 == 0 {
		raindrops = raindrops + "Pling"
	}
	if incomingNumber%5 == 0 {
		raindrops = raindrops + "Plang"
	}
	if incomingNumber%7 == 0 {
		raindrops = raindrops + "Plong"
	}
	if raindrops == "" {
		raindrops = strconv.Itoa(incomingNumber)
	}

	return raindrops
}

// The test program has a benchmark too.  How fast does your Convert convert?
