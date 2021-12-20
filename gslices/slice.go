package gslices

import (
	"constraints"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/life4/genesis/gerrors"
)

// Any returns true if f returns true for any element in arr
func Any[T any](items []T, f func(el T) bool) bool {
	for _, el := range items {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func All[T any](items []T, f func(el T) bool) bool {
	for _, el := range items {
		if !f(el) {
			return false
		}
	}
	return true
}

// Choice chooses a random element from the slice.
// If seed is zero, UNIX timestamp will be used.
func Choice[T any](items []T, seed int64) (T, error) {
	if len(items) == 0 {
		var tmp T
		return tmp, gerrors.ErrEmpty
	}

	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rand.Seed(seed)
	i := rand.Intn(len(items))
	return items[i], nil
}

// ChunkEvery returns slice of slices containing count elements each
func ChunkEvery[T any](items []T, count int) ([][]T, error) {
	chunks := make([][]T, 0)
	if count <= 0 {
		return chunks, gerrors.ErrNegativeValue
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
func Contains[T comparable](items []T, el T) bool {
	for _, val := range items {
		if val == el {
			return true
		}
	}
	return false
}

// Count return count of el occurrences in arr.
func Count[T comparable](items []T, el T) int {
	count := 0
	for _, val := range items {
		if val == el {
			count++
		}
	}
	return count
}

// CountBy returns how many times f returns true.
func CountBy[T any](items []T, f func(el T) bool) int {
	count := 0
	for _, el := range items {
		if f(el) {
			count++
		}
	}
	return count
}

// Cycle is an infinite loop over slice
func Cycle[T any](items []T) chan T {
	c := make(chan T, 1)
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
func Dedup[T comparable](items []T) []T {
	if len(items) == 0 {
		return items
	}

	result := make([]T, 0, len(items))
	prev := items[0]
	result = append(result, prev)
	for _, el := range items[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

// Delete deletes the first occurrence of the element from the slice
func Delete[T comparable](items []T, element T) []T {
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
func DeleteAll[T comparable](items []T, element T) []T {
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
func DeleteAt[T any](items []T, indices ...int) ([]T, error) {
	if len(indices) == 0 || len(items) == 0 {
		return items, nil
	}

	for _, index := range indices {
		if index >= len(items) {
			return items, gerrors.ErrOutOfRange
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
func DropEvery[T any](items []T, nth int, from int) ([]T, error) {
	if nth <= 0 {
		return items, gerrors.ErrNonPositiveValue
	}
	if from < 0 {
		return items, gerrors.ErrNegativeValue
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
func DropWhile[T any](items []T, f func(el T) bool) []T {
	for i, el := range items {
		if !f(el) {
			return items[i:]
		}
	}
	return []T{}
}

// Each calls f for every element from arr
func Each[T any](items []T, f func(el T)) {
	for _, el := range items {
		f(el)
	}
}

// EndsWith returns true if slice ends with the given suffix slice.
// If suffix is empty, it returns true.
func EndsWith[T comparable](items []T, suffix []T) bool {
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
func Equal[T comparable](items []T, other []T) bool {
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
func Filter[T any](items []T, f func(el T) bool) []T {
	result := make([]T, 0, len(items))
	for _, el := range items {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Find returns the first element for which f returns true
func Find[T any](items []T, f func(el T) bool) (T, error) {
	for _, el := range items {
		if f(el) {
			return el, nil
		}
	}
	var tmp T
	return tmp, gerrors.ErrNotFound
}

// FindIndex is like Find, but return element index instead of element itself.
// Returns -1 if element not found
func FindIndex[T any](items []T, f func(el T) bool) int {
	for i, el := range items {
		if f(el) {
			return i
		}
	}
	return -1
}

// Join concatenates elements of the slice to create a single string.
func Join[T any](items []T, sep string) string {
	strs := make([]string, 0, len(items))
	for _, el := range items {
		strs = append(strs, fmt.Sprintf("%v", el))
	}
	return strings.Join(strs, sep)
}

// InsertAt returns the slice with element inserted at given index.
func InsertAt[T any](items []T, index int, element T) ([]T, error) {
	result := make([]T, 0, len(items)+1)

	// insert at the end
	if index == len(items) {
		result = append(result, items...)
		result = append(result, element)
		return result, nil
	}

	if index > len(items) {
		return items, gerrors.ErrOutOfRange
	}
	if index < 0 {
		return items, gerrors.ErrNegativeValue
	}

	for i, el := range items {
		if i == index {
			result = append(result, element)
		}
		result = append(result, el)
	}
	return result, nil
}

// Intersperse inserts el between each element of arr
func Intersperse[T any](items []T, el T) []T {
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
func Last[T any](items []T) (T, error) {
	if len(items) == 0 {
		var tmp T
		return tmp, gerrors.ErrEmpty
	}
	return items[len(items)-1], nil
}

// Max returns the maximal element from arr
func Max[T constraints.Ordered](items []T) (T, error) {
	if len(items) == 0 {
		var tmp T
		return tmp, gerrors.ErrEmpty
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
func Min[T constraints.Ordered](items []T) (T, error) {
	if len(items) == 0 {
		var tmp T
		return tmp, gerrors.ErrEmpty
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
func Product[T any](items []T, repeat int) chan []T {
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
func product[T any](items []T, c chan []T, repeat int, left []T, pos int) {
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
func Reverse[T any](items []T) []T {
	if len(items) <= 1 {
		return items
	}
	result := make([]T, 0, len(items))
	for i := len(items) - 1; i >= 0; i-- {
		result = append(result, items[i])
	}
	return result
}

// Same returns true if all element in arr the same
func Same[T comparable](items []T) bool {
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
func Shuffle[T any](items []T, seed int64) []T {
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
func Sort[T constraints.Ordered](items []T) []T {
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
func Sorted[T constraints.Ordered](items []T) bool {
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
func Split[T comparable](items []T, sep T) [][]T {
	result := make([][]T, 0)
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
func StartsWith[T comparable](items []T, prefix []T) bool {
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
func Sum[T constraints.Ordered](items []T) T {
	var sum T
	for _, el := range items {
		sum += el
	}
	return sum
}

// TakeEvery returns slice of every nth elements
func TakeEvery[T any](items []T, nth int, from int) ([]T, error) {
	if nth <= 0 {
		return items, gerrors.ErrNonPositiveValue
	}
	if from < 0 {
		return items, gerrors.ErrNegativeValue
	}
	result := make([]T, 0, len(items))
	for i, el := range items {
		if (i+from)%nth == 0 {
			result = append(result, el)
		}
	}
	return result, nil
}

// TakeRandom returns slice of count random elements from the slice
func TakeRandom[T any](items []T, count int, seed int64) ([]T, error) {
	if count > len(items) {
		return nil, gerrors.ErrOutOfRange
	}
	if count <= 0 {
		return nil, gerrors.ErrNonPositiveValue
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
func TakeWhile[T any](items []T, f func(el T) bool) []T {
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
func ToChannel[T any](items []T) chan T {
	c := make(chan T, 1)
	go func() {
		for _, el := range items {
			c <- el
		}
		close(c)
	}()
	return c
}

// Uniq returns arr with only first occurrences of every element.
func Uniq[T comparable](items []T) []T {
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
func Window[T any](items []T, size int) ([][]T, error) {
	if size <= 0 {
		return nil, gerrors.ErrNonPositiveValue
	}
	result := make([][]T, 0, len(items)/size)
	for i := 0; i <= len(items)-size; i++ {
		chunk := items[i : i+size]
		result = append(result, chunk)
	}
	return result, nil
}

// Without returns the slice with filtered out element
func Without[T comparable](items []T, elements ...T) []T {
	result := make([]T, 0, len(items))
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