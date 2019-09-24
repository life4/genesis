# Channel.Sum

Sum returns sum of all elements from channel

## Source

```go
// Sum returns sum of all elements from channel
func (c Channel) Sum() T {
	var sum T
	for el := range c.data {
		sum += el
	}
	return sum
}
```