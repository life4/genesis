# Slice.Window

Window makes sliding window for a given slice: ({1,2,3}, 2) -> (1,2), (2,3)

## Source

```go
// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func (s Slice) Window(size int) [][]T {
	result := make([][]T, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}
```