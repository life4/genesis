# Sequence.Range

Range generates elements from start to end with given step

## Source

```go
// Range generates elements from start to end with given step
func (s Sequence) Range(start T, end T, step T) chan T {
	c := make(chan T, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}
```