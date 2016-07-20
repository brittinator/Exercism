package clock

import "fmt"

// must match `targetTestVersion` in clock_test.go.
const testVersion = 3
const MINUTES_IN_DAY = 24 * 60

// Clock is only in minutes
type Clock struct{ minute int }

func New(hr, min int) Clock {
	minutes := 60*hr + min
	return Clock{minutes}
}

func (c Clock) String() string {
	// prints out the clock time in hh:mm format as a string
	// pad the output with leading zeros if not double digit
	return fmt.Sprintf("%02d:%02d", Hour(c), Minute(c))
}

func (c Clock) Add(minutes int) Clock {
	c.minute = c.minute + minutes
	for c.minute < 0 {
		c.minute = c.minute + MINUTES_IN_DAY
	}
	return c
}

func Hour(c Clock) int {
	hour := c.minute / 60
	for hour > 23 {
		hour = hour - 24
	}
	return hour
}

func Minute(c Clock) int {
	return c.minute % 60
}
