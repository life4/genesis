# Channel.Reduce

Reduce applies f to acc and every element from channel and returns acc

## Source

```go
// Reduce applies f to acc and every element from channel and returns acc
func (c Channel) Reduce(acc G, f func(el T, acc G) G) G {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}
```