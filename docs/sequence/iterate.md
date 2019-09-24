# Sequence.Iterate

Iterate returns an infinite list of repeated applications of f to val

## Source

```go
// Iterate returns an infinite list of repeated applications of f to val
func (s Sequence) Iterate(val T, f func(val T) T) chan T {
	c := make(chan T, 1)
	go func() {
		for {
			c <- val
			val = f(val)
		}
	}()
	return c
}
```