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

// FromSlice converts a slice into a set.
func FromSlice[S ~[]K, K comparable](items S) map[K]Z {
	set := make(map[K]Z)
	for _, item := range items {
		set[item] = Z{}
	}
	return set
}

// Map applies the given function to each element of the set and returns a set of results.
func Map[S ~map[K]Z, K, R comparable](set S, f func(K) R) map[R]Z {
	result := make(map[R]Z)
	for item := range set {
		result[f(item)] = Z{}
	}
	return result
}

// Union returns a set containing all elements from all the given sets.
func UnionMany[S ~map[K]Z, K comparable](sets ...S) map[K]Z {
	result := make(map[K]Z)
	for _, set := range sets {
		for k := range set {
			result[k] = Z{}
		}
	}
	return result
}

// Union
// Intersection
// Difference
// SymmetricDifference
// EqualMany
// Filter
// Reduce

// Add
// Discard
// Pop
// Update
// Clear