# Slice.Split

Split splits arr by sep

## Source

```go
// Split splits arr by sep
func (s Slice) Split(sep T) [][]T {
	result := make([][]T, 0)
	curr := make([]T, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}
```