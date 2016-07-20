package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// IsLeapYear returns true or false on whether or not input year is a leap year
func IsLeapYear(year int) bool {
	if year%100 == 0 && year%400 == 0 {
		return true
	} else if year%4 == 0 && year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}
