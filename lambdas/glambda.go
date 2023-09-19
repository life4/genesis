package lambdas

import "github.com/life4/genesis/constraints"

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

// Default returns the default value of the same type as the given value.
//
// A few examples:
//
//   - 0 for int and float
//   - "" for string
//   - nil for slice
func Default[T any](T) T {
	var def T
	return def
}

// Max returns maximal value
func Max[T constraints.Ordered](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min returns minimal value
func Min[T constraints.Ordered](a T, b T) T {
	if a <= b {
		return a
	}
	return b
}
