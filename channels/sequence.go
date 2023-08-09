package channels

import (
	"context"

	"github.com/life4/genesis/constraints"
)

// Counter is like Range, but infinite.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the ctx is canceled.
// The returned channel is closed when this goroutine finishes.
//
// ⏸️ The returned channel is unbuffered.
func Counter[T constraints.Integer](ctx context.Context, start T, step T) chan T {
	c := make(chan T)
	go func() {
		defer close(c)
		for {
			select {
			case <-ctx.Done():
				return
			case c <- start:
				start += step
			}
		}
	}()
	return c
}

// Exponential generates elements from start with
// multiplication of value on factor on every step.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the ctx is canceled.
// The returned channel is closed when this goroutine finishes.
//
// ⏸️ The returned channel is unbuffered.
func Exponential[T constraints.Integer](ctx context.Context, start T, factor T) chan T {
	c := make(chan T)
	go func() {
		defer close(c)
		for {
			select {
			case <-ctx.Done():
				return
			case c <- start:
				start *= factor
			}
		}
	}()
	return c
}

// Iterate returns an infinite list of repeated applications of f to val.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the ctx is canceled.
// The returned channel is closed when this goroutine finishes.
//
// ⏸️ The returned channel is unbuffered.
func Iterate[T constraints.Integer](ctx context.Context, val T, f func(val T) T) chan T {
	c := make(chan T)
	go func() {
		defer close(c)
		for {
			select {
			case <-ctx.Done():
				return
			case c <- val:
				val = f(val)
			}
		}
	}()
	return c
}

// Range generates elements from start to end with given step.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the ctx is canceled.
// The returned channel is closed when this goroutine finishes.
//
// ⏸️ The returned channel is unbuffered.
func Range[T constraints.Integer](ctx context.Context, start T, end T, step T) chan T {
	c := make(chan T)
	pos := start <= end
	go func() {
		defer close(c)
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
	}()
	return c
}

// Repeat returns channel that produces val infinite times.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the ctx is canceled.
// The returned channel is closed when this goroutine finishes.
//
// ⏸️ The returned channel is unbuffered.
func Repeat[T constraints.Integer](ctx context.Context, val T) chan T {
	c := make(chan T)
	go func() {
		defer close(c)
		for {
			select {
			case <-ctx.Done():
				return
			case c <- val:
				continue
			}
		}
	}()
	return c
}

// Replicate returns channel that produces val n times.
//
// ⏹️ Internally, the function starts a goroutine.
// This goroutine finishes when the ctx is canceled.
// The returned channel is closed when this goroutine finishes.
//
// ⏸️ The returned channel is unbuffered.
func Replicate[T constraints.Integer](ctx context.Context, val T, n int) chan T {
	c := make(chan T)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			c <- val
		}
	}()
	return c
}
