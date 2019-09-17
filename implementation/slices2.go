package implementation

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkBy(arr []T, f func(el T) G) [][]T {
	chunks := make([][]T, 0)
	chunk := make([]T, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
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

// DedupBy returns a given slice without consecutive elements
// For which f returns the same result
func DedupBy(arr []T, f func(el T) G) []T {
	result := make([]T, 0, len(arr))

	prev := f(arr[0])
	result = append(result, arr[0])
	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

// GroupBy groups element from array by value returned by f
func GroupBy(arr []T, f func(el T) G) map[G][]T {
	result := make(map[G][]T)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]T, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func Map(arr []T, f func(el T) G) []G {
	result := make([]G, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func Reduce(arr []T, acc G, f func(el T, acc G) G) G {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhile(arr []T, acc G, f func(el T, acc G) (G, error)) (G, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func Scan(arr []T, acc G, f func(el T, acc G) G) []G {
	result := make([]G, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}
