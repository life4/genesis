package channels

import (
	"context"
	"reflect"
	"sync"
)

// ChunkEveryC returns channel with slices containing count elements each.
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

// FirstC selects the first available element from the given channels.
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

// MergeC merges multiple channels into one.
//
// The order in which elements are merged from different channels
// is not guaranteed.
//
// ⏹️ Internally, the function starts a goroutine per input channel.
// Each of these goroutines finishes either when their input channel is closed.
// or when the ctx context is cancelled.
// The returned channel is closed when all of these goroutines finish.
//
// ⏸️ The returned channel is unbuffered.
// The goroutines will be blocked and won't consume elements
// from the input channels until the value from the output channel
// is consumed by another goroutine.
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
		wg.Done()
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
