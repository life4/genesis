# Channel.Take

Take takes first n elements from channel c.

## Source

```go
// Take takes first n elements from channel c.
func (c Channel) Take(n int) []T {
	result := make([]T, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}
```