# Sequence.Repeat

Repeat returns channel that produces val infinite times

## Source

```go
// Repeat returns channel that produces val infinite times
func (s Sequence) Repeat(val T) chan T {
	c := make(chan T, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}
```