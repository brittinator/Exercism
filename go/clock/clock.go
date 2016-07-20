package clock

import "fmt"

const testVersion = 4
const minutesInADay = 24 * 60

// Clock is only in minutes
type Clock struct{ minute int }

// New creates a Clock struct that is in valid format,
// i.e. within the day 1440 minute frame
func New(hr, min int) Clock {
	minutes := 60*hr + min
	for minutes >= minutesInADay {
		minutes = minutes - minutesInADay
	}
	for minutes < 0 {
		minutes = minutes + minutesInADay
	}
	return Clock{minutes}
}

// String prints out the clock time in hh:mm format as a string
func (c Clock) String() string {
	// pad the output with leading zeros if not double digit
	return fmt.Sprintf("%02d:%02d", Hour(c), Minute(c))
}

func (c Clock) Add(minutes int) Clock {
	c.minute = c.minute + minutes
	for c.minute < 0 {
		c.minute = c.minute + minutesInADay
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
