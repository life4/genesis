package slices

import (
	"constraints"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Any returns true if f returns true for any element in arr
func Any[S ~[]T, T any](items S, f func(el T) bool) bool {
	for _, el := range items {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func All[S ~[]T, T any](items S, f func(el T) bool) bool {
	for _, el := range items {
		if !f(el) {
			return false
		}
	}
	return true
}

// Choice chooses a random element from the slice.
// If seed is zero, UNIX timestamp will be used.
func Choice[S ~[]T, T any](items S, seed int64) (T, error) {
	if len(items) == 0 {
		var tmp T
		return tmp, ErrEmpty
	}

	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rand.Seed(seed)
	i := rand.Intn(len(items))
	return items[i], nil
}

// ChunkEvery returns slice of slices containing count elements each
func ChunkEvery[S ~[]T, T any](items S, count int) ([]S, error) {
	chunks := make([]S, 0)
	if count <= 0 {
		return chunks, ErrNegativeValue
	}
	chunk := make([]T, 0, count)
	for i, el := range items {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]T, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks, nil
}

// Contains returns true if el in arr.
func Contains[S ~[]T, T comparable](items S, el T) bool {
	for _, val := range items {
		if val == el {
			return true
		}
	}
	return false
}

// Count return count of el occurrences in arr.
func Count[S ~[]T, T comparable](items S, el T) int {
	count := 0
	for _, val := range items {
		if val == el {
			count++
		}
	}
	return count
}

// CountFunc returns how many times f returns true.
func CountFunc[S ~[]T, T any](items S, f func(el T) bool) int {
	count := 0
	for _, el := range items {
		if f(el) {
			count++
		}
	}
	return count
}

// Copy creates a copy of the given slice.
func Copy[S ~[]T, T any](items S) S {
	if items == nil {
		return nil
	}
	var res S
	return append(res, items...)
}

// Cycle is an infinite loop over slice
func Cycle[S ~[]T, T any](items S) chan T {
	c := make(chan T)
	go func() {
		defer close(c)
		if len(items) == 0 {
			return
		}
		for {
			for _, val := range items {
				c <- val
			}
		}
	}()
	return c
}

// Dedup returns a given slice without consecutive duplicated elements
func Dedup[S ~[]T, T comparable](items S) S {
	if len(items) == 0 {
		return items
	}

	result := make(S, 1, len(items))
	prev := items[0]
	result[0] = prev
	for _, el := range items[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

// Delete deletes the first occurrence of the element from the slice
func Delete[S ~[]T, T comparable](items S, element T) S {
	if len(items) == 0 {
		return items
	}

	result := make([]T, 0, len(items))
	deleted := false
	for _, el := range items {
		if !deleted && el == element {
			deleted = true
			continue
		}
		result = append(result, el)
	}
	return result

}

// DeleteAll deletes all occurrences of the element from the slice
func DeleteAll[S ~[]T, T comparable](items S, element T) S {
	if len(items) == 0 {
		return items
	}

	result := make([]T, 0, len(items))
	for _, el := range items {
		if el == element {
			continue
		}
		result = append(result, el)
	}
	return result

}

// DeleteAt returns the slice without elements on given positions
func DeleteAt[S ~[]T, T any](items S, indices ...int) (S, error) {
	if len(indices) == 0 || len(items) == 0 {
		return items, nil
	}

	for _, index := range indices {
		if index >= len(items) {
			return items, ErrOutOfRange
		}
	}

	result := make([]T, 0, len(items)-1)
	for i, el := range items {
		add := true
		for _, index := range indices {
			if i == index {
				add = false
				break
			}
		}
		if add {
			result = append(result, el)
		}
	}
	return result, nil
}

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEvery[S ~[]T, T any](items S, nth int, from int) (S, error) {
	if nth <= 0 {
		return items, ErrNonPositiveValue
	}
	if from < 0 {
		return items, ErrNegativeValue
	}
	result := make([]T, 0, len(items)/nth)
	for i, el := range items {
		if (i+from)%nth != 0 {
			result = append(result, el)
		}
	}
	return result, nil
}

// DropWhile drops elements from arr while f returns true
func DropWhile[S ~[]T, T any](items S, f func(el T) bool) S {
	for i, el := range items {
		if !f(el) {
			return items[i:]
		}
	}
	return []T{}
}

// Each calls f for every element from arr
func Each[S ~[]T, T any](items S, f func(el T)) {
	for _, el := range items {
		f(el)
	}
}

// EndsWith returns true if slice ends with the given suffix slice.
// If suffix is empty, it returns true.
func EndsWith[S ~[]T, T comparable](items S, suffix S) bool {
	if len(suffix) > len(items) {
		return false
	}
	start := len(items) - len(suffix)
	for i, el := range suffix {
		if el != items[start+i] {
			return false
		}
	}
	return true
}

// Equal returns true if slices are equal
func Equal[S ~[]T, T comparable](items S, other S) bool {
	if len(items) != len(other) {
		return false
	}
	for i, el := range other {
		if items[i] != el {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func Filter[S ~[]T, T any](items S, f func(el T) bool) S {
	result := make([]T, 0, len(items))
	for _, el := range items {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Reject is like filter but it returns slice of T for which F returned false
func Reject[S ~[]T, T any](items S, f func(el T) bool) S {
	notF := func(el T) bool {
		return !f(el)
	}
	return Filter(items, notF)
}

// Find returns the first element for which f returns true
func Find[S ~[]T, T any](items S, f func(el T) bool) (T, error) {
	for _, el := range items {
		if f(el) {
			return el, nil
		}
	}
	var tmp T
	return tmp, ErrNotFound
}

// FindIndex is like Find, but return element index instead of element itself.
// Returns -1 if element not found
func FindIndex[S ~[]T, T any](items S, f func(el T) bool) int {
	for i, el := range items {
		if f(el) {
			return i
		}
	}
	return -1
}

// Grow increases the slice's by n elements.
// So, for cap(slice)=8 and n=2, the result will have cap at least 10.
// The function can be used to reduce allocations when inserting more elements
// into an existing slice.
func Grow[S ~[]T, T any](items S, n int) S {
	return append(items, make(S, n)...)[:len(items)]
}

// Shrink removes unused capacity from the slice.
func Shrink[S ~[]T, T any](items S) S {
	return items[:len(items):len(items)]
}

// Join concatenates elements of the slice to create a single string.
func Join[S ~[]T, T any](items S, sep string) string {
	strs := make([]string, 0, len(items))
	for _, el := range items {
		strs = append(strs, fmt.Sprintf("%v", el))
	}
	return strings.Join(strs, sep)
}

// Index returns the index of the first occurrence of item in items.
func Index[S []T, T comparable](items S, item T) (int, error) {
	for i, val := range items {
		if val == item {
			return i, nil
		}
	}
	return 0, ErrNotFound
}

// IndexFunc returns the first index in items for which f returns true.
func IndexFunc[S []T, T comparable](items S, f func(T) bool) (int, error) {
	for i, val := range items {
		if f(val) {
			return i, nil
		}
	}
	return 0, ErrNotFound
}

// InsertAt returns the items slice with the item inserted at the given index.
func InsertAt[S ~[]T, T any](items S, index int, item T) (S, error) {
	result := make([]T, 0, len(items)+1)

	// insert at the end
	if index == len(items) {
		result = append(result, items...)
		result = append(result, item)
		return result, nil
	}

	if index > len(items) {
		return items, ErrOutOfRange
	}
	if index < 0 {
		return items, ErrNegativeValue
	}

	for i, el := range items {
		if i == index {
			result = append(result, item)
		}
		result = append(result, el)
	}
	return result, nil
}

// Intersperse inserts el between each element of arr
func Intersperse[S ~[]T, T any](items S, el T) S {
	if len(items) == 0 {
		return items
	}
	result := make([]T, 0, len(items)*2-1)
	result = append(result, items[0])
	for _, val := range items[1:] {
		result = append(result, el, val)
	}
	return result
}

// Last returns the last element from the slice
func Last[S ~[]T, T any](items S) (T, error) {
	if len(items) == 0 {
		var tmp T
		return tmp, ErrEmpty
	}
	return items[len(items)-1], nil
}

// Max returns the maximal element from arr
func Max[S ~[]T, T constraints.Ordered](items S) (T, error) {
	if len(items) == 0 {
		var tmp T
		return tmp, ErrEmpty
	}

	max := items[0]
	for _, el := range items[1:] {
		if el > max {
			max = el
		}
	}
	return max, nil
}

// Min returns the minimal element from arr
func Min[S ~[]T, T constraints.Ordered](items S) (T, error) {
	if len(items) == 0 {
		var tmp T
		return tmp, ErrEmpty
	}

	min := items[0]
	for _, el := range items[1:] {
		if el < min {
			min = el
		}
	}
	return min, nil
}

// Permutations returns successive size-length permutations of elements from the slice.
// {1, 2, 3} -> {1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}
func Permutations[T any](items []T, size int) chan []T {
	c := make(chan []T, 1)
	go func() {
		if len(items) > 0 {
			permutations(items, c, size, []T{}, items)
		}
		close(c)
	}()
	return c
}

// permutations is a core implementation for Permutations
func permutations[T any](items []T, c chan []T, size int, left []T, right []T) {
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
		permutations(items, c, size, newLeft, newRight)
	}
}

// Product returns cortesian product of elements
// {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4}
func Product[S ~[]T, T any](items S, repeat int) chan []T {
	c := make(chan []T, 1)
	go func() {
		defer close(c)
		if repeat < 1 {
			return
		}
		product(items, c, repeat, []T{}, 0)
	}()
	return c
}

// product is a core implementation for Product
func product[S ~[]T, T any](items S, c chan []T, repeat int, left []T, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range items {
			result := make([]T, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range items {
		result := make([]T, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		product(items, c, repeat, result, pos+1)
	}
}

// Reverse returns given arr in reversed order
func Reverse[S ~[]T, T any](items S) S {
	if len(items) <= 1 {
		return items
	}
	result := make([]T, 0, len(items))
	for i := len(items) - 1; i >= 0; i-- {
		result = append(result, items[i])
	}
	return result
}

// Repeat repeats items slice n times.
func Repeat[S ~[]T, T any](items S, n int) S {
	result := make([]T, 0, len(items)*n)
	for i := 0; i < n; i++ {
		result = append(result, items...)
	}
	return result
}

// Same returns true if all element in arr the same
func Same[S ~[]T, T comparable](items S) bool {
	if len(items) <= 1 {
		return true
	}
	for i := 0; i < len(items)-1; i++ {
		if items[i] != items[i+1] {
			return false
		}
	}
	return true
}

// Shuffle in random order arr elements
func Shuffle[S ~[]T, T any](items S, seed int64) S {
	if len(items) <= 1 {
		return items
	}
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rand.Seed(seed)
	swap := func(i, j int) {
		items[i], items[j] = items[j], items[i]
	}
	rand.Shuffle(len(items), swap)
	return items
}

// Sort returns sorted slice
func Sort[S ~[]T, T constraints.Ordered](items S) S {
	if len(items) <= 1 {
		return items
	}
	less := func(i int, j int) bool {
		return items[i] < items[j]
	}
	sort.SliceStable(items, less)
	return items
}

// Sorted returns true if slice is sorted
func Sorted[S ~[]T, T constraints.Ordered](items S) bool {
	if len(items) <= 1 {
		return true
	}
	for i := 1; i < len(items); i++ {
		if items[i-1] > items[i] {
			return false
		}
	}
	return true
}

// Split splits arr by sep
func Split[S ~[]T, T comparable](items S, sep T) []S {
	result := make([]S, 0)
	curr := make([]T, 0)
	for _, el := range items {
		if el == sep {
			result = append(result, curr)
			curr = make([]T, 0)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWith[S ~[]T, T comparable](items S, prefix S) bool {
	if len(prefix) > len(items) {
		return false
	}
	for i, el := range prefix {
		if el != items[i] {
			return false
		}
	}
	return true
}

// Sum return sum of all elements from arr
func Sum[S ~[]T, T constraints.Ordered](items S) T {
	var sum T
	for _, el := range items {
		sum += el
	}
	return sum
}

// TakeEvery returns slice of every nth elements
func TakeEvery[S ~[]T, T any](items S, nth int, from int) (S, error) {
	if nth <= 0 {
		return items, ErrNonPositiveValue
	}
	if from < 0 {
		return items, ErrNegativeValue
	}
	result := make(S, 0, len(items))
	for i, el := range items {
		if (i+from)%nth == 0 {
			result = append(result, el)
		}
	}
	return result, nil
}

// TakeRandom returns slice of count random elements from the slice
func TakeRandom[S ~[]T, T any](items S, count int, seed int64) (S, error) {
	if count > len(items) {
		return nil, ErrOutOfRange
	}
	if count <= 0 {
		return nil, ErrNonPositiveValue
	}

	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rand.Seed(seed)
	swap := func(i, j int) {
		items[i], items[j] = items[j], items[i]
	}
	rand.Shuffle(len(items), swap)
	return items[:count], nil
}

// TakeWhile takes elements from arr while f returns true
func TakeWhile[S ~[]T, T any](items S, f func(el T) bool) S {
	result := make([]T, 0, len(items))
	for _, el := range items {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// ToChannel returns channel with elements from the slice
func ToChannel[S ~[]T, T any](items S) chan T {
	c := make(chan T)
	go func() {
		for _, el := range items {
			c <- el
		}
		close(c)
	}()
	return c
}

// Uniq returns arr with only first occurrences of every element.
func Uniq[S ~[]T, T comparable](items S) S {
	if len(items) <= 1 {
		return items
	}
	added := make(map[T]struct{})
	nothing := struct{}{}
	result := make([]T, 0, len(items))
	for _, el := range items {
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
func Window[S ~[]T, T any](items S, size int) ([]S, error) {
	if size <= 0 {
		return nil, ErrNonPositiveValue
	}
	result := make([]S, 0, len(items)/size)
	for i := 0; i <= len(items)-size; i++ {
		chunk := items[i : i+size]
		result = append(result, chunk)
	}
	return result, nil
}

// Without returns the slice with filtered out element
func Without[S ~[]T, T comparable](items S, elements ...T) S {
	result := make(S, 0, len(items))
	for _, el := range items {
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

// Wrap makes a single element slice out of the given value
func Wrap[T any](item T) []T {
	return []T{item}
}
