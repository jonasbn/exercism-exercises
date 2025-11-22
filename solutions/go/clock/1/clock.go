package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	c := Clock{h: 0, m: 0}

	c = c.calculate(h, m)

	return c
}

func (c Clock) Add(m int) Clock {
	c = c.calculate(0, m)

	return c
}

func (c Clock) Subtract(m int) Clock {
	// reversing sign
	c = c.calculate(0, m*-1)

	return c
}

func (c Clock) String() string {

	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) calculate(h, m int) Clock {

	m = (c.m + m)
	m += (c.h + h) * 60

	m = m % 1440

	if m < 0 {
		m += 1440
	}

	h = m / 60
	m -= 60 * h

	if h == 24 {
		h = 0
	}

	c.m = m
	c.h = h

	return c
}
