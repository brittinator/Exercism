// Package clause.
package gigasecond

import "time"

// Constant declarations
const (
	testVersion               = 4
	gigasecond  time.Duration = 1000000000 * time.Second
)

// AddGigasecond takes a birthdate and returns the gigasecond anniversary
func AddGigasecond(born time.Time) time.Time {
	gigaBirfday := born.Add(gigasecond)
	return gigaBirfday
}
