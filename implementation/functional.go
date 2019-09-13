package implementation

// T is a generic type
type T int8

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

// Map applies F to all elements in slice of T and returns slice of results
func Map(arr []T, f func(el T) T) []T {
	result := make([]T, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Each calls f for every element from arr
func Each(arr []T, f func(el T)) {
	for _, el := range arr {
		f(el)
	}
}

// Reduce applies F to acc and every element in slice of T and returns acc
func Reduce(arr []T, acc T, f func(el T, acc T) T) T {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func Identity(t T) T { return t }

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
