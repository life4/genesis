# Slice.ChunkBy

ChunkBy splits arr on every element for which f returns a new value.

## Source

```go
// ChunkBy splits arr on every element for which f returns a new value.
func (s Slice) ChunkBy(f func(el T) G) [][]T {
	chunks := make([][]T, 0)
	chunk := make([]T, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]T, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}
```