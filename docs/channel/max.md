# Channel.Max

Max returns the maximal element from channel

## Source

```go
// Max returns the maximal element from channel
func (c Channel) Max() T {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}
```