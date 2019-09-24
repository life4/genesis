# Slice.Shuffle

Shuffle in random order arr elements

## Source

```go
// Shuffle in random order arr elements
func (s Slice) Shuffle() []T {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}
```