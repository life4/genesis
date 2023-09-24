package iters

import c "github.com/life4/genesis/constraints"

type Iter[T any] interface {
	Next() (T, bool)
}

func FromChannel[T any](ch <-chan T) Iter[T] {
	return &iChannel[T]{ch}
}

type iChannel[T any] struct {
	ch <-chan T
}

func (i *iChannel[T]) Next() (T, bool) {
	v, ok := <-i.ch
	return v, ok
}

func FromSlice[S ~[]T, T any](slice S) Iter[T] {
	return &iSlice[T]{slice, 0}
}

type iSlice[T any] struct {
	slice []T
	next  int
}

func (i *iSlice[T]) Next() (T, bool) {
	if i.next >= len(i.slice) {
		return *new(T), false
	}
	v := i.slice[i.next]
	i.next += 1
	return v, true
}

func Map[T, R any](it Iter[T], f func(T) R) Iter[R] {
	return iMap[T, R]{it, f}
}

type iMap[T, R any] struct {
	iter Iter[T]
	f    func(T) R
}

func (i iMap[T, R]) Next() (R, bool) {
	val, more := i.iter.Next()
	if !more {
		var res R
		return res, false
	}
	res := i.f(val)
	return res, true
}

func Take[T any, I c.Integer](it Iter[T], n I) Iter[T] {
	return &iTake[T, I]{it, n}
}

type iTake[T any, I c.Integer] struct {
	iter Iter[T]
	left I
}

func (i *iTake[T, I]) Next() (T, bool) {
	if i.left <= 0 {
		var val T
		return val, false
	}
	i.left -= 1
	return i.iter.Next()
}

func ToSlice[T any](it Iter[T]) []T {
	res := make([]T, 0)
	for {
		val, more := it.Next()
		if !more {
			return res
		}
		res = append(res, val)
	}
}
