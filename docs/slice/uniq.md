# Slice.Uniq

Uniq returns arr with only first occurences of every element.

## Source

```go
// Uniq returns arr with only first occurences of every element.
func (s Slice) Uniq() []T {
	added := make(map[T]struct{})
	nothing := struct{}{}
	result := make([]T, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}
```