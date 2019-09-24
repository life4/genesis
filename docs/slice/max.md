# Slice.Max

Max returns the maximal element from arr

## Source

```go
// Max returns the maximal element from arr
func (s Slice) Max() T {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}
```