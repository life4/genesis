package channels

import (
	"context"
	"reflect"
	"sync"

	"github.com/life4/genesis/constraints"
)

// All returns true if f returns true for all elements in channel.
func AllC[T any](ctx context.Context, c <-chan T, f func(el T) bool) bool {
	for {
		select {
		case el, ok := <-c:
			if !ok {
				return true
			}
			if !f(el) {
				return false

			}
		case <-ctx.Done():
			return true
		}
	}
}

// Any returns true if f returns true for any element in channel.
func AnyC[T any](ctx context.Context, c <-chan T, f func(el T) bool) bool {
	for {
		select {
		case el, ok := <-c:
			if !ok {
				return false
			}
			if f(el) {
				return true

			}
		case <-ctx.Done():
			return false
		}
	}
}

// ChunkEveryC returns channel with slices containing count elements each.
func ChunkEveryC[T any](ctx context.Context, c <-chan T, count int) chan []T {
	chunks := make(chan []T, 1)
	go func() {
		defer close(chunks)
		chunk := make([]T, 0, count)
		i := 0
		for el := range WithContext(c, ctx) {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				select {
				case chunks <- chunk:
				case <-ctx.Done():
					return
				}
				chunk = make([]T, 0, count)
			}
		}
		if len(chunk) > 0 {
			select {
			case chunks <- chunk:
			case <-ctx.Done():
				return
			}
		}
	}()
	return chunks
}

// Count return count of el occurrences in channel.
func CountC[T comparable](ctx context.Context, c <-chan T, el T) int {
	count := 0
	for {
		select {
		case val, ok := <-c:
			if !ok {
				return count
			}
			if val == el {
				count++
			}
		case <-ctx.Done():
			return count
		}
	}
}

// Drop drops first n elements from channel c and returns a new channel with the rest.
// It returns channel do be unblocking. If you want array instead, wrap result into TakeAll.
func DropC[T any](ctx context.Context, c <-chan T, n int) chan T {
	result := make(chan T)
	go func() {
		defer close(result)
		i := 0
		for {
			select {
			case el, ok := <-c:
				if !ok {
					return
				}
				if i >= n {
					select {
					case result <- el:
					case <-ctx.Done():
						return
					}
				}
				i++
			case <-ctx.Done():
				return
			}
		}
	}()
	return result
}

// Each calls f for every element in the channel.
func EachC[T any](ctx context.Context, c <-chan T, f func(el T)) {
	for {
		select {
		case el, ok := <-c:
			if !ok {
				return
			}
			f(el)
		case <-ctx.Done():
			return
		}
	}
}

