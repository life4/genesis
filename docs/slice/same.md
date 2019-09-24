# Slice.Same

Same returns true if all element in arr the same

## Source

```go
// Same returns true if all element in arr the same
func (s Slice) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}
```