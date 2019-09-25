package implementation

import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Slice is a set of operations to work with slice
type Slice struct {
	Data []T
}

// Any returns true if f returns true for any element in arr
func (s Slice) Any(f func(el T) bool) bool {
	for _, el := range s.Data {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func (s Slice) All(f func(el T) bool) bool {
	for _, el := range s.Data {
		if !f(el) {
			return false
		}
	}
	return true
}

// ChunkBy splits arr on every element for which f returns a new value.
func (s Slice) ChunkBy(f func(el T) G) [][]T {
	chunks := make([][]T, 0)
	if len(s.Data) == 0 {
		return chunks
	}

	chunk := make([]T, 0)
	prev := f(s.Data[0])
	chunk = append(chunk, s.Data[0])

	for _, el := range s.Data[1:] {
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
	if count <= 0 {
		count = 1
	}
	chunks := make([][]T, 0)
	chunk := make([]T, 0, count)
	for i, el := range s.Data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]T, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Contains returns true if el in arr.
func (s Slice) Contains(el T) bool {
	for _, val := range s.Data {
		if val == el {
			return true
		}
	}
	return false
}

// Count return count of el occurences in arr.
func (s Slice) Count(el T) int {
	count := 0
	for _, val := range s.Data {
		if val == el {
			count++
		}
	}
	return count
}

// CountBy returns how many times f returns true.
func (s Slice) CountBy(f func(el T) bool) int {
	count := 0
	for _, el := range s.Data {
		if f(el) {
			count++
		}
	}
	return count
}

// Cycle is an infinite loop over slice
func (s Slice) Cycle() chan T {
	c := make(chan T, 1)
	go func() {
		defer close(c)
		if len(s.Data) == 0 {
			return
		}
		for {
			for _, val := range s.Data {
				c <- val
			}
		}
	}()
	return c
}

// Dedup returns a given slice without consecutive duplicated elements
func (s Slice) Dedup() []T {
	result := make([]T, 0, len(s.Data))
	if len(s.Data) == 0 {
		return result
	}

	prev := s.Data[0]
	result = append(result, prev)
	for _, el := range s.Data[1:] {
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
	result := make([]T, 0, len(s.Data))
	if len(s.Data) == 0 {
		return result
	}

	prev := f(s.Data[0])
	result = append(result, s.Data[0])
	for _, el := range s.Data[1:] {
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
	result := make([]T, 0, len(s.Data)/nth)
	for i, el := range s.Data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func (s Slice) DropWhile(f func(arr T) bool) []T {
	result := make([]T, 0, len(s.Data))
	for _, el := range s.Data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Each calls f for every element from arr
func (s Slice) Each(f func(el T)) {
	for _, el := range s.Data {
		f(el)
	}
}

// Equal returns true if slices are equal
func (s Slice) Equal(other []T) bool {
	if len(s.Data) != len(other) {
		return false
	}
	for i, el := range other {
		if s.Data[i] != el {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func (s Slice) Filter(f func(el T) bool) []T {
	result := make([]T, 0, len(s.Data))
	for _, el := range s.Data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Find returns the first element for which f returns true
func (s Slice) Find(def T, f func(el T) bool) T {
	for _, el := range s.Data {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func (s Slice) FindIndex(f func(el T) bool) int {
	for i, el := range s.Data {
		if f(el) {
			return i
		}
	}
	return -1
}

// Join concatenates elements of the slice to create a single string.
func (s Slice) Join(sep string) string {
	strs := make([]string, 0, len(s.Data))
	for _, el := range s.Data {
		strs = append(strs, string(el))
	}
	return strings.Join(strs, sep)
}

// GroupBy groups element from array by value returned by f
func (s Slice) GroupBy(f func(el T) G) map[G][]T {
	result := make(map[G][]T)
	for _, el := range s.Data {
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
	if len(s.Data) == 0 {
		tmp := make([]T, 0)
		return tmp
	}
	result := make([]T, 0, len(s.Data)*2-1)
	result = append(result, s.Data[0])
	for _, val := range s.Data[1:] {
		result = append(result, el, val)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func (s Slice) Map(f func(el T) G) []G {
	result := make([]G, 0, len(s.Data))
	for _, el := range s.Data {
		result = append(result, f(el))
	}
	return result
}

// Max returns the maximal element from arr
func (s Slice) Max() T {
	max := s.Data[0]
	for _, el := range s.Data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func (s Slice) Min() T {
	min := s.Data[0]
	for _, el := range s.Data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Permutations returns successive size-length permutations of elements from the slice.
// {1, 2, 3} -> {1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}
func (s Slice) Permutations(size int) chan []T {
	c := make(chan []T, 1)
	go func() {
		if len(s.Data) > 0 {
			s.permutations(c, size, []T{}, s.Data)
		}
		close(c)
	}()
	return c
}

// permutations is a core implementation for Permutations
func (s Slice) permutations(c chan []T, size int, left []T, right []T) {
	if len(left) == size || len(right) == 0 {
		c <- left
		return
	}

	for i, el := range right {
		newLeft := make([]T, 0, len(left)+1)
		newLeft = append(newLeft, left...)
		newLeft = append(newLeft, el)

		newRight := make([]T, 0, len(right)-1)
		for j, other := range right {
			if j != i {
				newRight = append(newRight, other)
			}
		}
		s.permutations(c, size, newLeft, newRight)
	}
}

// Product returns cortesian product of elements
// {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4}
func (s Slice) Product(repeat int) chan []T {
	c := make(chan []T, 1)
	go s.product(c, repeat, []T{}, 0)
	return c
}

// product is a core implementation for Product
func (s Slice) product(c chan []T, repeat int, left []T, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.Data {
			result := make([]T, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.Data {
		result := make([]T, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

// Reduce applies F to acc and every element in slice of T and returns acc
func (s Slice) Reduce(acc G, f func(el T, acc G) G) G {
	for _, el := range s.Data {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func (s Slice) ReduceWhile(acc G, f func(el T, acc G) (G, error)) (G, error) {
	for _, el := range s.Data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Reverse returns given arr in reversed order
func (s Slice) Reverse() []T {
	result := make([]T, 0, len(s.Data))
	for i := len(s.Data) - 1; i >= 0; i-- {
		result = append(result, s.Data[i])
	}
	return result
}

// Same returns true if all element in arr the same
func (s Slice) Same() bool {
	for i := 0; i < len(s.Data)-1; i++ {
		if s.Data[i] != s.Data[i+1] {
			return false
		}
	}
	return true
}

// Scan is like Reduce, but returns slice of f results
func (s Slice) Scan(acc G, f func(el T, acc G) G) []G {
	result := make([]G, 0, len(s.Data))
	for _, el := range s.Data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// Shuffle in random order arr elements
func (s Slice) Shuffle() []T {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
	}
	rand.Shuffle(len(s.Data), swap)
	return s.Data
}

// Sort returns sorted slice
func (s Slice) Sort() []T {
	less := func(i int, j int) bool {
		return s.Data[i] < s.Data[j]
	}
	sort.SliceStable(s.Data, less)
	return s.Data
}

// Sorted returns true if slice is sorted
func (s Slice) Sorted() bool {
	for i := 1; i < len(s.Data); i++ {
		if s.Data[i-1] > s.Data[i] {
			return false
		}
	}
	return true
}

// Split splits arr by sep
func (s Slice) Split(sep T) [][]T {
	result := make([][]T, 0)
	curr := make([]T, 0)
	for _, el := range s.Data {
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
	if len(prefix) > len(s.Data) {
		return false
	}
	for i, el := range prefix {
		if el != s.Data[i] {
			return false
		}
	}
	return true
}

// Sum return sum of all elements from arr
func (s Slice) Sum() T {
	var sum T
	for _, el := range s.Data {
		sum += el
	}
	return sum
}

// TakeWhile takes elements from arr while f returns true
func (s Slice) TakeWhile(f func(el T) bool) []T {
	result := make([]T, 0, len(s.Data))
	for _, el := range s.Data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// ToChannel returns channel with elements from the slice
func (s Slice) ToChannel() chan T {
	c := make(chan T, 1)
	go func() {
		for _, el := range s.Data {
			c <- el
		}
		close(c)
	}()
	return c
}

// Uniq returns arr with only first occurences of every element.
func (s Slice) Uniq() []T {
	added := make(map[T]struct{})
	nothing := struct{}{}
	result := make([]T, 0, len(s.Data))
	for _, el := range s.Data {
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
	result := make([][]T, 0, len(s.Data)/size)
	for i := 0; i <= len(s.Data)-size; i++ {
		chunk := s.Data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

// Without returns the slice with filtered out element
func (s Slice) Without(elements ...T) []T {
	result := make([]T, 0, len(s.Data))
	for _, el := range s.Data {
		allowed := true
		for _, other := range elements {
			if el == other {
				allowed = false
			}
		}
		if allowed {
			result = append(result, el)
		}
	}
	return result
}
