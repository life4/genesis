package implementation

// T is a generic type
type T struct{}

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
