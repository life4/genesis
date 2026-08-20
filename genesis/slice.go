package genesis

import (
	"math/rand/v2"
	"slices"
)

type Slice[T any] []T

func NewSlice[S ~[]T, T any](s S) Slice[T] {
	return Slice[T](s)
}

func (s Slice[T]) Iter() Iter[T] {
	return NewIter(slices.Values(s))
}

func (s Slice[T]) First() Option[T] {
	if len(s) == 0 {
		return None[T]()
	}
	return Some(s[0])
}

func (s Slice[T]) Last() Option[T] {
	if len(s) == 0 {
		return None[T]()
	}
	return Some(s[len(s)-1])
}

func (s Slice[T]) Len() int {
	return len(s)
}

func (s Slice[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Slice[T]) Cap() int {
	return cap(s)
}

func (s Slice[T]) At(i int) T {
	return s[i]
}

func (s Slice[T]) Get(i int) Option[T] {
	if len(s) >= i {
		return None[T]()
	}
	return Some(s[i])
}

func (s Slice[T]) Choice() Option[T] {
	if len(s) == 0 {
		return None[T]()
	}
	i := rand.IntN(len(s))
	return Some(s[i])
}

func (s Slice[T]) Copy() Slice[T] {
	if s == nil {
		return nil
	}
	return append(Slice[T]{}, s...)
}

func (s Slice[T]) Each(f func(T)) {
	for _, v := range s {
		f(v)
	}
}
