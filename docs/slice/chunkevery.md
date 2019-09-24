# Slice.ChunkEvery

ChunkEvery returns slice of slices containing count elements each

## Source

```go
// ChunkEvery returns slice of slices containing count elements each
func (s Slice) ChunkEvery(count int) [][]T {
	chunks := make([][]T, 0)
	chunk := make([]T, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]T, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}
```