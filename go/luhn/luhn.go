package luhn

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

// Valid uses the Luhn formula to determine if a number is valid.
func Valid(input string) bool {
	var scrubbedInput []int64
	// strip spaces
	for _, r := range input {
		if unicode.IsSpace(r) {
			continue
		}

		if !unicode.IsDigit(r) {
			return false
		}

		num, err := strconv.Atoi(string(r))
		if err != nil {
			log.Fatal("error converting", err)
		}
		scrubbedInput = append(scrubbedInput, int64(num))
	}
	// len 1 or 0 invalid
	if len(scrubbedInput) <= 1 {
		return false
	}

	fmt.Println("scrubbed> ", scrubbedInput)

	// double every 2nd starting @ back
	for i := len(scrubbedInput) - 2; i >= 0; i = i - 2 {
		scrubbedInput[i] = double(scrubbedInput[i])
	}
	var sum int64
	for _, num := range scrubbedInput {
		sum += num
	}

	if sum%10 == 0 {
		// true if divide 10 remainder zero
		return true
	}
	return false
}

// double is a special doubling function, see README for more details.
func double(n int64) int64 {
	// if result > 9, minus 9
	n = n * 2
	if n > 9 {
		n = n - 9
	}
	return n
}
