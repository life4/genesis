package sets

import (
	c "github.com/life4/genesis/constraints"
)

// Z is a zero type used as the map value for sets.
type Z = struct{}

func Contains[S ~map[K]Z, K comparable](set S, item K) bool {
	var found bool
	_, found = set[item]
	return found
}

// Copy returns a shallow copy of the given set.
func Copy[S ~map[K]Z, K comparable](set S) S {
	result := make(S)
	for k := range set {
		result[k] = Z{}
	}
	return result
}

// Equal returns true if both given sets have the same elements.
func Equal[S1, S2 ~map[K]Z, K comparable](first S1, second S2) bool {
	if len(first) != len(second) {
		return false
	}
	for k := range second {
		_, found := first[k]
		if !found {
			return false
		}
	}
	return true
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

func Sum[S ~map[K]Z, K c.Integer | c.Float](set S) K {
	var result K
	for item := range set {
		result += item
	}
	return result
}
