package hamming

import "errors"

const testVersion = 4

// Distance returns the Hamming Distance between two strings of equal length.
func Distance(a, b string) (ham int, err error) {
	if len(a) != len(b) {
		err = errors.New("Strings must be of equal length.")
		return -1, err
	}
	if a == b || len(a) == 0 || len(b) == 0 {
		// short circuit loop if two loops are the same or length of zero
		ham = 0
	} else {
		i := 0
		for i < len(a) {
			if a[i] != b[i] {
				ham++
			}
			i++
		}
	}
	return
}
