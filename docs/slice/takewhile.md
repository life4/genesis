# Slice.TakeWhile

TakeWhile takes elements from arr while f returns true

## Source

```go
// TakeWhile takes elements from arr while f returns true
func (s Slice) TakeWhile(f func(el T) bool) []T {
	result := make([]T, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}
```