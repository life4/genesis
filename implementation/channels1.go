package implementation

// Count is like Range, but infinite
func Count(start T, step T) chan T {
	c := make(chan T, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

// Cycle is an infinite loop over arr
func Cycle(arr []T) chan T {
	c := make(chan T, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Exponential generates elements from start with
// multiplication of value on factor on every step
func Exponential(start T, factor T) chan T {
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
func Range(start T, end T, step T) chan T {
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
func Repeat(val T) chan T {
	c := make(chan T, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func Take(c chan T, n int) []T {
	result := make([]T, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAll(c chan T) []T {
	result := make([]T, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}
