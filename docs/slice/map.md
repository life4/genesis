# Slice.Map

Map applies F to all elements in slice of T and returns slice of results

## Source

```go
// Map applies F to all elements in slice of T and returns slice of results
func (s Slice) Map(f func(el T) G) []G {
	result := make([]G, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}
```