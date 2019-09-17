package implementation

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
	result := make([]T, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}
