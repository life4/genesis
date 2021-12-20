package gnumbers

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
