# Channel.Filter

Filter returns a new channel with elements from input channel for which f returns true

## Source

```go
// Filter returns a new channel with elements from input channel
// for which f returns true
func (c Channel) Filter(f func(el T) bool) chan T {
	result := make(chan T, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}
```