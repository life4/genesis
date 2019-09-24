# Slice.Find

Find returns the first element for which f returns true

## Source

```go
// Find returns the first element for which f returns true
func (s Slice) Find(def T, f func(el T) bool) T {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}
```