package sets

// This file contains functions returning bool.

// Contains returns true if the given set contains the given element.
func Contains[S ~map[K]Z, K comparable](set S, item K) bool {
	var found bool
	_, found = set[item]
	return found
}

// Disjoint returns true if the two given sets don't have any common elements.
func Disjoint[S1, S2 ~map[K]Z, K comparable](first S1, second S2) bool {
	for k := range second {
		_, common := first[k]
		if common {
			return false
		}
	}
	return true
}

// DisjointMany returns true if the given sets don't have any common elements.
func DisjointMany[S ~map[K]Z, K comparable](sets ...S) bool {
	union := make(map[K]Z)
	for _, set := range sets {
		for k := range set {
			_, seen := union[k]
			if seen {
				return false
			}
			union[k] = Z{}
		}
	}
	return true
}

// Empty returns true if the set has no elements
func Empty[S ~map[K]Z, K comparable](set S) bool {
	return len(set) == 0
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

// Intersect returns true if the two given sets have at least one common element.
func Intersect[S1, S2 ~map[K]Z, K comparable](first S1, second S2) bool {
	for k := range second {
		_, common := first[k]
		if common {
			return true
		}
	}
	return false
}

// Subset returns true if the first set is a subset of the second one.
//
// One set is called a subset of another if all its elements are included in that set.
//
// This function is the same as [Superset] but with inversed argument order.
// Which one to use is a matter of readability.
func Subset[S1, S2 ~map[K]Z, K comparable](small S1, big S2) bool {
	for k := range small {
		_, found := big[k]
		if !found {
			return false
		}
	}
	return true
}

// Superset returns true if the first set is a superset of the second one.
//
// One set is called a superset of another if it includes all elements of that set.
//
// This function is the same as [Subset] but with inversed argument order.
// Which one to use is a matter of readability.
func Superset[S1, S2 ~map[K]Z, K comparable](big S1, small S2) bool {
	for k := range small {
		_, found := big[k]
		if !found {
			return false
		}
	}
	return true
}
