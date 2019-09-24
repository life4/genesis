package implementation

// Sequence is a set of operations to generate sequences
type Sequence struct {
	Data chan T
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

// Iterate returns an infinite list of repeated applications of f to val
func (s Sequence) Iterate(val T, f func(val T) T) chan T {
	c := make(chan T, 1)
	go func() {
		for {
			c <- val
			val = f(val)
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

// Replicate returns channel that produces val n times
func (s Sequence) Replicate(val T, n int) chan T {
	c := make(chan T, 1)
	go func() {
		for i := 0; i < n; i++ {
			c <- val
		}
	}()
	return c
}
