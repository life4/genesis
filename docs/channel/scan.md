# Channel.Scan

Scan is like Reduce, but returns slice of f results

## Source

```go
// Scan is like Reduce, but returns slice of f results
func (c Channel) Scan(acc G, f func(el T, acc G) G) chan G {
	result := make(chan G, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}
```