package implementation

// Slices is a set of operations to work with slice of slices
type Slices struct {
	Data [][]T
}

// Concat concatenates given slices into a single slice.
func (s Slices) Concat() []T {
	result := make([]T, 0)
	for _, arr := range s.Data {
		result = append(result, arr...)
	}
	return result
}

// Product returns cortesian product of elements
// {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4}
func (s Slices) Product() chan []T {
	c := make(chan []T, 1)
	go s.product(c, []T{}, 0)
	return c
}

// product is a core implementation of Product
func (s Slices) product(c chan []T, left []T, pos int) {
	// iterate over the last array
	if pos == len(s.Data)-1 {
		for _, el := range s.Data[pos] {
			result := make([]T, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.Data[pos] {
		result := make([]T, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

// Zip returns chan of arrays of elements from given arrs on the same position.
func (s Slices) Zip() chan []T {
	if len(s.Data) == 0 {
		result := make(chan []T)
		close(result)
		return result
	}

	size := len(s.Data[0])
	for _, arr := range s.Data[1:] {
		if len(arr) < size {
			size = len(arr)
		}
	}

	result := make(chan []T, 1)
	go func() {
		for i := 0; i < size; i++ {
			chunk := make([]T, 0, len(s.Data))
			for _, arr := range s.Data {
				chunk = append(chunk, arr[i])
			}
			result <- chunk
		}
		close(result)
	}()
	return result
}
