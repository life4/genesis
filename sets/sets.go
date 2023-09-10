package sets

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
