# Slice.Sorted

Sorted returns true if slice is sorted

## Source

```go
// Sorted returns true if slice is sorted
func (s Slice) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}
```