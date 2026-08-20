package genesis

// ComparableSlice is a version of [Slice] with [comparable] elements.
type ComparableSlice[S ~[]T, T comparable] struct {
	Slice[S, T]
}

// NewSlice wraps an existing slice into [Slice].
func NewComparableSlice[S ~[]T, T comparable](s S) ComparableSlice[S, T] {
	return ComparableSlice[S, T]{Slice[S, T]{s}}
}

// Contains returns true if el is in items.
func (s ComparableSlice[S, T]) Contains(el T) bool {
	for _, val := range s.s {
		if val == el {
			return true
		}
	}
	return false
}

// Count return the count of el occurrences in items.
func (s ComparableSlice[S, T]) Count(el T) int {
	count := 0
	for _, val := range s.s {
		if val == el {
			count++
		}
	}
	return count
}

// Dedup returns items without consecutive duplicated elements.
func (s ComparableSlice[S, T]) Dedup() ComparableSlice[S, T] {
	if len(s.s) == 0 {
		return s
	}

	result := make(S, 1, len(s.s))
	prev := s.s[0]
	result[0] = prev
	for _, el := range s.s[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return NewComparableSlice(result)
}

// Delete deletes the first occurrence of the element from items.
func (s ComparableSlice[S, T]) Delete(element T) ComparableSlice[S, T] {
	if len(s.s) == 0 {
		return s
	}

	result := make(S, 0, len(s.s))
	deleted := false
	for _, el := range s.s {
		if !deleted && el == element {
			deleted = true
			continue
		}
		result = append(result, el)
	}
	return NewComparableSlice(result)
}
