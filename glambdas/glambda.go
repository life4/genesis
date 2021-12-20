package glambdas

import "constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Abs returns positive value
func Abs[T Number](a T) T {
	if a <= 0 {
		return -a
	}
	return a
}

// Min returns minimal value
func Min[T constraints.Ordered](a T, b T) T {
	if a <= b {
		return a
	}
	return b
}

// Max returns maximal value
func Max[T constraints.Ordered](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

// EqualTo returns lambda that checks if an item is equal to the given value.
func EqualTo[T comparable](a T) func(T) bool {
	return func(b T) bool {
		return a == b
	}
}

// LessThan returns lambda that checks if an item is less than the given value.
func LessThan[T constraints.Ordered](a T) func(T) bool {
	return func(b T) bool {
		return b < a
	}
}

// Not negates the result of a lambda
func Not[T constraints.Ordered](f func(T) bool) func(T) bool {
	return func(item T) bool {
		return !f(item)
	}
}

// Zero checks if the given number is zero
func Zero[T Number](n T) bool {
	return n == 0
}

// Empty checks if the given slcie is empty
func Empty[T any](items []T) bool {
	return len(items) == 0
}
