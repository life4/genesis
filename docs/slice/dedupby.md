# Slice.DedupBy

DedupBy returns a given slice without consecutive elements For which f returns the same result

## Source

```go
// DedupBy returns a given slice without consecutive elements
// For which f returns the same result
func (s Slice) DedupBy(f func(el T) G) []T {
	result := make([]T, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}
```