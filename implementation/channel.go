package implementation

import (
	"sync"
)

// Channel is a set of operations with channel
type Channel struct {
	data chan T
}

// Any returns true if f returns true for any element in channel
func (c Channel) Any(f func(el T) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in channel
func (c Channel) All(f func(el T) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

// ChunkEvery returns channel with slices containing count elements each
func (c Channel) ChunkEvery(count int) chan []T {
	chunks := make(chan []T, 1)
	go func() {
		chunk := make([]T, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]T, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
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
		close(result)
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
		close(result)
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
		close(result)
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

// Tee returns 2 channels with elements from the input channel
func (c Channel) Tee(count int) []chan T {
	channels := make([]chan T, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan T, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan T) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

// ToSlice returns slice with all elements from channel.
func (c Channel) ToSlice() []T {
	result := make([]T, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}
