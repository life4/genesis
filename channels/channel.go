package channels

import (
	"context"
	"reflect"
	"sync"

	"github.com/life4/genesis/constraints"
)

// Any returns true if f returns true for any element in channel.
func Any[T any](c <-chan T, f func(el T) bool) bool {
	for el := range c {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in channel.
func All[T any](c <-chan T, f func(el T) bool) bool {
	for el := range c {
		if !f(el) {
			return false
		}
	}
	return true
}

// ChunkEvery returns channel with slices containing count elements each.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the input channel is closed.
// The returned channel is closed when this goroutine finishes.
//
// ⏸️ The returned channel is unbuffered.
// The goroutine will be blocked and won't consume elements
// from the input channel until the value from the output channel
// is consumed by another goroutine.
func ChunkEvery[T any](c <-chan T, count int) chan []T {
	chunks := make(chan []T, 1)
	go func() {
		defer close(chunks)
		chunk := make([]T, 0, count)
		i := 0
		for el := range c {
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
	}()
	return chunks
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

// First selects the first available element from the given channels.
//
// The function returns in one of the following cases:
//
//  1. One of the given channels is closed. In this case,
//     [ErrClosed] is returned.
//  2. ⏹️ The ctx context is canceled. In this case,
//     the cancelation reason is returned as an error.
//  3. One of the given channels returns a value. In this case,
//     the error is nil.
//
// If all channels are non-closed and empty and ctx is not canceled,
// the function will block and wait for one of the above to occur.
//
// If multiple messages are received at the same time,
// only one is chosen via a uniform pseudo-random selection (see below).
// Other messages will stay in their queues untocuhed.
//
// In no scenario any messages are lost. A message is either
// returned by a function or stays in the queue.
//
// # 😱 Errors
//
//   - [ErrEmpty]: no channels are passed into the function.
//   - [ErrClosed]: a channel was closed.
//   - Another: cancelation cause returned by ctx.Err().
//
// # 🍽 What is starvation
//
// Imagine that you iterate through a list of channels and return an element
// from the first one that is available. Since the order of channels is always
// the same, if a value is always available in a multiple channels, you'll be
// selecting only from the first one. In other words, you never select from other
// channels as long as the first one always has a value for you. It might result
// in a situation when out of multiplt jobs you start only one is not blocked.
//
// This situation is called "[starvation]". To avoid that, the element in
// select statement, according to the [Go specification],
// "is chosen via a uniform pseudo-random selection".
//
// This function preserves that behavior by using internally
// a dynamically created (using [reflect]) select statement.
//
// [starvation]: https://en.wikipedia.org/wiki/Starvation_(computer_science)
// [Go specification]: https://go.dev/ref/spec#Select_statements
func First[T any](ctx context.Context, cs ...<-chan T) (T, error) {
	// try to use a regular select if a small number of channels is passed
	switch len(cs) {
	case 0:
		return *new(T), ErrEmpty
	case 1:
		select {
		case v, ok := <-cs[0]:
			if !ok {
				return *new(T), ErrClosed
			}
			return v, nil
		case <-ctx.Done():
			return *new(T), ctx.Err()
		}
	}

	cases := make([]reflect.SelectCase, 0, len(cs)+1)
	cases = append(cases, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(ctx.Done()),
	})
	for _, c := range cs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(c),
		})
	}
	chosen, rval, ok := reflect.Select(cases)
	if chosen == 0 {
		var v T
		return v, ctx.Err()
	}
	val := rval.Interface().(T)
	if !ok {
		return val, ErrClosed
	}
	return val, nil
}

// Map applies f to all elements from channel and returns channel with results.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the input channel is closed.
// The returned channel is closed when this goroutine finishes.
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

// Pop reads a value from the channel (with context).
//
// The function is blocking. It will wait and return
// in one of the following conditions:
//
//  1. ⏹️ The context is canceled.
//  2. ⏹️ The channel is closed.
//  3. There is a value pushed into the channel.
//
// In the first two cases, the second return value (called "more" or "ok") is "false".
// Otherwise, if a value is succesfully pulled from the channel, it is "true".
//
// Reads from nil channels block forever. So, if a nil channel is passed,
// the function will exit only when the context is canceled.
func Pop[T any](ctx context.Context, c <-chan T) (T, bool) {
	select {
	case v, more := <-c:
		return v, more
	case <-ctx.Done():
		return *new(T), false
	}
}

// Push writes the value into the channel (with context).
//
// ⚠️ Experimental! Behavior of this function might change in the future
// or the function can be removed altogether.
// It's not clear yet what's the best approach for when the target channel is closed.
// By default, Go panics in this case, which might be not good in some situations.
// Also, using this function might cause situations when the canceled context
// will be ignored by the target function instead of exiting.
func Push[T any](ctx context.Context, c chan<- T, v T) {
	select {
	case c <- v:
	case <-ctx.Done():
	}
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

// WithContext creates an echo channel of the given one that can be canceled with ctx.
//
// This can be useful in 2 scenarios:
//
//  1. To be able to cancel any function in this package
//     without closing the original channel.
//  2. For simpler iteration through channels with support for cancellation.
//     This pattern is descirbed in the "Concurrency in Go" book
//     in "The or-done-channel" chapter (page 119).
//
// ⏹️ Internally, the function starts a goroutine
// that copies values from the input channel into the output one.
// This goroutine finishes when the input channel is closed
// or the ctx context is canceled.
// The returned channel is closed when this goroutine finishes.
//
// ⏸️ The returned channel is unbuffered.
// The goroutine will be blocked and won't consume elements
// from the input channel until the value from the output channel
// is consumed by another goroutine.
func WithContext[T any](c <-chan T, ctx context.Context) chan T {
	result := make(chan T)
	go func() {
		defer close(result)
		for {
			select {
			case <-ctx.Done():
				return
			case val, more := <-c:
				if !more {
					return
				}
				select {
				case result <- val:
				case <-ctx.Done():
				}
			}
		}
	}()
	return result
}
