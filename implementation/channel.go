package implementation

// Channel is a set of operations with channel
type Channel struct {
	data chan T
}

// Count return count of el occurences in channel.
func (c Channel) Count(el T) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

// Drop drops first n elements from channel c and returns a new channel with the rest.
// It returns channel do be unblocking. If you want array instead, wrap result into TakeAll.
func (c Channel) Drop(n int) chan T {
	result := make(chan T, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

// Each calls f for every element in the channel
func (c Channel) Each(f func(el T)) {
	for el := range c.data {
		f(el)
	}
}

// Filter returns a new channel with elements from input channel
// for which f returns true
func (c Channel) Filter(f func(el T) bool) chan T {
	result := make(chan T, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
	}()
	return result
}

// Map applies f to all elements from channel and returns channel with results
func (c Channel) Map(f func(el T) G) chan G {
	result := make(chan G, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
	}()
	return result
}

// Max returns the maximal element from channel
func (c Channel) Max() T {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from channel
func (c Channel) Min() T {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies f to acc and every element from channel and returns acc
func (c Channel) Reduce(acc G, f func(el T, acc G) G) G {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

// Scan is like Reduce, but returns slice of f results
func (c Channel) Scan(acc G, f func(el T, acc G) G) chan G {
	result := make(chan G, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
	}()
	return result
}

// Sum returns sum of all elements from channel
func (c Channel) Sum() T {
	var sum T
	for el := range c.data {
		sum += el
	}
	return sum
}

// Take takes first n elements from channel c.
func (c Channel) Take(n int) []T {
	result := make([]T, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

// ToSlice returns slice with all elements from channel.
func (c Channel) ToSlice() []T {
	result := make([]T, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}
