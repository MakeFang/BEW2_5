package clock

import (
	"fmt"
)

const minutesInDay = 1440

type Clock int

func New(hour, minute int) Clock {
	c := Clock((hour*60+minute)%minutesInDay)
	if c < 0 {
		c += minutesInDay
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return New(0, int(c)-minutes)
}
