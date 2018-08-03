package luhn

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

// Valid uses the Luhn formula to determine if a number is valid.
func Valid(input string) bool {
	// len 1 or 0 invalid
	if len(input) <= 1 {
		return false
	}
	// trims leading & trailing spaces
	var scrubbedInput []int64
	// strip spaces
	for _, r := range input {
		if unicode.IsSpace(r) {
			continue
		}
		// make sure all digits
		if unicode.IsDigit(r) {
			num, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal("error converting", err)
			}
			scrubbedInput = append(scrubbedInput, int64(num))
		} else {
			return false
		}
	}
	// len 1 or 0 invalid
	if len(scrubbedInput) <= 1 {
		return false
	}

	fmt.Println("scrubbed> ", scrubbedInput)

	// double every 2nd starting @ back
	for i := len(scrubbedInput) - 2; i >= 0; i = i - 2 {
		// if result > 9, minus 9
		num := scrubbedInput[i]
		num = num * 2
		if num > 9 {
			num = num - 9
		}
		fmt.Printf("i is : %v, scrubbed: %v, num: %v\n", i, scrubbedInput[i], num)
		scrubbedInput[i] = num
	}
	var sum int64
	// sum everything
	for _, num := range scrubbedInput {
		sum += num
	}

	if sum%10 == 0 {
		// if divide 10 true
		return true
	}
	// mod 10

	return false
}
