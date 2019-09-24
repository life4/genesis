# Slice.Scan

Scan is like Reduce, but returns slice of f results

## Source

```go
// Scan is like Reduce, but returns slice of f results
func (s Slice) Scan(acc G, f func(el T, acc G) G) []G {
	result := make([]G, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}
```