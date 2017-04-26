package accumulate

const testVersion = 1

// Accumulate fans out a function to each component in the slice
func Accumulate(input []string, command func(string) string) []string {
	for i, n := range input {
		input[i] = command(n)
	}
	return input
}
