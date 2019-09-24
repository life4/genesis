# Channel.Count

Count return count of el occurences in channel.

## Source

```go
// Count return count of el occurences in channel.
func (c Channel) Count(el T) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}
```