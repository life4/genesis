package channels

import (
	"context"
	"sync"

	"github.com/life4/genesis/constraints"
)

// Any returns true if f returns true for any element in channel.
//
// ⏹️ The function returns either when f returns true
// or when the channel is closed. If you want to be able
// to stop the function without closing the channel,
// wrap the channel into [WithContext].
func Any[T any](c <-chan T, f func(el T) bool) bool {
	for el := range c {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in channel.
//
// ⏹️ The function returns either when f returns false
// or when the channel is closed. If you want to be able
// to stop the function without closing the channel,
// wrap the channel into [WithContext].
func All[T any](c <-chan T, f func(el T) bool) bool {
	for el := range c {
		if !f(el) {
			return false
		}
	}
	return true
}

// BufferSize returns how many messages a channel can hold before being blocked.
//
// When you push that many messages in the channel, the next attempt
// to write in it will block until another goroutine reads a message.
//
// For unbuffered channels the result is 0.
func BufferSize[T any](c <-chan T) int {
	return cap(c)
}

func ChunkEvery[T any](c <-chan T, count int) chan []T {
	return ChunkEveryC(context.Background(), c, count)
}

// Close safely closes the given channel.
//
// This is a safer version of built-in close. The built-in close function
// will panic if the given channel is already closed or nil.
// This function in both cases returns false.
//
// There is a reason why the built-in function might panic.
// The best practice of working with channels is to close the channel
// only within the context ("owner") that created it.
// In such cases, you can be 100% sure that the channel is non-nil
// and closed only once. However, if you found yourself in a different
// situation and no refactoring can guarantee these rules,
// the safer Close function is here to help.
func Close[T any](c chan<- T) bool {
	defer func() {
		_ = recover()
	}()
	close(c)
	return true
}

// Count return count of el occurrences in channel.
//
// ⏹️ The function returns only when the channel is closed.
// If you want to be able to stop the function
// without closing the channel, wrap the channel into [WithContext].
func Count[T comparable](c <-chan T, el T) int {
	count := 0
	for val := range c {
		if val == el {
			count++
		}
	}
	return count
}

// Drop drops first n elements from channel c and returns a new channel with the rest.
// It returns channel do be unblocking. If you want array instead, wrap result into TakeAll.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the input channel is closed.
// The returned channel is closed when this goroutine finishes.
//
// 🐞 BUG: The goroutine might not be cleaned up if
// the input channel is closed but the goroutine is blocked
// in attempt to write into the output channel.
// To avoid the issue, make sure to consume all messages
// from the output channel. In a future release, the function
// might be changed to accept a context for better cancelation.
//
// ⏸️ The returned channel is unbuffered.
// The goroutine will be blocked and won't consume elements
// from the input channel until the value from the output channel
// is consumed by another goroutine.
func Drop[T any](c <-chan T, n int) chan T {
	result := make(chan T)
	go func() {
		defer close(result)
		i := 0
		for el := range c {
			if i >= n {
				result <- el
			}
			i++
		}
	}()
	return result
}

// Each calls f for every element in the channel.
func Each[T any](c <-chan T, f func(el T)) {
	for el := range c {
		f(el)
	}
}

// Filter returns a new channel with elements from input channel
// for which f returns true.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the input channel is closed.
// The returned channel is closed when this goroutine finishes.
//
// 🐞 BUG: The goroutine might not be cleaned up if
// the input channel is closed but the goroutine is blocked
// in attempt to write into the output channel.
// To avoid the issue, make sure to consume all messages
// from the output channel. In a future release, the function
// might be changed to accept a context for better cancelation.
//
// ⏸️ The returned channel is unbuffered.
// The goroutine will be blocked and won't consume elements
// from the input channel until the value from the output channel
// is consumed by another goroutine.
func Filter[T any](c <-chan T, f func(el T) bool) chan T {
	result := make(chan T)
	go func() {
		defer close(result)
		for el := range c {
			if f(el) {
				result <- el
			}
		}
	}()
	return result
}

// First is an alias for [FirstC] without a context.
func First[T any](cs ...<-chan T) (T, error) {
	return FirstC(context.Background(), cs...)
}

// Given a channel of channels of values, return a channel of values.
//
// This pattern is described in the "Concurrency in Go" book
// as "the Bridge channel" pattern. It might be useful when you design
// your system as a pipeline consisting of goroutines connected through
// channels and you have steps producing new steps.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the input channel is closed.
// The returned channel is closed when this goroutine finishes.
//
// 🐞 BUG: The goroutine might not be cleaned up if
// the input channel is closed but the goroutine is blocked
// in attempt to write into the output channel.
// To avoid the issue, make sure to consume all messages
// from the output channel. In a future release, the function
// might be changed to accept a context for better cancelation.
//
// ⏸️ The returned channel is unbuffered.
// The goroutine will be blocked and won't consume elements
// from the input channel until the value from the output channel
// is consumed by another goroutine.
func Flatten[T any](c <-chan <-chan T) chan T {
	res := make(chan T)
	go func() {
		defer close(res)
		for stream := range c {
			for v := range stream {
				res <- v
			}
		}
	}()
	return res
}

// IsEmpty returns true if there are no messages in the channel.
//
// For unbuffered channels, the result is always true.
func IsEmpty[T any](c <-chan T) bool {
	return len(c) == 0
}

// IsFull returns true if the channel's buffer is full.
//
// Attempts to write into a full channel will block
// until another goroutine reads a message from it.
//
// For unbuffered channels, the result is always true.
func IsFull[T any](c <-chan T) bool {
	return len(c) == cap(c)
}

// Map applies f to all elements from channel and returns channel with results.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the input channel is closed.
// The returned channel is closed when this goroutine finishes.
//
// 🐞 BUG: The goroutine might not be cleaned up if
// the input channel is closed but the goroutine is blocked
// in attempt to write into the output channel.
// To avoid the issue, make sure to consume all messages
// from the output channel. In a future release, the function
// might be changed to accept a context for better cancelation.
//
// ⏸️ The returned channel is unbuffered.
// The goroutine will be blocked and won't consume elements
// from the input channel until the value from the output channel
// is consumed by another goroutine.
func Map[T any, G any](c <-chan T, f func(el T) G) chan G {
	result := make(chan G, 1)
	go func() {
		defer close(result)
		for el := range c {
			result <- f(el)
		}
	}()
	return result
}

// Max returns the maximal element from channel.
//
// ⏹️ The function is blocked until the channel is closed.
//
// 😱 If a channel is closed without any elements being emitted, `ErrEmpty` is returned.
func Max[T constraints.Ordered](c <-chan T) (T, error) {
	max, ok := <-c
	if !ok {
		return max, ErrEmpty
	}
	for el := range c {
		if el > max {
			max = el
		}
	}
	return max, nil
}

// Merge is an alias for [MergeC] without a context.
func Merge[T any](cs ...<-chan T) chan T {
	return MergeC(context.Background(), cs...)
}

// Min returns the minimal element from channel.
//
// ⏹️ The function is blocked until the channel is closed.
//
// 😱 If a channel is closed without any elements being emitted, `ErrEmpty` is returned.
func Min[T constraints.Ordered](c <-chan T) (T, error) {
	min, ok := <-c
	if !ok {
		return min, ErrEmpty
	}
	for el := range c {
		if el < min {
			min = el
		}
	}
	return min, nil
}

// Reduce applies f to acc and every element from channel and returns acc.
func Reduce[T any, G any](c <-chan T, acc G, f func(el T, acc G) G) G {
	for el := range c {
		acc = f(el, acc)
	}
	return acc
}

// Scan is like Reduce, but returns slice of f results.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the input channel is closed.
// The returned channel is closed when this goroutine finishes.
//
// 🐞 BUG: The goroutine might not be cleaned up if
// the input channel is closed but the goroutine is blocked
// in attempt to write into the output channel.
// To avoid the issue, make sure to consume all messages
// from the output channel. In a future release, the function
// might be changed to accept a context for better cancelation.
//
// ⏸️ The returned channel is unbuffered.
// The goroutine will be blocked and won't consume elements
// from the input channel until the value from the output channel
// is consumed by another goroutine.
func Scan[T any, G any](c <-chan T, acc G, f func(el T, acc G) G) chan G {
	result := make(chan G, 1)
	go func() {
		defer close(result)
		for el := range c {
			acc = f(el, acc)
			result <- acc
		}
	}()
	return result
}

// Sum returns sum of all elements from channel.
func Sum[T constraints.Ordered](c <-chan T) T {
	var sum T
	for el := range c {
		sum += el
	}
	return sum
}

// Take takes first count elements from the channel.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the input channel is closed.
// The returned channel is closed when this goroutine finishes.
//
// 🐞 BUG: The goroutine might not be cleaned up if
// the input channel is closed but the goroutine is blocked
// in attempt to write into the output channel.
// To avoid the issue, make sure to consume all messages
// from the output channel. In a future release, the function
// might be changed to accept a context for better cancelation.
//
// ⏸️ The returned channel is unbuffered.
// The goroutine will be blocked and won't consume elements
// from the input channel until the value from the output channel
// is consumed by another goroutine.
func Take[T any](c <-chan T, count int) chan T {
	result := make(chan T)
	go func() {
		defer close(result)
		if count <= 0 {
			return
		}
		i := 0
		for el := range c {
			result <- el
			i++
			if i == count {
				return
			}
		}
	}()
	return result
}

// Tee returns "count" number of channels with elements from the input channel.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the input channel is closed.
// The returned channels are closed when this goroutine finishes.
//
// 🐞 BUG: The goroutine might not be cleaned up if
// the input channel is closed but the goroutine is blocked
// in attempt to write into the output channel.
// To avoid the issue, make sure to consume all messages
// from the output channel. In a future release, the function
// might be changed to accept a context for better cancelation.
//
// ⏸️ The returned channels are unbuffered.
// The goroutine will be blocked and won't consume elements
// from the input channel until the value from all the output channels
// is consumed by another goroutine(s).
func Tee[T any](c <-chan T, count int) []chan T {
	channels := make([]chan T, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan T))
	}
	go func() {
		defer func() {
			for _, ch := range channels {
				close(ch)
			}
		}()

		for el := range c {
			wg := sync.WaitGroup{}
			putInto := func(ch chan T) {
				defer wg.Done()
				ch <- el
			}
			wg.Add(count)
			for _, ch := range channels {
				go putInto(ch)
			}
			wg.Wait()
		}
	}()
	return channels
}

// ToSlice returns slice with all elements from channel.
func ToSlice[T any](c <-chan T) []T {
	result := make([]T, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// WithBuffer creates an echo channel of the given one with the given buffer size.
//
// This function effectively makes writes into the given channel non-blocking
// until the buffer size of pending messages is reached, assuming that all reads
// will be done only from the channel that the function returns.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the input channel is closed.
// The returned channel is closed when this goroutine finishes.
//
// 🐞 BUG: The goroutine might not be cleaned up if
// the input channel is closed but the goroutine is blocked
// in attempt to write into the output channel.
// To avoid the issue, make sure to consume all messages
// from the output channel (or at least up to the buffer size).
// In a future release, the function
// might be changed to accept a context for better cancelation.
func WithBuffer[T any](c <-chan T, bufSize int) chan T {
	result := make(chan T, bufSize)
	go func() {
		defer close(result)
		for el := range c {
			result <- el
		}
	}()
	return result
}
