# Slice.Contains

Contains returns true if el in arr.

## Source

```go
// Contains returns true if el in arr.
func (s Slice) Contains(el T) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}
```