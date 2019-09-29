package implementation

// Pair is a set of functions for 2 values that you can pass into reduce-like funcs
type Pair struct {
	// empty
}

// Min returns minimal value
func (Pair) Min(a T, b T) T {
	if a <= b {
		return a
	}
	return b
}

// Max returns maximal value
func (Pair) Max(a T, b T) T {
	if a > b {
		return a
	}
	return b
}
