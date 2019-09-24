# Slice.Cycle

Cycle is an infinite loop over slice

## Source

```go
// Cycle is an infinite loop over slice
func (s Slice) Cycle() chan T {
	c := make(chan T, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}
```