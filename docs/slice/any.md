# Slice.Any

Any returns true if f returns true for any element in arr

## Source

```go
// Any returns true if f returns true for any element in arr
func (s Slice) Any(f func(el T) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}
```