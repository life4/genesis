package gnumbers

import "constraints"

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
