# Slice.Dedup

Dedup returns a given slice without consecutive duplicated elements

## Source

```go
// Dedup returns a given slice without consecutive duplicated elements
func (s Slice) Dedup() []T {
	result := make([]T, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}
```