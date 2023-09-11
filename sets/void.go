package sets

// The file contains functions that don't return anything (or only return an error).

// Add adds the given element in the set.
//
// If the element already in the set, nothing happens.
//
// The set is modified in place.
func Add[S ~map[K]Z, K comparable](set S, value K) {
	set[value] = Z{}
}

// Clear removes all elements from the set.
//
// The set is modified in place.
func Clear[S ~map[K]Z, K comparable](set S) {
	for k := range set {
		delete(set, k)
	}
}

// Discard removes the given element from the set.
//
// If the element already not in the set, nothing happens.
//
// The set is modified in place.
func Discard[S ~map[K]Z, K comparable](set S, value K) {
	delete(set, value)
}

// Pop removes an element from the set and returns that element.
//
// If the set is empty, [ErrEmpty] is returned.
//
// The set is modified in place.
func Pop[S ~map[K]Z, K comparable](set S) (K, error) {
	for k := range set {
		delete(set, k)
		return k, nil
	}
	var k K
	return k, ErrEmpty
}

// Update adds elements from the values set into the target set.
//
// The set is modified in place.
func Update[S1, S2 ~map[K]Z, K comparable](target S1, values S2) {
	for k := range values {
		target[k] = Z{}
	}
}
