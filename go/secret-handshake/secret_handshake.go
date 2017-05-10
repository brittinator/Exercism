package secret

const testVersion = 1

// Handshake translates base-10 digits into their binary representation,
// then returns secret codewords based on the translation
func Handshake(input uint) (output []string) {
	// bitwise AND operation to see if there's something in the first bit
	// 1 is 00001 in binary
	// 1 = wink
	if input&1 > 0 {
		output = append(output, "wink")
	}
	// 2 is 00010 in binary
	// 10 = double blink
	if input&2 > 0 {
		output = append(output, "double blink")
	}
	// 4 is 00100 in binary
	// 100 = close your eyes
	if input&4 > 0 {
		output = append(output, "close your eyes")
	}
	// 8 is 01000 in binary
	// 1000 = jump
	if input&8 > 0 {
		output = append(output, "jump")
	}
	// 16 is 10000 in binary
	// 10000 = Reverse the order of the operations in the secret handshake.
	if input&16 > 0 {
		output = reverse(output)
	}
	return
}

func reverse(input []string) []string {
	output := make([]string, len(input))
	for i, codeWord := range input {
		output[len(input)-1-i] = codeWord
	}
	return output
}
