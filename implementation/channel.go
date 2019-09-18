package implementation

// Channel is a set of operations with channel
type Channel struct {
	data chan T
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
