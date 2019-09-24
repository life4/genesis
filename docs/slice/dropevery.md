# Slice.DropEvery

DropEvery returns a slice of every nth element in the enumerable dropped, starting with the first element.

## Source

```go
// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func (s Slice) DropEvery(nth int) []T {
	result := make([]T, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}
```