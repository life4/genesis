package implementation

import (
	"sync"
)

// Channel is a set of operations with channel
type Channel struct {
	Data chan T
}

// Any returns true if f returns true for any element in channel
func (c Channel) Any(f func(el T) bool) bool {
	for el := range c.Data {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in channel
func (c Channel) All(f func(el T) bool) bool {
	for el := range c.Data {
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
		for el := range c.Data {
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

// Count return count of el occurrences in channel.
func (c Channel) Count(el T) int {
	count := 0
	for val := range c.Data {
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
		for el := range c.Data {
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
	for el := range c.Data {
		f(el)
	}
}

// Filter returns a new channel with elements from input channel
// for which f returns true
func (c Channel) Filter(f func(el T) bool) chan T {
	result := make(chan T, 1)
	go func() {
		for el := range c.Data {
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
		for el := range c.Data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

// Max returns the maximal element from channel
func (c Channel) Max() (T, error) {
	max, ok := <-c.Data
	if !ok {
		return max, ErrEmpty
	}
	for el := range c.Data {
		if el > max {
			max = el
		}
	}
	return max, nil
}

// Min returns the minimal element from channel
func (c Channel) Min() (T, error) {
	min, ok := <-c.Data
	if !ok {
		return min, ErrEmpty
	}
	for el := range c.Data {
		if el < min {
			min = el
		}
	}
	return min, nil
}

// Reduce applies f to acc and every element from channel and returns acc
func (c Channel) Reduce(acc G, f func(el T, acc G) G) G {
	for el := range c.Data {
		acc = f(el, acc)
	}
	return acc
}

// Scan is like Reduce, but returns slice of f results
func (c Channel) Scan(acc G, f func(el T, acc G) G) chan G {
	result := make(chan G, 1)
	go func() {
		for el := range c.Data {
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
	for el := range c.Data {
		sum += el
	}
	return sum
}

// Take takes first count elements from the channel.
func (c Channel) Take(count int) chan T {
	result := make(chan T, 1)
	go func() {
		defer close(result)
		if count <= 0 {
			return
		}
		i := 0
		for el := range c.Data {
			result <- el
			i++
			if i == count {
				return
			}
		}
	}()
	return result
}

// Tee returns 2 channels with elements from the input channel
func (c Channel) Tee(count int) []chan T {
	channels := make([]chan T, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan T, 1))
	}
	go func() {
		for el := range c.Data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan T) {
				defer wg.Done()
				ch <- el
			}
			wg.Add(count)
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
	for val := range c.Data {
		result = append(result, val)
	}
	return result
}
