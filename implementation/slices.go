package implementation

// Slices is a set of operations to work with slice of slices
type Slices struct {
	data [][]T
}

// Concat concatenates given slices into a single slice.
func (s Slices) Concat() []T {
	result := make([]T, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

// Zip returns array of arrays of elements from given arrs
// on the same position
func (s Slices) Zip() [][]T {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]T, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]T, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}
