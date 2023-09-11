package sets

// This file contains functions returning not a bool or a set. Usually, a set element.

import (
	c "github.com/life4/genesis/constraints"
)

// Z is a zero type used as the map value for sets.
type Z = struct{}

func Max[S ~map[K]Z, K c.Ordered](set S) (K, error) {
	var max K
	first := true
	for k := range set {
		if first {
			max = k
			first = false
			continue
		}
		if k > max {
			max = k
		}
	}
	if first {
		return max, ErrEmpty
	}
	return max, nil
}

func Min[S ~map[K]Z, K c.Ordered](set S) (K, error) {
	var min K
	first := true
	for k := range set {
		if first {
			min = k
			first = false
			continue
		}
		if k < min {
			min = k
		}
	}
	if first {
		return min, ErrEmpty
	}
	return min, nil
}

// Reduce applies the function to acc and every set element and returns the acc.
func Reduce[S ~map[K]Z, K comparable, R any](set S, acc R, f func(K, R) R) R {
	for k := range set {
		acc = f(k, acc)
	}
	return acc
}

// Sum returns the sum of all elements of the set.
//
// If the set is empty, 0 is returned.
func Sum[S ~map[K]Z, K c.Integer | c.Float](set S) K {
	var result K
	for item := range set {
		result += item
	}
	return result
}

// ToSlice converts the set to a slice.
//
// The order of elements in the resulting set is semi-random.
func ToSlice[S ~map[K]Z, K comparable](set S) []K {
	result := make([]K, 0, len(set))
	for k := range set {
		result = append(result, k)
	}
	return result
}
