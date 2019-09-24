# Sequence.Count

Count is like Range, but infinite

## Source

```go
// Count is like Range, but infinite
func (s Sequence) Count(start T, step T) chan T {
	c := make(chan T, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}
```