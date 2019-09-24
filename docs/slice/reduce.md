# Slice.Reduce

Reduce applies F to acc and every element in slice of T and returns acc

## Source

```go
// Reduce applies F to acc and every element in slice of T and returns acc
func (s Slice) Reduce(acc G, f func(el T, acc G) G) G {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}
```