package implementation

// Counter is like Range, but infinite
func Counter(start T, step T) chan T {
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

// Product returns cortesian product of elements
// {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4}
func Product(arrs ...[]T) chan []T {
	c := make(chan []T, 1)
	go product(c, arrs, []T{}, 0)
	return c
}

func product(c chan []T, arrs [][]T, left []T, pos int) {

	// iterate over the last array
	if pos == len(arrs)-1 {
		for _, el := range arrs[pos] {
			result := make([]T, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range arrs[pos] {
		result := make([]T, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		product(c, arrs, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
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
