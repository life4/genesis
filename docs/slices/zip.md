# Slices.Zip

Zip returns array of arrays of elements from given arrs on the same position

## Source

```go
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
```