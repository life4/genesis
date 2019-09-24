# Sequence.Exponential

Exponential generates elements from start with multiplication of value on factor on every step

## Source

```go
// Exponential generates elements from start with
// multiplication of value on factor on every step
func (s Sequence) Exponential(start T, factor T) chan T {
	c := make(chan T, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}
```