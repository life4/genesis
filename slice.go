package genesis

import (
	"iter"
	"math/rand/v2"
)

// Slices is a generic wrapper around built-in slice type.
//
// Can be constructed from an existing slice using [NewSlice].
type Slice[S ~[]T, T any] struct {
	s S
}

// NewSlice wraps an existing slice into [Slice].
func NewSlice[S ~[]T, T any](s S) Slice[S, T] {
	return Slice[S, T]{s}
}

func (s Slice[S, T]) Unwrap() S {
	return s.s
}

func (s Slice[S, T]) ToSlice() []T {
	return s.s
}

// Choice chooses a random element from the slice.
func (s Slice[S, T]) Choice() (T, error) {
	if len(s.s) == 0 {
		var tmp T
		return tmp, ErrEmpty
	}
	i := rand.IntN(len(s.s))
	return s.s[i], nil
}

// ChunkEvery splits items into groups of length count.
//
// If items can't be split evenly, the final chunk will contain the remaining elements.
func (s Slice[S, T]) ChunkEvery(count int) (Slice[[]S, S], error) {
	chunks := make([]S, 0)
	if count <= 0 {
		return Slice[[]S, S]{chunks}, ErrNegativeValue
	}
	chunk := make([]T, 0, count)
	for i, el := range s.s {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]T, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return Slice[[]S, S]{chunks}, nil
}

// Copy creates a copy of the slice.
func (s Slice[S, T]) Copy() Slice[S, T] {
	if s.s == nil {
		return s
	}
	var res S
	return Slice[S, T]{append(res, s.s...)}
}

// Cycle is an infinite loop over items.
//
// Doesn't iterate if the slice is empty.
func (s Slice[S, T]) Cycle() iter.Seq[T] {
	return func(yield func(T) bool) {
		if len(s.s) == 0 {
			return
		}
		for {
			for _, val := range s.s {
				more := yield(val)
				if !more {
					return
				}
			}
		}
	}
}
