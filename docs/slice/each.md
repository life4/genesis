# Slice.Each

Each calls f for every element from arr

## Source

```go
// Each calls f for every element from arr
func (s Slice) Each(f func(el T)) {
	for _, el := range s.data {
		f(el)
	}
}
```