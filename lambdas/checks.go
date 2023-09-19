package lambdas

import "github.com/life4/genesis/constraints"

// EqualTo returns lambda that checks if an item is equal to the given value.
func EqualTo[T comparable](a T) func(T) bool {
	return func(b T) bool {
		return a == b
	}
}

// IsDefault checks if the given value is the default for this type.
//
// A few examples:
//
//   - 0 for int and float
//   - "" for string
//   - nil for slice
func IsDefault[T comparable](value T) bool {
	var def T
	return value == def
}

// IsEmpty checks if the given slice is empty
func IsEmpty[T any](items []T) bool {
	return len(items) == 0
}

func IsNaN[T comparable](value T) bool {
	return value != value
}

func IsNil[T any](value *T) bool {
	return value == nil
}

// IsNotDefault checks if the given value is not the default for this type.
//
// A few examples:
//
//   - 0 for int and float
//   - "" for string
//   - nil for slice
func IsNotDefault[T comparable](value T) bool {
	var def T
	return value != def
}

// Empty checks if the given slice is not empty
func IsNotEmpty[T any](items []T) bool {
	return len(items) != 0
}

func IsNotNaN[T comparable](value T) bool {
	return value == value
}

func IsNotNil[T any](value *T) bool {
	return value != nil
}

// IsNotZero checks if the given number is not zero
func IsNotZero[T Number](n T) bool {
	return n != 0
}

// IsZero checks if the given number is zero
func IsZero[T Number](n T) bool {
	return n == 0
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
