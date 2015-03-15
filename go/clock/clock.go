// package main
package clock

import (
	"fmt"
)

var (
	TestVersion = 1
)

type Clock struct {
	minutes int
}

func New(hrs, mins int) Clock {
	return militaryTime(Clock{onlyMinutes(hrs, mins)})
}

func militaryTime(c Clock) Clock {
	dayMinutes := 24 * 60
	if c.minutes < 0 {
		c.minutes = dayMinutes + (c.minutes % dayMinutes)
	} else {
		c.minutes = c.minutes % dayMinutes
	}
	return c
}

func (c Clock) Add(mins int) Clock {
	c.minutes += mins
	return militaryTime(c)
}

func (c *Clock) String() string {
	hours, minutes := hoursMinutes(0, c.minutes)
	return doubleDigit(fmt.Sprintf("%d", hours)) + ":" + doubleDigit(fmt.Sprintf("%d", minutes))
}

func hoursMinutes(hrs, mins int) (int, int) {
	num := (60 * hrs) + mins
	remainder := num % 60
	return (num - (remainder)) / 60, remainder
}

func onlyMinutes(hrs, mins int) int {
	return (60 * hrs) + mins
}

func doubleDigit(s string) string {
	need := 2 - len(s)
	if need > 0 {
		return "0" + s
	} else {
		return s
	}
}
