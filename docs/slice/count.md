# Slice.Count

Count return count of el occurences in arr.

## Source

```go
// Count return count of el occurences in arr.
func (s Slice) Count(el T) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}
```