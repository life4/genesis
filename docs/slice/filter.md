# Slice.Filter

Filter returns slice of T for which F returned true

## Source

```go
// Filter returns slice of T for which F returned true
func (s Slice) Filter(f func(el T) bool) []T {
	result := make([]T, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}
```