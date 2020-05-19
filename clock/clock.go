package clock

import (
	"fmt"
	"strconv"
	"time"
)

// Clock struct with minute and hour
type Clock struct {
	minute int
	hour   int
}

// New returns new instace of clock
func New(hour, minute int) Clock {
	realTime := time.Date(0, time.January, 1, hour, minute, 0, 0, time.UTC)

	return Clock{
		minute: realTime.Minute(),
		hour:   realTime.Hour(),
	}
}

// Add given minutes to clock
func (c Clock) Add(minutes int) Clock {

	t := getSomeTime(c.hour, c.minute+minutes)

	return Clock{
		minute: t.Minute(),
		hour:   t.Hour(),
	}
}

// Subtract given minutes from clock
func (c Clock) Subtract(minutes int) Clock {
	t := getSomeTime(c.hour, c.minute-minutes)

	return Clock{
		minute: t.Minute(),
		hour:   t.Hour(),
	}
}

// String stringifies clock in "01:30" format
func (c Clock) String() string {
	shour := strconv.Itoa(c.hour)
	sminute := strconv.Itoa(c.minute)

	if c.hour < 10 {
		shour = fmt.Sprintf("0%s", strconv.Itoa(c.hour))
	}

	if c.minute < 10 {
		sminute = fmt.Sprintf("0%s", strconv.Itoa(c.minute))
	}
	return fmt.Sprintf("%s:%s", shour, sminute)
}

func getSomeTime(hour, minute int) time.Time {
	return time.Date(0, time.January, 1, hour, minute, 0, 0, time.UTC)
}
