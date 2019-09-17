package implementation

// Filter returns slice of T for which F returned true
func Filter(arr []T, f func(el T) bool) []T {
	result := make([]T, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Each calls f for every element from arr
func Each(arr []T, f func(el T)) {
	for _, el := range arr {
		f(el)
	}
}

// Any returns true if f returns true for any element in arr
func Any(arr []T, f func(el T) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func All(arr []T, f func(el T) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// ChunkEvery returns slice of slices containing count elements each
func ChunkEvery(arr []T, count int) [][]T {
	chunks := make([][]T, 0)
	chunk := make([]T, 0)
	for i, el := range arr {
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

// Concat concatenates given slices into a single slice.
func Concat(arrs ...[]T) []T {
	result := make([]T, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Contains returns true if el in arr.
func Contains(arr []T, el T) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
}

// Dedup returns a given slice without consecutive duplicated elements
func Dedup(arr []T) []T {
	result := make([]T, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
}

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEvery(arr []T, nth int) []T {
	result := make([]T, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhile(arr []T, f func(arr T) bool) []T {
	result := make([]T, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Find returns the first element for which f returns true
func Find(arr []T, def T, f func(el T) bool) T {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndex(arr []T, f func(el T) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func Intersperse(arr []T, el T) []T {
	result := make([]T, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func Max(arr []T) T {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func Min(arr []T) T {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reverse returns given arr in reversed order
func Reverse(arr []T) []T {
	result := make([]T, 0, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	return result
}

// Same returns true if all element in arr the same
func Same(arr []T) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWith(arr []T, prefix []T) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// TakeWhile takes elements from arr while f returns true
func TakeWhile(arr []T, f func(el T) bool) []T {
	result := make([]T, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func Window(arr []T, size int) [][]T {
	result := make([][]T, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
	}
	return result
}
