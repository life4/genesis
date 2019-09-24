# Channel.All

All returns true if f returns true for all elements in channel

## Source

```go
// All returns true if f returns true for all elements in channel
func (c Channel) All(f func(el T) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}
```