package collatzconjecture

import (
	"errors"
)

/*
Take any positive integer n. If n is even, divide n by 2 to get n / 2. If n is
odd, multiply n by 3 and add 1 to get 3n + 1. Repeat the process indefinitely.
The conjecture states that no matter which number you start with, you will
always reach 1 eventually.
*/

// CollatzConjecture returns the number of steps it takes to reach one.
func CollatzConjecture(input int) (int, error) {
	if input < 1 {
		return 0, errors.New("input value must be a postive integer")
	}

	var steps int
	for {
		if input == 1 {
			return steps, nil
		}
		input = collatz(input)
		steps++
	}
}

func collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return n*3 + 1
}
