# Slice.Intersperse

Intersperse inserts el between each element of arr

## Source

```go
// Intersperse inserts el between each element of arr
func (s Slice) Intersperse(el T) []T {
	result := make([]T, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}
```