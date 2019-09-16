package implementation

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2(arr []T, f func(el T) G) []G {
	result := make([]G, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce2 applies F to acc and every element in slice of T and returns acc
func Reduce2(arr []T, acc G, f func(el T, acc G) G) G {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile2 is like Reduce, but stops when f returns error
func ReduceWhile2(arr []T, acc G, f func(el T, acc G) (G, error)) (G, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan2 is like Reduce2, but returns slice of f results
func Scan2(arr []T, acc G, f func(el T, acc G) G) []G {
	result := make([]G, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkBy(arr []T, f func(el T) G) [][]T {
	chunks := make([][]T, 0)
	chunk := make([]T, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		chunk = append(chunk, el)
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]T, 0)
			prev = curr
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
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