// EchoC moves messages from one channel to the other.
//
// If you want to move messages from multiple channels into one, use [MergeC] instead.
//
// If you want to move messages from one channel into multiple, use [TeeC] instead.
func EchoC[T any](ctx context.Context, from <-chan T, to chan<- T) {
	for {
		select {
		case el, ok := <-from:
			if !ok {
				return
			}
			select {
			case to <- el:
			case <-ctx.Done():
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

// Filter returns a new channel with elements from input channel for which f returns true.
func FilterC[T any](ctx context.Context, c <-chan T, f func(el T) bool) chan T {
	result := make(chan T)
	go func() {
		defer close(result)
		for {
			select {
			case el, ok := <-c:
				if !ok {
					return
				}
				if f(el) {
					select {
					case result <- el:
					case <-ctx.Done():
						return
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return result
}

// FirstC selects the first available element from the given channels.
//
// The function returns in one of the following cases:
//
//  1. One of the given channels is closed. In this case,
//     [ErrClosed] is returned.
//  2. The ctx context is canceled. In this case,
//     the cancelation reason is returned as an error.
//  3. One of the given channels returns a value. In this case,
//     the error is nil.
//
// If all channels are non-closed and empty and ctx is not canceled,
// the function will block and wait for one of the above to occur.
//
// If a message available in multiple channels, only one is chosen
// via a uniform pseudo-random selection to avoid [starvation].
//
// # 😱 Errors
//
//   - [ErrEmpty]: no channels are passed into the function.
//   - [ErrClosed]: a channel was closed.
//   - Another: cancelation cause returned by ctx.Err().
//
// [starvation]: https://en.wikipedia.org/wiki/Starvation_(computer_science)
func FirstC[T any](ctx context.Context, cs ...<-chan T) (T, error) {
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

// Given a channel of channels of values, return a channel of values.
//
// This pattern is described in the "Concurrency in Go" book
// as "the Bridge channel" pattern. It might be useful when you design
// your system as a pipeline consisting of goroutines connected through
// channels and you have steps producing new steps.
func FlattenC[T any](ctx context.Context, c <-chan <-chan T) chan T {
	result := make(chan T)
	go func() {
		defer close(result)
		for {
			select {
			case stream, ok := <-c:
				if !ok {
					return
				}
				EchoC(ctx, stream, result)
			case <-ctx.Done():
				return
			}
		}
	}()
	return result
}

// Map applies f to all elements from channel and returns channel with results.
func MapC[T any, G any](ctx context.Context, c <-chan T, f func(el T) G) chan G {
	result := make(chan G, 1)
	go func() {
		defer close(result)
		for {
			select {
			case el, ok := <-c:
				if !ok {
					return
				}
				select {
				case result <- f(el):
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return result
}

// Max returns the maximal element from channel.
func MaxC[T constraints.Ordered](ctx context.Context, c <-chan T) (T, error) {
	select {
	case max, ok := <-c:
		if !ok {
			return max, ErrEmpty
		}
		for {
			select {
			case el, ok := <-c:
				if !ok {
					return max, nil
				}
				if el > max {
					max = el
				}
			case <-ctx.Done():
				return max, ctx.Err()
			}
		}
	case <-ctx.Done():
		var max T
		return max, ctx.Err()
	}
}

// MergeC merges multiple channels into one.
//
// The order in which elements are merged from different channels
// is not guaranteed.
func MergeC[T any](ctx context.Context, cs ...<-chan T) chan T {
	res := make(chan T)

	// make sure the result channel is closed
	// when all input channels are closed
	wg := sync.WaitGroup{}
	wg.Add(len(cs))
	go func() {
		wg.Wait()
		close(res)
	}()

	echo := func(c <-chan T) {
		defer wg.Done()
		for {
			select {
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case res <- v:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}

	for _, c := range cs {
		go echo(c)
	}
	return res
}

// Min returns the minimal element from channel.
func MinC[T constraints.Ordered](ctx context.Context, c <-chan T) (T, error) {
	select {
	case min, ok := <-c:
		if !ok {
			return min, ErrEmpty
		}
		for {
			select {
			case el, ok := <-c:
				if !ok {
					return min, nil
				}
				if el < min {
					min = el
				}
			case <-ctx.Done():
				return min, ctx.Err()
			}
		}
	case <-ctx.Done():
		var min T
		return min, ctx.Err()
	}
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
func ReduceC[T any, G any](ctx context.Context, c <-chan T, acc G, f func(el T, acc G) G) G {
	for {
		select {
		case el, ok := <-c:
			if !ok {
				return acc
			}
			acc = f(el, acc)
		case <-ctx.Done():
			return acc
		}
	}
}

// Scan is like Reduce, but returns slice of f results.
func ScanC[T any, G any](ctx context.Context, c <-chan T, acc G, f func(el T, acc G) G) chan G {
	result := make(chan G, 1)
	go func() {
		defer close(result)
		for {
			select {
			case el, ok := <-c:
				if !ok {
					return
				}
				acc = f(el, acc)
				select {
				case result <- acc:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return result
}

// Sum returns sum of all elements from channel.
func SumC[T constraints.Ordered](ctx context.Context, c <-chan T) T {
	var sum T
	for {
		select {
		case el, ok := <-c:
			if !ok {
				return sum
			}
			sum += el
		case <-ctx.Done():
			return sum
		}
	}
}

// Take takes first count elements from the channel.
func TakeC[T any](ctx context.Context, c <-chan T, count int) chan T {
	result := make(chan T)
	go func() {
		defer close(result)
		if count <= 0 {
			return
		}
		i := 0
		for {
			select {
			case el, ok := <-c:
				if !ok {
					return
				}
				select {
				case result <- el:
				case <-ctx.Done():
					return
				}
				i++
				if i == count {
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return result
}

// Tee returns "count" number of channels with elements from the input channel.
func TeeC[T any](ctx context.Context, c <-chan T, count int) []chan T {
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

		for {
			select {
			case el, ok := <-c:
				if !ok {
					return
				}
				wg := sync.WaitGroup{}
				putInto := func(ch chan T) {
					defer wg.Done()
					select {
					case ch <- el:
					case <-ctx.Done():
					}
				}
				wg.Add(count)
				for _, ch := range channels {
					go putInto(ch)
				}
				wg.Wait()
			case <-ctx.Done():
				return
			}
		}
	}()
	return channels
}

// ToSlice returns slice with all elements from channel.
func ToSliceC[T any](ctx context.Context, c <-chan T) []T {
	result := make([]T, 0)
	for {
		select {
		case el, ok := <-c:
			if !ok {
				return result
			}
			result = append(result, el)
		case <-ctx.Done():
			return result
		}
	}
}

// WithBuffer creates an echo channel of the given one with the given buffer size.
//
// This function effectively makes writes into the given channel non-blocking
// until the buffer size of pending messages is reached, assuming that all reads
// will be done only from the channel that the function returns.
func WithBufferC[T any](ctx context.Context, c <-chan T, bufSize int) chan T {
	result := make(chan T, bufSize)
	go func() {
		defer close(result)
		EchoC(ctx, c, result)
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
