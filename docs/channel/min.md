# Channel.Min

Min returns the minimal element from channel

## Source

```go
// Min returns the minimal element from channel
func (c Channel) Min() T {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}
```