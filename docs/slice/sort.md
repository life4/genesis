# Slice.Sort

Sort returns sorted slice

## Source

```go
// Sort returns sorted slice
func (s Slice) Sort() []T {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}
```