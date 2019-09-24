# Slice.Reverse

Reverse returns given arr in reversed order

## Source

```go
// Reverse returns given arr in reversed order
func (s Slice) Reverse() []T {
	result := make([]T, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}
```