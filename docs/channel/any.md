# Channel.Any

Any returns true if f returns true for any element in channel

## Source

```go
// Any returns true if f returns true for any element in channel
func (c Channel) Any(f func(el T) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}
```