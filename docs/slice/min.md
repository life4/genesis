# Slice.Min

Min returns the minimal element from arr

## Source

```go
// Min returns the minimal element from arr
func (s Slice) Min() T {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}
```