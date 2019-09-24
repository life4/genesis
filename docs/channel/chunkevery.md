# Channel.ChunkEvery

ChunkEvery returns channel with slices containing count elements each

## Source

```go
// ChunkEvery returns channel with slices containing count elements each
func (c Channel) ChunkEvery(count int) chan []T {
	chunks := make(chan []T, 1)
	go func() {
		chunk := make([]T, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]T, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}
```