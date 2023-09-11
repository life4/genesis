package sets

// This file contains functions returning a set.

// Copy returns a shallow copy of the given set.
func Copy[S ~map[K]Z, K comparable](set S) S {
	result := make(S)
	for k := range set {
		result[k] = Z{}
	}
	return result
}

// Difference returns a set containing elements that appear in the first set but not in the second.
func Difference[S1, S2 ~map[K]Z, K comparable](first S1, second S2) map[K]Z {
	result := make(map[K]Z)
	for k := range first {
		_, found := second[k]
		if !found {
			result[k] = Z{}
		}
	}
	return result
}

// Filter returns elements of the set for which the given function returns true.
func Filter[S ~map[K]Z, K comparable](set S, f func(K) bool) S {
	result := make(S)
	for k := range set {
		if f(k) {
			result[k] = Z{}
		}
	}
	return result
}

// FromSlice converts a slice into a set.
func FromSlice[S ~[]K, K comparable](items S) map[K]Z {
	set := make(map[K]Z)
	for _, item := range items {
		set[item] = Z{}
	}
	return set
}

// Intersection returns a set containing elements that appear in both of the given sets.
func Intersection[S1, S2 ~map[K]Z, K comparable](first S1, second S2) map[K]Z {
	result := make(map[K]Z)
	for k := range second {
		_, found := first[k]
		if found {
			result[k] = Z{}
		}
	}
	return result
}

// Map applies the given function to each element of the set and returns a set of results.
func Map[S ~map[K]Z, K, R comparable](set S, f func(K) R) map[R]Z {
	result := make(map[R]Z)
	for item := range set {
		result[f(item)] = Z{}
	}
	return result
}

// SymmetricDifference returns a set containing elements that appear only in one set but not both.
func SymmetricDifference[S1, S2 ~map[K]Z, K comparable](first S1, second S2) map[K]Z {
	result := make(map[K]Z)
	for k := range first {
		_, found := second[k]
		if !found {
			result[k] = Z{}
		}
	}
	for k := range second {
		_, found := first[k]
		if !found {
			result[k] = Z{}
		}
	}
	return result
}

// Union returns a set containing all elements from both of the given sets.
func Union[S1, S2 ~map[K]Z, K comparable](first S1, second S2) map[K]Z {
	result := make(map[K]Z)
	for k := range first {
		result[k] = Z{}
	}
	for k := range second {
		result[k] = Z{}
	}
	return result
}

// UnionMany returns a set containing all elements from all the given sets.
func UnionMany[S ~map[K]Z, K comparable](sets ...S) map[K]Z {
	result := make(map[K]Z)
	for _, set := range sets {
		for k := range set {
			result[k] = Z{}
		}
	}
	return result
}

// Add
// Discard
// Pop
// Update
// Clear
