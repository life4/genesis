package implementation

import (
	"math/rand"
	"time"
)

// Slice is a set of operations to work with slice
type Slice struct {
	data []T
}

// Any returns true if f returns true for any element in arr
func (s Slice) Any(f func(el T) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func (s Slice) All(f func(el T) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

// ChunkBy splits arr on every element for which f returns a new value.
func (s Slice) ChunkBy(f func(el T) G) [][]T {
	chunks := make([][]T, 0)
	chunk := make([]T, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]T, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// ChunkEvery returns slice of slices containing count elements each
func (s Slice) ChunkEvery(count int) [][]T {
	chunks := make([][]T, 0)
	chunk := make([]T, 0)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]T, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Contains returns true if el in arr.
func (s Slice) Contains(el T) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

// Count return count of el occurences in arr.
func (s Slice) Count(el T) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

// Dedup returns a given slice without consecutive duplicated elements
func (s Slice) Dedup() []T {
	result := make([]T, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

// DedupBy returns a given slice without consecutive elements
// For which f returns the same result
func (s Slice) DedupBy(f func(el T) G) []T {
	result := make([]T, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func (s Slice) DropEvery(nth int) []T {
	result := make([]T, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func (s Slice) DropWhile(f func(arr T) bool) []T {
	result := make([]T, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Each calls f for every element from arr
func (s Slice) Each(f func(el T)) {
	for _, el := range s.data {
		f(el)
	}
}

// Filter returns slice of T for which F returned true
func (s Slice) Filter(f func(el T) bool) []T {
	result := make([]T, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Find returns the first element for which f returns true
func (s Slice) Find(def T, f func(el T) bool) T {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func (s Slice) FindIndex(f func(el T) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

// GroupBy groups element from array by value returned by f
func (s Slice) GroupBy(f func(el T) G) map[G][]T {
	result := make(map[G][]T)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]T, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Intersperse inserts el between each element of arr
func (s Slice) Intersperse(el T) []T {
	result := make([]T, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func (s Slice) Map(f func(el T) G) []G {
	result := make([]G, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

// Max returns the maximal element from arr
func (s Slice) Max() T {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func (s Slice) Min() T {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func (s Slice) Reduce(acc G, f func(el T, acc G) G) G {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func (s Slice) ReduceWhile(acc G, f func(el T, acc G) (G, error)) (G, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Reverse returns given arr in reversed order
func (s Slice) Reverse() []T {
	result := make([]T, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

// Same returns true if all element in arr the same
func (s Slice) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

// Scan is like Reduce2, but returns slice of f results
func (s Slice) Scan(acc G, f func(el T, acc G) G) []G {
	result := make([]G, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// Shuffle in random order arr elements
func (s Slice) Shuffle() []T {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

// Split splits arr by sep
func (s Slice) Split(sep T) [][]T {
	result := make([][]T, 0)
	curr := make([]T, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func (s Slice) StartsWith(prefix []T) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

// Sum return sum of all elements from arr
func (s Slice) Sum() T {
	var sum T
	for _, el := range s.data {
		sum += el
	}
	return sum
}

// TakeWhile takes elements from arr while f returns true
func (s Slice) TakeWhile(f func(el T) bool) []T {
	result := make([]T, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Uniq returns arr with only first occurences of every element.
func (s Slice) Uniq() []T {
	added := make(map[T]struct{})
	nothing := struct{}{}
	result := make([]T, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func (s Slice) Window(size int) [][]T {
	result := make([][]T, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}
