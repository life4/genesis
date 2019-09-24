# Slice.StartsWith

StartsWith returns true if slice starts with the given prefix slice. If prefix is empty, it returns true.

## Source

```go
// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func (s Slice) StartsWith(prefix []T) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}
```