# Sequence.Replicate

Replicate returns channel that produces val n times

## Source

```go
// Replicate returns channel that produces val n times
func (s Sequence) Replicate(val T, n int) chan T {
	c := make(chan T, 1)
	go func() {
		for i := 0; i < n; i++ {
			c <- val
		}
	}()
	return c
}
```