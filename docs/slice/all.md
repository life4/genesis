# Slice.All

All returns true if f returns true for all elements in arr

## Source

```go
// All returns true if f returns true for all elements in arr
func (s Slice) All(f func(el T) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}
```