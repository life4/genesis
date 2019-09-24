# Slices.Concat

Concat concatenates given slices into a single slice.

## Source

```go
// Concat concatenates given slices into a single slice.
func (s Slices) Concat() []T {
	result := make([]T, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}
```