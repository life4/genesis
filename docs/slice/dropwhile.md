# Slice.DropWhile

DropWhile drops elements from arr while f returns true

## Source

```go
// DropWhile drops elements from arr while f returns true
func (s Slice) DropWhile(f func(arr T) bool) []T {
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