package implementation

// Sequence is a set of operations to generate sequences
type Sequence struct {
	data chan T
}

// Count is like Range, but infinite
func (s Sequence) Count(start T, step T) chan T {
	c := make(chan T, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

// Exponential generates elements from start with
// multiplication of value on factor on every step
func (s Sequence) Exponential(start T, factor T) chan T {
	c := make(chan T, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}

// Range generates elements from start to end with given step
func (s Sequence) Range(start T, end T, step T) chan T {
	c := make(chan T, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func (s Sequence) Repeat(val T) chan T {
	c := make(chan T, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}
