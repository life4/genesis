# Slice.FindIndex

FindIndex is like Find, but return element index instead of element itself Returns -1 if element is not found

## Source

```go
// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func (s Slice) FindIndex(f func(el T) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}
```