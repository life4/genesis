# Channel.Map

Map applies f to all elements from channel and returns channel with results

## Source

```go
// Map applies f to all elements from channel and returns channel with results
func (c Channel) Map(f func(el T) G) chan G {
	result := make(chan G, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}
```