# Channel.Each

Each calls f for every element in the channel

## Source

```go
// Each calls f for every element in the channel
func (c Channel) Each(f func(el T)) {
	for el := range c.data {
		f(el)
	}
}
```